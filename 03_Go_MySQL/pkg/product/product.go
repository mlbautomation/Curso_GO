package product

import "time"

// Model of product
type Model struct {
	ID           uint
	Name         string
	Observations string
	Price        int
	CreateAt     time.Time
	UpadateAt    time.Time
}

type Models []Model

/*
type Storage interface {
	Migrate() error
	// aqui ver porque se trabajan con punteros
	// Create(Model) error
	// Update(Model) error
	// GetAll() (Models, error)
	// GetByID(uint) (Model, error)
	// Delete(uint) error
}

type Service struct {
	storage Storage
}

func NewService(s Storage) *Service {
	return &Service{s}
}

func (s *Service) Migrate() error {
	return s.storage.Migrate()
}
*/
