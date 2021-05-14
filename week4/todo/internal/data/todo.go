package data

import (
	"sync/atomic"
	"time"
	"todo/internal/biz"
)

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

func (t *TodoRepo) DeleteTodo(todoID int64) error {
	todos := &t.data.todoStore.todos
	for i := range *todos {
		if (*todos)[i].ID == todoID {
			(*todos)[i], (*todos)[len(*todos)-1] = (*todos)[len(*todos)-1], (*todos)[i]
			*todos = (*todos)[:len(*todos)-1]
			return nil
		}
	}
	return nil
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
