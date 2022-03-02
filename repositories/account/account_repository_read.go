package account

import (
	"julo-case-study/models/db"

	"github.com/gocraft/dbr"
)

func GetAccountByID(sess *dbr.Session, userID string) (*db.Account, error) {
	var account *db.Account

	query := sess.Select("*").
		From(db.Account{}.TableName()).
		Where("id::text = ?", userID)

	err := query.LoadOne(&account)
	if err != nil {
		return nil, err
	}

	return account, nil
}
