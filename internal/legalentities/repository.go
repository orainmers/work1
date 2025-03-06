package legalentities

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/krisch/crm-backend/domain" // Импортируй пакет с доменной моделью
	"gorm.io/gorm"
)

// Repository предоставляет методы для работы с сущностями LegalEntity в БД.
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
		Model(&domain.LegalEntity{}).   // Указываем модель
		Where("uuid = ?", entity.UUID). // Условие для поиска по UUID
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
