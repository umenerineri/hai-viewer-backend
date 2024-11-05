package handler

import (
	"net/http"

	"github.com/umenerineri/hai-viewer-backend/api/_pkg/config"
	"github.com/umenerineri/hai-viewer-backend/api/_pkg/middleware"
	"github.com/umenerineri/hai-viewer-backend/api/_pkg/ogen"
	"github.com/umenerineri/hai-viewer-backend/api/_pkg/presentation/controller"
)

func Handler(w http.ResponseWriter, req *http.Request) {
	// Initialize the ogen server with your handlers
	hdl, err := ogen.NewServer(
		&controller.HaiHandler{},
		&config.HaiSecurityHandler{},
	)

	if err != nil {
		http.Error(w, "Failed to initialize server: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Wrap your handler with middleware
	handler := middleware.EnableCORS(middleware.LoggingMiddleware(hdl))

	// Serve the HTTP request
	handler.ServeHTTP(w, req)
}
