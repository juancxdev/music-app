package artists

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

func newArtistsOrclRepository(db *sqlx.DB, user *models.Users, txID string) *orcl {
	return &orcl{
		DB:   db,
		user: user,
		TxID: txID,
	}
}

// Create registra en la BD
func (s *orcl) create(m *Artists) error {
	date := time.Now()
	m.UpdatedAt = date
	m.CreatedAt = date
	const osqlInsert = `INSERT INTO public.artists (id ,name, surname, nationality, created_at, updated_at)  VALUES (:id ,:name, :surname, :nationality,:created_at, :updated_at) `
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
func (s *orcl) update(m *Artists) error {
	date := time.Now()
	m.UpdatedAt = date
	const osqlUpdate = `UPDATE public.artists SET name = :name, surname = :surname, nationality = :nationality, updated_at = :updated_at WHERE id = :id  `
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
	const osqlDelete = `DELETE FROM public.artists WHERE id = :id `
	m := Artists{ID: id}
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
func (s *orcl) getByID(id string) (*Artists, error) {
	const osqlGetByID = `SELECT id , name, surname, nationality, created_at, updated_at FROM public.artists WHERE id = :1 `
	mdl := Artists{}
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
func (s *orcl) getAll() ([]*Artists, error) {
	var ms []*Artists
	const osqlGetAll = ` SELECT id , name, surname, nationality, created_at, updated_at FROM public.artists `

	err := s.DB.Select(&ms, osqlGetAll)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}
