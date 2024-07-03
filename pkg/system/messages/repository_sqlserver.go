package messages

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"music-app/internal/logger"
	"music-app/internal/models"
)

// sqlServer estructura de conexión a la BD de mssql
type sqlserver struct {
	DB   *sqlx.DB
	user *models.Users
	TxID string
}

func NewMessageSqlServerRepository(db *sqlx.DB, user *models.Users, txID string) *sqlserver {
	return &sqlserver{
		DB:   db,
		user: user,
		TxID: txID,
	}
}

// Create registra en la BD
func (s *sqlserver) Create(m *Message) error {
	m.IdUser = s.user.ID
	const sqlInsert = `INSERT INTO cfg.messages (id ,spa, eng, type_message, id_user, created_at, updated_at) VALUES (:id ,:spa, :eng, :type_message, :id_user, GetDate(), GetDate()) `
	_, err := s.DB.NamedExec(sqlInsert, &m)
	if err != nil {
		logger.Error.Printf(s.TxID, " - couldn't insert Message: %v", err)
		return err
	}
	return nil
}

// Update actualiza un registro en la BD
func (s *sqlserver) Update(m *Message) error {
	m.IdUser = s.user.ID
	const sqlUpdate = `UPDATE cfg.messages SET spa = :spa, eng = :eng, type_message = :type_message,id_user = :id_user, updated_at = GetDate() WHERE id = :id `
	rs, err := s.DB.NamedExec(sqlUpdate, &m)
	if err != nil {
		logger.Error.Printf(s.TxID, " - couldn't update Message: %v", err)
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("ecatch:108")
	}
	return nil
}

// Delete elimina un registro de la BD
func (s *sqlserver) Delete(id int) error {
	const sqlDelete = `DELETE FROM cfg.messages WHERE id = :id `
	m := Message{ID: id}
	rs, err := s.DB.NamedExec(sqlDelete, &m)
	if err != nil {
		logger.Error.Printf(s.TxID, " - couldn't delete Message: %v", err)
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("ecatch:108")
	}
	return nil
}

// GetByID consulta un registro por su ID
func (s *sqlserver) GetByID(id int) (*Message, error) {
	const sqlGetByID = `SELECT convert(nvarchar(50), id) id , spa, eng, type_message, created_at, updated_at FROM cfg.messages  WITH (NOLOCK)  WHERE id = @id AND is_delete = 0`
	mdl := Message{}
	err := s.DB.Get(&mdl, sqlGetByID, sql.Named("id", id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		logger.Error.Printf(s.TxID, " - couldn't execute GetByID Message: %v", err)
		return &mdl, err
	}
	return &mdl, nil
}

// GetAll consulta todos los registros de la BD
func (s *sqlserver) GetAll() ([]*Message, error) {
	var ms []*Message
	const sqlGetAll = `SELECT convert(nvarchar(50), id) id , spa, eng, type_message, created_at, updated_at FROM cfg.messages WITH (NOLOCK) WHERE is_delete = 0 ORDER BY id`

	err := s.DB.Select(&ms, sqlGetAll)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		logger.Error.Printf(s.TxID, " - couldn't execute GetAll cfg.messages: %v", err)
		return ms, err
	}
	return ms, nil
}
