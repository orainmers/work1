package legalentities

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/krisch/crm-backend/domain" // Импортируем доменную модель
)

// Service содержит бизнес-логику для работы с LegalEntity и BankAccount.
type Service struct {
	repo *Repository
}

// NewService возвращает новый экземпляр Service.
func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

// CreateLegalEntity создаёт новую запись LegalEntity.
func (s *Service) CreateLegalEntity(ctx context.Context, name string) (uuid.UUID, error) {
	// Используем доменную модель
	entity := &domain.LegalEntity{
		UUID:      uuid.New(), // Генерируем новый UUID
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Передаем доменную модель в репозиторий
	err := s.repo.Create(ctx, entity)
	return entity.UUID, err
}

// GetAllLegalEntities возвращает список всех LegalEntity (не удалённых).
func (s *Service) GetAllLegalEntities(ctx context.Context) ([]domain.LegalEntity, error) {
	// Получаем все сущности через репозиторий
	return s.repo.GetAll(ctx)
}

// GetLegalEntity возвращает конкретную LegalEntity по UUID.
func (s *Service) GetLegalEntity(ctx context.Context, id uuid.UUID) (domain.LegalEntity, error) {
	// Получаем сущность по UUID через репозиторий
	return s.repo.GetByUUID(ctx, id)
}

// UpdateLegalEntity обновляет поля существующей записи LegalEntity.
func (s *Service) UpdateLegalEntity(ctx context.Context, id uuid.UUID, newName string) error {
	// Получаем существующую сущность через репозиторий
	entity, err := s.repo.GetByUUID(ctx, id)
	if err != nil {
		return err
	}

	// Обновляем поля
	entity.Name = newName
	entity.UpdatedAt = time.Now()

	// Сохраняем обновленную сущность в репозитории
	err = s.repo.Update(ctx, &entity)
	return err
}

// DeleteLegalEntity «мягко» удаляет LegalEntity (soft delete).
func (s *Service) DeleteLegalEntity(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}

// GetAllBankAccounts возвращает все банковские аккаунты, связанные с юридическим лицом.
func (s *Service) GetAllBankAccounts(ctx context.Context, legalEntityUUID uuid.UUID) ([]domain.BankAccount, error) {
	// Получаем все банковские аккаунты для юридического лица
	return s.repo.GetAllBankAccounts(ctx, legalEntityUUID)
}

// CreateBankAccount создаёт новый банковский аккаунт для юридического лица.
func (s *Service) CreateBankAccount(ctx context.Context, bankAccount *domain.BankAccount) (uuid.UUID, error) {
	// Генерируем новый UUID для банковского аккаунта
	bankAccount.UUID = uuid.New()
	bankAccount.CreatedAt = time.Now()
	bankAccount.UpdatedAt = time.Now()

	// Передаем доменную модель банковского аккаунта в репозиторий
	err := s.repo.CreateBankAccount(ctx, bankAccount)
	return bankAccount.UUID, err
}

// DeleteBankAccount «мягко» удаляет банковский аккаунт (soft delete).
func (s *Service) DeleteBankAccount(ctx context.Context, bankAccountUUID uuid.UUID) error {
	return s.repo.DeleteBankAccount(ctx, bankAccountUUID)
}

// UpdateBankAccount обновляет банковский аккаунт.
func (s *Service) UpdateBankAccount(ctx context.Context, bankAccount *domain.BankAccount) error {
	// Обновляем время изменения
	bankAccount.UpdatedAt = time.Now()

	// Передаем обновленную модель банковского аккаунта в репозиторий
	return s.repo.UpdateBankAccount(ctx, bankAccount)
}
