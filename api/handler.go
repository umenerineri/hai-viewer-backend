package main

import (
	"net/http"

	"github.com/umenerineri/hai-viewer-backend/config"
	"github.com/umenerineri/hai-viewer-backend/middleware"
	"github.com/umenerineri/hai-viewer-backend/ogen"
	"github.com/umenerineri/hai-viewer-backend/presentation/controller"
)

var Handler http.Handler

func init() {
	hdl, err := ogen.NewServer(
		&controller.HaiHandler{},
		&config.HaiSecurityHandler{},
	)

	if err != nil {
		panic("Failed to initialize server: " + err.Error())
	}

	Handler = middleware.EnableCORS(middleware.LoggingMiddleware(hdl))
}
