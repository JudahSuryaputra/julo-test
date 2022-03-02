package db

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID        uuid.UUID  `db:"id" json:"id"`
	OwnedBy   uuid.UUID  `db:"owned_by" json:"owned_by"`
	Status    string     `db:"status" json:"status"`
	EnabledAt *time.Time `db:"enabled_at" json:"enabled_at,omitempty"`
	Balance   int        `db:"balance" json:"balance"`
}

func (c Account) TableName() string {
	return "accounts"
}
