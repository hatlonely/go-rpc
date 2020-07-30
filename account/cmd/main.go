package main

import (
	"context"
	"net/http"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"

	account "github.com/hatlonely/go-rpc/account/api/gen/go/api"
)

type AccountService struct{}

func (s *AccountService) SignIn(ctx context.Context, req *account.SignInReq) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}

func main() {
	mux := runtime.NewServeMux()
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	if err := account.RegisterSignInServiceHandlerServer(ctx, mux, &AccountService{}); err != nil {
		panic(err)
	}
	if err := http.ListenAndServe(":80", mux); err != nil {
		panic(err)
	}
}
