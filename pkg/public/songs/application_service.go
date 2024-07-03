package songs

import (
	"fmt"

	"github.com/asaskevich/govalidator"
	"music-app/internal/logger"
	"music-app/internal/models"
)

type PortsServerSongs interface {
	CreateSongs(id string, name string, artist string, album string) (*Songs, int, error)
	UpdateSongs(id string, name string, artist string, album string) (*Songs, int, error)
	DeleteSongs(id string) (int, error)
	GetSongsByID(id string) (*Songs, int, error)
	GetAllSongs() ([]*Songs, error)
}

type service struct {
	repository ServicesSongsRepository
	user       *models.Users
	txID       string
}

func NewSongsService(repository ServicesSongsRepository, user *models.Users, TxID string) PortsServerSongs {
	return &service{repository: repository, user: user, txID: TxID}
}

func (s *service) CreateSongs(id string, name string, artist string, album string) (*Songs, int, error) {
	m := NewSongs(id, name, artist, album)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}

	if err := s.repository.create(m); err != nil {
		if err.Error() == "ecatch:108" {
			return m, 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't create Songs :", err)
		return m, 3, err
	}
	return m, 29, nil
}

func (s *service) UpdateSongs(id string, name string, artist string, album string) (*Songs, int, error) {
	m := NewSongs(id, name, artist, album)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}
	if err := s.repository.update(m); err != nil {
		logger.Error.Println(s.txID, " - couldn't update Songs :", err)
		return m, 18, err
	}
	return m, 29, nil
}

func (s *service) DeleteSongs(id string) (int, error) {
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

func (s *service) GetSongsByID(id string) (*Songs, int, error) {
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

func (s *service) GetAllSongs() ([]*Songs, error) {
	return s.repository.getAll()
}
