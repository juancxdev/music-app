package albums

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

type ServicesAlbumsRepository interface {
	create(m *Albums) error
	update(m *Albums) error
	delete(id string) error
	getByID(id string) (*Albums, error)
	getAll() ([]*Albums, error)
}

func FactoryStorage(db *sqlx.DB, user *models.Users, txID string) ServicesAlbumsRepository {
	var s ServicesAlbumsRepository
	engine := db.DriverName()
	switch engine {
	case SqlServer:
		return newAlbumsSqlServerRepository(db, user, txID)
	case Postgresql:
		return newAlbumsPsqlRepository(db, user, txID)
	case Oracle:
		return newAlbumsOrclRepository(db, user, txID)
	default:
		logger.Error.Println("el motor de base de datos no está implementado.", engine)
	}
	return s
}
