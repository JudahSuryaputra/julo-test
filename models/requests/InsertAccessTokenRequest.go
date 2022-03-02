package requests

import "github.com/google/uuid"

type InsertAccessTokenRequest struct {
	AccountID uuid.UUID `json:"account_id"`
	Token     *string   `json:"token"`
}
