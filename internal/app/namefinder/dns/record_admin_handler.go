package dns

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/lucidstacklabs/namefinder/internal/pkg/auth"
	"net/http"
	"strconv"
	"time"
)

type RecordAdminHandler struct {
	router        *gin.Engine
	authenticator *auth.Authenticator
	recordService *RecordService
}

func NewRecordAdminHandler(router *gin.Engine, authenticator *auth.Authenticator, recordService *RecordService) *RecordAdminHandler {
	return &RecordAdminHandler{router: router, authenticator: authenticator, recordService: recordService}
}

func (h *RecordAdminHandler) Register() {

	h.router.POST("/admin/api/v1/namespaces/:namespaceID/records", func(c *gin.Context) {
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

		var req RecordAdditionRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": err.Error(),
			})

			return
		}

		namespaceID := c.Param("namespaceID")

		record, err := h.recordService.Add(ctx, namespaceID, &req, ActorTypeAdmin, aa.ID)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": err.Error(),
			})

			return
		}

		c.JSON(http.StatusCreated, record)
	})

	h.router.GET("/admin/api/v1/namespaces/:namespaceID/records", func(c *gin.Context) {
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

	h.router.GET("/admin/api/v1/namespaces/:namespaceID/records/:recordID", func(c *gin.Context) {
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
}
