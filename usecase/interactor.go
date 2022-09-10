package usecase

import (
	"context"
	"go-architecture/config"
	"go-architecture/domain"
	"go-architecture/repository"
)

type Usecase struct {
	RespositoryInteractor repository.RespositoryInteractor
	Config                *config.ShareConfig
}

func NewUsecase(repo repository.RespositoryInteractor, config config.ShareConfig) UsecaseInteractor {
	return &Usecase{
		RespositoryInteractor: repo,
		Config: &config,
	}
}

type UsecaseInteractor interface {
	CreateUser(ctx context.Context, request domain.User) (err error)
	GetUserById(ctx context.Context, request domain.User) (result domain.User, err error)
}
