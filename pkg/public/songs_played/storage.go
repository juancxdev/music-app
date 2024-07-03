package songs_played

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

type ServicesSongsPlayedRepository interface {
	create(m *SongsPlayed) error
	update(m *SongsPlayed) error
	delete(id string) error
	getByID(id string) (*SongsPlayed, error)
	getAll() ([]*SongsPlayed, error)
}

func FactoryStorage(db *sqlx.DB, user *models.Users, txID string) ServicesSongsPlayedRepository {
	var s ServicesSongsPlayedRepository
	engine := db.DriverName()
	switch engine {
	case SqlServer:
		return newSongsPlayedSqlServerRepository(db, user, txID)
	case Postgresql:
		return newSongsPlayedPsqlRepository(db, user, txID)
	case Oracle:
		return newSongsPlayedOrclRepository(db, user, txID)
	default:
		logger.Error.Println("el motor de base de datos no está implementado.", engine)
	}
	return s
}
