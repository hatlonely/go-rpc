package grpcex

import (
	"context"
	"fmt"
	"runtime/debug"
	"strings"
	"time"

	"github.com/hatlonely/go-kit/logger"
	"github.com/hatlonely/go-kit/validator"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

func GetRequestIDFromContext(ctx context.Context) string {
	var requestID string
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if vals, ok := md["x-request-id"]; ok {
			requestID = strings.Join(vals, ",")
		}
	}
	return requestID
}

func WithGrpcLogger(log *logger.Logger) grpc.ServerOption {
	return grpc.UnaryInterceptor(func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		var requestID, remoteIP string
		if md, ok := metadata.FromIncomingContext(ctx); ok {
			if vals, ok := md["x-request-id"]; ok {
				requestID = strings.Join(vals, ",")
			}
			if requestID == "" {
				requestID = uuid.NewV4().String()
				md.Set("x-request-id", requestID)
			}
			if vals, ok := md["x-remote-addr"]; ok {
				remoteIP = strings.Split(strings.Join(vals, ","), ":")[0]
			}
		}

		var res interface{}
		var err error
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
			md, _ := metadata.FromIncomingContext(ctx)
			c, _ := md["ctx"]
			log.Info(map[string]interface{}{
				"requestID": requestID,
				"remoteIP":  remoteIP,
				"clientIP":  clientIP,
				"method":    info.FullMethod,
				"req":       req,
				"ctx":       c,
				"res":       res,
				"err":       err.Error(),
				"errStack":  fmt.Sprintf("%+v", err),
				"resTimeMs": time.Now().Sub(ts).Milliseconds(),
			})
			_ = grpc.SendHeader(ctx, metadata.New(map[string]string{
				"x-request-id": requestID,
			}))
		}()

		if err = validator.Validate(req); err != nil {
			err = NewError(err, codes.InvalidArgument, requestID, "InvalidArgument", err.Error())
		} else {
			res, err = handler(ctx, req)
		}

		switch e := err.(type) {
		case *Error:
			return res, e.ToStatus().Err()
		}
		return res, NewInternalError(err, requestID).ToStatus().Err()
	})
}
