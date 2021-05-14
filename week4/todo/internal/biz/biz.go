package biz

import (
	"github.com/google/wire"
	"time"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewTodoUsecase)

type Todo struct {
	Id        int64
	Content   string
	Timestamp time.Time
}

type TodoUsecase struct {
	repo TodoRepo
}

func NewTodoUsecase(repo TodoRepo) *TodoUsecase {
	return &TodoUsecase{repo: repo}
}

func (t TodoUsecase) AddTodo(content string) error {
	if err := t.repo.AddTodo(content); err != nil {
		return err
	}
	return nil
}

func (t TodoUsecase) ListTodos() ([]Todo, error) {
	todos, err := t.repo.ListTodos()
	if err != nil {
		return nil, err
	}
	return todos, nil
}

type TodoRepo interface {
	AddTodo(content string) error
	ListTodos() ([]Todo, error)
}
