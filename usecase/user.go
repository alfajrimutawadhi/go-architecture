package usecase

import (
	"context"
	"go-architecture/common"
	"go-architecture/domain"

	"github.com/google/uuid"
)

func (usecase *Usecase) CreateUser(ctx context.Context, request domain.User) (id uuid.UUID, err error) {
	id, err = usecase.RespositoryInteractor.CreateUser(ctx, request)
	if err != nil {
		return uuid.Nil, common.WrapError(err, "usecase", "CreateUser")
	}
	return
}

func (usecase *Usecase) GetUserById(ctx context.Context, request domain.User) (result domain.User, err error) {
	result, err = usecase.RespositoryInteractor.GetUserById(ctx, request)
	if err != nil {
		return domain.User{}, common.WrapError(err, "usecase", "GetUserById")
	}
	return
}
