package repository_mock

import (
	"context"
	"go-architecture/domain"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	mock.Mock
}

func (_m *RepositoryMock) CreateUser(ctx context.Context, request domain.User) (uuid.UUID, error) {
	ret := _m.Called(ctx, request)

	var r0 uuid.UUID
	if rf, ok := ret.Get(0).(func(context.Context, domain.User) uuid.UUID); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Get(0).(uuid.UUID)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, domain.User) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *RepositoryMock) GetUserById(ctx context.Context, request domain.User) (domain.User, error) {
	ret := _m.Called(ctx, request)

	var r0 domain.User
	if rf, ok := ret.Get(0).(func(context.Context, domain.User) domain.User); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, domain.User) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
