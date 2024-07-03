package play_lists

import (
	"fmt"

	"github.com/asaskevich/govalidator"
	"music-app/internal/logger"
	"music-app/internal/models"
)

type PortsServerPlaylists interface {
	CreatePlaylists(id string, name string, user int) (*Playlists, int, error)
	UpdatePlaylists(id string, name string, user int) (*Playlists, int, error)
	DeletePlaylists(id string) (int, error)
	GetPlaylistsByID(id string) (*Playlists, int, error)
	GetAllPlaylists() ([]*Playlists, error)
}

type service struct {
	repository ServicesPlaylistsRepository
	user       *models.Users
	txID       string
}

func NewPlaylistsService(repository ServicesPlaylistsRepository, user *models.Users, TxID string) PortsServerPlaylists {
	return &service{repository: repository, user: user, txID: TxID}
}

func (s *service) CreatePlaylists(id string, name string, user int) (*Playlists, int, error) {
	m := NewPlaylists(id, name, user)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}

	if err := s.repository.create(m); err != nil {
		if err.Error() == "ecatch:108" {
			return m, 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't create Playlists :", err)
		return m, 3, err
	}
	return m, 29, nil
}

func (s *service) UpdatePlaylists(id string, name string, user int) (*Playlists, int, error) {
	m := NewPlaylists(id, name, user)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}
	if err := s.repository.update(m); err != nil {
		logger.Error.Println(s.txID, " - couldn't update Playlists :", err)
		return m, 18, err
	}
	return m, 29, nil
}

func (s *service) DeletePlaylists(id string) (int, error) {
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

func (s *service) GetPlaylistsByID(id string) (*Playlists, int, error) {
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

func (s *service) GetAllPlaylists() ([]*Playlists, error) {
	return s.repository.getAll()
}
