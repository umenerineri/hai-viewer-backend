package config

import (
	"context"
	"errors"

	ogen "github.com/umenerineri/hai-viewer-backend/api/_pkg/ogen"
)

type HaiSecurityHandler struct{}

func (m *HaiSecurityHandler) HandleApiKeyAuth(ctx context.Context, operationName string, auth ogen.ApiKeyAuth) (context.Context, error) {
	// 認証ロジックをここに実装
	cfg, err := Load()
	if err != nil {
		return nil, err
	}
	validAPIKey := cfg.Server.Api
	if auth.APIKey == validAPIKey {
		return ctx, nil
	}
	return nil, errors.New("invalid API key")
}
