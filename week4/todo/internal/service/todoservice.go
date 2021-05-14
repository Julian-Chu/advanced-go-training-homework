package service

import (
	"context"

	pb "todo/api/todo/v1"
)

type TodoServiceService struct {
	pb.UnimplementedTodoServiceServer
}

func NewTodoServiceService() pb.TodoServiceServer {
	return &TodoServiceService{}
}

func (s *TodoServiceService) CreateTodo(ctx context.Context, req *pb.CreateTodoRequest) (*pb.CreateTodoReply, error) {
	return &pb.CreateTodoReply{}, nil
}
func (s *TodoServiceService) DeleteTodo(ctx context.Context, req *pb.DeleteTodoRequest) (*pb.DeleteTodoReply, error) {
	return &pb.DeleteTodoReply{}, nil
}
func (s *TodoServiceService) ListTodo(ctx context.Context, req *pb.ListTodoRequest) (*pb.ListTodoReply, error) {
	return &pb.ListTodoReply{}, nil
}
