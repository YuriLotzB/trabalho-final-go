package handlers

import (
	"net/http"
	"strconv"

	"api-gin/models"
	"api-gin/services"
	"github.com/gin-gonic/gin"
)

type TurmaHandler struct {
	service         *services.TurmaService
	alocacaoService *services.AlocacaoService
}

func NewTurmaHandler(service *services.TurmaService, alocacaoService *services.AlocacaoService) *TurmaHandler {
	return &TurmaHandler{service: service, alocacaoService: alocacaoService}
}

func (h *TurmaHandler) Criar(c *gin.Context) {
	var req models.Turma
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "JSON inválido"})
		return
	}
	criada, err := h.service.Criar(req)
	if err != nil {
		responderErro(c, err)
		return
	}
	c.JSON(http.StatusCreated, criada)
}
func (h *TurmaHandler) Listar(c *gin.Context) {
	itens, err := h.service.Listar()
	if err != nil {
		responderErro(c, err)
		return
	}
	c.JSON(http.StatusOK, itens)
}
func (h *TurmaHandler) BuscarPorID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "id inválido"})
		return
	}
	turma, err := h.service.BuscarPorID(id)
	if err != nil {
		responderErro(c, err)
		return
	}
	c.JSON(http.StatusOK, turma)
}

type matriculaRequest struct {
	AlunoID int `json:"aluno_id" binding:"required"`
}

func (h *TurmaHandler) MatricularAluno(c *gin.Context) {
	turmaID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "id inválido"})
		return
	}
	var req matriculaRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "aluno_id é obrigatório"})
		return
	}
	turma, err := h.service.MatricularAluno(turmaID, req.AlunoID)
	if err != nil {
		responderErro(c, err)
		return
	}
	c.JSON(http.StatusOK, turma)
}
func (h *TurmaHandler) ListarAlunos(c *gin.Context) {
	turmaID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "id inválido"})
		return
	}
	alunos, err := h.service.ListarAlunos(turmaID)
	if err != nil {
		responderErro(c, err)
		return
	}
	c.JSON(http.StatusOK, alunos)
}

type alocacaoRequest struct {
	SalaID     int    `json:"sala_id" binding:"required"`
	DiaSemana  string `json:"dia_semana" binding:"required"`
	HoraInicio string `json:"hora_inicio" binding:"required"`
	HoraFim    string `json:"hora_fim" binding:"required"`
}

func (h *TurmaHandler) AlocarSala(c *gin.Context) {
	turmaID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "id inválido"})
		return
	}
	var req alocacaoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "dados de alocação inválidos"})
		return
	}
	alocacao, err := h.alocacaoService.Alocar(turmaID, req.SalaID, req.DiaSemana, req.HoraInicio, req.HoraFim)
	if err != nil {
		responderErro(c, err)
		return
	}
	c.JSON(http.StatusCreated, alocacao)
}
