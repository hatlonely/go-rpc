package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/hatlonely/go-kit/binding"
	"github.com/hatlonely/go-kit/cli"
	"github.com/hatlonely/go-kit/config"
	"github.com/hatlonely/go-kit/flag"
	"github.com/hatlonely/go-kit/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"

	account "github.com/hatlonely/go-rpc/rpc-account/api/gen/go/api"
	"github.com/hatlonely/go-rpc/rpc-account/internal/service"
	"github.com/hatlonely/go-rpc/rpc-account/pkg/grpcex"
)

var Version string

type Options struct {
	flag.Options

	Http struct {
		Port int
	}
	Grpc struct {
		Port int
	}

	Redis   cli.RedisOptions
	Mysql   cli.MySQLOptions
	Email   cli.EmailOptions
	Account service.AccountServiceOptions
}

func Must(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	var options Options
	Must(flag.Struct(&options))
	Must(flag.Parse())
	if options.Help {
		fmt.Println(flag.Usage())
		return
	}
	if options.Version {
		fmt.Println(Version)
		return
	}

	if options.ConfigPath == "" {
		options.ConfigPath = "config/account.json"
	}
	cfg, err := config.NewSimpleFileConfig(options.ConfigPath)
	if err != nil {
		panic(err)
	}
	Must(binding.Bind(&options, flag.Instance(), binding.NewEnvGetter(binding.WithEnvPrefix("ACCOUNT")), cfg))

	grpcLog, err := logger.NewLoggerWithConfig(cfg.Sub("logger.grpc"))
	if err != nil {
		panic(err)
	}
	infoLog, err := logger.NewLoggerWithConfig(cfg.Sub("logger.info"))
	if err != nil {
		panic(err)
	}

	redisCli, err := cli.NewRedisWithOptions(&options.Redis)
	if err != nil {
		panic(err)
	}
	mysqlCli, err := cli.NewMysqlWithOptions(&options.Mysql)
	if err != nil {
		panic(err)
	}
	emailCli := cli.NewEmailWithOptions(&options.Email)

	svc, err := service.NewAccountService(
		mysqlCli, redisCli, emailCli,
		service.WithAccountExpiration(options.Account.AccountExpiration),
		service.WithCaptchaExpiration(options.Account.CaptchaExpiration),
	)
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer(
		grpcex.WithGrpcLogger(grpcLog),
		grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
			MinTime:             5 * time.Second, // If a client pings more than once every 5 seconds, terminate the connection
			PermitWithoutStream: true,            // Allow pings even when there are no active streams
		}),
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle:     15 * time.Second, // If a client is idle for 15 seconds, send a GOAWAY
			MaxConnectionAge:      30 * time.Second, // If any connection is alive for more than 30 seconds, send a GOAWAY
			MaxConnectionAgeGrace: 5 * time.Second,  // Allow 5 seconds for pending RPCs to complete before forcibly closing connections
			Time:                  5 * time.Second,  // Ping the client if it is idle for 5 seconds to ensure the connection is still active
			Timeout:               1 * time.Second,  // Wait 1 second for the ping ack before assuming the connection is dead
		}),
	)
	account.RegisterAccountServiceServer(server, svc)
	address, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%v", options.Grpc.Port))
	if err != nil {
		panic(err)
	}

	go func() {
		if err := server.Serve(address); err != nil {
			panic(err)
		}
	}()

	mux := runtime.NewServeMux(
		grpcex.WithMuxMetadata(),
		grpcex.WithMuxIncomingHeaderMatcher(),
		grpcex.WithMuxOutgoingHeaderMatcher(),
		grpcex.WithMuxProtoErrorHandler(),
	)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	if err := account.RegisterAccountServiceHandlerFromEndpoint(
		ctx, mux, fmt.Sprintf("0.0.0.0:%v", options.Grpc.Port), []grpc.DialOption{grpc.WithInsecure()},
	); err != nil {
		panic(err)
	}

	infoLog.Info(options)

	if err := http.ListenAndServe(fmt.Sprintf(":%v", options.Http.Port), handlers.CombinedLoggingHandler(os.Stdout, mux)); err != nil {
		panic(err)
	}
}
