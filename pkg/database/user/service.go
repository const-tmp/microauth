package user

import (
	"github.com/gofrs/uuid"
	"github.com/h1ght1me/auth-micro/config"
	"gorm.io/gorm"
)

type Service struct {
	DB     *gorm.DB
	Config *config.Config
}

func NewService(db *gorm.DB, cfg *config.Config) *Service {
	return &Service{
		DB:     db,
		Config: cfg,
	}
}

func (s *Service) User(id uuid.UUID) (*User, error) {
	var t User
	tx := s.DB.First(&t, id)
	return &t, tx.Error
}

func (s *Service) GetByName(name string) (*User, error) {
	var t User
	tx := s.DB.Where("name = ?", name).First(&t)
	return &t, tx.Error
}

func (s *Service) CreateUser(t *User) error {
	tx := s.DB.Create(t)
	return tx.Error
}

func (s *Service) Users() ([]*User, error) {
	var tasks []*User
	tx := s.DB.Find(&tasks)
	return tasks, tx.Error
}

func (s *Service) DeleteUser(id uuid.UUID) error {
	task := &User{ID: id}
	s.DB.Delete(&task)
	return nil
}
