// Package httpserver implements HTTP server.
package httpserver

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

const (
	_defaultReadTimeout     = 20 * time.Second
	_defaultWriteTimeout    = 20 * time.Second
	_defaultPort            = "7070"
	_defaultShutdownTimeout = 20 * time.Second
)

type Server struct {
	server          *http.Server
	notify          chan error
	shutdownTimeout time.Duration
}

func New(handler http.Handler, port string, opts ...Option) *Server {
	if port == "" {
		port = _defaultPort
	}
	httpServer := &http.Server{
		Handler:      handler,
		ReadTimeout:  _defaultReadTimeout,
		WriteTimeout: _defaultWriteTimeout,
		Addr:         fmt.Sprintf(":%s", port),
	}

	s := &Server{
		server:          httpServer,
		notify:          make(chan error, 1),
		shutdownTimeout: _defaultShutdownTimeout,
	}

	for _, opt := range opts {
		opt(s)
	}

	s.start()

	return s
}

func (s *Server) start() {
	go func() {
		s.notify <- s.server.ListenAndServe()
		close(s.notify)
	}()
}

func (s *Server) Notify() <-chan error {
	return s.notify
}

func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	return s.server.Shutdown(ctx)
}
