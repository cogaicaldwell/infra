package utils

import (
	"context"
	"errors"

	"google.golang.org/grpc/metadata"

	"coginfra/interceptors"
	"coginfra/storage"
)

func ValidateAPIKey(ctx context.Context, s storage.Storage) (bool, error) {
	apiKey, err := extractAPIKeyFromHeader(ctx)
	if err != nil {
		return false, err
	}
	// Logger.Println(apiKey)
	return s.ValidateApiKey(apiKey)
}

func extractAPIKeyFromHeader(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", errors.New("cannot extract metadata")
	} else {
		return md.Get(interceptors.XAPIKey)[0], nil
	}
}
