package service

import (
	"context"
	"todo/internal/biz"

	pb "todo/api/todo/v1"
)

type TodoServiceService struct {
	pb.UnimplementedTodoServiceServer
	todo *biz.TodoUsecase
}

func NewTodoServiceService(todo *biz.TodoUsecase) pb.TodoServiceServer {
	return &TodoServiceService{todo: todo}
}

func (s *TodoServiceService) CreateTodo(ctx context.Context, req *pb.CreateTodoRequest) (*pb.CreateTodoReply, error) {
	s.todo.AddTodo(req.Todo)
	return &pb.CreateTodoReply{}, nil
}
func (s *TodoServiceService) DeleteTodo(ctx context.Context, req *pb.DeleteTodoRequest) (*pb.DeleteTodoReply, error) {
	return &pb.DeleteTodoReply{}, nil
}
func (s *TodoServiceService) ListTodo(ctx context.Context, req *pb.ListTodoRequest) (*pb.ListTodoReply, error) {
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
