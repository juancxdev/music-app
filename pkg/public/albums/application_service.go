package albums

import (
	"fmt"
	"time"

	"github.com/asaskevich/govalidator"
	"music-app/internal/logger"
	"music-app/internal/models"
)

type PortsServerAlbums interface {
	CreateAlbums(id string, name string, artist string, releaseDate time.Time) (*Albums, int, error)
	UpdateAlbums(id string, name string, artist string, releaseDate time.Time) (*Albums, int, error)
	DeleteAlbums(id string) (int, error)
	GetAlbumsByID(id string) (*Albums, int, error)
	GetAllAlbums() ([]*Albums, error)
}

type service struct {
	repository ServicesAlbumsRepository
	user       *models.Users
	txID       string
}

func NewAlbumsService(repository ServicesAlbumsRepository, user *models.Users, TxID string) PortsServerAlbums {
	return &service{repository: repository, user: user, txID: TxID}
}

func (s *service) CreateAlbums(id string, name string, artist string, releaseDate time.Time) (*Albums, int, error) {
	m := NewAlbums(id, name, artist, releaseDate)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}

	if err := s.repository.create(m); err != nil {
		if err.Error() == "ecatch:108" {
			return m, 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't create Albums :", err)
		return m, 3, err
	}
	return m, 29, nil
}

func (s *service) UpdateAlbums(id string, name string, artist string, releaseDate time.Time) (*Albums, int, error) {
	m := NewAlbums(id, name, artist, releaseDate)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}
	if err := s.repository.update(m); err != nil {
		logger.Error.Println(s.txID, " - couldn't update Albums :", err)
		return m, 18, err
	}
	return m, 29, nil
}

func (s *service) DeleteAlbums(id string) (int, error) {
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

func (s *service) GetAlbumsByID(id string) (*Albums, int, error) {
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

func (s *service) GetAllAlbums() ([]*Albums, error) {
	return s.repository.getAll()
}
