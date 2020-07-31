package main

import (
	"context"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"

	account "github.com/hatlonely/go-rpc/account/api/gen/go/api"
	"github.com/hatlonely/go-rpc/account/internal/service"
)

func main() {
	mux := runtime.NewServeMux()
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	if err := account.RegisterSignInServiceHandlerServer(ctx, mux, &service.AccountService{}); err != nil {
		panic(err)
	}
	if err := http.ListenAndServe(":80", mux); err != nil {
		panic(err)
	}
}
