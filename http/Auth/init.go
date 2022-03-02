package Auth

import (
	"encoding/json"
	"errors"
	"julo-case-study/http/enum"
	"julo-case-study/http/formatter"
	"julo-case-study/http/utils"
	"julo-case-study/models/requests"
	"julo-case-study/repositories/access_token"
	"julo-case-study/repositories/account"
	"net/http"

	"github.com/gocraft/dbr"
)

type InitAccount struct {
	DBConn *dbr.Connection
}

func (c InitAccount) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var request requests.InitAccountRequest
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&request)
	if err != nil {
		formatter.FAIL(w, http.StatusBadRequest, errors.New(enum.ErrMissingData))
		return
	}

	sess := c.DBConn.NewSession(nil)

	userID, err := account.CreateAccount(sess, request)
	if err != nil {
		formatter.FAIL(w, http.StatusBadRequest, err)
		return
	}

	token, err := utils.EncodeAuthToken(*userID)
	if err != nil {
		formatter.FAIL(w, http.StatusInternalServerError, err)
		return
	}

	insertAccessTokenRequest := requests.InsertAccessTokenRequest{
		AccountID: *userID,
		Token:     &token,
	}
	err = access_token.InsertAccessToken(sess, insertAccessTokenRequest)
	if err != nil {
		formatter.FAIL(w, http.StatusInternalServerError, err)
		return
	}

	formatter.SUCCESS(w, http.StatusCreated, struct {
		Token string `json:"token"`
	}{
		Token: token,
	})
	return
}
