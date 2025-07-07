package dns

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/lucidstacklabs/namefinder/internal/app/namefinder/namespace"
	"github.com/lucidstacklabs/namefinder/internal/pkg/auth"
	"net/http"
	"strconv"
	"time"
)

type RecordHandler struct {
	router              *gin.Engine
	authenticator       *auth.Authenticator
	apiKeyAccessService *namespace.ApiKeyAccessService
	recordService       *RecordService
}

func NewRecordHandler(router *gin.Engine, authenticator *auth.Authenticator, apiKeyAccessService *namespace.ApiKeyAccessService, recordService *RecordService) *RecordHandler {
	return &RecordHandler{router: router, authenticator: authenticator, apiKeyAccessService: apiKeyAccessService, recordService: recordService}
}

func (h *RecordHandler) Register() {

	h.router.POST("/api/v1/namespaces/:namespaceID/records", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c, time.Second*5)
		defer cancel()

		apiKey, err := h.authenticator.ValidateApiKeyContext(c, ctx)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": err.Error(),
			})

			return
		}

		namespaceID := c.Param("namespaceID")

		hasPermission, err := h.apiKeyAccessService.HasPermission(ctx, namespaceID, apiKey.ID, namespace.ActionCreate)

		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{
				"success": false,
				"message": err.Error(),
			})

			return
		}

		if !hasPermission {
			c.JSON(http.StatusForbidden, gin.H{
				"success": false,
				"message": "Permission denied",
			})

			return
		}

		var req RecordAdditionRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": err.Error(),
			})

			return
		}

		record, err := h.recordService.Add(ctx, namespaceID, &req, ActorTypeApiKey, apiKey.ID)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": err.Error(),
			})

			return
		}

		c.JSON(http.StatusCreated, record)
	})

	h.router.GET("/api/v1/namespaces/:namespaceID/records", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c, time.Second*5)
		defer cancel()

		apiKey, err := h.authenticator.ValidateApiKeyContext(c, ctx)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": err.Error(),
			})

			return
		}

		namespaceID := c.Param("namespaceID")

		hasPermission, err := h.apiKeyAccessService.HasPermission(ctx, namespaceID, apiKey.ID, namespace.ActionRead)

		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{
				"success": false,
				"message": err.Error(),
			})

			return
		}

		if !hasPermission {
			c.JSON(http.StatusForbidden, gin.H{
				"success": false,
				"message": "Permission denied",
			})

			return
		}

		page := c.DefaultQuery("page", "0")
		size := c.DefaultQuery("size", "50")
		pageInt, err := strconv.ParseInt(page, 10, 64)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": err.Error(),
			})

			return
		}

		sizeInt, err := strconv.ParseInt(size, 10, 64)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": err.Error(),
			})

			return
		}

		records, err := h.recordService.List(ctx, namespaceID, pageInt, sizeInt)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, records)
	})

	h.router.GET("/api/v1/namespaces/:namespaceID/records/:recordID", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c, time.Second*5)
		defer cancel()

		apiKey, err := h.authenticator.ValidateApiKeyContext(c, ctx)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": err.Error(),
			})

			return
		}

		namespaceID := c.Param("namespaceID")

		hasPermission, err := h.apiKeyAccessService.HasPermission(ctx, namespaceID, apiKey.ID, namespace.ActionRead)

		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{
				"success": false,
				"message": err.Error(),
			})

			return
		}

		if !hasPermission {
			c.JSON(http.StatusForbidden, gin.H{
				"success": false,
				"message": "Permission denied",
			})

			return
		}

		recordID := c.Param("recordID")
		record, err := h.recordService.Get(ctx, namespaceID, recordID)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, record)
	})

	h.router.PUT("/api/v1/namespaces/:namespaceID/records/:recordID", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c, time.Second*5)
		defer cancel()

		apiKey, err := h.authenticator.ValidateApiKeyContext(c, ctx)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": err.Error(),
			})

			return
		}

		namespaceID := c.Param("namespaceID")

		hasPermission, err := h.apiKeyAccessService.HasPermission(ctx, namespaceID, apiKey.ID, namespace.ActionUpdate)

		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{
				"success": false,
				"message": err.Error(),
			})

			return
		}

		if !hasPermission {
			c.JSON(http.StatusForbidden, gin.H{
				"success": false,
				"message": "Permission denied",
			})

			return
		}

		recordID := c.Param("recordID")

		var req RecordUpdateRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": err.Error(),
			})

			return
		}

		record, err := h.recordService.Update(ctx, namespaceID, recordID, &req)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, record)
	})
}
