package service

import (
	"encoding/json"
	"go-architecture/domain"
	"go-architecture/service/helper"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (handler *HttpHandler) Register(ctx *fiber.Ctx) error {
	var user domain.User
	req := ctx.Body()
	if err := json.Unmarshal(req, &user); err != nil {
		return ctx.Status(500).JSON(helper.BaseApiResponse{
			Status:  500,
			Message: helper.InternalServerErrorMessage,
		})
	}

	if err := handler.Usecase.CreateUser(ctx.Context(), user); err != nil {
		return ctx.Status(500).JSON(helper.BaseApiResponse{
			Status:  500,
			Message: helper.InternalServerErrorMessage,
		})
	}

	return ctx.Status(200).JSON(helper.BaseApiResponse{
		Status:  200,
		Message: "Success create user",
	})
}

func (handler *HttpHandler) GetUserByEmail(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	request := domain.User{
		Id: uuid.MustParse(id),
	}
	result, err := handler.Usecase.GetUserById(ctx.Context(), request)
	if err != nil {
		return ctx.Status(500).JSON(helper.BaseApiResponse{
			Status:  500,
			Message: helper.InternalServerErrorMessage,
		})
	} else if result.Id == uuid.Nil {
		return ctx.Status(400).JSON(helper.BaseApiResponse{
			Status:  500,
			Message: helper.BadRequestMessage,
		})
	}

	return ctx.Status(200).JSON(helper.BaseApiResponse{
		Status:  200,
		Message: "Success get user",
		Data:    result,
	})
}
