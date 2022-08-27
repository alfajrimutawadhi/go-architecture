package usecase

import (
	"context"
	"go-architecture/common"
	"go-architecture/domain"
)

func (repo *Repository) CreateUser(ctx context.Context, request domain.User) (err error) {
	if err = repo.RespositoryInteractor.CreateUser(ctx, request); err != nil {
		common.WrapError(err, "usecase", "CreateUser")
		return
	}
	return
}

func (repo *Repository) GetUserById(ctx context.Context, request domain.User) (result domain.User, err error) {
	result, err = repo.RespositoryInteractor.GetUserById(ctx, request)
	if err != nil {
		common.WrapError(err, "usecase", "GetUserByEmail")
		return
	}
	return
}