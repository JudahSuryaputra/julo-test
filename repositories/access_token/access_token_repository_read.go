package access_token

import (
	"fmt"
	"julo-case-study/models/db"

	"github.com/gocraft/dbr"
)

func GetUserAccessToken(sess *dbr.Session, accessToken *string) (*db.AccessToken, error) {
	var userAccessToken db.AccessToken

	query := sess.Select("*").
		From(db.AccessToken{}.TableName()).
		Where("token = ?", accessToken)

	err := query.LoadOne(&userAccessToken)
	if err != nil {
		if err != dbr.ErrNotFound {
			fmt.Println(err)
			return &userAccessToken, err
		}
		return nil, err
	}

	return &userAccessToken, nil
}
