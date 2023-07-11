// Code generated by mockery v2.31.4. DO NOT EDIT.

package repository_mocks

import (
	model "belajar-golang-mock/model"
	context "context"

	mock "github.com/stretchr/testify/mock"

	sql "database/sql"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// FindAllCats provides a mock function with given fields: ctx, tx
func (_m *Repository) FindAllCats(ctx context.Context, tx *sql.Tx) ([]*model.Cat, error) {
	ret := _m.Called(ctx, tx)

	var r0 []*model.Cat
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *sql.Tx) ([]*model.Cat, error)); ok {
		return rf(ctx, tx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *sql.Tx) []*model.Cat); ok {
		r0 = rf(ctx, tx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Cat)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *sql.Tx) error); ok {
		r1 = rf(ctx, tx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertCat provides a mock function with given fields: ctx, tx, cat
func (_m *Repository) InsertCat(ctx context.Context, tx *sql.Tx, cat model.Cat) (int64, error) {
	ret := _m.Called(ctx, tx, cat)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *sql.Tx, model.Cat) (int64, error)); ok {
		return rf(ctx, tx, cat)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *sql.Tx, model.Cat) int64); ok {
		r0 = rf(ctx, tx, cat)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *sql.Tx, model.Cat) error); ok {
		r1 = rf(ctx, tx, cat)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
