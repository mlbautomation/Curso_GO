package pkg

type Storage interface {
	Migrate() error
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
