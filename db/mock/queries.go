// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/DMV-Nicolas/tinygram/db/mongo (interfaces: Querier)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	reflect "reflect"

	db "github.com/DMV-Nicolas/tinygram/db/mongo"
	gomock "github.com/golang/mock/gomock"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
	mongo "go.mongodb.org/mongo-driver/mongo"
)

// MockQuerier is a mock of Querier interface.
type MockQuerier struct {
	ctrl     *gomock.Controller
	recorder *MockQuerierMockRecorder
}

// MockQuerierMockRecorder is the mock recorder for MockQuerier.
type MockQuerierMockRecorder struct {
	mock *MockQuerier
}

// NewMockQuerier creates a new mock instance.
func NewMockQuerier(ctrl *gomock.Controller) *MockQuerier {
	mock := &MockQuerier{ctrl: ctrl}
	mock.recorder = &MockQuerierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQuerier) EXPECT() *MockQuerierMockRecorder {
	return m.recorder
}

// CreateLike mocks base method.
func (m *MockQuerier) CreateLike(arg0 context.Context, arg1 db.CreateLikeParams) (*mongo.InsertOneResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLike", arg0, arg1)
	ret0, _ := ret[0].(*mongo.InsertOneResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateLike indicates an expected call of CreateLike.
func (mr *MockQuerierMockRecorder) CreateLike(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLike", reflect.TypeOf((*MockQuerier)(nil).CreateLike), arg0, arg1)
}

// CreatePost mocks base method.
func (m *MockQuerier) CreatePost(arg0 context.Context, arg1 db.CreatePostParams) (*mongo.InsertOneResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePost", arg0, arg1)
	ret0, _ := ret[0].(*mongo.InsertOneResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePost indicates an expected call of CreatePost.
func (mr *MockQuerierMockRecorder) CreatePost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePost", reflect.TypeOf((*MockQuerier)(nil).CreatePost), arg0, arg1)
}

// CreateUser mocks base method.
func (m *MockQuerier) CreateUser(arg0 context.Context, arg1 db.CreateUserParams) (*mongo.InsertOneResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(*mongo.InsertOneResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockQuerierMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockQuerier)(nil).CreateUser), arg0, arg1)
}

// DeleteLike mocks base method.
func (m *MockQuerier) DeleteLike(arg0 context.Context, arg1 primitive.ObjectID) (*mongo.DeleteResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteLike", arg0, arg1)
	ret0, _ := ret[0].(*mongo.DeleteResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteLike indicates an expected call of DeleteLike.
func (mr *MockQuerierMockRecorder) DeleteLike(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteLike", reflect.TypeOf((*MockQuerier)(nil).DeleteLike), arg0, arg1)
}

// DeletePost mocks base method.
func (m *MockQuerier) DeletePost(arg0 context.Context, arg1 primitive.ObjectID) (*mongo.DeleteResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePost", arg0, arg1)
	ret0, _ := ret[0].(*mongo.DeleteResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeletePost indicates an expected call of DeletePost.
func (mr *MockQuerierMockRecorder) DeletePost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePost", reflect.TypeOf((*MockQuerier)(nil).DeletePost), arg0, arg1)
}

// DeleteUser mocks base method.
func (m *MockQuerier) DeleteUser(arg0 context.Context, arg1 primitive.ObjectID) (*mongo.DeleteResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", arg0, arg1)
	ret0, _ := ret[0].(*mongo.DeleteResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockQuerierMockRecorder) DeleteUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockQuerier)(nil).DeleteUser), arg0, arg1)
}

// GetLike mocks base method.
func (m *MockQuerier) GetLike(arg0 context.Context, arg1 primitive.ObjectID) (db.Like, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLike", arg0, arg1)
	ret0, _ := ret[0].(db.Like)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLike indicates an expected call of GetLike.
func (mr *MockQuerierMockRecorder) GetLike(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLike", reflect.TypeOf((*MockQuerier)(nil).GetLike), arg0, arg1)
}

// GetPost mocks base method.
func (m *MockQuerier) GetPost(arg0 context.Context, arg1 string, arg2 interface{}) (db.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPost", arg0, arg1, arg2)
	ret0, _ := ret[0].(db.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPost indicates an expected call of GetPost.
func (mr *MockQuerierMockRecorder) GetPost(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPost", reflect.TypeOf((*MockQuerier)(nil).GetPost), arg0, arg1, arg2)
}

// GetUser mocks base method.
func (m *MockQuerier) GetUser(arg0 context.Context, arg1 string, arg2 interface{}) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1, arg2)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockQuerierMockRecorder) GetUser(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockQuerier)(nil).GetUser), arg0, arg1, arg2)
}

// ListLikes mocks base method.
func (m *MockQuerier) ListLikes(arg0 context.Context, arg1 db.ListLikesParams) ([]db.Like, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListLikes", arg0, arg1)
	ret0, _ := ret[0].([]db.Like)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListLikes indicates an expected call of ListLikes.
func (mr *MockQuerierMockRecorder) ListLikes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListLikes", reflect.TypeOf((*MockQuerier)(nil).ListLikes), arg0, arg1)
}

// ListPosts mocks base method.
func (m *MockQuerier) ListPosts(arg0 context.Context, arg1 db.ListPostsParams) ([]db.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPosts", arg0, arg1)
	ret0, _ := ret[0].([]db.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPosts indicates an expected call of ListPosts.
func (mr *MockQuerierMockRecorder) ListPosts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPosts", reflect.TypeOf((*MockQuerier)(nil).ListPosts), arg0, arg1)
}

// ListUsers mocks base method.
func (m *MockQuerier) ListUsers(arg0 context.Context, arg1 db.ListUsersParams) ([]db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUsers", arg0, arg1)
	ret0, _ := ret[0].([]db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUsers indicates an expected call of ListUsers.
func (mr *MockQuerierMockRecorder) ListUsers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUsers", reflect.TypeOf((*MockQuerier)(nil).ListUsers), arg0, arg1)
}

// UpdatePost mocks base method.
func (m *MockQuerier) UpdatePost(arg0 context.Context, arg1 db.UpdatePostParams) (*mongo.UpdateResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePost", arg0, arg1)
	ret0, _ := ret[0].(*mongo.UpdateResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdatePost indicates an expected call of UpdatePost.
func (mr *MockQuerierMockRecorder) UpdatePost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePost", reflect.TypeOf((*MockQuerier)(nil).UpdatePost), arg0, arg1)
}

// UpdateUser mocks base method.
func (m *MockQuerier) UpdateUser(arg0 context.Context, arg1 db.UpdateUserParams) (*mongo.UpdateResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", arg0, arg1)
	ret0, _ := ret[0].(*mongo.UpdateResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockQuerierMockRecorder) UpdateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockQuerier)(nil).UpdateUser), arg0, arg1)
}
