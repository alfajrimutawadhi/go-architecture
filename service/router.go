package service

import "github.com/gofiber/fiber/v2"

func (handler *HttpHandler) Router(app *fiber.App) {
	// user
	user := app.Group("/user")
	user.Post("/register", handler.Register)
	user.Get("/:id", handler.GetUserByEmail)
}