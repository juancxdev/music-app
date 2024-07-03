package artists

import (
	"fmt"

	"github.com/asaskevich/govalidator"
	"music-app/internal/logger"
	"music-app/internal/models"
)

type PortsServerArtists interface {
	CreateArtists(id string, name string, surname string, nationality string) (*Artists, int, error)
	UpdateArtists(id string, name string, surname string, nationality string) (*Artists, int, error)
	DeleteArtists(id string) (int, error)
	GetArtistsByID(id string) (*Artists, int, error)
	GetAllArtists() ([]*Artists, error)
}

type service struct {
	repository ServicesArtistsRepository
	user       *models.Users
	txID       string
}

func NewArtistsService(repository ServicesArtistsRepository, user *models.Users, TxID string) PortsServerArtists {
	return &service{repository: repository, user: user, txID: TxID}
}

func (s *service) CreateArtists(id string, name string, surname string, nationality string) (*Artists, int, error) {
	m := NewArtists(id, name, surname, nationality)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}

	if err := s.repository.create(m); err != nil {
		if err.Error() == "ecatch:108" {
			return m, 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't create Artists :", err)
		return m, 3, err
	}
	return m, 29, nil
}

func (s *service) UpdateArtists(id string, name string, surname string, nationality string) (*Artists, int, error) {
	m := NewArtists(id, name, surname, nationality)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}
	if err := s.repository.update(m); err != nil {
		logger.Error.Println(s.txID, " - couldn't update Artists :", err)
		return m, 18, err
	}
	return m, 29, nil
}

func (s *service) DeleteArtists(id string) (int, error) {
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

func (s *service) GetArtistsByID(id string) (*Artists, int, error) {
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

func (s *service) GetAllArtists() ([]*Artists, error) {
	return s.repository.getAll()
}
