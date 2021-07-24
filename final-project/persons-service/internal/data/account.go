package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"persons-service/internal/biz"
)

type accountRepo struct {
	log      *log.Helper
	accounts map[string]*biz.Account
}

func newAccountRepo(data *Data, logger log.Logger) biz.AccountRepo {
	return &accountRepo{log: log.NewHelper(log.With(logger, "module", "data/person")), accounts: make(map[string]*biz.Account)}
}

func (a *accountRepo) CreateAccount(ctx context.Context, username string) (*biz.Account, error) {
	if _, ok := a.accounts[username]; ok {
		return nil, biz.ErrAccountExited
	}
	account := &biz.Account{
		Username: username,
		ID:       uuid.NewString(),
	}
	a.accounts[username] = account
	return account, nil
}

func (a *accountRepo) GetAccount(ctx context.Context, username string) (*biz.Account, error) {
	if account, ok := a.accounts[username]; ok {
		return account, nil
	}
	return nil, biz.ErrAccountNotFound
}
