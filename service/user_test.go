package service_test

import (
	"context"
	"errors"
	"go-architecture/config"
	"go-architecture/domain"
	"go-architecture/service"
	usecase_mock "go-architecture/usecase/mock"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func Test_Register(t *testing.T) {
	m := new(usecase_mock.UsecaseMock)
	app := fiber.New()
	app.Use(func(c *fiber.Ctx) error {
		return c.Next()
	})

	type want struct {
		err error
	}

	testCases := []struct {
		name  string
		want  want
		patch func()
	}{
		{
			name: "When_Register_ExpectBeError",
			want: want{
				err: fiber.ErrBadRequest,
			},
			patch: func() {
				app.Post("/register", func(c *fiber.Ctx) error {
					return c.SendString("request")
				})
				m.On("CreateUser", context.Background(), domain.User{}).Return(errors.New("error")).Once()
			},
		},
	}
	for _, tt := range testCases {
		svc := service.NewHttpHandler(m, config.ShareConfig{})
		t.Run(tt.name, func(t *testing.T) {
			tt.patch()
			if err := svc.Register(&fiber.Ctx{}); err != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
