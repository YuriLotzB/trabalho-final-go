package main

import (
	"github.com/gin-gonic/gin"

	"api-gin/handlers"
	"api-gin/repositories"
	"api-gin/routes"
	"api-gin/services"
)

const version = "1.0.0"

func main() {
	r := gin.New()
	r.Use(gin.Recovery())
	salaRepo := repositories.NewInMemorySalaRepository()
	alunoRepo := repositories.NewInMemoryAlunoRepository()
	turmaRepo := repositories.NewInMemoryTurmaRepository()
	alocacaoRepo := repositories.NewInMemoryAlocacaoRepository()
	salaService := services.NewSalaService(salaRepo)
	alunoService := services.NewAlunoService(alunoRepo)
	alocacaoService := services.NewAlocacaoService(turmaRepo, salaRepo, alocacaoRepo, alunoRepo)
	turmaService := services.NewTurmaService(turmaRepo, alunoRepo, salaRepo, alocacaoRepo)
	healthHandler := handlers.NewHealthHandler(version)
	salaHandler := handlers.NewSalaHandler(salaService)
	alunoHandler := handlers.NewAlunoHandler(alunoService)
	turmaHandler := handlers.NewTurmaHandler(turmaService, alocacaoService)
	v1 := r.Group("/api/v1")
	routes.Register(v1, healthHandler, salaHandler, alunoHandler, turmaHandler)
	_ = r
}
