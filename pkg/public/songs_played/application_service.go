package songs_played

import (
	"fmt"
	"time"

	"github.com/asaskevich/govalidator"
	"music-app/internal/logger"
	"music-app/internal/models"
)

type PortsServerSongsPlayed interface {
	CreateSongsPlayed(id string, user int, song int, date time.Time) (*SongsPlayed, int, error)
	UpdateSongsPlayed(id string, user int, song int, date time.Time) (*SongsPlayed, int, error)
	DeleteSongsPlayed(id string) (int, error)
	GetSongsPlayedByID(id string) (*SongsPlayed, int, error)
	GetAllSongsPlayed() ([]*SongsPlayed, error)
}

type service struct {
	repository ServicesSongsPlayedRepository
	user       *models.Users
	txID       string
}

func NewSongsPlayedService(repository ServicesSongsPlayedRepository, user *models.Users, TxID string) PortsServerSongsPlayed {
	return &service{repository: repository, user: user, txID: TxID}
}

func (s *service) CreateSongsPlayed(id string, user int, song int, date time.Time) (*SongsPlayed, int, error) {
	m := NewSongsPlayed(id, user, song, date)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}

	if err := s.repository.create(m); err != nil {
		if err.Error() == "ecatch:108" {
			return m, 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't create SongsPlayed :", err)
		return m, 3, err
	}
	return m, 29, nil
}

func (s *service) UpdateSongsPlayed(id string, user int, song int, date time.Time) (*SongsPlayed, int, error) {
	m := NewSongsPlayed(id, user, song, date)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}
	if err := s.repository.update(m); err != nil {
		logger.Error.Println(s.txID, " - couldn't update SongsPlayed :", err)
		return m, 18, err
	}
	return m, 29, nil
}

func (s *service) DeleteSongsPlayed(id string) (int, error) {
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

func (s *service) GetSongsPlayedByID(id string) (*SongsPlayed, int, error) {
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

func (s *service) GetAllSongsPlayed() ([]*SongsPlayed, error) {
	return s.repository.getAll()
}
