//go:build wireinject
// +build wireinject

package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"

	envpkg "github.com/lk153/go-template-gen/internal/env"
	userspkg "github.com/lk153/go-template-gen/internal/users"
	yamlpkg "github.com/lk153/go-template-gen/internal/yaml"
)

func InitializeRouter(r *gin.Engine) userspkg.Controller {
	panic(wire.Build(userspkg.UserSet))
}

func InitEnv() (yamlpkg.YamlProcessor, error) {
	wire.Build(yamlpkg.InitYamlProcessor, envpkg.InitEnvVars)
	return yamlpkg.YamlProcessor{}, nil
}

func Startup() {
	r := gin.Default()
	InitEnv()
	InitializeRouter(r)
	srv := &http.Server{
		Addr:    ":8888",
		Handler: r.Handler(),
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listen on: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server shutdown with error:", err)
	}

	select {
	case <-ctx.Done():
		log.Println("Timeout!!!")
	}
	log.Println("Server exiting")
}
