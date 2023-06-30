package repository_test

import (
	"context"
	"errors"
	"go-architecture/config"
	"go-architecture/domain"
	"go-architecture/repository"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func Test_CreateUser(t *testing.T) {
	var mockErr = errors.New("error")
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mockId := uuid.New()

	type request struct {
		ctx context.Context
		req domain.User
	}

	type want struct {
		err bool
	}

	testCases := []struct {
		name  string
		args  request
		want  want
		patch func()
	}{
		{
			name: "When_CreateUser_ExpectBeError",
			args: request{
				ctx: context.Background(),
				req: domain.User{},
			},
			want: want{
				err: true,
			},
			patch: func() {
				mock.ExpectExec("INSERT").WillReturnError(mockErr)
			},
		},
		{
			name: "When_CreateUser_ExpectNotBeError",
			args: request{
				ctx: context.Background(),
				req: domain.User{
					Id:       mockId,
					Name:     "test",
					Email:    "test@gmail.com",
					Password: "test",
				},
			},
			want: want{
				err: false,
			},
			patch: func() {
				mock.ExpectExec("INSERT").WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
	}
	for _, tt := range testCases {
		repo := repository.NewRepository(sqlx.NewDb(db, "sqlmock"), config.ShareConfig{})
		t.Run(tt.name, func(t *testing.T) {
			tt.patch()
			_, err = repo.CreateUser(tt.args.ctx, tt.args.req)
			if tt.want.err {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func Test_GetUserById(t *testing.T) {
	var mockErr = errors.New("error")
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	type request struct {
		ctx context.Context
		req domain.User
	}

	type want struct {
		result domain.User
		err    bool
	}

	testCases := []struct {
		name  string
		args  request
		want  want
		patch func()
	}{
		{
			name: "When_GetUserById_ExpectBeError",
			args: request{
				ctx: context.Background(),
				req: domain.User{},
			},
			want: want{
				result: domain.User{},
				err:    true,
			},
			patch: func() {
				mock.ExpectQuery("SELECT").WillReturnError(mockErr)
			},
		},
		{
			name: "When_GetUserById_ErrorRowScan_ExpectBeError",
			args: request{
				ctx: context.Background(),
				req: domain.User{},
			},
			want: want{
				result: domain.User{},
				err:    true,
			},
			patch: func() {
				mock.ExpectQuery("SELECT").WillReturnRows(sqlmock.NewRows([]string{}))
			},
		},
		{
			name: "When_GetUserById_ExpectNotBeError",
			args: request{
				ctx: context.Background(),
				req: domain.User{},
			},
			want: want{
				result: domain.User{},
				err:    false,
			},
			patch: func() {
				mock.ExpectQuery("SELECT").WillReturnRows(sqlmock.NewRows([]string{}).AddRow())
			},
		},
	}
	for _, tt := range testCases {
		repo := repository.NewRepository(sqlx.NewDb(db, "sqlmock"), config.ShareConfig{})
		t.Run(tt.name, func(t *testing.T) {
			tt.patch()
			result, err := repo.GetUserById(tt.args.ctx, tt.args.req)
			assert.Equal(t, tt.want.result, result)
			if tt.want.err {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
