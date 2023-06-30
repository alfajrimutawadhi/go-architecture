package repository

import (
	"context"
	"go-architecture/common"
	"go-architecture/domain"
	"go-architecture/domain/constant"

	"github.com/google/uuid"
)

func (repo *Repository) CreateUser(ctx context.Context, request domain.User) (err error) {
	if request.Id == uuid.Nil {
		request.Id = uuid.New()
	}
	request.CreatedAt = common.CurrentTime()
	request.DBAesKey = repo.Config.DB.DBAesKey
	_, err = repo.DB.NamedExecContext(ctx, constant.INSERT_USER, request)
	if err != nil {
		return common.WrapError(err, "sqlx", "ExecContext")
	}

	return
}

func (repo *Repository) GetUserById(ctx context.Context, request domain.User) (result domain.User, err error) {
	request.DBAesKey = repo.Config.DB.DBAesKey
	rows, err := repo.DB.NamedQueryContext(ctx, constant.SELECT_USER_BY_ID, request)
	if err != nil {
		return domain.User{}, common.WrapError(err, "sqlx", "NamedQueryContext")
	}
	rows.Next()

	if err = rows.StructScan(&result); err != nil {
		return domain.User{}, common.WrapError(err, "sqlx", "Scan")
	}
	return
}
