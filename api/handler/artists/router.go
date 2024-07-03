package artists

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func RouterArtists(app *fiber.App, db *sqlx.DB, txID string) {
	h := handlerArtists{db: db, txID: txID}
	api := app.Group("/api")
	v1 := api.Group("/v1")
	artists := v1.Group("/artists")
	artists.Post("/", h.CreateArtists)
	artists.Put("/", h.UpdateArtists)
	artists.Get("/", h.GetAllArtists)
	artists.Get("/:id", h.GetArtistsByID)
	artists.Delete("/:id", h.DeleteArtists)
}
