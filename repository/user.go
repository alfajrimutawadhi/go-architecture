package repository

import (
	"context"
	"go-architecture/common"
	"go-architecture/domain"
	"go-architecture/domain/constant"

	"github.com/google/uuid"
)

func (db *Database) CreateUser(ctx context.Context, request domain.User) (err error) {
	if request.Id == uuid.Nil {
		request.Id = uuid.New()
	}
	request.CreatedAt = common.CurrentTime()
	request.DBAesKey = db.Config.DB.DBAesKey
	_, err = db.NamedExecContext(ctx, constant.INSERT_USER, request)
	if err != nil {
		common.WrapError(err, "sqlx", "ExecContext")
		return
	}

	return
}

func (db *Database) GetUserById(ctx context.Context, request domain.User) (result domain.User, err error) {
	request.DBAesKey = db.Config.DB.DBAesKey
	rows, err := db.NamedQueryContext(ctx, constant.SELECT_USER_BY_ID, request)
	if err != nil {
		common.WrapError(err, "sqlx", "NamedQueryContext")
		return
	}
	rows.Next()
	
	if err = rows.StructScan(&result); err != nil {
		common.WrapError(err, "sqlx", "Scan")
		return
	}
	return
}
