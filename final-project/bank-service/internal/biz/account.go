package biz

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
)

var ErrAccountExited = errors.New("account existed")
var ErrAccountNotFound = errors.New("account  not found")

type Account struct {
	Username string
	ID       string
}

type AccountRepo interface {
	CreateAccount(ctx context.Context, username string) (*Account, error)
	GetAccount(ctx context.Context, username string) (*Account, error)
}

type AccountUsercase struct {
	repo AccountRepo
	log  *log.Helper
}

func NewAccountUsercase(repo AccountRepo, logger log.Logger) *AccountUsercase {
	return &AccountUsercase{repo: repo, log: log.NewHelper(log.With(logger, "module", "usecase/account"))}
}

func (uc AccountUsercase) CreateAccount(ctx context.Context, username string) (*Account, error) {
	return uc.repo.CreateAccount(ctx, username)
}
