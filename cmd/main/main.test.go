package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/online_marketplace/internal/config"
	"github.com/online_marketplace/internal/handler"
	"github.com/online_marketplace/internal/registry"
)

func main_test() {
	config := config.Init()
	appCtx := registry.NewServiceContext(*config)
	fmt.Printf("Loaded Configuration: %+v\n", config)

	router := handler.NewRouter(appCtx)

	// Server address
	serverAddr := fmt.Sprintf("%s:%d", config.Server.Http.Host, config.Server.Http.Port)

	// HTTP server
	srv := &http.Server{
		Addr:    serverAddr,
		Handler: router,
	}

	// Graceful shutdown
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Printf("Server started on %s", serverAddr)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")

}
