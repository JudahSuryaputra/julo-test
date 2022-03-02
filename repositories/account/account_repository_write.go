package account

import (
	"julo-case-study/http/enum"
	"julo-case-study/models/db"
	"julo-case-study/models/requests"
	"time"

	"github.com/gocraft/dbr"
	"github.com/google/uuid"
)

type CreateAccountRequest struct {
	OwnedBy uuid.UUID `db:"owned_by" json:"owned_by"`
}

func CreateAccount(sess *dbr.Session, request requests.InitAccountRequest) (*uuid.UUID, error) {
	var userID *uuid.UUID
	account := CreateAccountRequest{
		OwnedBy: request.CustomerXID,
	}

	columns := []string{"owned_by"}

	err := sess.InsertInto(db.Account{}.TableName()).
		Columns(columns...).
		Record(account).
		Returning("id").
		Load(&userID)
	if err != nil {
		return nil, err
	}

	return userID, nil
}

func EnableWalletByID(sess *dbr.Session, userID string, now time.Time) error {
	data := make(map[string]interface{})

	data["status"] = enum.WalletEnabled
	data["enabled_at"] = now

	_, err := sess.Update(db.Account{}.TableName()).
		Where("id = ?", userID).
		SetMap(data).
		Exec()
	if err != nil {
		return err
	}

	return nil
}
