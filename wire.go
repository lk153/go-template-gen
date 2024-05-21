//go:build wireinject
// +build wireinject

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"

	envpkg "github.com/lk153/go-template-gen/internal/env"
	userspkg "github.com/lk153/go-template-gen/internal/users"
	yamlpkg "github.com/lk153/go-template-gen/internal/yaml"
)

func InitializeRouter() (userspkg.Controller, error) {
	wire.Build(userspkg.UserSet)
	return userspkg.Controller{}, nil
}

func InitEnv() (yamlpkg.YamlProcessor, error) {
	wire.Build(yamlpkg.InitYamlProcessor, envpkg.InitEnvVars)
	return yamlpkg.YamlProcessor{}, nil
}

func Startup() {
	r := gin.Default()
	InitEnv()
	ctr, err := InitializeRouter()
	if err != nil {
		fmt.Println("SERVER ERROR:", err)
		os.Exit(1)
	}

	r.GET("/users", ctr.GetUsers)
	if err = r.Run(":8888"); err != nil {
		log.Fatalf("impossible to start server: %s", err)
	}
}
