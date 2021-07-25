package data

import (
	accounts "bankservice/api/accountsservice"
	"bankservice/internal/biz"
	"context"
	"github.com/Shopify/sarama"
	"github.com/go-kratos/kratos/v2/log"
)

type personRepo struct {
	//data *Data
	log           *log.Helper
	persons       map[string]*biz.Person
	accountClient accounts.AccountClient
	producer      sarama.AsyncProducer
}

func (p *personRepo) GetPerson(ctx context.Context, username string) (*biz.Person, error) {
	if person, ok := p.persons[username]; ok {
		reply, err := p.accountClient.GetAccount(ctx, &accounts.GetAccountRequest{Username: username})
		if err != nil {
			return nil, err
		}
		person.AccountID = reply.GetAccountID()
		return person, nil
	}
	return nil, biz.ErrUserNotFound
}

func NewPersonRepo(data *Data, logger log.Logger, client accounts.AccountClient, producer sarama.AsyncProducer) biz.PersonRepo {
	return &personRepo{log: log.NewHelper(log.With(logger, "module", "data/person")), persons: make(map[string]*biz.Person), accountClient: client, producer: producer}
}

func (p *personRepo) CreatePerson(_ context.Context, person *biz.Person) (err error) {
	if _, ok := p.persons[person.Username]; ok {
		return biz.ErrUserExisted
	}
	//accountReply, err := p.accountClient.CreateAccount(ctx, &accounts.CreateAccountRequest{Username: person.Username})
	p.producer.Input() <- &sarama.ProducerMessage{
		Topic: "applications",
		Value: sarama.StringEncoder(person.Username),
	}
	if err != nil {
		return err
	}
	p.persons[person.Username] = person
	return nil
}

func NewKafkaProducer() sarama.AsyncProducer {
	c := sarama.NewConfig()
	p, err := sarama.NewAsyncProducer([]string{"localhost:9092"}, c)
	if err != nil {
		panic(err)
	}
	return p
}
