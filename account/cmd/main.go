package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/hatlonely/go-kit/binding"
	"github.com/hatlonely/go-kit/config"
	"github.com/hatlonely/go-kit/flag"

	account "github.com/hatlonely/go-rpc/account/api/gen/go/api"
	"github.com/hatlonely/go-rpc/account/internal/service"
)

var Version string

type Options struct {
	Help    bool   `flag:"--help,-h; default: false; usage: show help info"`
	Version bool   `flag:"--version,-v; default: false; usage: show version"`
	Port    string `bind:"port" flag:"--port; usage: service port"`
}

func main() {
	var options Options
	conf, err := config.NewConfigWithBaseFile("config/base.json")
	if err != nil {
		panic(err)
	}
	if err := flag.Struct(&options); err != nil {
		panic(err)
	}
	if err := flag.Parse(); err != nil {
		panic(err)
	}
	if options.Help {
		fmt.Println(flag.Usage())
		return
	}
	if options.Version {
		fmt.Println(Version)
		return
	}

	if err := binding.Bind(&options, flag.Instance(), conf); err != nil {
		panic(err)
	}

	mux := runtime.NewServeMux()
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	if err := account.RegisterSignInServiceHandlerServer(ctx, mux, &service.AccountService{}); err != nil {
		panic(err)
	}
	if err := http.ListenAndServe(fmt.Sprintf(":%v", options.Port), mux); err != nil {
		panic(err)
	}
}
