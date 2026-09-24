package handlers

import (
	"net/http"
	"strconv"

	"api-gin/models"
	"api-gin/services"
	"github.com/gin-gonic/gin"
)

type SalaHandler struct{ service *services.SalaService }

func NewSalaHandler(service *services.SalaService) *SalaHandler {
	return &SalaHandler{service: service}
}

func (h *SalaHandler) Criar(c *gin.Context) {
	var req models.Sala
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

func (h *SalaHandler) Listar(c *gin.Context) {
	itens, err := h.service.Listar()
	if err != nil {
		responderErro(c, err)
		return
	}
	c.JSON(http.StatusOK, itens)
}

func (h *SalaHandler) BuscarPorID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "id inválido"})
		return
	}
	sala, err := h.service.BuscarPorID(id)
	if err != nil {
		responderErro(c, err)
		return
	}
	c.JSON(http.StatusOK, sala)
}
