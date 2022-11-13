package core

import (
	"apirest/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func RouterCore(app *fiber.App, txID string) {
	h := handlerCore{TxID: txID}
	api := app.Group("/api")
	v1 := api.Group("/v1")
	v1.Post("/user-by-id", middleware.JWTProtected(), h.GetUserByID)
	v1.Post("/user-by-email", middleware.JWTProtected(), h.GetUserByEmail)
	v1.Post("/create-user", h.CreateUser)
	v1.Post("/delete-user", middleware.JWTProtected(), h.DeleteUser)
	v1.Post("/login", h.Login)
}
