// Code generated by mockery v2.17.0. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/Allen-Career-Institute/test-and-assessment-commons/pkg/data/entity"

	mock "github.com/stretchr/testify/mock"
)

// StudentTestResultRepository is an autogenerated mock type for the StudentTestResultRepository type
type StudentTestResultRepository struct {
	mock.Mock
}

type StudentTestResultRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *StudentTestResultRepository) EXPECT() *StudentTestResultRepository_Expecter {
	return &StudentTestResultRepository_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, e
func (_m *StudentTestResultRepository) Create(ctx context.Context, e *entity.StudentTestResultEntity) error {
	ret := _m.Called(ctx, e)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.StudentTestResultEntity) error); ok {
		r0 = rf(ctx, e)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StudentTestResultRepository_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type StudentTestResultRepository_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - e *entity.StudentTestResultEntity
func (_e *StudentTestResultRepository_Expecter) Create(ctx interface{}, e interface{}) *StudentTestResultRepository_Create_Call {
	return &StudentTestResultRepository_Create_Call{Call: _e.mock.On("Create", ctx, e)}
}

func (_c *StudentTestResultRepository_Create_Call) Run(run func(ctx context.Context, e *entity.StudentTestResultEntity)) *StudentTestResultRepository_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*entity.StudentTestResultEntity))
	})
	return _c
}

func (_c *StudentTestResultRepository_Create_Call) Return(_a0 error) *StudentTestResultRepository_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

// FindByID provides a mock function with given fields: ctx, id
func (_m *StudentTestResultRepository) FindByID(ctx context.Context, id string) (*entity.StudentTestResultEntity, error) {
	ret := _m.Called(ctx, id)

	var r0 *entity.StudentTestResultEntity
	if rf, ok := ret.Get(0).(func(context.Context, string) *entity.StudentTestResultEntity); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.StudentTestResultEntity)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StudentTestResultRepository_FindByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindByID'
type StudentTestResultRepository_FindByID_Call struct {
	*mock.Call
}

// FindByID is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *StudentTestResultRepository_Expecter) FindByID(ctx interface{}, id interface{}) *StudentTestResultRepository_FindByID_Call {
	return &StudentTestResultRepository_FindByID_Call{Call: _e.mock.On("FindByID", ctx, id)}
}

func (_c *StudentTestResultRepository_FindByID_Call) Run(run func(ctx context.Context, id string)) *StudentTestResultRepository_FindByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *StudentTestResultRepository_FindByID_Call) Return(_a0 *entity.StudentTestResultEntity, _a1 error) *StudentTestResultRepository_FindByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// List provides a mock function with given fields: ctx, offset, limit
func (_m *StudentTestResultRepository) List(ctx context.Context, offset uint, limit int) ([]*entity.StudentTestResultEntity, error) {
	ret := _m.Called(ctx, offset, limit)

	var r0 []*entity.StudentTestResultEntity
	if rf, ok := ret.Get(0).(func(context.Context, uint, int) []*entity.StudentTestResultEntity); ok {
		r0 = rf(ctx, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.StudentTestResultEntity)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint, int) error); ok {
		r1 = rf(ctx, offset, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StudentTestResultRepository_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type StudentTestResultRepository_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
//   - offset uint
//   - limit int
func (_e *StudentTestResultRepository_Expecter) List(ctx interface{}, offset interface{}, limit interface{}) *StudentTestResultRepository_List_Call {
	return &StudentTestResultRepository_List_Call{Call: _e.mock.On("List", ctx, offset, limit)}
}

func (_c *StudentTestResultRepository_List_Call) Run(run func(ctx context.Context, offset uint, limit int)) *StudentTestResultRepository_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint), args[2].(int))
	})
	return _c
}

func (_c *StudentTestResultRepository_List_Call) Return(_a0 []*entity.StudentTestResultEntity, _a1 error) *StudentTestResultRepository_List_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Update provides a mock function with given fields: ctx, e
func (_m *StudentTestResultRepository) Update(ctx context.Context, e *entity.StudentTestResultEntity) (*entity.StudentTestResultEntity, error) {
	ret := _m.Called(ctx, e)

	var r0 *entity.StudentTestResultEntity
	if rf, ok := ret.Get(0).(func(context.Context, *entity.StudentTestResultEntity) *entity.StudentTestResultEntity); ok {
		r0 = rf(ctx, e)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.StudentTestResultEntity)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *entity.StudentTestResultEntity) error); ok {
		r1 = rf(ctx, e)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StudentTestResultRepository_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type StudentTestResultRepository_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - e *entity.StudentTestResultEntity
func (_e *StudentTestResultRepository_Expecter) Update(ctx interface{}, e interface{}) *StudentTestResultRepository_Update_Call {
	return &StudentTestResultRepository_Update_Call{Call: _e.mock.On("Update", ctx, e)}
}

func (_c *StudentTestResultRepository_Update_Call) Run(run func(ctx context.Context, e *entity.StudentTestResultEntity)) *StudentTestResultRepository_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*entity.StudentTestResultEntity))
	})
	return _c
}

func (_c *StudentTestResultRepository_Update_Call) Return(_a0 *entity.StudentTestResultEntity, _a1 error) *StudentTestResultRepository_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type mockConstructorTestingTNewStudentTestResultRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewStudentTestResultRepository creates a new instance of StudentTestResultRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewStudentTestResultRepository(t mockConstructorTestingTNewStudentTestResultRepository) *StudentTestResultRepository {
	mock := &StudentTestResultRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}