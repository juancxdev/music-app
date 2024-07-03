package songs

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

type ServicesSongsRepository interface {
	create(m *Songs) error
	update(m *Songs) error
	delete(id string) error
	getByID(id string) (*Songs, error)
	getAll() ([]*Songs, error)
}

func FactoryStorage(db *sqlx.DB, user *models.Users, txID string) ServicesSongsRepository {
	var s ServicesSongsRepository
	engine := db.DriverName()
	switch engine {
	case SqlServer:
		return newSongsSqlServerRepository(db, user, txID)
	case Postgresql:
		return newSongsPsqlRepository(db, user, txID)
	case Oracle:
		return newSongsOrclRepository(db, user, txID)
	default:
		logger.Error.Println("el motor de base de datos no está implementado.", engine)
	}
	return s
}
