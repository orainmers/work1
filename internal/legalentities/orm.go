package legalentities

import (
	"time"

	"github.com/google/uuid"
	"github.com/krisch/crm-backend/domain" // Импортируем доменную модель
)

// LegalEntity представляет ORM-модель для юридических лиц.
type LegalEntity struct {
	UUID      uuid.UUID  `json:"uuid" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name      string     `json:"name" gorm:"not null;type:varchar(255)"`
	CreatedAt time.Time  `json:"created_at" gorm:"<-:create;autoCreateTime"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt *time.Time `json:"deleted_at"`
}

// BankAccount представляет ORM-модель для банковских аккаунтов.
type BankAccount struct {
	UUID                 uuid.UUID  `json:"uuid" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	BIC                  string     `json:"bic" gorm:"not null;type:varchar(255)"`
	BankName             string     `json:"bank_name" gorm:"not null;type:varchar(255)"`
	Address              string     `json:"address" gorm:"type:varchar(255)"`
	CorrespondentAccount string     `json:"correspondent_account" gorm:"type:varchar(255)"`
	AccountNumber        string     `json:"account_number" gorm:"not null;type:varchar(255)"`
	Currency             string     `json:"currency" gorm:"type:varchar(10)"`
	Comment              string     `json:"comment" gorm:"type:text"`
	LegalEntityUUID      uuid.UUID  `json:"legal_entity_uuid" gorm:"type:uuid;not null"`
	CreatedAt            time.Time  `json:"created_at" gorm:"<-:create;autoCreateTime"`
	UpdatedAt            time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt            *time.Time `json:"deleted_at"`
}

// TableName указывает, какую таблицу в БД будет использовать ORM.
func (LegalEntity) TableName() string {
	return "legal_entities"
}

// TableName указывает, какую таблицу в БД будет использовать ORM для банковских аккаунтов.
func (BankAccount) TableName() string {
	return "bank_accounts"
}

// ToDomain конвертирует ORM-модель BankAccount в доменную модель domain.BankAccount.
func (ba *BankAccount) ToDomain() *domain.BankAccount {
	return &domain.BankAccount{
		UUID:                 ba.UUID,
		BIC:                  ba.BIC,
		BankName:             ba.BankName,
		Address:              ba.Address,
		CorrespondentAccount: ba.CorrespondentAccount,
		AccountNumber:        ba.AccountNumber,
		Currency:             ba.Currency,
		Comment:              ba.Comment,
		LegalEntityUUID:      ba.LegalEntityUUID,
		CreatedAt:            ba.CreatedAt,
		UpdatedAt:            ba.UpdatedAt,
		DeletedAt:            ba.DeletedAt,
	}
}
