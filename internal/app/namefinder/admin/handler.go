package admin

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Handler struct {
	router  *gin.Engine
	service *Service
}

func NewHandler(router *gin.Engine, service *Service) *Handler {
	return &Handler{router: router, service: service}
}

func (h *Handler) Register() {

	h.router.POST("/api/v1/admins/init", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c, time.Second*5)
		defer cancel()

		var req InitRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": err.Error(),
			})
			return
		}

		admin, err := h.service.Init(ctx, &req)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusCreated, admin)
	})
}
