package grpcex

import (
	"context"
	"strings"

	"google.golang.org/grpc/metadata"
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
