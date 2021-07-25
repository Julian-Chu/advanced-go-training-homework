package data

import (
	accounts "bankservice/api/accountsservice"
	"bankservice/internal/biz"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type personRepo struct {
	//data *Data
	log           *log.Helper
	persons       map[string]*biz.Person
	accountClient accounts.AccountClient
}

func (p *personRepo) GetPerson(ctx context.Context, username string) (*biz.Person, error) {
	if person, ok := p.persons[username]; ok {
		return person, nil
	}
	return nil, biz.ErrUserNotFound
}

func NewPersonRepo(data *Data, logger log.Logger, client accounts.AccountClient) biz.PersonRepo {
	return &personRepo{log: log.NewHelper(log.With(logger, "module", "data/person")), persons: make(map[string]*biz.Person), accountClient: client}
}

func (p *personRepo) CreatePerson(ctx context.Context, person *biz.Person) (accountID string, err error) {
	if _, ok := p.persons[person.Username]; ok {
		return "", biz.ErrUserExisted
	}
	accountReply, err := p.accountClient.CreateAccount(ctx, &accounts.CreateAccountRequest{Username: person.Username})
	if err != nil {
		return "", err
	}
	p.persons[person.Username] = person
	return accountReply.AccountID, nil
}
