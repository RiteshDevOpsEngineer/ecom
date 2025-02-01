// In middleware.go
package middleware

import (
	"net/http"

	"github.com/RiteshDevOpsEngineer/ecom/config"
	"github.com/RiteshDevOpsEngineer/ecom/internal/core/domain"

	"github.com/gin-gonic/gin"
)

func MaintenanceMiddleware(cfg *config.Container) gin.HandlerFunc {
	return func(c *gin.Context) {

		if cfg.MaintenanceMode == "true" {
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"error": gin.H{
					"status":  http.StatusServiceUnavailable,
					"message": domain.ErrMaintenance.Error(),
				},
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
