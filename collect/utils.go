package collect

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Healthz is used to test if the service is alive
func Healthz(c *gin.Context) {
	c.JSON(http.StatusOK, "I'm alive")
}
