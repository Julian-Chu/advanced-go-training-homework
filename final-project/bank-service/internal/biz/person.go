package biz

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
)

var ErrUserExisted = errors.New("user existed")
var ErrUserNotFound = errors.New("user not found")

type Person struct {
	Username  string
	Email     string
	AccountID string
}

type PersonRepo interface {
	CreatePerson(ctx context.Context, person *Person) error
	GetPerson(ctx context.Context, username string) (*Person, error)
}

type PersonUsercase struct {
	repo PersonRepo
	log  *log.Helper
}

func NewPersonUsercase(repo PersonRepo, logger log.Logger) *PersonUsercase {
	return &PersonUsercase{repo: repo, log: log.NewHelper(log.With(logger, "module", "usecase/person"))}
}

func (uc *PersonUsercase) Create(ctx context.Context, p *Person) error {
	return uc.repo.CreatePerson(ctx, p)
}

func (uc *PersonUsercase) Get(ctx context.Context, username string) (*Person, error) {
	return uc.repo.GetPerson(ctx, username)
}

func (uc PersonUsercase) IsUserNotFound(err error) bool {
	if err == ErrUserNotFound {
		return true
	}
	return false
}
