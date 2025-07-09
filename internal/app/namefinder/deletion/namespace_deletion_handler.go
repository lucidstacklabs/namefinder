package deletion

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/lucidstacklabs/namefinder/internal/app/namefinder/dns"
	"github.com/lucidstacklabs/namefinder/internal/app/namefinder/namespace"
	"github.com/lucidstacklabs/namefinder/internal/pkg/auth"
	"net/http"
	"time"
)

type NamespaceDeletionHandler struct {
	router              *gin.Engine
	authenticator       *auth.Authenticator
	namespaceService    *namespace.Service
	apiKeyAccessService *namespace.ApiKeyAccessService
	recordService       *dns.RecordService
}

func NewNamespaceDeletionHandler(router *gin.Engine, authenticator *auth.Authenticator, namespaceService *namespace.Service, apiKeyAccessService *namespace.ApiKeyAccessService, recordService *dns.RecordService) *NamespaceDeletionHandler {
	return &NamespaceDeletionHandler{router: router, authenticator: authenticator, namespaceService: namespaceService, apiKeyAccessService: apiKeyAccessService, recordService: recordService}
}

func (h *NamespaceDeletionHandler) Register() {

	h.router.DELETE("/api/v1/namespaces/:namespaceID", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c, time.Second*5)
		defer cancel()

		_, err := h.authenticator.ValidateAdminContext(c)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": err.Error(),
			})

			return
		}

		namespaceID := c.Param("namespaceID")

		namespaceDetails, err := h.namespaceService.Delete(ctx, namespaceID)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": err.Error(),
			})

			return
		}

		err = h.apiKeyAccessService.DeleteByNamespaceID(ctx, namespaceID)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": err.Error(),
			})

			return
		}

		err = h.recordService.DeleteByNamespaceID(ctx, namespaceID)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, namespaceDetails)
	})
}
