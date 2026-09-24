package handlers

import (
	"errors"
	"net/http"

	"api-gin/services"
	"github.com/gin-gonic/gin"
)

func responderErro(c *gin.Context, err error) {
	switch {
	case errors.Is(err, services.ErrNaoEncontrado):
		c.JSON(http.StatusNotFound, gin.H{"erro": "recurso não localizado"})
	case errors.Is(err, services.ErrDuplicado):
		c.JSON(http.StatusConflict, gin.H{"erro": "duplicidade"})
	case errors.Is(err, services.ErrCapacidadeInsuficiente):
		c.JSON(http.StatusUnprocessableEntity, gin.H{"erro": "capacidade insuficiente"})
	case errors.Is(err, services.ErrConflitoAgenda):
		c.JSON(http.StatusConflict, gin.H{"erro": "conflito de agenda"})
	case errors.Is(err, services.ErrValidacao):
		c.JSON(http.StatusBadRequest, gin.H{"erro": "dados inválidos"})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "erro interno do servidor"})
	}
}
