package dns

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/lucidstacklabs/namefinder/internal/pkg/auth"
	"net/http"
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
}
