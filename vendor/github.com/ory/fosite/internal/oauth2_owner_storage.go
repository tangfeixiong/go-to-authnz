// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/ory/fosite/handler/oauth2 (interfaces: ResourceOwnerPasswordCredentialsGrantStorage)

package internal

import (
	context "context"

	gomock "github.com/golang/mock/gomock"
	fosite "github.com/ory/fosite"
)

// Mock of ResourceOwnerPasswordCredentialsGrantStorage interface
type MockResourceOwnerPasswordCredentialsGrantStorage struct {
	ctrl     *gomock.Controller
	recorder *_MockResourceOwnerPasswordCredentialsGrantStorageRecorder
}

// Recorder for MockResourceOwnerPasswordCredentialsGrantStorage (not exported)
type _MockResourceOwnerPasswordCredentialsGrantStorageRecorder struct {
	mock *MockResourceOwnerPasswordCredentialsGrantStorage
}

func NewMockResourceOwnerPasswordCredentialsGrantStorage(ctrl *gomock.Controller) *MockResourceOwnerPasswordCredentialsGrantStorage {
	mock := &MockResourceOwnerPasswordCredentialsGrantStorage{ctrl: ctrl}
	mock.recorder = &_MockResourceOwnerPasswordCredentialsGrantStorageRecorder{mock}
	return mock
}

func (_m *MockResourceOwnerPasswordCredentialsGrantStorage) EXPECT() *_MockResourceOwnerPasswordCredentialsGrantStorageRecorder {
	return _m.recorder
}

func (_m *MockResourceOwnerPasswordCredentialsGrantStorage) Authenticate(_param0 context.Context, _param1 string, _param2 string) error {
	ret := _m.ctrl.Call(_m, "Authenticate", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockResourceOwnerPasswordCredentialsGrantStorageRecorder) Authenticate(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Authenticate", arg0, arg1, arg2)
}

func (_m *MockResourceOwnerPasswordCredentialsGrantStorage) CreateAccessTokenSession(_param0 context.Context, _param1 string, _param2 fosite.Requester) error {
	ret := _m.ctrl.Call(_m, "CreateAccessTokenSession", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockResourceOwnerPasswordCredentialsGrantStorageRecorder) CreateAccessTokenSession(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateAccessTokenSession", arg0, arg1, arg2)
}

func (_m *MockResourceOwnerPasswordCredentialsGrantStorage) CreateRefreshTokenSession(_param0 context.Context, _param1 string, _param2 fosite.Requester) error {
	ret := _m.ctrl.Call(_m, "CreateRefreshTokenSession", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockResourceOwnerPasswordCredentialsGrantStorageRecorder) CreateRefreshTokenSession(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateRefreshTokenSession", arg0, arg1, arg2)
}

func (_m *MockResourceOwnerPasswordCredentialsGrantStorage) DeleteAccessTokenSession(_param0 context.Context, _param1 string) error {
	ret := _m.ctrl.Call(_m, "DeleteAccessTokenSession", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockResourceOwnerPasswordCredentialsGrantStorageRecorder) DeleteAccessTokenSession(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteAccessTokenSession", arg0, arg1)
}

func (_m *MockResourceOwnerPasswordCredentialsGrantStorage) DeleteRefreshTokenSession(_param0 context.Context, _param1 string) error {
	ret := _m.ctrl.Call(_m, "DeleteRefreshTokenSession", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockResourceOwnerPasswordCredentialsGrantStorageRecorder) DeleteRefreshTokenSession(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteRefreshTokenSession", arg0, arg1)
}

func (_m *MockResourceOwnerPasswordCredentialsGrantStorage) GetAccessTokenSession(_param0 context.Context, _param1 string, _param2 fosite.Session) (fosite.Requester, error) {
	ret := _m.ctrl.Call(_m, "GetAccessTokenSession", _param0, _param1, _param2)
	ret0, _ := ret[0].(fosite.Requester)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockResourceOwnerPasswordCredentialsGrantStorageRecorder) GetAccessTokenSession(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetAccessTokenSession", arg0, arg1, arg2)
}

func (_m *MockResourceOwnerPasswordCredentialsGrantStorage) GetRefreshTokenSession(_param0 context.Context, _param1 string, _param2 fosite.Session) (fosite.Requester, error) {
	ret := _m.ctrl.Call(_m, "GetRefreshTokenSession", _param0, _param1, _param2)
	ret0, _ := ret[0].(fosite.Requester)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockResourceOwnerPasswordCredentialsGrantStorageRecorder) GetRefreshTokenSession(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetRefreshTokenSession", arg0, arg1, arg2)
}
