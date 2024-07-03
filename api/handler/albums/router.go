package albums

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func RouterAlbums(app *fiber.App, db *sqlx.DB, txID string) {
	h := handlerAlbums{db: db, txID: txID}
	api := app.Group("/api")
	v1 := api.Group("/v1")
	albums := v1.Group("/albums")
	albums.Post("/", h.CreateAlbums)
	albums.Put("/", h.UpdateAlbums)
	albums.Get("/", h.GetAllAlbums)
	albums.Get("/:id", h.GetAlbumsByID)
	albums.Delete("/:id", h.DeleteAlbums)
}
