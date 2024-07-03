package songs_play_list

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

type ServicesSongsOfPlaylistRepository interface {
	create(m *SongsOfPlaylist) error
	update(m *SongsOfPlaylist) error
	delete(id string) error
	getByID(id string) (*SongsOfPlaylist, error)
	getAll() ([]*SongsOfPlaylist, error)
}

func FactoryStorage(db *sqlx.DB, user *models.Users, txID string) ServicesSongsOfPlaylistRepository {
	var s ServicesSongsOfPlaylistRepository
	engine := db.DriverName()
	switch engine {
	case SqlServer:
		return newSongsOfPlaylistSqlServerRepository(db, user, txID)
	case Postgresql:
		return newSongsOfPlaylistPsqlRepository(db, user, txID)
	case Oracle:
		return newSongsOfPlaylistOrclRepository(db, user, txID)
	default:
		logger.Error.Println("el motor de base de datos no está implementado.", engine)
	}
	return s
}
