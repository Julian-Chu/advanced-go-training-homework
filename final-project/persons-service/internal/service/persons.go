package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"persons-service/api/persons-service"
	"persons-service/internal/biz"
)

type PersonsService struct {
	persons.UnimplementedPersonServer

	uc  *biz.PersonUsercase
	log *log.Helper
}

func (p *PersonsService) CreatePerson(ctx context.Context, request *persons.CreatePersonRequest) (*persons.CreatePersonReply, error) {
	person := &biz.Person{
		Username: request.Username,
		Email:    request.Email,
	}
	err := p.uc.Create(ctx, person)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (p *PersonsService) UpdatePerson(ctx context.Context, request *persons.UpdatePersonRequest) (*persons.UpdatePersonReply, error) {
	panic("implement me")
}

func (p *PersonsService) DeletePerson(ctx context.Context, request *persons.DeletePersonRequest) (*persons.DeletePersonReply, error) {
	panic("implement me")
}

func (p *PersonsService) GetPerson(ctx context.Context, request *persons.GetPersonRequest) (*persons.GetPersonReply, error) {
	person, err := p.uc.Get(ctx, request.GetEmail())
	if err != nil {
		if p.uc.IsUserNotFound(err) {
			return nil, errors.New(400, err.Error(), err.Error())
		}
		return nil, err
	}
	rep := &persons.GetPersonReply{
		Username: person.Username,
		Email:    person.Email,
		Account:  nil,
	}
	return rep, nil
}

func NewPersonsService(us *biz.PersonUsercase, logger log.Logger) *PersonsService {
	return &PersonsService{uc: us, log: log.NewHelper(log.With(logger, "module", "service/person"))}
}
