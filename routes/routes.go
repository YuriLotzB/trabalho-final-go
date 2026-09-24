package routes

import (
	"api-gin/handlers"
	"github.com/gin-gonic/gin"
)

func Register(v1 *gin.RouterGroup, health *handlers.HealthHandler, sala *handlers.SalaHandler, aluno *handlers.AlunoHandler, turma *handlers.TurmaHandler) {
	v1.GET("/health", health.Check)
	v1.POST("/salas", sala.Criar)
	v1.GET("/salas", sala.Listar)
	v1.GET("/salas/:id", sala.BuscarPorID)
	v1.POST("/alunos", aluno.Criar)
	v1.GET("/alunos", aluno.Listar)
	v1.GET("/alunos/:id", aluno.BuscarPorID)
	v1.POST("/turmas", turma.Criar)
	v1.GET("/turmas", turma.Listar)
	v1.GET("/turmas/:id", turma.BuscarPorID)
	v1.POST("/turmas/:id/alunos", turma.MatricularAluno)
	v1.GET("/turmas/:id/alunos", turma.ListarAlunos)
	v1.POST("/turmas/:id/alocar", turma.AlocarSala)
}
