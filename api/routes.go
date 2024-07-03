package api

import (
	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"music-app/api/handler/albums"
	"music-app/api/handler/artists"
	"music-app/api/handler/play_lists"
	"music-app/api/handler/songs"
	"music-app/api/handler/songs_play_list"
	"music-app/api/handler/songs_played"
	"music-app/api/handler/users"
)

// Se cargan los loggerHttp, y los allowedOrigins (registrosHttp) (permisos de origen)
func routes(loggerHttp bool, allowedOrigins string, db *sqlx.DB) *fiber.App {
	app := fiber.New()

	prometheus := fiberprometheus.New("MUSIC API REST")
	prometheus.RegisterAt(app, "/metrics")

	app.Get("/doc/*", swagger.New(swagger.Config{
		URL:         "/doc/doc.json",
		DeepLinking: false,
	}))

	app.Use(recover.New())
	app.Use(prometheus.Middleware)
	app.Use(cors.New(cors.Config{
		AllowOrigins: allowedOrigins,
		AllowHeaders: "Origin, X-Requested-With, Content-Type, Accept, Authorization, signature",
		AllowMethods: "GET,POST,PUT,PATCH,DELETE",
	}))

	if loggerHttp {
		app.Use(logger.New())
	}

	TxID := uuid.New().String()

	loadRoutes(app, TxID, db)

	return app
}

// Aqui se cargan las direcciones o las ubicaciones de las funciones Handler
func loadRoutes(app *fiber.App, TxID string, db *sqlx.DB) {
	albums.RouterAlbums(app, db, TxID)
	artists.RouterArtists(app, db, TxID)
	play_lists.RouterPlaylists(app, db, TxID)
	songs.RouterSongs(app, db, TxID)
	songs_play_list.RouterSongsOfPlaylist(app, db, TxID)
	songs_played.RouterSongsPlayed(app, db, TxID)
	users.RouterUsers(app, db, TxID)
}
