package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct{ Version string }

func NewHealthHandler(version string) *HealthHandler { return &HealthHandler{Version: version} }

func (h *HealthHandler) Check(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "healthy", "timestamp": time.Now(), "version": h.Version})
}
