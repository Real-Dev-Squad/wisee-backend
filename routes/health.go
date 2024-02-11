package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

func HealthRoute(ctx *gin.Context, db *bun.DB) {
	err := db.PingContext(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "unhealthy, database connection failed",
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "healthy, systems functional",
	})
}
