package users

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

type ServicesUsersRepository interface {
	create(m *Users) error
	update(m *Users) error
	delete(id string) error
	getByID(id string) (*Users, error)
	getAll() ([]*Users, error)
}

func FactoryStorage(db *sqlx.DB, user *models.Users, txID string) ServicesUsersRepository {
	var s ServicesUsersRepository
	engine := db.DriverName()
	switch engine {
	case SqlServer:
		return newUsersSqlServerRepository(db, user, txID)
	case Postgresql:
		return newUsersPsqlRepository(db, user, txID)
	case Oracle:
		return newUsersOrclRepository(db, user, txID)
	default:
		logger.Error.Println("el motor de base de datos no está implementado.", engine)
	}
	return s
}
