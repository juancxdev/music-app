package songs

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func RouterSongs(app *fiber.App, db *sqlx.DB, txID string) {
	h := handlerSongs{db: db, txID: txID}
	api := app.Group("/api")
	v1 := api.Group("/v1")
	songs := v1.Group("/songs")
	songs.Post("/", h.CreateSongs)
	songs.Put("/", h.UpdateSongs)
	songs.Get("/", h.GetAllSongs)
	songs.Get("/:id", h.GetSongsByID)
	songs.Delete("/:id", h.DeleteSongs)
}
