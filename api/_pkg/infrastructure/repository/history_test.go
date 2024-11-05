package impl_repository_test

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	config "github.com/umenerineri/hai-viewer-backend/api/_pkg/config"
	"github.com/umenerineri/hai-viewer-backend/api/_pkg/domain/entity/history"
	impl_repository "github.com/umenerineri/hai-viewer-backend/api/_pkg/infrastructure/repository"
)

func TestImplHistoryRepository_Integration(t *testing.T) {
	t.Parallel()
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get current working directory: %v", err)
	}
	t.Cleanup(func() {
		err := os.Chdir(cwd)
		if err != nil {
			t.Fatalf("failed to get current working directory: %v", err)
		}
	})
	err = os.Chdir("../..")
	if err != nil {
		t.Fatalf("failed to get current working directory: %v", err)
	}
	ctx := context.Background()

	// Firestoreクライアントの初期化
	app, err := config.InitializeApp()
	if err != nil {
		t.Fatalf("failed to initialize Firebase app: %v", err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		t.Fatalf("failed to initialize Firestore client: %v", err)
	}
	defer client.Close()

	historyRepo, err := impl_repository.NewImplHistoryRepository(ctx)
	if err != nil {
		t.Fatalf("failed to get history repo: %v", err)
	}

	// テストデータの作成

	/** NOTE: ここでを書き換える場合はuser_test.goのtestIdも書き換える */
	version := 0
	/** */
	historyData := history.NewHistory(version)

	// Createのテスト
	err = historyRepo.Create(*historyData)
	assert.NoError(t, err)
}
