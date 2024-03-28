package routes

import (
	"net/http"

	"github.com/Real-Dev-Squad/wisee-backend/src/utils/logger"
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

func HealthRoutes(rg *gin.RouterGroup, db *bun.DB) {
	healthCheck := rg.Group("/health")

	healthCheck.GET("", func(ctx *gin.Context) {
		err := db.Ping()

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "error",
			})

			logger.Error(err)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})
}
