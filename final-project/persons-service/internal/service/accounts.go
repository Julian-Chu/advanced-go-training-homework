package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"persons-service/internal/biz"

	pb "persons-service/api/accountsservice"
)

type AccountsService struct {
	pb.UnimplementedAccountServer
	uc  *biz.AccountUsercase
	log *log.Helper
}

func NewAccountsService(uc *biz.AccountUsercase, logger log.Logger) *AccountsService {
	return &AccountsService{uc: uc, log: log.NewHelper(log.With(logger, "module", "service/account"))}
}

func (s *AccountsService) CreateAccount(ctx context.Context, req *pb.CreateAccountRequest) (*pb.CreateAccountReply, error) {
	account, err := s.uc.CreateAccount(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	return &pb.CreateAccountReply{
		AccountID: account.ID,
	}, nil
}
func (s *AccountsService) UpdateAccount(ctx context.Context, req *pb.UpdateAccountRequest) (*pb.UpdateAccountReply, error) {
	return &pb.UpdateAccountReply{}, nil
}
func (s *AccountsService) DeleteAccount(ctx context.Context, req *pb.DeleteAccountRequest) (*pb.DeleteAccountReply, error) {
	return &pb.DeleteAccountReply{}, nil
}
func (s *AccountsService) GetAccount(ctx context.Context, req *pb.GetAccountRequest) (*pb.GetAccountReply, error) {
	return &pb.GetAccountReply{}, nil
}
func (s *AccountsService) ListAccount(ctx context.Context, req *pb.ListAccountRequest) (*pb.ListAccountReply, error) {
	return &pb.ListAccountReply{}, nil
}
