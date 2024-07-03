package artists

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"music-app/internal/models"
)

// psql estructura de conexión a la BD de postgresql
type psql struct {
	DB   *sqlx.DB
	user *models.Users
	TxID string
}

func newArtistsPsqlRepository(db *sqlx.DB, user *models.Users, txID string) *psql {
	return &psql{
		DB:   db,
		user: user,
		TxID: txID,
	}
}

// Create registra en la BD
func (s *psql) create(m *Artists) error {
	date := time.Now()
	m.UserCreator = s.user.ID
	m.UpdatedAt = date
	m.CreatedAt = date
	const psqlInsert = `INSERT INTO public.artists (id ,name, surname, nationality, user_creator, created_at, updated_at) VALUES (:id ,:name, :surname, :nationality, :user_creator,:created_at, :updated_at) `
	rs, err := s.DB.NamedExec(psqlInsert, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("ecatch:108")
	}
	return nil
}

// Update actualiza un registro en la BD
func (s *psql) update(m *Artists) error {
	date := time.Now()
	m.UpdatedAt = date
	m.UserCreator = s.user.ID
	const psqlUpdate = `UPDATE public.artists SET name = :name, surname = :surname, nationality = :nationality, user_creator = :user_creator, updated_at = :updated_at WHERE id = :id `
	rs, err := s.DB.NamedExec(psqlUpdate, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("ecatch:108")
	}
	return nil
}

// Delete elimina un registro de la BD
func (s *psql) delete(id string) error {

	// Logical delete
	const psqlLogicalDelete = `UPDATE public.artists SET is_deleted = true, user_deleter = :user_deleter, deleted_at = now() WHERE id = :id`
	m := Artists{ID: id, UserDeleter: &s.user.ID}
	rs, err := s.DB.NamedExec(psqlLogicalDelete, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("ecatch:108")
	}

	return nil
}

// GetByID consulta un registro por su ID
func (s *psql) getByID(id string) (*Artists, error) {
	const psqlGetByID = `SELECT id , name, surname, nationality, is_deleted, user_deleter, deleted_at, user_creator, created_at, updated_at FROM public.artists WHERE id = $1 `
	mdl := Artists{}
	err := s.DB.Get(&mdl, psqlGetByID, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return &mdl, err
	}
	return &mdl, nil
}

// GetAll consulta todos los registros de la BD
func (s *psql) getAll() ([]*Artists, error) {
	var ms []*Artists
	const psqlGetAll = ` SELECT id , name, surname, nationality, is_deleted, user_deleter, deleted_at, user_creator, created_at, updated_at FROM public.artists `

	err := s.DB.Select(&ms, psqlGetAll)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}
