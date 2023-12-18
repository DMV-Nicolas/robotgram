package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	mockdb "github.com/DMV-Nicolas/tinygram/db/mock"
	db "github.com/DMV-Nicolas/tinygram/db/mongo"
	"github.com/DMV-Nicolas/tinygram/token"
	"github.com/DMV-Nicolas/tinygram/util"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestCreatePostAPI(t *testing.T) {
	user, _ := randomUser(t)
	post := randomPost(t, user.ID)
	result := &mongo.InsertOneResult{InsertedID: post.ID}

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
				"images":      post.Images,
				"videos":      post.Videos,
				"description": post.Description,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, time.Minute)
			}, buildStubs: func(querier *mockdb.MockQuerier) {
				querier.EXPECT().
					GetUser(gomock.Any(), gomock.Eq("_id"), gomock.Eq(user.ID)).
					Times(1).
					Return(user, nil)

				arg := db.CreatePostParams{
					UserID:      user.ID,
					Images:      post.Images,
					Videos:      post.Videos,
					Description: post.Description,
				}

				querier.EXPECT().
					CreatePost(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return(result, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusCreated, recorder.Code)
				requireBodyMatchResult(t, recorder.Body, result)
			},
		},
		{
			name: "InternalError",
			body: map[string]any{
				"images":      post.Images,
				"videos":      post.Videos,
				"description": post.Description,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, time.Minute)
			}, buildStubs: func(querier *mockdb.MockQuerier) {
				querier.EXPECT().
					GetUser(gomock.Any(), gomock.Eq("_id"), gomock.Eq(user.ID)).
					Times(1).
					Return(user, nil)
				querier.EXPECT().
					CreatePost(gomock.Any(), gomock.Any()).
					Times(1).
					Return(nil, mongo.ErrClientDisconnected)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "IncorrectBodyTypes",
			body: map[string]any{
				"images":      "incorrect",
				"videos":      "incorrect",
				"description": 100,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, time.Minute)
			}, buildStubs: func(querier *mockdb.MockQuerier) {
				querier.EXPECT().GetUser(gomock.Any(), gomock.Any(), gomock.Any()).Times(0)
				querier.EXPECT().CreatePost(gomock.Any(), gomock.Any()).Times(0)
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
			server := newTestServer(t, queries)
			recorder := httptest.NewRecorder()

			url := "/posts"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			request.Header.Add("Content-Type", "application/json")
			require.NoError(t, err)

			tc.setupAuth(t, request, server.tokenMaker)
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func randomPost(t *testing.T, userID primitive.ObjectID) db.Post {
	return db.Post{
		ID:          util.RandomID(),
		UserID:      userID,
		Images:      util.RandomImages(1),
		Videos:      util.RandomVideos(1),
		Description: util.RandomDescription(100),
	}
}
