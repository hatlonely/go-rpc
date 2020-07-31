package service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"

	account "github.com/hatlonely/go-rpc/account/api/gen/go/api"
)

type AccountService struct{}

func (s *AccountService) SignIn(ctx context.Context, req *account.SignInReq) (*account.SignInRes, error) {
	return &account.SignInRes{
		Token: "hello world",
	}, nil
}

func (s *AccountService) SignUp(ctx context.Context, req *account.SignUpReq) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}

func (s *AccountService) SignOut(ctx context.Context, req *account.SignOutReq) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}
