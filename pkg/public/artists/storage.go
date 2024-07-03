package artists

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

type ServicesArtistsRepository interface {
	create(m *Artists) error
	update(m *Artists) error
	delete(id string) error
	getByID(id string) (*Artists, error)
	getAll() ([]*Artists, error)
}

func FactoryStorage(db *sqlx.DB, user *models.Users, txID string) ServicesArtistsRepository {
	var s ServicesArtistsRepository
	engine := db.DriverName()
	switch engine {
	case SqlServer:
		return newArtistsSqlServerRepository(db, user, txID)
	case Postgresql:
		return newArtistsPsqlRepository(db, user, txID)
	case Oracle:
		return newArtistsOrclRepository(db, user, txID)
	default:
		logger.Error.Println("el motor de base de datos no está implementado.", engine)
	}
	return s
}
