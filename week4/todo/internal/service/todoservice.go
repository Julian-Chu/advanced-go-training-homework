package service

import (
	"context"
	"todo/internal/biz"

	pb "todo/api/todo/v1"
)

type TodoService struct {
	pb.UnimplementedTodoServiceServer
	todo *biz.TodoUsecase
}

func NewTodoService(todo *biz.TodoUsecase) *TodoService {
	return &TodoService{todo: todo}
}

func (s *TodoService) CreateTodo(ctx context.Context, req *pb.CreateTodoRequest) (*pb.CreateTodoReply, error) {
	s.todo.AddTodo(req.Todo)
	return &pb.CreateTodoReply{}, nil
}
func (s *TodoService) DeleteTodo(ctx context.Context, req *pb.DeleteTodoRequest) (*pb.DeleteTodoReply, error) {
	return &pb.DeleteTodoReply{}, nil
}
func (s *TodoService) ListTodo(ctx context.Context, req *pb.ListTodoRequest) (*pb.ListTodoReply, error) {
	todos, err := s.todo.ListTodos()
	if err != nil {
		return nil, err
	}
	todoMessages := make([]*pb.Todo, len(todos))
	for i := range todos {
		todoMessages[i] = &pb.Todo{
			Id:      todos[i].ID,
			Content: todos[i].Content,
		}
	}
	return &pb.ListTodoReply{
		Todos: todoMessages,
	}, nil
}
