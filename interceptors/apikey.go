package interceptors

import (
	"context"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type CtxApiKeyTpe int

const (
	// XAPIKey is a key for getting request id.
	XAPIKey                    = "X-API-Key"
	unknownAPIKey              = "<unknown>"
	CtxApiKey     CtxApiKeyTpe = iota + 1
)

// APIKeyInterceptor is a interceptor of access control list.
func APIKeyInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		requestID := apiKeyFromContext(ctx)
		ctx = context.WithValue(ctx, CtxApiKey, requestID)
		return handler(ctx, req)
	}
}

func apiKeyFromContext(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return unknownAPIKey
	}

	key := strings.ToLower(XAPIKey)
	header, ok := md[key]
	if !ok || len(header) == 0 {
		return unknownAPIKey
	}

	apiKey := header[0]
	if apiKey == "" {
		return unknownAPIKey
	}

	return apiKey
}
