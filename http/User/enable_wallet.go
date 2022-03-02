package User

import (
	"errors"
	"julo-case-study/http/enum"
	"julo-case-study/http/formatter"
	"julo-case-study/http/utils"
	"julo-case-study/models/responses"
	"julo-case-study/repositories/account"
	"net/http"
	"time"

	"github.com/gocraft/dbr"
)

type EnableWallet struct {
	DBConn *dbr.Connection
}

func (c EnableWallet) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

	if currentAccount.Status == enum.WalletEnabled {
		formatter.FAIL(w, http.StatusBadRequest, errors.New(enum.ErrAlreadyEnabled))
		return
	}

	now := time.Now()
	err = account.EnableWalletByID(sess, accountID, now)
	if err != nil {
		formatter.FAIL(w, http.StatusInternalServerError, err)
		return
	}

	currentAccount.Status = enum.WalletEnabled
	currentAccount.EnabledAt = &now

	data := responses.WalletResponse{
		Wallet: currentAccount,
	}

	formatter.SUCCESS(w, http.StatusCreated, data)
	return
}
