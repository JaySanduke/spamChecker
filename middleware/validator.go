package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ValidateJSON(obj interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := c.ShouldBindJSON(obj); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":  "Invalid request format",
				"detail": err.Error(),
			})
			c.Abort()
			return
		}

		c.Set("validatedBody", obj)
		c.Next()
	}
}
