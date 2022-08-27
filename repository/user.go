package repository

import (
	"context"
	"go-architecture/common"
	"go-architecture/domain"
	"go-architecture/domain/constant"

	"github.com/google/uuid"
)

func (db *Database) CreateUser(ctx context.Context, request domain.User) (err error) {
	request.Id = uuid.New()
	request.CreatedAt = common.CurrentTime()
	request.DBAesKey = db.Config.DB.DBAesKey
	result, err := db.NamedExecContext(ctx, constant.INSERT_USER, request)
	if err != nil {
		common.WrapError(err, "sqlx", "ExecContext")
		return
	}

	_, err = result.LastInsertId()
	if err != nil {
		common.WrapError(err, "sqlx", "LastInsertId")
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
	err = rows.StructScan(&result)
	// err = rows.Scan(
	// 	&result.Id,
	// 	&result.Name,
	// 	&result.Email,
	// 	&result.Password,
	// 	&result.CreatedAt,
	// 	&result.UpdatedAt,
	// )
	if err != nil {
		common.WrapError(err, "sqlx", "Scan")
		return
	}
	return
}
