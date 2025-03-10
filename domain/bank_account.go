package domain

import (
	"time"

	"github.com/google/uuid"
)

type BankAccount struct {
	UUID                 uuid.UUID  `validate:"uuid"`
	BIC                  string     `validate:"lte=255,gte=1" ru:"БИК"`
	BankName             string     `validate:"lte=255,gte=1" ru:"Название банка"`
	Address              string     `validate:"lte=255" ru:"Адрес банка"`
	CorrespondentAccount string     `validate:"lte=255" ru:"Корреспондентский счет"`
	AccountNumber        string     `validate:"lte=255,gte=1" ru:"Номер счета"`
	Currency             string     `validate:"lte=10" ru:"Валюта"`
	Comment              string     `validate:"lte=500" ru:"Комментарий"`
	LegalEntityUUID      uuid.UUID  `validate:"uuid" ru:"UUID юридического лица"`
	CreatedAt            time.Time  `validate:"required"`
	UpdatedAt            time.Time  `validate:"required"`
	DeletedAt            *time.Time `validate:"omitempty"`
}
