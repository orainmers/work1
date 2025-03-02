package legalentities

import (
	"context"

	"github.com/google/uuid"
)

// Service содержит бизнес-логику для работы с LegalEntity.
type Service struct {
	repo *Repository
}

// NewService возвращает новый экземпляр Service.
func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

// CreateLegalEntity создаёт новую запись LegalEntity.
func (s *Service) CreateLegalEntity(ctx context.Context, name string) (uuid.UUID, error) {
	entity := &LegalEntity{
		Name: name,
	}
	err := s.repo.Create(ctx, entity)
	return entity.UUID, err
}

// GetAllLegalEntities возвращает список всех LegalEntity (не удалённых).
func (s *Service) GetAllLegalEntities(ctx context.Context) ([]LegalEntity, error) {
	return s.repo.GetAll(ctx)
}

// GetLegalEntity возвращает конкретную LegalEntity по UUID.
func (s *Service) GetLegalEntity(ctx context.Context, id uuid.UUID) (LegalEntity, error) {
	return s.repo.GetByUUID(ctx, id)
}

// UpdateLegalEntity обновляет поля существующей записи LegalEntity.
func (s *Service) UpdateLegalEntity(ctx context.Context, id uuid.UUID, newName string) error {
	// Сначала получаем существующую запись
	entity, err := s.repo.GetByUUID(ctx, id)
	if err != nil {
		return err
	}

	// Меняем поля
	entity.Name = newName

	// Сохраняем изменения
	return s.repo.Update(ctx, &entity)
}

// DeleteLegalEntity «мягко» удаляет LegalEntity (soft delete).
func (s *Service) DeleteLegalEntity(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}
