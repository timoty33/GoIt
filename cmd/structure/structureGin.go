package structure

import (
	"fmt"
	"goit/utils"
	"os"
	"path/filepath"
)

const (
	permPasta   = 0755
	permArquivo = 0644
)

func StructureGin(nomeProjeto string) error {
	mainGoContent := `package main

import (
	"` + nomeProjeto + `/internal/server"
)

func main() {
	s := server.New()
	s.Run()
}
`
	serverGoContent := `package server

import (
	"github.com/gin-gonic/gin"
	"` + nomeProjeto + `/internal/routes"
)

type Server struct {
	engine *gin.Engine
}

func New() *Server {
	r := gin.Default()
	routes.RegisterRoutes(r)

	return &Server{engine: r}
}

func (s *Server) Run() {
	s.engine.Run()
}
`
	handlerGoContent := `package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

// Toda vez que você ver // goit: <-- Isso é um marcador e não pode ser removido (esse comentário pode)
// goit:add-handlers-here
`
	routesGoContent := `package routes

import (
	"github.com/gin-gonic/gin"
	"` + nomeProjeto + `/internal/handler"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/ping", handler.Ping)

	// Toda vez que você ver // goit: <-- Isso é um marcador e não pode ser removido (esse comentário pode)
	// goit:add-routes-here
}
`
	dtoGoContent := `package dto

// Toda vez que você ver // goit: <-- Isso é um marcador e não pode ser removido (esse comentário pode)
// goit:add-dtos-here
`
	middlewareGoContent := `package middleware

// Toda vez que você ver // goit: <-- Isso é um marcador e não pode ser removido (esse comentário pode)
// goit:add-middlewares-here
`
	gitignoreContent := `.goit.config.json`

	err := os.Mkdir(nomeProjeto, permPasta)
	if err != nil {
		return fmt.Errorf("erro ao criar pasta do projeto %s: %w", nomeProjeto, err)
	}

	projectLayout := map[string]string{
		"cmd/main.go":                       mainGoContent,
		"internal/server/server.go":         serverGoContent,
		"internal/dto/dto.go":               dtoGoContent,
		"internal/handler/handler.go":       handlerGoContent,
		"internal/middleware/middleware.go": middlewareGoContent,
		"internal/routes/routes.go":         routesGoContent,
		"internal/config/":                  "",
		"README.md":                         "# " + nomeProjeto + "\n",
		".gitignore":                        gitignoreContent,
	}

	for path, content := range projectLayout {
		fullPath := filepath.Join(nomeProjeto, path)

		dir := filepath.Dir(fullPath)
		if err := os.MkdirAll(dir, permPasta); err != nil {
			return fmt.Errorf("erro ao criar diretório %s: %w", dir, err)
		}

		if content != "" {
			if err := os.WriteFile(fullPath, []byte(content), permArquivo); err != nil {
				return fmt.Errorf("erro ao criar arquivo %s: %w", fullPath, err)
			}
		}
	}

	configs := utils.Config{
		RoutesFile: "internal/routes/routes.go",

		HandlersFolder: "internal/handler",
		HandlersFile:   "internal/handler/handler.go",

		MiddlewaresFolder: "internal/middleware",
		MiddlewaresFile:   "internal/middleware/middleware.go",

		DtoFolder: "internal/dto",
		DtoFile:   "internal/dto/dto.go",

		ServicesFolder:   "internal/service",
		ModelsFolder:     "internal/model",
		MigrationsFolder: "internal/migrations",
		DatabaseFolder:   "internal/database",

		Framework:   "gin",
		ProjectName: nomeProjeto,
	}

	if err := utils.SaveJsonConfig(configs, nomeProjeto); err != nil {
		return fmt.Errorf("erro ao salvar configurações: %w", err)
	}

	return nil
}
