package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

//REF: https://qiita.com/takehanKosuke/items/1b17ade882b50cf2d737

// マッピング用の構造体
type Config struct {
	Server   Server
	Firebase Firebase
}

type Server struct {
	Api string
}

type Firebase struct {
	Bucket        string
	Database      string
	ProjectID     string
	StorageExpKey string
}

func Load() (*Config, error) {

	// .env ファイルが存在する場合にのみ読み込む
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	cfg := &Config{
		Server: Server{
			Api: os.Getenv("X_API_KEY"),
		},
		Firebase: Firebase{
			Bucket:        os.Getenv("FIREBASE_BUCKET"),
			Database:      os.Getenv("FIREBASE_DATABASE"),
			ProjectID:     os.Getenv("FIREBASE_PROJECT_ID"),
			StorageExpKey: os.Getenv("FIREBASE_STORAGE_EXP_KEY"),
		},
	}

	hoge := os.Getenv("X_API_KEY")
	fmt.Println(hoge)

	if cfg.Server.Api == "" || cfg.Firebase.Bucket == "" || cfg.Firebase.Database == "" || cfg.Firebase.ProjectID == "" || cfg.Firebase.StorageExpKey == "" {
		return nil, fmt.Errorf("環境変数の設定が不足しています")
	}

	return cfg, nil
}
