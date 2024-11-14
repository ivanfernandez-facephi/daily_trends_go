// Code generated by mockery v2.46.3. DO NOT EDIT.

package mocks

import (
	context "context"
	domain "daily-trends/go/internal/feeds/domain"

	mock "github.com/stretchr/testify/mock"
)

// FeedRepository is an autogenerated mock type for the FeedRepository type
type FeedRepository struct {
	mock.Mock
}

// FindById provides a mock function with given fields: ctx, id
func (_m *FeedRepository) FindById(ctx context.Context, id domain.FeedId) (*domain.Feed, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for FindById")
	}

	var r0 *domain.Feed
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.FeedId) (*domain.Feed, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.FeedId) *domain.Feed); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Feed)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.FeedId) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: ctx, feed
func (_m *FeedRepository) Save(ctx context.Context, feed *domain.Feed) error {
	ret := _m.Called(ctx, feed)

	if len(ret) == 0 {
		panic("no return value specified for Save")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Feed) error); ok {
		r0 = rf(ctx, feed)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewFeedRepository creates a new instance of FeedRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewFeedRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *FeedRepository {
	mock := &FeedRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}