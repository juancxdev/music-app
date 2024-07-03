package play_lists

import (
	"github.com/jmoiron/sqlx"

	"music-app/internal/logger"
	"music-app/internal/models"
)

const (
	Postgresql = "postgres"
	SqlServer  = "sqlserver"
	Oracle     = "oci8"
)

type ServicesPlaylistsRepository interface {
	create(m *Playlists) error
	update(m *Playlists) error
	delete(id string) error
	getByID(id string) (*Playlists, error)
	getAll() ([]*Playlists, error)
}

func FactoryStorage(db *sqlx.DB, user *models.Users, txID string) ServicesPlaylistsRepository {
	var s ServicesPlaylistsRepository
	engine := db.DriverName()
	switch engine {
	case SqlServer:
		return newPlaylistsSqlServerRepository(db, user, txID)
	case Postgresql:
		return newPlaylistsPsqlRepository(db, user, txID)
	case Oracle:
		return newPlaylistsOrclRepository(db, user, txID)
	default:
		logger.Error.Println("el motor de base de datos no está implementado.", engine)
	}
	return s
}
