package handlers

import (
	"net/http"
	"strconv"

	"api-gin/models"
	"api-gin/services"
	"github.com/gin-gonic/gin"
)

type AlunoHandler struct{ service *services.AlunoService }

func NewAlunoHandler(service *services.AlunoService) *AlunoHandler {
	return &AlunoHandler{service: service}
}

func (h *AlunoHandler) Criar(c *gin.Context) {
	var req models.Aluno
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "JSON inválido"})
		return
	}
	criado, err := h.service.Criar(req)
	if err != nil {
		responderErro(c, err)
		return
	}
	c.JSON(http.StatusCreated, criado)
}
func (h *AlunoHandler) Listar(c *gin.Context) {
	itens, err := h.service.Listar()
	if err != nil {
		responderErro(c, err)
		return
	}
	c.JSON(http.StatusOK, itens)
}
func (h *AlunoHandler) BuscarPorID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "id inválido"})
		return
	}
	aluno, err := h.service.BuscarPorID(id)
	if err != nil {
		responderErro(c, err)
		return
	}
	c.JSON(http.StatusOK, aluno)
}
