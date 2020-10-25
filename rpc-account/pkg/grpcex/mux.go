package grpcex

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func WithMuxMetadata() runtime.ServeMuxOption {
	return runtime.WithMetadata(func(ctx context.Context, req *http.Request) metadata.MD {
		requestID := req.Header.Get("X-Request-Id")
		if requestID == "" {
			return metadata.Pairs("x-remote-addr", req.RemoteAddr, "x-request-id", requestID)
		} else {
			return metadata.Pairs("x-remote-addr", req.RemoteAddr)
		}
	})
}

func WithMuxIncomingHeaderMatcher() runtime.ServeMuxOption {
	return runtime.WithIncomingHeaderMatcher(func(key string) (string, bool) {
		switch key {
		case "X-Request-Id":
			return "x-request-id", true
		default:
			return runtime.DefaultHeaderMatcher(key)
		}
	})
}

func WithMuxOutgoingHeaderMatcher() runtime.ServeMuxOption {
	return runtime.WithOutgoingHeaderMatcher(func(key string) (string, bool) {
		switch key {
		case "X-Request-Id", "x-request-id":
			return "x-request-id", true
		default:
			return runtime.DefaultHeaderMatcher(key)
		}
	})
}

func WithMuxProtoErrorHandler() runtime.ServeMuxOption {
	return runtime.WithProtoErrorHandler(func(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, writer http.ResponseWriter, request *http.Request, err error) {
		requestID := GetRequestIDFromContext(ctx)

		s := status.Convert(err)
		if len(s.Details()) >= 1 {
			if e, ok := s.Details()[0].(*EInfo); ok {
				writer.WriteHeader(int(e.Status))
				buf, _ := json.Marshal(e)
				_, _ = writer.Write(buf)
				return
			}
		}
		e := NewInternalError(err, requestID).Info
		e.Status = int64(runtime.HTTPStatusFromCode(s.Code()))
		e.Code = http.StatusText(int(e.Status))
		writer.WriteHeader(int(e.Status))
		buf, _ := json.Marshal(e)
		_, _ = writer.Write(buf)
	})
}
