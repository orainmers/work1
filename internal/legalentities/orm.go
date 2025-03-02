package legalentities

import (
	"time"

	"github.com/google/uuid"
)

// LegalEntity представляет ORM-модель для юридических лиц
// с полями: uuid, name, created_at, updated_at, deleted_at.
type LegalEntity struct {
	UUID      uuid.UUID  `json:"uuid" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name      string     `json:"name" gorm:"not null;type:varchar(255)"`
	CreatedAt time.Time  `json:"created_at" gorm:"<-:create;autoCreateTime"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt *time.Time `json:"deleted_at"`
}

// TableName указывает, какую таблицу в БД будет использовать ORM.
func (LegalEntity) TableName() string {
	return "legal_entities"
}
