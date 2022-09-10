package usecase

import (
	"context"
	"go-architecture/common"
	"go-architecture/domain"
)

func (usecase *Usecase) CreateUser(ctx context.Context, request domain.User) (err error) {
	if err = usecase.RespositoryInteractor.CreateUser(ctx, request); err != nil {
		common.WrapError(err, "usecase", "CreateUser")
		return
	}
	return
}

func (usecase *Usecase) GetUserById(ctx context.Context, request domain.User) (result domain.User, err error) {
	result, err = usecase.RespositoryInteractor.GetUserById(ctx, request)
	if err != nil {
		common.WrapError(err, "usecase", "GetUserById")
		return
	}
	return
}