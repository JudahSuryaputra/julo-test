package access_token

import (
	"julo-case-study/models/db"
	"julo-case-study/models/requests"

	"github.com/gocraft/dbr"
)

func InsertAccessToken(sess *dbr.Session, r requests.InsertAccessTokenRequest) error {
	accessToken := db.AccessToken{
		AccountID: r.AccountID,
		Token:     r.Token,
	}

	columns := []string{
		"account_id",
		"token",
	}

	_, err := sess.InsertInto(db.AccessToken{}.TableName()).
		Columns(columns...).
		Record(accessToken).
		Exec()
	if err != nil {
		return err
	}

	return nil
}
