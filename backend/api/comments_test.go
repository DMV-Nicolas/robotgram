package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	mockdb "github.com/DMV-Nicolas/robotgram/backend/db/mock"
	db "github.com/DMV-Nicolas/robotgram/backend/db/mongo"
	"github.com/DMV-Nicolas/robotgram/backend/token"
	"github.com/DMV-Nicolas/robotgram/backend/util"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestCreateCommentAPI(t *testing.T) {
	user, _ := randomUser(t)
	post := randomPost(t, primitive.NewObjectID())
	comment := randomComment(t, user.ID, post.ID)
	result := &mongo.InsertOneResult{InsertedID: comment.ID}

	testCases := []struct {
		name          string
		body          map[string]any
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		buildStubs    func(store *mockdb.MockQuerier)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: map[string]any{
				"target_id": post.ID.Hex(),
				"content":   comment.Content,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, time.Minute)
			}, buildStubs: func(querier *mockdb.MockQuerier) {
				arg := db.CreateCommentParams{
					UserID:   user.ID,
					TargetID: post.ID,
					Content:  comment.Content,
				}

				querier.EXPECT().
					CreateComment(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return(result, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusCreated, recorder.Code)
				requireBodyMatchInsertOneResult(t, recorder.Body, result)
			},
		},
		{
			name: "InternalError",
			body: map[string]any{
				"target_id": post.ID.Hex(),
				"content":   comment.Content,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, time.Minute)
			}, buildStubs: func(querier *mockdb.MockQuerier) {
				querier.EXPECT().
					CreateComment(gomock.Any(), gomock.Any()).
					Times(1).
					Return(nil, mongo.ErrClientDisconnected)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "InvalidTargetID",
			body: map[string]any{
				"target_id": "qwertyuiopasdfghjklñzxcv",
				"content":   comment.Content,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, time.Minute)
			}, buildStubs: func(querier *mockdb.MockQuerier) {
				querier.EXPECT().
					CreateComment(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "TargetIDLenIsNot24",
			body: map[string]any{
				"target_id": "X-x",
				"content":   comment.Content,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, time.Minute)
			}, buildStubs: func(querier *mockdb.MockQuerier) {
				querier.EXPECT().
					CreateComment(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			queries := mockdb.NewMockQuerier(ctrl)
			tc.buildStubs(queries)

			// marshal data body to json
			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			// start test server and send request
			server := newTestServer(t, queries, util.RandomPassword(32))
			recorder := httptest.NewRecorder()

			url := "/v1/comments"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			request.Header.Add("Content-Type", "application/json")
			require.NoError(t, err)

			tc.setupAuth(t, request, server.tokenMaker)
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func TestListCommentsAPI(t *testing.T) {
	offset, limit := 5, 10
	post := randomPost(t, primitive.NewObjectID())
	comments := make([]db.Comment, limit-offset)
	for i := 0; i < limit-offset; i++ {
		comments[i] = randomComment(t, primitive.NewObjectID(), post.ID)
	}

	testCases := []struct {
		name          string
		query         map[string]any
		buildStubs    func(store *mockdb.MockQuerier)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			query: map[string]any{
				"target_id": post.ID.Hex(),
				"offset":    offset,
				"limit":     limit,
			},
			buildStubs: func(querier *mockdb.MockQuerier) {
				arg := db.ListCommentsParams{
					TargetID: post.ID,
					Offset:   int64(offset),
					Limit:    int64(limit),
				}

				querier.EXPECT().
					ListComments(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return(comments, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchComments(t, recorder.Body, comments)
			},
		},
		{
			name: "InternalError",
			query: map[string]any{
				"target_id": post.ID.Hex(),
				"offset":    offset,
				"limit":     limit,
			},
			buildStubs: func(querier *mockdb.MockQuerier) {
				querier.EXPECT().
					ListComments(gomock.Any(), gomock.Any()).
					Times(1).
					Return(nil, mongo.ErrClientDisconnected)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "InvalidTargetID",
			query: map[string]any{
				"target_id": "qwertyuiopasdfghjklñzxcv",
				"offset":    offset,
				"limit":     limit,
			},
			buildStubs: func(querier *mockdb.MockQuerier) {
				querier.EXPECT().
					ListComments(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "TargetIDLenIsNot24",
			query: map[string]any{
				"target_id": "s-s",
				"offset":    offset,
				"limit":     limit,
			},
			buildStubs: func(querier *mockdb.MockQuerier) {
				querier.EXPECT().
					ListComments(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			queries := mockdb.NewMockQuerier(ctrl)
			tc.buildStubs(queries)

			// start test server and send request
			server := newTestServer(t, queries, util.RandomPassword(32))
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/v1/comments/%v", tc.query["target_id"])
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			request.Header.Add("Content-Type", "application/json")

			q := request.URL.Query()
			q.Add("offset", fmt.Sprint(tc.query["offset"]))
			q.Add("limit", fmt.Sprint(tc.query["limit"]))
			request.URL.RawQuery = q.Encode()

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func TestUpdateCommentAPI(t *testing.T) {
	user, _ := randomUser(t)
	comment := randomComment(t, user.ID, primitive.NewObjectID())
	result := &mongo.UpdateResult{
		MatchedCount:  1,
		ModifiedCount: 1,
		UpsertedCount: 0,
		UpsertedID:    0,
	}

	testCases := []struct {
		name          string
		body          map[string]any
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		buildStubs    func(store *mockdb.MockQuerier)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: map[string]any{
				"id":      comment.ID,
				"content": comment.Content,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, time.Minute)
			}, buildStubs: func(querier *mockdb.MockQuerier) {
				querier.EXPECT().
					GetComment(gomock.Any(), gomock.Eq(comment.ID)).
					Times(1).
					Return(comment, nil)

				arg := db.UpdateCommentParams{
					ID:      comment.ID,
					Content: comment.Content,
				}

				querier.EXPECT().
					UpdateComment(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return(result, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchUpdateResult(t, recorder.Body, result)
			},
		},
		{
			name: "InternalError",
			body: map[string]any{
				"id":      comment.ID,
				"content": comment.Content,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, time.Minute)
			}, buildStubs: func(querier *mockdb.MockQuerier) {
				querier.EXPECT().
					GetComment(gomock.Any(), gomock.Eq(comment.ID)).
					Times(1).
					Return(comment, nil)
				querier.EXPECT().
					UpdateComment(gomock.Any(), gomock.Any()).
					Times(1).
					Return(nil, mongo.ErrClientDisconnected)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "CommentNotFound",
			body: map[string]any{
				"id":      comment.ID,
				"content": comment.Content,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, time.Minute)
			}, buildStubs: func(querier *mockdb.MockQuerier) {
				querier.EXPECT().
					GetComment(gomock.Any(), gomock.Eq(comment.ID)).
					Times(1).
					Return(comment, nil)
				querier.EXPECT().
					UpdateComment(gomock.Any(), gomock.Any()).
					Times(1).
					Return(nil, mongo.ErrNoDocuments)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name: "NonCommentOwner",
			body: map[string]any{
				"id":      comment.ID,
				"content": comment.Content,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, primitive.NewObjectID(), time.Minute)
			}, buildStubs: func(querier *mockdb.MockQuerier) {
				querier.EXPECT().
					GetComment(gomock.Any(), gomock.Eq(comment.ID)).
					Times(1).
					Return(comment, nil)
				querier.EXPECT().
					UpdateComment(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "InvalidComment",
			body: map[string]any{
				"id":      "qwertyuiopasdfghjklñzxcv",
				"content": comment.Content,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, primitive.NewObjectID(), time.Minute)
			}, buildStubs: func(querier *mockdb.MockQuerier) {
				querier.EXPECT().
					GetComment(gomock.Any(), gomock.Any()).
					Times(0)
				querier.EXPECT().
					UpdateComment(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "CommentIDIsNot24",
			body: map[string]any{
				"id":      ":O",
				"content": comment.Content,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, primitive.NewObjectID(), time.Minute)
			}, buildStubs: func(querier *mockdb.MockQuerier) {
				querier.EXPECT().
					GetComment(gomock.Any(), gomock.Any()).
					Times(0)
				querier.EXPECT().
					UpdateComment(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			queries := mockdb.NewMockQuerier(ctrl)
			tc.buildStubs(queries)

			// marshal data body to json
			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			// start test server and send request
			server := newTestServer(t, queries, util.RandomPassword(32))
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/v1/comments/%v", tc.body["id"])
			request, err := http.NewRequest(http.MethodPut, url, bytes.NewReader(data))
			require.NoError(t, err)

			request.Header.Add("Content-Type", "application/json")

			tc.setupAuth(t, request, server.tokenMaker)
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func TestDeleteCommentAPI(t *testing.T) {
	user, _ := randomUser(t)
	comment := randomComment(t, user.ID, primitive.NewObjectID())
	result := &mongo.DeleteResult{
		DeletedCount: 1,
	}

	tests := []struct {
		name          string
		id            any
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		buildStubs    func(store *mockdb.MockQuerier)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			id:   comment.ID.Hex(),
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, time.Minute)
			},
			buildStubs: func(querier *mockdb.MockQuerier) {
				querier.EXPECT().
					GetComment(gomock.Any(), gomock.Eq(comment.ID)).
					Times(1).
					Return(comment, nil)
				querier.EXPECT().
					DeleteComment(gomock.Any(), gomock.Eq(comment.ID)).
					Times(1).
					Return(result, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchDeleteResult(t, recorder.Body, result)
			},
		},
		{
			name: "InternalError",
			id:   comment.ID.Hex(),
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, time.Minute)
			},
			buildStubs: func(querier *mockdb.MockQuerier) {
				querier.EXPECT().
					GetComment(gomock.Any(), gomock.Eq(comment.ID)).
					Times(1).
					Return(comment, nil)
				querier.EXPECT().
					DeleteComment(gomock.Any(), gomock.Any()).
					Times(1).
					Return(nil, mongo.ErrClientDisconnected)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "CommentNotFound",
			id:   comment.ID.Hex(),
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, time.Minute)
			},
			buildStubs: func(querier *mockdb.MockQuerier) {
				querier.EXPECT().
					GetComment(gomock.Any(), gomock.Eq(comment.ID)).
					Times(1).
					Return(comment, nil)
				querier.EXPECT().
					DeleteComment(gomock.Any(), gomock.Any()).
					Times(1).
					Return(nil, mongo.ErrNoDocuments)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name: "NonCommentOwner",
			id:   comment.ID.Hex(),
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, primitive.NewObjectID(), time.Minute)
			},
			buildStubs: func(querier *mockdb.MockQuerier) {
				querier.EXPECT().
					GetComment(gomock.Any(), gomock.Eq(comment.ID)).
					Times(1).
					Return(comment, nil)
				querier.EXPECT().
					DeleteComment(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "InvalidComment",
			id:   "qwertyuiopasdfghjklñzxcv",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, primitive.NewObjectID(), time.Minute)
			},
			buildStubs: func(querier *mockdb.MockQuerier) {
				querier.EXPECT().
					GetComment(gomock.Any(), gomock.Any()).
					Times(0)
				querier.EXPECT().
					DeleteComment(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "CommentLenIsNot24",
			id:   "D:",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, primitive.NewObjectID(), time.Minute)
			},
			buildStubs: func(querier *mockdb.MockQuerier) {
				querier.EXPECT().
					GetComment(gomock.Any(), gomock.Any()).
					Times(0)
				querier.EXPECT().
					DeleteComment(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "ValidCommentNotFound",
			id:   comment.ID.Hex(),
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, primitive.NewObjectID(), time.Minute)
			},
			buildStubs: func(querier *mockdb.MockQuerier) {
				querier.EXPECT().
					GetComment(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.Comment{}, mongo.ErrNoDocuments)
				querier.EXPECT().
					DeleteComment(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name: "ValidCommentInternalError",
			id:   comment.ID.Hex(),
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, primitive.NewObjectID(), time.Minute)
			},
			buildStubs: func(querier *mockdb.MockQuerier) {
				querier.EXPECT().
					GetComment(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.Comment{}, mongo.ErrClientDisconnected)
				querier.EXPECT().
					DeleteComment(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			queries := mockdb.NewMockQuerier(ctrl)
			tc.buildStubs(queries)

			// start test server and send request
			server := newTestServer(t, queries, util.RandomPassword(32))
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/v1/comments/%v", tc.id)
			request, err := http.NewRequest(http.MethodDelete, url, nil)
			require.NoError(t, err)

			request.Header.Add("Content-Type", "application/json")

			tc.setupAuth(t, request, server.tokenMaker)
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func randomComment(t *testing.T, userID primitive.ObjectID, targetID primitive.ObjectID) db.Comment {
	return db.Comment{
		ID:        primitive.NewObjectID(),
		UserID:    userID,
		TargetID:  targetID,
		Content:   util.RandomString(10),
		CreatedAt: time.Now(),
	}
}
