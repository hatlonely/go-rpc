package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/hatlonely/go-kit/binding"
	"github.com/hatlonely/go-kit/cli"
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

	Redis cli.RedisOptions `bind:"redis"`
	Mysql cli.MySQLOptions `bind:"mysql"`
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

	if err := binding.Bind(&options, flag.Instance(), binding.NewEnvGetter(), conf); err != nil {
		panic(err)
	}

	redis, err := cli.NewRedisWithOptions(&options.Redis)
	if err != nil {
		panic(err)
	}
	mysql, err := cli.NewMysqlWithOptions(&options.Mysql)
	if err != nil {
		panic(err)
	}

	service := service.NewAccountService(mysql, redis)

	mux := runtime.NewServeMux()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	if err := account.RegisterSignInServiceHandlerServer(ctx, mux, service); err != nil {
		panic(err)
	}
	if err := http.ListenAndServe(fmt.Sprintf(":%v", options.Port), mux); err != nil {
		panic(err)
	}
}
