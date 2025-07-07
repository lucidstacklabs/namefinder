package namespace

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/lucidstacklabs/namefinder/internal/pkg/auth"
	"net/http"
	"time"
)

type ApiKeyAccessHandler struct {
	router        *gin.Engine
	authenticator *auth.Authenticator
	service       *ApiKeyAccessService
}

func NewApiKeyAccessHandler(router *gin.Engine, authenticator *auth.Authenticator, service *ApiKeyAccessService) *ApiKeyAccessHandler {
	return &ApiKeyAccessHandler{router: router, authenticator: authenticator, service: service}
}

func (h *ApiKeyAccessHandler) Register() {

	h.router.POST("/api/v1/namespaces/:namespaceID/api-keys", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c, time.Second*5)
		defer cancel()

		aa, err := h.authenticator.ValidateAdminContext(c)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": err.Error(),
			})
			return
		}

		namespaceID := c.Param("namespaceID")

		var req ApiKeyAccessRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": err.Error(),
			})

			return
		}

		err = h.service.Add(ctx, namespaceID, &req, aa.ID)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "api key access added successfully to namespace",
		})
	})
}
