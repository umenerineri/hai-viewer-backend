package controller_test

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	ogen "github.com/umenerineri/hai-viewer-backend/ogen"
	controller "github.com/umenerineri/hai-viewer-backend/presentation/controller"
)

func TestViewHandler(t *testing.T) {
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

	// コントローラーのハンドラを作成
	h := &controller.HaiHandler{}

	// 実行
	res, err := h.APIHandlerViewGet(ctx)
	assert.NoError(t, err, "ImageGenerationPost failed")
	assert.NotNil(t, res, "Response should not be nil")

	// レスポンスの検証
	switch v := res.(type) {
	case *ogen.APIHandlerViewGetOK:
		t.Logf("%d", len(v.Result))
		assert.NotEmpty(t, v.Result[0].URL, "URL should not be empty")
	case *ogen.APIHandlerViewGetBadRequest:
		t.Errorf("expected success but got bad request: %s", v.Error.Value)
	default:
		t.Errorf("unexpected response type: %T", v)
	}
}
