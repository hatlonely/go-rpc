package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/hatlonely/go-kit/bind"
	"github.com/hatlonely/go-kit/config"
	"github.com/hatlonely/go-kit/flag"
	"github.com/hatlonely/go-kit/logger"
	"github.com/hatlonely/go-kit/refx"
	"github.com/hatlonely/go-kit/rpcx"
	"google.golang.org/grpc"

	"github.com/hatlonely/go-rpc/rpc-cicd/api/gen/go/api"
	"github.com/hatlonely/go-rpc/rpc-cicd/internal/cli"
	"github.com/hatlonely/go-rpc/rpc-cicd/internal/service"
)

var Version string

type Options struct {
	flag.Options

	Http struct {
		Port int
		Cors rpcx.CORSOptions
	}
	Grpc struct {
		Port int
	}

	Service service.Options

	Mongo cli.MongoOptions

	Logger struct {
		Info logger.Options
		Grpc logger.Options
	}
}

func Must(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	var options Options
	Must(flag.Struct(&options, refx.WithCamelName()))
	Must(flag.Parse(flag.WithJsonVal()))
	if options.Help {
		fmt.Println(flag.Usage())
		return
	}
	if options.Version {
		fmt.Println(Version)
		return
	}

	if options.ConfigPath == "" {
		options.ConfigPath = "config/go-rpc-cicd.json"
	}
	cfg, err := config.NewConfigWithSimpleFile(options.ConfigPath)
	Must(err)
	Must(bind.Bind(&options, []bind.Getter{
		flag.Instance(), bind.NewEnvGetter(bind.WithEnvPrefix("ACCOUNT")), cfg,
	}, refx.WithCamelName()))

	grpcLog, err := logger.NewLoggerWithOptions(&options.Logger.Grpc)
	Must(err)
	infoLog, err := logger.NewLoggerWithOptions(&options.Logger.Info)
	Must(err)

	mongoCli, err := cli.NewMongoWithOptions(&options.Mongo)
	Must(err)

	svc, err := service.NewCICDServiceWithOptions(mongoCli, &options.Service)
	Must(err)
	rpcServer := grpc.NewServer(rpcx.WithGRPCDecorator(grpcLog))
	api.RegisterCICDServiceServer(rpcServer, svc)

	go func() {
		address, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%v", options.Grpc.Port))
		Must(err)
		Must(rpcServer.Serve(address))
	}()

	muxServer := runtime.NewServeMux(
		rpcx.WithMuxMetadata(),
		rpcx.WithMuxIncomingHeaderMatcher(),
		rpcx.WithMuxOutgoingHeaderMatcher(),
		rpcx.WithMuxProtoErrorHandler(),
	)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	Must(api.RegisterCICDServiceHandlerFromEndpoint(
		ctx, muxServer, fmt.Sprintf("0.0.0.0:%v", options.Grpc.Port), []grpc.DialOption{grpc.WithInsecure()},
	))
	infoLog.Info(options)
	Must(http.ListenAndServe(fmt.Sprintf(":%v", options.Http.Port), rpcx.CORSWithOptions(handlers.CombinedLoggingHandler(os.Stdout, muxServer), &options.Http.Cors)))
}
