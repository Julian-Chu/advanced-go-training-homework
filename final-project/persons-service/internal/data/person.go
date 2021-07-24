package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"persons-service/internal/biz"
)

type personRepo struct {
	//data *Data
	log     *log.Helper
	persons map[string]*biz.Person
}

func (p *personRepo) GetPerson(ctx context.Context, username string) (*biz.Person, error) {
	if person, ok := p.persons[username]; ok {
		return person, nil
	}
	return nil, biz.ErrUserNotFound
}

func NewPersonRepo(data *Data, logger log.Logger) biz.PersonRepo {
	return &personRepo{log: log.NewHelper(log.With(logger, "module", "data/person")), persons: make(map[string]*biz.Person)}
}

func (p *personRepo) CreatePerson(ctx context.Context, person *biz.Person) error {
	if _, ok := p.persons[person.Username]; ok {
		return biz.ErrUserExisted
	}
	p.persons[person.Username] = person
	return nil
}
