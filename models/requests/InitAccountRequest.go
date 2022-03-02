package requests

import "github.com/google/uuid"

type InitAccountRequest struct {
	CustomerXID uuid.UUID `json:"customer_xid"`
}
