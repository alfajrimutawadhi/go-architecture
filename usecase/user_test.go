package usecase_test

import (
	"context"
	"errors"
	"go-architecture/config"
	"go-architecture/domain"
	repository_mock "go-architecture/repository/mock"
	"go-architecture/usecase"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_CreateUser(t *testing.T) {
	m := new(repository_mock.RepositoryMock)

	type args struct {
		ctx     context.Context
		request domain.User
	}

	type want struct {
		err bool
	}

	testCases := []struct {
		name  string
		args  args
		want  want
		patch func()
	}{
		{
			name: "When_CreateUser_ExpectBeError",
			args: args{
				ctx:     context.Background(),
				request: domain.User{},
			},
			want: want{
				err: true,
			},
			patch: func() {
				m.On("CreateUser", context.Background(), domain.User{}).Return(errors.New("error")).Once()
			},
		},
		{
			name: "When_CreateUser_ExpectNotBeError",
			args: args{
				ctx:     context.Background(),
				request: domain.User{},
			},
			want: want{
				err: false,
			},
			patch: func() {
				m.On("CreateUser", context.Background(), domain.User{}).Return(nil).Once()
			},
		},
	}
	for _, tt := range testCases {
		usecase := usecase.NewUsecase(m, config.ShareConfig{})
		t.Run(tt.name, func(t *testing.T) {
			tt.patch()
			err := usecase.CreateUser(tt.args.ctx, tt.args.request)
			if tt.want.err {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func Test_GetuserById(t *testing.T) {
	m := new(repository_mock.RepositoryMock)
	id := uuid.New()

	type args struct {
		ctx     context.Context
		request domain.User
	}

	type want struct {
		result domain.User
		err    bool
	}

	testCases := []struct {
		name  string
		args  args
		want  want
		patch func()
	}{
		{
			name: "When_GetUserById_ExpectBeError",
			args: args{
				ctx:     context.Background(),
				request: domain.User{},
			},
			want: want{
				result: domain.User{},
				err:    true,
			},
			patch: func() {
				m.On("GetUserById", context.Background(), domain.User{}).Return(domain.User{}, errors.New("error")).Once()
			},
		},
		{
			name: "When_GetUserById_ExpectNotBeError",
			args: args{
				ctx:     context.Background(),
				request: domain.User{},
			},
			want: want{
				result: domain.User{
					Id:    id,
					Name:  "test",
					Email: "test@gmail.com",
				},
				err: false,
			},
			patch: func() {
				m.On("GetUserById", context.Background(), domain.User{}).Return(domain.User{
					Id:    id,
					Name:  "test",
					Email: "test@gmail.com",
				}, nil).Once()
			},
		},
	}
	for _, tt := range testCases {
		uc := usecase.NewUsecase(m, config.ShareConfig{})
		t.Run(tt.name, func(t *testing.T) {
			tt.patch()
			result, err := uc.GetUserById(tt.args.ctx, tt.args.request)
			assert.Equal(t, tt.want.result, result)
			if tt.want.err {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
