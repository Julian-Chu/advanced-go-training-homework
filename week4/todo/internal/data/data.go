package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"sync/atomic"
	"time"
	"todo/internal/biz"
	"todo/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewTodoRepo)

// Data .
type Data struct {
	todoStore *InMemoryTodoStore
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		logger.Log("msg", "closing the data resources")
	}
	store := NewInMemoryTodoStore()
	return &Data{todoStore: store}, cleanup, nil
}

type InMemoryTodoStore struct {
	todos  []biz.Todo
	todoID int64
}

func NewInMemoryTodoStore() *InMemoryTodoStore {
	return &InMemoryTodoStore{todos: make([]biz.Todo, 0, 1024), todoID: 0}
}

type TodoRepo struct {
	data *Data
}

func (t *TodoRepo) AddTodo(content string) error {
	todoID := atomic.AddInt64(&t.data.todoStore.todoID, 1)
	newTodo := biz.Todo{
		ID:        todoID,
		Content:   content,
		Timestamp: time.Now(),
	}
	t.data.todoStore.todos = append(t.data.todoStore.todos, newTodo)
	return nil
}

func (t TodoRepo) ListTodos() ([]biz.Todo, error) {
	return t.data.todoStore.todos, nil
}

func NewTodoRepo(data *Data) biz.TodoRepo {
	return &TodoRepo{data: data}
}
