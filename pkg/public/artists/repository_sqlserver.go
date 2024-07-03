package artists

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"music-app/internal/models"
)

// sqlServer estructura de conexión a la BD de mssql
type sqlserver struct {
	DB   *sqlx.DB
	user *models.Users
	TxID string
}

func newArtistsSqlServerRepository(db *sqlx.DB, user *models.Users, txID string) *sqlserver {
	return &sqlserver{
		DB:   db,
		user: user,
		TxID: txID,
	}
}

// Create registra en la BD
func (s *sqlserver) create(m *Artists) error {
	date := time.Now()
	m.UpdatedAt = date
	m.CreatedAt = date
	m.UserCreator = s.user.ID
	const sqlInsert = `INSERT INTO public.artists (id ,name, surname, nationality, user_creator, created_at, updated_at) VALUES (:id ,:name, :surname, :nationality:user_creator, :created_at, :updated_at) `
	rs, err := s.DB.NamedExec(sqlInsert, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("ecatch:108")
	}
	return nil
}

// Update actualiza un registro en la BD
func (s *sqlserver) update(m *Artists) error {
	date := time.Now()
	m.UpdatedAt = date
	const sqlUpdate = `UPDATE public.artists SET name = :name, surname = :surname, nationality = :nationality, user_creator = :user_creator, updated_at = :updated_at WHERE id = :id `
	rs, err := s.DB.NamedExec(sqlUpdate, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("ecatch:108")
	}
	return nil
}

// Delete elimina un registro de la BD
func (s *sqlserver) delete(id string) error {

	// Logical delete
	date := time.Now()
	const psqlLogicalDelete = `UPDATE public.artists SET is_deleted = 1, user_deleter = :user_deleter, deleted_at = :deleted_at WHERE id = :id`
	m := Artists{ID: id, UserDeleter: &s.user.ID, DeletedAt: &date}
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
func (s *sqlserver) getByID(id string) (*Artists, error) {
	const sqlGetByID = `SELECT convert(nvarchar(50), id) id , name, surname, nationality, is_deleted, user_deleter, deleted_at, user_creator, created_at, updated_at FROM public.artists  WITH (NOLOCK)  WHERE id = @id `
	mdl := Artists{}
	err := s.DB.Get(&mdl, sqlGetByID, sql.Named("id", id))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return &mdl, err
	}
	return &mdl, nil
}

// GetAll consulta todos los registros de la BD
func (s *sqlserver) getAll() ([]*Artists, error) {
	var ms []*Artists
	const sqlGetAll = `SELECT convert(nvarchar(50), id) id , name, surname, nationality, is_deleted, user_deleter, deleted_at, user_creator, created_at, updated_at FROM public.artists  WITH (NOLOCK) `

	err := s.DB.Select(&ms, sqlGetAll)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}
