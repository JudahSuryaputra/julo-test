package utils

import (
	"errors"
	"julo-case-study/http/enum"
	"julo-case-study/models/requests"
	"julo-case-study/repositories/access_token"
	"net/http"
	"strings"

	"github.com/gocraft/dbr"
	"github.com/google/uuid"
)

func CheckAuthorization(r *http.Request, sess *dbr.Session) error {
	authorization := r.Header.Get("Authorization")
	accessToken := strings.Split(authorization, "Token ")
	if len(accessToken) != 2 {
		return errors.New(enum.UnauthorizedUser)
	}
	currentToken, err := access_token.GetUserAccessToken(sess, &accessToken[1])
	if currentToken == nil || err != nil {
		return errors.New(enum.UnauthorizedUser)
	}

	return nil
}

func InsertAccessToken(sess *dbr.Session, accountID uuid.UUID, accessToken string) error {
	insertAccessTokenRequest := requests.InsertAccessTokenRequest{
		AccountID: accountID,
		Token:     &accessToken,
	}

	err := access_token.InsertAccessToken(sess, insertAccessTokenRequest)
	if err != nil {
		return err
	}

	return nil
}
