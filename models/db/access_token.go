package db

import (
	"time"

	"github.com/google/uuid"
)

type AccessToken struct {
	ID        int       `db:"id" json:"id"`
	AccountID uuid.UUID `db:"account_id" json:"account_id"`
	Token     *string   `db:"token" json:"token,omitempty"`
	CreatedAt time.Time `db:"created_at" json:"-"`
}

func (c AccessToken) TableName() string {
	return "access_tokens"
}
