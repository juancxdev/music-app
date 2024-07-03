package play_lists

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"music-app/internal/models"
)

// Orcl estructura de conexión a la BD de Oracle
type orcl struct {
	DB   *sqlx.DB
	user *models.Users
	TxID string
}

func newPlaylistsOrclRepository(db *sqlx.DB, user *models.Users, txID string) *orcl {
	return &orcl{
		DB:   db,
		user: user,
		TxID: txID,
	}
}

// Create registra en la BD
func (s *orcl) create(m *Playlists) error {
	date := time.Now()
	m.UpdatedAt = date
	m.CreatedAt = date
	const osqlInsert = `INSERT INTO public.play_lists (id ,name, user, created_at, updated_at)  VALUES (:id ,:name, :user,:created_at, :updated_at) `
	rs, err := s.DB.NamedExec(osqlInsert, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("ecatch:108")
	}
	return nil
}

// Update actualiza un registro en la BD
func (s *orcl) update(m *Playlists) error {
	date := time.Now()
	m.UpdatedAt = date
	const osqlUpdate = `UPDATE public.play_lists SET name = :name, user = :user, updated_at = :updated_at WHERE id = :id  `
	rs, err := s.DB.NamedExec(osqlUpdate, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("ecatch:108")
	}
	return nil
}

// Delete elimina un registro de la BD
func (s *orcl) delete(id string) error {
	const osqlDelete = `DELETE FROM public.play_lists WHERE id = :id `
	m := Playlists{ID: id}
	rs, err := s.DB.NamedExec(osqlDelete, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("ecatch:108")
	}
	return nil
}

// GetByID consulta un registro por su ID
func (s *orcl) getByID(id string) (*Playlists, error) {
	const osqlGetByID = `SELECT id , name, user, created_at, updated_at FROM public.play_lists WHERE id = :1 `
	mdl := Playlists{}
	err := s.DB.Get(&mdl, osqlGetByID, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return &mdl, err
	}
	return &mdl, nil
}

// GetAll consulta todos los registros de la BD
func (s *orcl) getAll() ([]*Playlists, error) {
	var ms []*Playlists
	const osqlGetAll = ` SELECT id , name, user, created_at, updated_at FROM public.play_lists `

	err := s.DB.Select(&ms, osqlGetAll)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}
