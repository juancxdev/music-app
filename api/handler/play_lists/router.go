package play_lists

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func RouterPlaylists(app *fiber.App, db *sqlx.DB, txID string) {
	h := handlerPlaylists{db: db, txID: txID}
	api := app.Group("/api")
	v1 := api.Group("/v1")
	playlists := v1.Group("/playlists")
	playlists.Post("/", h.CreatePlaylists)
	playlists.Put("/", h.UpdatePlaylists)
	playlists.Get("/", h.GetAllPlaylists)
	playlists.Get("/:id", h.GetPlaylistsByID)
	playlists.Delete("/:id", h.DeletePlaylists)
}
