package messages

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"music-app/internal/logger"
	"music-app/internal/models"
)

// psql estructura de conexión a la BD de postgresql
type psql struct {
	DB   *sqlx.DB
	user *models.Users
	TxID string
}

func NewMessagePsqlRepository(db *sqlx.DB, user *models.Users, txID string) *psql {
	return &psql{
		DB:   db,
		user: user,
		TxID: txID,
	}
}

// Create registra en la BD
func (s *psql) Create(m *Message) error {
	const sqlInsert = `INSERT INTO public.messages (type_msg, msg_es, msg_en, user_creator, created_at, updated_at) 
              VALUES (:type_msg, :msg_es, :msg_en, :user_creator, :created_at, :updated_at)`
	_, err := s.DB.NamedExec(sqlInsert, &m)
	if err != nil {
		logger.Error.Printf(s.TxID, " - couldn't insert Message: %v", err)
		return err
	}
	return nil
}

// Update actualiza un registro en la BD
func (s *psql) Update(m *Message) error {
	m.IdUser = s.user.ID
	const sqlUpdate = `UPDATE public.messages SET type_msg=:type_msg, msg_es=:msg_es, msg_en=:msg_en, 
              user_creator=:user_creator, updated_at=:updated_at WHERE id=:id`
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
func (s *psql) Delete(id int) error {
	const sqlDelete = `DELETE FROM public.messages WHERE id = :id `
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
func (s *psql) GetByID(id int) (*Message, error) {
	const sqlGetByID = `SELECT id, type_msg, msg_es, msg_en, user_creator, created_at, updated_at FROM public.messages WHERE id=$1`
	mdl := Message{}
	err := s.DB.Get(&mdl, sqlGetByID, id)
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
func (s *psql) GetAll() ([]*Message, error) {
	var ms []*Message
	const sqlGetAll = `SELECT id, type_msg, msg_es, msg_en, user_creator, created_at, updated_at FROM public.messages`

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
