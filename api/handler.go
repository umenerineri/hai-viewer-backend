package handler

import (
	"net/http"

	"github.com/umenerineri/hai-viewer-backend/config"
	"github.com/umenerineri/hai-viewer-backend/middleware"
	"github.com/umenerineri/hai-viewer-backend/ogen"
	"github.com/umenerineri/hai-viewer-backend/presentation/controller"
)

// Handler は http.Handler を返す関数として定義
func Handler(w http.ResponseWriter, r *http.Request) {
	hdl, err := ogen.NewServer(
		&controller.HaiHandler{},
		&config.HaiSecurityHandler{},
	)

	if err != nil {
		http.Error(w, "Failed to initialize server: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// ハンドラーをラップしてリクエストを処理
	middleware.EnableCORS(middleware.LoggingMiddleware(hdl)).ServeHTTP(w, r)
}
