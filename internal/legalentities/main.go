package legalentities

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/krisch/crm-backend/domain"
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
	entity := &domain.LegalEntity{
		UUID:      uuid.New(),
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := s.repo.Create(ctx, entity)
	return entity.UUID, err
}

// GetAllLegalEntities возвращает список всех LegalEntity (не удалённых).
func (s *Service) GetAllLegalEntities(ctx context.Context) ([]domain.LegalEntity, error) {
	return s.repo.GetAll(ctx)
}

// GetLegalEntity возвращает конкретную LegalEntity по UUID.
func (s *Service) GetLegalEntity(ctx context.Context, id uuid.UUID) (domain.LegalEntity, error) {
	return s.repo.GetByUUID(ctx, id)
}

// UpdateLegalEntity обновляет поля существующей записи LegalEntity.
func (s *Service) UpdateLegalEntity(ctx context.Context, id uuid.UUID, newName string) error {
	entity, err := s.repo.GetByUUID(ctx, id)
	if err != nil {
		return err
	}

	entity.Name = newName
	entity.UpdatedAt = time.Now()

	return s.repo.Update(ctx, &entity)
}

// DeleteLegalEntity «мягко» удаляет LegalEntity (soft delete).
func (s *Service) DeleteLegalEntity(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}

// GetAllBankAccounts возвращает список всех банковских счетов.
// Если передан legalEntityUUID == uuid.Nil, возвращает все счета без фильтрации.
func (s *Service) GetAllBankAccounts(ctx context.Context, legalEntityUUID uuid.UUID) ([]domain.BankAccount, error) {
	if legalEntityUUID == uuid.Nil {
		return s.repo.GetAllBankAccounts(ctx, uuid.Nil) // Получаем все счета
	}
	return s.repo.GetAllBankAccounts(ctx, legalEntityUUID) // Получаем счета только для конкретного юр. лица
}

// GetBankAccountByUUID возвращает один банковский счет по его UUID.
func (s *Service) GetBankAccountByUUID(ctx context.Context, bankAccountUUID uuid.UUID) (domain.BankAccount, error) {
	return s.repo.GetBankAccountByUUID(ctx, bankAccountUUID)
}

// CreateBankAccount создаёт новый банковский аккаунт для юридического лица.
func (s *Service) CreateBankAccount(ctx context.Context, bankAccount *domain.BankAccount) (uuid.UUID, error) {
	bankAccount.UUID = uuid.New()
	bankAccount.CreatedAt = time.Now()
	bankAccount.UpdatedAt = time.Now()

	err := s.repo.CreateBankAccount(ctx, bankAccount)
	return bankAccount.UUID, err
}

// DeleteBankAccount «мягко» удаляет банковский аккаунт (soft delete).
func (s *Service) DeleteBankAccount(ctx context.Context, bankAccountUUID uuid.UUID) error {
	return s.repo.DeleteBankAccount(ctx, bankAccountUUID)
}

// UpdateBankAccount обновляет банковский аккаунт.
func (s *Service) UpdateBankAccount(ctx context.Context, bankAccount *domain.BankAccount) error {
	bankAccount.UpdatedAt = time.Now()
	return s.repo.UpdateBankAccount(ctx, bankAccount)
}

// GetAllBankAccountsByLegalEntityUUID возвращает банковские счета для конкретного юридического лица.
func (s *Service) GetAllBankAccountsByLegalEntityUUID(ctx context.Context, legalEntityUUID uuid.UUID) ([]domain.BankAccount, error) {
	return s.repo.GetAllBankAccountsByLegalEntityUUID(ctx, legalEntityUUID)
}
