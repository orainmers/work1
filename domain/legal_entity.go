package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type LegalEntity struct {
	UUID      uuid.UUID `validate:"uuid"`
	Name      string    `validate:"lte=100,gte=1"  ru:"название"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func (l *LegalEntity) ChangeName(name string) error {
	if len(name) < 1 || len(name) > 100 {
		return errors.New("название от 1 до 100 символов")
	}

	l.Name = name

	return nil
}

func NewLegalEntityUUID(uid uuid.UUID) *LegalEntity {
	legalentity := &LegalEntity{
		UUID: uid,
	}

	return legalentity
}
