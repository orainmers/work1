package legalentities

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/krisch/crm-backend/domain"
	"gorm.io/gorm"
)

// Repository предоставляет методы для работы с LegalEntity и BankAccount в БД.
type Repository struct {
	db *gorm.DB
}

// NewRepository возвращает новый экземпляр репозитория.
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

// Create добавляет новую сущность LegalEntity в БД.
func (r *Repository) Create(ctx context.Context, entity *domain.LegalEntity) error {
	return r.db.WithContext(ctx).Create(entity).Error
}

// GetByUUID возвращает LegalEntity по его UUID.
func (r *Repository) GetByUUID(ctx context.Context, id uuid.UUID) (domain.LegalEntity, error) {
	var entity domain.LegalEntity
	err := r.db.WithContext(ctx).
		Where("uuid = ?", id).
		First(&entity).Error
	return entity, err
}

// GetAll возвращает все LegalEntity, у которых нет метки удалённости (deleted_at).
func (r *Repository) GetAll(ctx context.Context) ([]domain.LegalEntity, error) {
	var entities []domain.LegalEntity
	err := r.db.WithContext(ctx).
		Where("deleted_at IS NULL").
		Find(&entities).Error
	return entities, err
}

// Update обновляет существующую запись LegalEntity.
func (r *Repository) Update(ctx context.Context, entity *domain.LegalEntity) error {
	return r.db.WithContext(ctx).
		Model(&domain.LegalEntity{}).
		Where("uuid = ?", entity.UUID).
		Updates(map[string]interface{}{
			"name":       entity.Name,
			"updated_at": entity.UpdatedAt,
		}).Error
}

// Delete помечает LegalEntity как удалённую (soft delete), устанавливая deleted_at.
func (r *Repository) Delete(ctx context.Context, id uuid.UUID) error {
	now := time.Now()
	result := r.db.WithContext(ctx).
		Model(&domain.LegalEntity{}).
		Where("uuid = ?", id).
		Where("deleted_at IS NULL").
		Update("deleted_at", &now)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

// GetAllBankAccounts возвращает банковские счета.
// Если передан legalEntityUUID == uuid.Nil, возвращает все счета без фильтрации.
func (r *Repository) GetAllBankAccounts(ctx context.Context, legalEntityUUID uuid.UUID) ([]domain.BankAccount, error) {
	var bankAccounts []domain.BankAccount
	query := r.db.WithContext(ctx).Where("deleted_at IS NULL")

	if legalEntityUUID != uuid.Nil {
		query = query.Where("legal_entity_uuid = ?", legalEntityUUID)
	}

	err := query.Find(&bankAccounts).Error
	return bankAccounts, err
}

// GetBankAccountByUUID возвращает один банковский счет по его UUID.
func (r *Repository) GetBankAccountByUUID(ctx context.Context, bankAccountUUID uuid.UUID) (domain.BankAccount, error) {
	var bankAccount domain.BankAccount
	err := r.db.WithContext(ctx).
		Where("uuid = ? AND deleted_at IS NULL", bankAccountUUID).
		First(&bankAccount).Error
	return bankAccount, err
}

// CreateBankAccount добавляет новый банковский аккаунт для юридического лица.
func (r *Repository) CreateBankAccount(ctx context.Context, bankAccount *domain.BankAccount) error {
	return r.db.WithContext(ctx).Create(bankAccount).Error
}

// DeleteBankAccount помечает банковский аккаунт как удалённый (soft delete), устанавливая deleted_at.
func (r *Repository) DeleteBankAccount(ctx context.Context, bankAccountUUID uuid.UUID) error {
	now := time.Now()
	result := r.db.WithContext(ctx).
		Model(&domain.BankAccount{}).
		Where("uuid = ?", bankAccountUUID).
		Where("deleted_at IS NULL").
		Update("deleted_at", &now)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

// UpdateBankAccount обновляет банковский аккаунт.
func (r *Repository) UpdateBankAccount(ctx context.Context, bankAccount *domain.BankAccount) error {
	return r.db.WithContext(ctx).
		Model(&domain.BankAccount{}).
		Where("uuid = ?", bankAccount.UUID).
		Updates(map[string]interface{}{
			"bic":                   bankAccount.BIC,
			"bank_name":             bankAccount.BankName,
			"address":               bankAccount.Address,
			"correspondent_account": bankAccount.CorrespondentAccount,
			"account_number":        bankAccount.AccountNumber,
			"currency":              bankAccount.Currency,
			"comment":               bankAccount.Comment,
			"updated_at":            bankAccount.UpdatedAt,
		}).Error
}

// GetAllBankAccountsByLegalEntityUUID получает все банковские счета для конкретного юридического лица.
func (r *Repository) GetAllBankAccountsByLegalEntityUUID(ctx context.Context, legalEntityUUID uuid.UUID) ([]domain.BankAccount, error) {
	var bankAccounts []domain.BankAccount
	err := r.db.WithContext(ctx).
		Where("legal_entity_uuid = ? AND deleted_at IS NULL", legalEntityUUID).
		Find(&bankAccounts).Error
	return bankAccounts, err
}
