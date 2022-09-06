package service

import "github.com/gofiber/fiber/v2"

func (handler *HttpHandler) Router(app *fiber.App) {
	// user
	app.Post("user/register", handler.Register)
	app.Get("user/:id", handler.GetUserById)
}