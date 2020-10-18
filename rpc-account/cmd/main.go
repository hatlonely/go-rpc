package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"runtime/debug"
	"time"

	"github.com/gorilla/handlers"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/hatlonely/go-kit/binding"
	"github.com/hatlonely/go-kit/cli"
	"github.com/hatlonely/go-kit/config"
	"github.com/hatlonely/go-kit/flag"
	"github.com/hatlonely/go-kit/logger"
	"github.com/hatlonely/go-kit/validator"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"

	account "github.com/hatlonely/go-rpc/rpc-account/api/gen/go/api"
	"github.com/hatlonely/go-rpc/rpc-account/internal/service"
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

	cfg, err := config.NewConfigWithBaseFile("config/base.json")
	if err != nil {
		panic(err)
	}
	Must(binding.Bind(&options, flag.Instance(), binding.NewEnvGetter(), cfg))

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

	service, err := service.NewAccountService(
		mysqlCli, redisCli, emailCli,
		service.WithAccountExpiration(options.Account.AccountExpiration),
		service.WithCaptchaExpiration(options.Account.CaptchaExpiration),
	)
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer(
		grpc.UnaryInterceptor(func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (res interface{}, err error) {
			ts := time.Now()
			defer func() {
				if perr := recover(); perr != nil {
					err = errors.Wrap(fmt.Errorf("%v\n%v", string(debug.Stack()), perr), "panic")
				}
				p, ok := peer.FromContext(ctx)
				clientIP := ""
				if ok && p != nil {
					clientIP = p.Addr.String()
				}
				grpcLog.Info(map[string]interface{}{
					"client":    clientIP,
					"url":       info.FullMethod,
					"req":       req,
					"res":       res,
					"err":       err,
					"resTimeNs": time.Now().Sub(ts).Nanoseconds(),
				})
			}()

			if err = validator.Validate(req); err != nil {
				return nil, status.Error(codes.InvalidArgument, err.Error())
			}

			return handler(ctx, req)
		}),
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
	account.RegisterAccountServiceServer(server, service)
	address, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%v", options.Grpc.Port))
	if err != nil {
		panic(err)
	}

	go func() {
		if err := server.Serve(address); err != nil {
			panic(err)
		}
	}()

	mux := runtime.NewServeMux()
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
