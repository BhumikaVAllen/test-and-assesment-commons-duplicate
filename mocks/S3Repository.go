// Code generated by mockery v2.17.0. DO NOT EDIT.

package mocks

import (
	request "github.com/Allen-Career-Institute/test-and-assessment-commons/pkg/data/filestore/request"
	mock "github.com/stretchr/testify/mock"
)

// S3Repository is an autogenerated mock type for the S3Repository type
type S3Repository struct {
	mock.Mock
}

type S3Repository_Expecter struct {
	mock *mock.Mock
}

func (_m *S3Repository) EXPECT() *S3Repository_Expecter {
	return &S3Repository_Expecter{mock: &_m.Mock}
}

// DownloadFile provides a mock function with given fields: _a0
func (_m *S3Repository) DownloadFile(_a0 *request.S3Request) ([]byte, error) {
	ret := _m.Called(_a0)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(*request.S3Request) []byte); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*request.S3Request) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// S3Repository_DownloadFile_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DownloadFile'
type S3Repository_DownloadFile_Call struct {
	*mock.Call
}

// DownloadFile is a helper method to define mock.On call
//   - _a0 *request.S3Request
func (_e *S3Repository_Expecter) DownloadFile(_a0 interface{}) *S3Repository_DownloadFile_Call {
	return &S3Repository_DownloadFile_Call{Call: _e.mock.On("DownloadFile", _a0)}
}

func (_c *S3Repository_DownloadFile_Call) Run(run func(_a0 *request.S3Request)) *S3Repository_DownloadFile_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*request.S3Request))
	})
	return _c
}

func (_c *S3Repository_DownloadFile_Call) Return(_a0 []byte, _a1 error) *S3Repository_DownloadFile_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GeneratePreSignedURl provides a mock function with given fields: _a0
func (_m *S3Repository) GeneratePreSignedURl(_a0 *request.S3Request) (*string, error) {
	ret := _m.Called(_a0)

	var r0 *string
	if rf, ok := ret.Get(0).(func(*request.S3Request) *string); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*request.S3Request) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// S3Repository_GeneratePreSignedURl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GeneratePreSignedURl'
type S3Repository_GeneratePreSignedURl_Call struct {
	*mock.Call
}

// GeneratePreSignedURl is a helper method to define mock.On call
//   - _a0 *request.S3Request
func (_e *S3Repository_Expecter) GeneratePreSignedURl(_a0 interface{}) *S3Repository_GeneratePreSignedURl_Call {
	return &S3Repository_GeneratePreSignedURl_Call{Call: _e.mock.On("GeneratePreSignedURl", _a0)}
}

func (_c *S3Repository_GeneratePreSignedURl_Call) Run(run func(_a0 *request.S3Request)) *S3Repository_GeneratePreSignedURl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*request.S3Request))
	})
	return _c
}

func (_c *S3Repository_GeneratePreSignedURl_Call) Return(_a0 *string, _a1 error) *S3Repository_GeneratePreSignedURl_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// UploadFile provides a mock function with given fields: _a0
func (_m *S3Repository) UploadFile(_a0 *request.S3Request) (*string, error) {
	ret := _m.Called(_a0)

	var r0 *string
	if rf, ok := ret.Get(0).(func(*request.S3Request) *string); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*request.S3Request) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// S3Repository_UploadFile_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UploadFile'
type S3Repository_UploadFile_Call struct {
	*mock.Call
}

// UploadFile is a helper method to define mock.On call
//   - _a0 *request.S3Request
func (_e *S3Repository_Expecter) UploadFile(_a0 interface{}) *S3Repository_UploadFile_Call {
	return &S3Repository_UploadFile_Call{Call: _e.mock.On("UploadFile", _a0)}
}

func (_c *S3Repository_UploadFile_Call) Run(run func(_a0 *request.S3Request)) *S3Repository_UploadFile_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*request.S3Request))
	})
	return _c
}

func (_c *S3Repository_UploadFile_Call) Return(_a0 *string, _a1 error) *S3Repository_UploadFile_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type mockConstructorTestingTNewS3Repository interface {
	mock.TestingT
	Cleanup(func())
}

// NewS3Repository creates a new instance of S3Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewS3Repository(t mockConstructorTestingTNewS3Repository) *S3Repository {
	mock := &S3Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
