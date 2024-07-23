package rest

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/online_marketplace/helper/errors"
	"github.com/online_marketplace/internal/config"
)

type (
	// Strategy Pattern
	// RunOption defines the method to customize a Server.
	RunOption func(*Server)

	// StartOption defines the method to customize http server.

	// A Server is a http server.
	Server struct {
		config config.ServerConfig
		route  *gin.Engine
		srv    *http.Server
	}
)
// factory method pattern
func MustNewServer(config config.ServerConfig, route *gin.Engine, opts ...RunOption) *Server {
	server, err := NewServer(config, route, opts...)
	if err != nil {
		panic(err)
	}

	return server
}

// factory method pattern
func NewServer(config config.ServerConfig, route *gin.Engine, opts ...RunOption) (*Server, error) {

	server := &Server{
		config: config,
		route:  route,
	}

	for _, opt := range opts {
		opt(server)
	}

	return server, nil
}

// Start starts the Server.
// Graceful shutdown is enabled by default.
// Use proc.SetTimeToForceQuit to customize the graceful shutdown period.
func (s *Server) Start() {
	// Server address
	serverAddr := fmt.Sprintf("%s:%d", s.config.Http.Host, s.config.Http.Port)

	// HTTP server
	s.srv = &http.Server{
		Addr:    serverAddr,
		Handler: s.route,
	}

	if err := s.srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("Server failed to start: %v", err)
	}
}

// Stop stops the Server.
func (s *Server) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.srv.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown: %v", err)
	} else {
		log.Println("Server gracefully stopped")
	}

}
