package config

import (
	"context"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

func InitializeApp() (*firebase.App, error) {
	cfg, err := Load()
	if err != nil {
		return nil, err
	}
	config := &firebase.Config{
		ProjectID:     cfg.Firebase.ProjectID,
		StorageBucket: cfg.Firebase.Bucket,
	}
	credentials := cfg.Firebase.StorageExpKey
	opt := option.WithCredentialsJSON([]byte(credentials))
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		return nil, err
	}
	return app, nil
}
