package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	mockdb "github.com/DMV-Nicolas/robotgram/db/mock"
	db "github.com/DMV-Nicolas/robotgram/db/mongo"
	"github.com/DMV-Nicolas/robotgram/token"
	"github.com/DMV-Nicolas/robotgram/util"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestRefreshTokenAPI(t *testing.T) {
	tokenSymmetricKey := util.RandomPassword(32)
	maker, err := token.NewPasetoMaker(tokenSymmetricKey)
	require.NoError(t, err)

	user, _ := randomUser(t)
	session := randomSession(t, user.ID, time.Minute, false, maker)
	expiredSession := randomSession(t, user.ID, -time.Minute, false, maker)
	pepitoSession := randomSession(t, primitive.NewObjectID(), time.Minute, false, maker)
	blockedSession := randomSession(t, user.ID, time.Minute, true, maker)

	testCases := []struct {
		name          string
		body          map[string]any
		buildStubs    func(store *mockdb.MockQuerier)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: map[string]any{
				"refresh_token": session.RefreshToken,
			},
			buildStubs: func(querier *mockdb.MockQuerier) {
				querier.EXPECT().
					GetSession(gomock.Any(), gomock.Eq(session.ID)).
					Times(1).
					Return(session, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name: "MistmachedSessionToken",
			body: map[string]any{
				"refresh_token": session.RefreshToken,
			},
			buildStubs: func(querier *mockdb.MockQuerier) {
				querier.EXPECT().
					GetSession(gomock.Any(), gomock.Eq(session.ID)).
					Times(1).
					Return(expiredSession, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "IncorrectSessionUser",
			body: map[string]any{
				"refresh_token": session.RefreshToken,
			},
			buildStubs: func(querier *mockdb.MockQuerier) {
				querier.EXPECT().
					GetSession(gomock.Any(), gomock.Eq(session.ID)).
					Times(1).
					Return(pepitoSession, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "BlockedSession",
			body: map[string]any{
				"refresh_token": blockedSession.RefreshToken,
			},
			buildStubs: func(querier *mockdb.MockQuerier) {
				querier.EXPECT().
					GetSession(gomock.Any(), gomock.Eq(blockedSession.ID)).
					Times(1).
					Return(blockedSession, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "SessionNotFound",
			body: map[string]any{
				"refresh_token": session.RefreshToken,
			},
			buildStubs: func(querier *mockdb.MockQuerier) {
				querier.EXPECT().
					GetSession(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.Session{}, mongo.ErrNoDocuments)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name: "InternalError",
			body: map[string]any{
				"refresh_token": session.RefreshToken,
			},
			buildStubs: func(querier *mockdb.MockQuerier) {
				querier.EXPECT().
					GetSession(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.Session{}, mongo.ErrClientDisconnected)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "InvalidToken",
			body: map[string]any{
				"refresh_token": "IAmAnInvalidToken:(",
			},
			buildStubs: func(querier *mockdb.MockQuerier) {
				querier.EXPECT().
					GetSession(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "NoToken",
			body: map[string]any{
				"refresh_token": "",
			},
			buildStubs: func(querier *mockdb.MockQuerier) {
				querier.EXPECT().
					GetSession(gomock.Any(), gomock.Any()).
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
			server := newTestServer(t, queries, tokenSymmetricKey)
			recorder := httptest.NewRecorder()

			url := "/token/refresh"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			request.Header.Add("Content-Type", "application/json")
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func randomSession(t *testing.T, userID primitive.ObjectID, duration time.Duration, isBlocked bool, tokenMaker token.Maker) db.Session {

	refreshToken, refreshPayload, err := tokenMaker.CreateToken(userID, duration)
	require.NoError(t, err)
	require.NotEmpty(t, refreshToken)
	require.NotEmpty(t, refreshPayload)

	return db.Session{
		ID:           refreshPayload.ID,
		UserID:       userID,
		RefreshToken: refreshToken,
		UserAgent:    "",
		ClientIP:     "",
		IsBlocked:    isBlocked,
		ExpiresAt:    refreshPayload.ExpiresAt,
	}
}
