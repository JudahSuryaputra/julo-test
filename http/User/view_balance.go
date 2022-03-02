package User

import (
	"errors"
	"julo-case-study/http/enum"
	"julo-case-study/http/formatter"
	"julo-case-study/http/utils"
	"julo-case-study/models/responses"
	"julo-case-study/repositories/account"
	"net/http"

	"github.com/gocraft/dbr"
)

type GetBalance struct {
	DBConn *dbr.Connection
}

func (c GetBalance) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	accountID := r.Context().Value("ID").(string)

	sess := c.DBConn.NewSession(nil)

	err := utils.CheckAuthorization(r, sess)
	if err != nil {
		formatter.FAIL(w, http.StatusForbidden, err)
		return
	}

	currentAccount, err := account.GetAccountByID(sess, accountID)
	if err != nil {
		formatter.FAIL(w, http.StatusForbidden, err)
		return
	}

	if currentAccount.Status == enum.WalletDisabled {
		formatter.FAIL(w, http.StatusNotFound, errors.New(enum.WalletDisabled))
		return
	}

	data := responses.WalletResponse{
		Wallet: currentAccount,
	}

	formatter.SUCCESS(w, http.StatusCreated, data)
	return
}
