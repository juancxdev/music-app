package songs_play_list

import (
	"fmt"

	"github.com/asaskevich/govalidator"
	"music-app/internal/logger"
	"music-app/internal/models"
)

type PortsServerSongsOfPlaylist interface {
	CreateSongsOfPlaylist(id string, playlist int, song int) (*SongsOfPlaylist, int, error)
	UpdateSongsOfPlaylist(id string, playlist int, song int) (*SongsOfPlaylist, int, error)
	DeleteSongsOfPlaylist(id string) (int, error)
	GetSongsOfPlaylistByID(id string) (*SongsOfPlaylist, int, error)
	GetAllSongsOfPlaylist() ([]*SongsOfPlaylist, error)
}

type service struct {
	repository ServicesSongsOfPlaylistRepository
	user       *models.Users
	txID       string
}

func NewSongsOfPlaylistService(repository ServicesSongsOfPlaylistRepository, user *models.Users, TxID string) PortsServerSongsOfPlaylist {
	return &service{repository: repository, user: user, txID: TxID}
}

func (s *service) CreateSongsOfPlaylist(id string, playlist int, song int) (*SongsOfPlaylist, int, error) {
	m := NewSongsOfPlaylist(id, playlist, song)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}

	if err := s.repository.create(m); err != nil {
		if err.Error() == "ecatch:108" {
			return m, 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't create SongsOfPlaylist :", err)
		return m, 3, err
	}
	return m, 29, nil
}

func (s *service) UpdateSongsOfPlaylist(id string, playlist int, song int) (*SongsOfPlaylist, int, error) {
	m := NewSongsOfPlaylist(id, playlist, song)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}
	if err := s.repository.update(m); err != nil {
		logger.Error.Println(s.txID, " - couldn't update SongsOfPlaylist :", err)
		return m, 18, err
	}
	return m, 29, nil
}

func (s *service) DeleteSongsOfPlaylist(id string) (int, error) {
	if !govalidator.IsUUID(id) {
		logger.Error.Println(s.txID, " - don't meet validations:", fmt.Errorf("id isn't uuid"))
		return 15, fmt.Errorf("id isn't uuid")
	}

	if err := s.repository.delete(id); err != nil {
		if err.Error() == "ecatch:108" {
			return 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't update row:", err)
		return 20, err
	}
	return 28, nil
}

func (s *service) GetSongsOfPlaylistByID(id string) (*SongsOfPlaylist, int, error) {
	if !govalidator.IsUUID(id) {
		logger.Error.Println(s.txID, " - don't meet validations:", fmt.Errorf("id isn't uuid"))
		return nil, 15, fmt.Errorf("id isn't uuid")
	}
	m, err := s.repository.getByID(id)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn`t getByID row:", err)
		return nil, 22, err
	}
	return m, 29, nil
}

func (s *service) GetAllSongsOfPlaylist() ([]*SongsOfPlaylist, error) {
	return s.repository.getAll()
}
