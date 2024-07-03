package users

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func RouterUsers(app *fiber.App, db *sqlx.DB, txID string) {
	h := handlerUsers{db: db, txID: txID}
	api := app.Group("/api")
	v1 := api.Group("/v1")
	users := v1.Group("/users")
	users.Post("/", h.CreateUsers)
	users.Put("/", h.UpdateUsers)
	users.Get("/", h.GetAllUsers)
	users.Get("/:id", h.GetUsersByID)
	users.Delete("/:id", h.DeleteUsers)
}
