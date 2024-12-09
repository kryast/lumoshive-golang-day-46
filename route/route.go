package route

import (
	"day-46/controller"
	"day-46/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/register", controller.RegisterUser)
	router.POST("/login", middleware.RateLimiter(), controller.LoginUser)

	protected := router.Group("/products", middleware.JWTMiddleware(), middleware.IPWhitelist())
	protected.GET("/:id", controller.GetProductByID)
}
