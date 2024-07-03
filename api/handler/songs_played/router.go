package songs_played

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func RouterSongsPlayed(app *fiber.App, db *sqlx.DB, txID string) {
	h := handlerSongsPlayed{db: db, txID: txID}
	api := app.Group("/api")
	v1 := api.Group("/v1")
	songsPlayed := v1.Group("/songsplayed")
	songsPlayed.Post("/", h.CreateSongsPlayed)
	songsPlayed.Put("/", h.UpdateSongsPlayed)
	songsPlayed.Get("/", h.GetAllSongsPlayed)
	songsPlayed.Get("/:id", h.GetSongsPlayedByID)
	songsPlayed.Delete("/:id", h.DeleteSongsPlayed)
}
