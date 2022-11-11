package core

import (
	"github.com/gofiber/fiber/v2"
)

func RouterCore(app *fiber.App, txID string) {
	h := handlerCore{TxID: txID}
	api := app.Group("/api")
	v1 := api.Group("/v1")
	v1.Post("/user-by-id", h.GetUserByID)
	v1.Post("/user-by-email", h.GetUserByEmail)
	v1.Post("/create-user", h.CreateUser)
}
