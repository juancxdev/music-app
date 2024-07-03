package messages

import (
	"github.com/jmoiron/sqlx"

	"music-app/internal/logger"
	"music-app/internal/models"
)

const (
	Postgresql = "postgres"
	SqlServer  = "sqlserver"
)

type ServicesMessageRepository interface {
	Create(m *Message) error
	Update(m *Message) error
	Delete(id int) error
	GetByID(id int) (*Message, error)
	GetAll() ([]*Message, error)
}

func FactoryStorage(db *sqlx.DB, user *models.Users, txID string) ServicesMessageRepository {
	var s ServicesMessageRepository
	engine := db.DriverName()
	switch engine {
	case SqlServer:
		return NewMessageSqlServerRepository(db, user, txID)
	case Postgresql:
		return NewMessagePsqlRepository(db, user, txID)
	default:
		logger.Error.Println("el motor de base de datos no está implementado.", engine)
	}
	return s
}
