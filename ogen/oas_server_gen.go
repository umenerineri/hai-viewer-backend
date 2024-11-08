// Code generated by ogen, DO NOT EDIT.

package ogen

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// APIHandlerViewGet implements GET /api/handler/view operation.
	//
	// Viewer Page for human AI drawings.
	//
	// GET /api/handler/view
	APIHandlerViewGet(ctx context.Context) (APIHandlerViewGetRes, error)
	// NewError creates *ErrRespStatusCode from error returned by handler.
	//
	// Used for common default response.
	NewError(ctx context.Context, err error) *ErrRespStatusCode
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h   Handler
	sec SecurityHandler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, sec SecurityHandler, opts ...ServerOption) (*Server, error) {
	s, err := newServerConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		sec:        sec,
		baseServer: s,
	}, nil
}
