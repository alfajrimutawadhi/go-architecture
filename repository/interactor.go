package repository

import (
	"context"
	"go-architecture/config"
	"go-architecture/domain"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	DB *sqlx.DB
	Config *config.ShareConfig
}

func NewRepository(db *sqlx.DB, config config.ShareConfig) RespositoryInteractor {
	return &Repository{
		DB :db,
		Config: &config,
	}
}

type RespositoryInteractor interface {
	CreateUser(ctx context.Context, request domain.User) (id uuid.UUID, err error)
	GetUserById(ctx context.Context, request domain.User) (result domain.User, err error)
}
