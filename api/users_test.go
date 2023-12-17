package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	mockdb "github.com/DMV-Nicolas/tinygram/db/mock"
	db "github.com/DMV-Nicolas/tinygram/db/mongo"
	"github.com/DMV-Nicolas/tinygram/util"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type eqCreateUserParamsMatcher struct {
	arg      db.CreateUserParams
	password string
}

func (e eqCreateUserParamsMatcher) Matches(x interface{}) bool {
	arg, ok := x.(db.CreateUserParams)
	if !ok {
		return false
	}

	err := util.CheckPassword(e.password, arg.HashedPassword)
	if err != nil {
		return false
	}

	e.arg.HashedPassword = arg.HashedPassword
	return reflect.DeepEqual(e.arg, arg)
}

func (e eqCreateUserParamsMatcher) String() string {
	return fmt.Sprintf("matches arg %v and password %v", e.arg, e.password)
}

func TestCreateUserAPI(t *testing.T) {
	user, password := randomUser(t)
	result := &mongo.InsertOneResult{InsertedID: user.ID}
	testCases := []struct {
		name          string
		body          map[string]any
		buildStubs    func(store *mockdb.MockQuerier)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: map[string]any{
				"username":  user.Username,
				"password":  password,
				"full_name": user.FullName,
				"email":     user.Email,
				"avatar":    user.Avatar,
				"gender":    user.Gender,
			},
			buildStubs: func(querier *mockdb.MockQuerier) {
				arg := db.CreateUserParams{
					Username:       user.Username,
					HashedPassword: user.HashedPassword,
					FullName:       user.FullName,
					Email:          user.Email,
					Avatar:         user.Avatar,
					Gender:         user.Gender,
				}

				querier.EXPECT().
					CreateUser(gomock.Any(), eqCreateUserParamsMatcher{arg, password}).
					Times(1).
					Return(result, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusCreated, recorder.Code)
				requireBodyMatchResult(t, recorder.Body, result)
			},
		},
		{
			name: "UsernameTaken",
			body: map[string]any{
				"username":  user.Username,
				"password":  password,
				"full_name": user.FullName,
				"email":     user.Email,
				"avatar":    user.Avatar,
				"gender":    user.Gender,
			},
			buildStubs: func(querier *mockdb.MockQuerier) {
				querier.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Times(1).
					Return(nil, db.ErrUsernameTaken)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "EmailTaken",
			body: map[string]any{
				"username":  user.Username,
				"password":  password,
				"full_name": user.FullName,
				"email":     user.Email,
				"avatar":    user.Avatar,
				"gender":    user.Gender,
			},
			buildStubs: func(querier *mockdb.MockQuerier) {
				querier.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Times(1).
					Return(nil, db.ErrEmailTaken)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "InternalError",
			body: map[string]any{
				"username":  user.Username,
				"password":  password,
				"full_name": user.FullName,
				"email":     user.Email,
				"avatar":    user.Avatar,
				"gender":    user.Gender,
			},
			buildStubs: func(querier *mockdb.MockQuerier) {
				querier.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Times(1).
					Return(nil, mongo.ErrClientDisconnected)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "TooLongPassword",
			body: map[string]any{
				"username":  user.Username,
				"password":  util.RandomPassword(100),
				"full_name": user.FullName,
				"email":     user.Email,
				"avatar":    user.Avatar,
				"gender":    user.Gender,
			},
			buildStubs: func(querier *mockdb.MockQuerier) {
				querier.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "BodyRequired",
			body: map[string]any{},
			buildStubs: func(querier *mockdb.MockQuerier) {
				querier.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "IncorrectBodyTypes",
			body: map[string]any{
				"username": 5,
			},
			buildStubs: func(querier *mockdb.MockQuerier) {
				querier.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
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
			server := newTestServer(t, queries)
			recorder := httptest.NewRecorder()

			url := "/users"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			request.Header.Add("Content-Type", "application/json")
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func TestLoginUserAPI(t *testing.T) {
	user, password := randomUser(t)
	testCases := []struct {
		name          string
		body          map[string]any
		buildStubs    func(store *mockdb.MockQuerier)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK-UsernameLogin",
			body: map[string]any{
				"username": user.Username,
				"email":    "",
				"password": password,
			},
			buildStubs: func(querier *mockdb.MockQuerier) {
				querier.EXPECT().
					GetUser(gomock.Any(), gomock.Eq("username"), gomock.Eq(user.Username)).
					Times(1).
					Return(user, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name: "OK-EmailLogin",
			body: map[string]any{
				"username": "",
				"email":    user.Email,
				"password": password,
			},
			buildStubs: func(querier *mockdb.MockQuerier) {
				querier.EXPECT().
					GetUser(gomock.Any(), gomock.Eq("email"), gomock.Eq(user.Email)).
					Times(1).
					Return(user, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name: "IncorrectPassword",
			body: map[string]any{
				"username": user.Username,
				"email":    "",
				"password": "incorrect-password",
			},
			buildStubs: func(querier *mockdb.MockQuerier) {
				querier.EXPECT().
					GetUser(gomock.Any(), gomock.Eq("username"), gomock.Eq(user.Username)).
					Times(1).
					Return(user, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "UserNotFound",
			body: map[string]any{
				"username": user.Username,
				"email":    "",
				"password": password,
			},
			buildStubs: func(querier *mockdb.MockQuerier) {
				querier.EXPECT().
					GetUser(gomock.Any(), gomock.Eq("username"), gomock.Any()).
					Times(1).
					Return(db.User{}, mongo.ErrNoDocuments)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name: "InternalError",
			body: map[string]any{
				"username": user.Username,
				"email":    "",
				"password": password,
			},
			buildStubs: func(querier *mockdb.MockQuerier) {
				querier.EXPECT().
					GetUser(gomock.Any(), gomock.Eq("username"), gomock.Any()).
					Times(1).
					Return(db.User{}, mongo.ErrClientDisconnected)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "NoUsernameOrEmailProvided",
			body: map[string]any{
				"username": "",
				"email":    "",
				"password": password,
			},
			buildStubs: func(querier *mockdb.MockQuerier) {
				querier.EXPECT().
					GetUser(gomock.Any(), gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "IncorrectBodyTypes",
			body: map[string]any{
				"username": 100,
				"email":    "",
				"password": password,
			},
			buildStubs: func(querier *mockdb.MockQuerier) {
				querier.EXPECT().
					GetUser(gomock.Any(), gomock.Any(), gomock.Any()).
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
			server := newTestServer(t, queries)
			recorder := httptest.NewRecorder()

			url := "/users/login"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			request.Header.Add("Content-Type", "application/json")
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func TestGetUserAPI(t *testing.T) {
	user, _ := randomUser(t)
	testCases := []struct {
		name          string
		username      any
		buildStubs    func(store *mockdb.MockQuerier)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:     "OK",
			username: user.Username,
			buildStubs: func(querier *mockdb.MockQuerier) {
				querier.EXPECT().
					GetUser(gomock.Any(), gomock.Eq("username"), gomock.Eq(user.Username)).
					Times(1).
					Return(user, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchUser(t, recorder.Body, user)
			},
		},
		{
			name:     "UserNotFound",
			username: user.Username,
			buildStubs: func(querier *mockdb.MockQuerier) {
				querier.EXPECT().
					GetUser(gomock.Any(), gomock.Eq("username"), gomock.Any()).
					Times(1).
					Return(db.User{}, mongo.ErrNoDocuments)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:     "InternalError",
			username: user.Username,
			buildStubs: func(querier *mockdb.MockQuerier) {
				querier.EXPECT().
					GetUser(gomock.Any(), gomock.Eq("username"), gomock.Any()).
					Times(1).
					Return(db.User{}, mongo.ErrClientDisconnected)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:     "NonAlphanumericUsername",
			username: "Fiesta Pagana!",
			buildStubs: func(querier *mockdb.MockQuerier) {
				querier.EXPECT().
					GetUser(gomock.Any(), gomock.Any(), gomock.Any()).
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
			server := newTestServer(t, queries)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/users/%v", tc.username)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			request.Header.Add("Content-Type", "application/json")
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func requireBodyMatchResult(t *testing.T, body *bytes.Buffer, result *mongo.InsertOneResult) {
	bodyResult := new(mongo.InsertOneResult)
	err := json.NewDecoder(body).Decode(bodyResult)
	require.NoError(t, err)
	require.NotEmpty(t, bodyResult)

	bodyInsertedID, err := primitive.ObjectIDFromHex(bodyResult.InsertedID.(string))
	require.NoError(t, err)

	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	require.True(t, ok)

	require.Equal(t, insertedID, bodyInsertedID)
}

func requireBodyMatchUser(t *testing.T, body *bytes.Buffer, user db.User) {
	bodyResult := new(db.User)
	err := json.NewDecoder(body).Decode(bodyResult)
	require.NoError(t, err)
	require.NotEmpty(t, bodyResult)

	require.Equal(t, bodyResult.ID, user.ID)
	require.Equal(t, bodyResult.Username, user.Username)
	require.Equal(t, bodyResult.Email, user.Email)
	require.Equal(t, bodyResult.Description, user.Description)
	require.Equal(t, bodyResult.FullName, user.FullName)
	require.Equal(t, bodyResult.Avatar, user.Avatar)

	require.WithinDuration(t, bodyResult.CreatedAt, user.CreatedAt, time.Second)
}

func randomUser(t *testing.T) (db.User, string) {
	password := util.RandomPassword(16)
	hashedPassword, err := util.HashPassword(password)
	require.NoError(t, err)

	return db.User{
		ID:             primitive.NewObjectID(),
		Username:       util.RandomUsername(),
		HashedPassword: hashedPassword,
		FullName:       util.RandomUsername(),
		Email:          util.RandomEmail(),
		Avatar:         util.RandomPassword(10),
		Description:    util.RandomPassword(5),
		Gender:         "male",
		CreatedAt:      time.Now(),
	}, password
}
