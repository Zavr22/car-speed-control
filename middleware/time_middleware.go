package middleware

import (
	"github.com/Zavr22/car-speed-control/config"
	"github.com/Zavr22/car-speed-control/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AccessTimeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cfg, err := config.LoadConfig()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load config"})
			c.Abort()
			return
		}

		startTime, err := cfg.GetStartTime()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid start time format"})
			c.Abort()
			return
		}

		endTime, err := cfg.GetEndTime()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid end time format"})
			c.Abort()
			return
		}

		if !utils.IsWithinAccessHours(startTime, endTime) {
			c.JSON(http.StatusForbidden, gin.H{"error": "access restricted to specific hours"})
			c.Abort()
			return
		}

		c.Next()
	}
}
