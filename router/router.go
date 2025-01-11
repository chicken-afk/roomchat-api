package router

import (
	"goboilerplate/domains/auths"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Router(r *gin.Engine) *gin.Engine {
	/* Cors */
	allowedOrigins := "*"
	logrus.Warn("allow cors origin:", allowedOrigins)
	defaultConfig := cors.DefaultConfig()
	defaultConfig.AllowOrigins = []string{allowedOrigins}
	defaultConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	r.Use(cors.New(defaultConfig))

	// Prefix v1 for versioning
	routeV1 := r.Group("/api/v1")

	/* Router */
	routeV1.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": true,
			"message": "System is running",
		})
	})
	// Add routes from other modules
	auths.Router(routeV1)

	return r
}
