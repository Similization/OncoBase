package server

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

// Run starts the http server with the given host and port,
// and serves the handler.
//
// Parameters:
// - host: The host address on which the server will listen.
// - port: The port on which the server will listen.
// - handler: The http.Handler to use for serving requests.
//
// Returns:
// - error: An error if the server fails to start.
func (s *Server) Run(host string, port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           host + ":" + port, //nolint:goimports
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	// ListenAndServe starts an HTTP server with the given handler.
	// It returns an error if the server fails to start.
	return s.httpServer.ListenAndServe()
}

// Shutdown gracefully shuts down the server without interrupting any
// active connections. If the provided context expires before the
// shutdown is complete, Shutdown returns the context's error, otherwise
// it returns any error returned from closing the server's listener.
//
// The context is used to limit the amount of time spent shutting down.
// A zero context.Context will not limit the shutdown.
//
// The Shutdown method does not close the provided handler. Therefore,
// the caller is responsible for closing the handler after Shutdown
// returns.
//
// Parameters:
// - ctx: The context.Context used to limit the amount of time spent shutting down.
//
// Returns:
// - error: An error if the server fails to shutdown, otherwise nil.
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
