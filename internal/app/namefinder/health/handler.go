package health

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type CheckHandler struct {
	router *gin.Engine
}

func NewCheckHandler(router *gin.Engine) *CheckHandler {
	return &CheckHandler{
		router: router,
	}
}

func (h *CheckHandler) Register() {

	h.router.GET("/health", func(c *gin.Context) {
		c.JSONP(http.StatusOK, &CheckResponse{Status: "ok"})
	})
}
