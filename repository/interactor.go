package repository

import (
	"context"
	"go-architecture/config"
	"go-architecture/domain"

	"github.com/jmoiron/sqlx"
)

type Database struct {
	*sqlx.DB
	Config *config.ShareConfig
}

func NewRepository(db *sqlx.DB, config config.ShareConfig) RespositoryInteractor {
	return &Database{
		db,
		&config,
	}
}

type RespositoryInteractor interface {
	CreateUser(ctx context.Context, request domain.User) (err error)
	GetUserById(ctx context.Context, request domain.User) (result domain.User, err error)
}
