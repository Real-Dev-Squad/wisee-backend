package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func userRoutes(rg *gin.RouterGroup){
	users :=  rg.Group("/users")

	users.GET("", func(c * gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"message": "users",
	})
	})
}
