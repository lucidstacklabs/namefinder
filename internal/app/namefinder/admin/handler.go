package admin

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/lucidstacklabs/namefinder/internal/pkg/auth"
	"net/http"
	"time"
)

type Handler struct {
	router        *gin.Engine
	authenticator *auth.Authenticator
	service       *Service
}

func NewHandler(router *gin.Engine, authenticator *auth.Authenticator, service *Service) *Handler {
	return &Handler{router: router, authenticator: authenticator, service: service}
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

	h.router.POST("/api/v1/admins/token", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c, time.Second*5)
		defer cancel()

		var req TokenRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": err.Error(),
			})
			return
		}

		token, err := h.service.GetToken(ctx, &req)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, token)
	})
}
