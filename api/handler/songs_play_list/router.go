package songs_play_list

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func RouterSongsOfPlaylist(app *fiber.App, db *sqlx.DB, txID string) {
	h := handlerSongsOfPlaylist{db: db, txID: txID}
	api := app.Group("/api")
	v1 := api.Group("/v1")
	songsOfPlaylist := v1.Group("/songs-play-list")
	songsOfPlaylist.Post("/", h.CreateSongsOfPlaylist)
	songsOfPlaylist.Put("/", h.UpdateSongsOfPlaylist)
	songsOfPlaylist.Get("/", h.GetAllSongsOfPlaylist)
	songsOfPlaylist.Get("/:id", h.GetSongsOfPlaylistByID)
	songsOfPlaylist.Delete("/:id", h.DeleteSongsOfPlaylist)
}
