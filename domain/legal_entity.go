package domain

import (
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
