package songs_play_list

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

func newSongsOfPlaylistSqlServerRepository(db *sqlx.DB, user *models.Users, txID string) *sqlserver {
	return &sqlserver{
		DB:   db,
		user: user,
		TxID: txID,
	}
}

// Create registra en la BD
func (s *sqlserver) create(m *SongsOfPlaylist) error {
	date := time.Now()
	m.UpdatedAt = date
	m.CreatedAt = date
	m.UserCreator = s.user.ID
	const sqlInsert = `INSERT INTO public.songs_play_list (id ,playlist, song, user_creator, created_at, updated_at) VALUES (:id ,:playlist, :song:user_creator, :created_at, :updated_at) `
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
func (s *sqlserver) update(m *SongsOfPlaylist) error {
	date := time.Now()
	m.UpdatedAt = date
	const sqlUpdate = `UPDATE public.songs_play_list SET playlist = :playlist, song = :song, user_creator = :user_creator, updated_at = :updated_at WHERE id = :id `
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
	const psqlLogicalDelete = `UPDATE public.songs_play_list SET is_deleted = 1, user_deleter = :user_deleter, deleted_at = :deleted_at WHERE id = :id`
	m := SongsOfPlaylist{ID: id, UserDeleter: &s.user.ID, DeletedAt: &date}
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
func (s *sqlserver) getByID(id string) (*SongsOfPlaylist, error) {
	const sqlGetByID = `SELECT convert(nvarchar(50), id) id , playlist, song, is_deleted, user_deleter, deleted_at, user_creator, created_at, updated_at FROM public.songs_play_list  WITH (NOLOCK)  WHERE id = @id `
	mdl := SongsOfPlaylist{}
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
func (s *sqlserver) getAll() ([]*SongsOfPlaylist, error) {
	var ms []*SongsOfPlaylist
	const sqlGetAll = `SELECT convert(nvarchar(50), id) id , playlist, song, is_deleted, user_deleter, deleted_at, user_creator, created_at, updated_at FROM public.songs_play_list  WITH (NOLOCK) `

	err := s.DB.Select(&ms, sqlGetAll)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}
