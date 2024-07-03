package songs_played

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

func newSongsPlayedPsqlRepository(db *sqlx.DB, user *models.Users, txID string) *psql {
	return &psql{
		DB:   db,
		user: user,
		TxID: txID,
	}
}

// Create registra en la BD
func (s *psql) create(m *SongsPlayed) error {
	date := time.Now()
	m.UserCreator = s.user.ID
	m.UpdatedAt = date
	m.CreatedAt = date
	const psqlInsert = `INSERT INTO public.songs_played (id ,user, song, date, user_creator, created_at, updated_at) VALUES (:id ,:user, :song, :date, :user_creator,:created_at, :updated_at) `
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
func (s *psql) update(m *SongsPlayed) error {
	date := time.Now()
	m.UpdatedAt = date
	m.UserCreator = s.user.ID
	const psqlUpdate = `UPDATE public.songs_played SET user = :user, song = :song, date = :date, user_creator = :user_creator, updated_at = :updated_at WHERE id = :id `
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
	const psqlLogicalDelete = `UPDATE public.songs_played SET is_deleted = true, user_deleter = :user_deleter, deleted_at = now() WHERE id = :id`
	m := SongsPlayed{ID: id, UserDeleter: &s.user.ID}
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
func (s *psql) getByID(id string) (*SongsPlayed, error) {
	const psqlGetByID = `SELECT id , user, song, date, is_deleted, user_deleter, deleted_at, user_creator, created_at, updated_at FROM public.songs_played WHERE id = $1 `
	mdl := SongsPlayed{}
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
func (s *psql) getAll() ([]*SongsPlayed, error) {
	var ms []*SongsPlayed
	const psqlGetAll = ` SELECT id , user, song, date, is_deleted, user_deleter, deleted_at, user_creator, created_at, updated_at FROM public.songs_played `

	err := s.DB.Select(&ms, psqlGetAll)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}
