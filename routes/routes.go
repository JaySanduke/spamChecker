package routes

import (
	"spamChecker/controllers"
	"spamChecker/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	r.GET("/heathz", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	api := r.Group("/api")
	{
		user := api.Group("/user")
		user.POST("/register", middleware.ValidateJSON(&controllers.RegisterInput{}), controllers.Register)
		user.POST("/login", controllers.Login)

		private := api.Group("/")
		private.Use(middleware.JWTAuthMiddleware())
		{
			private.POST("/spam/mark", controllers.MarkSpam)
			private.GET("/user/search", controllers.Search)
			private.GET("/user/profile/:phone", controllers.GetProfile)
		}
	}
}
