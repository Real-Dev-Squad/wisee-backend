package routes

import (
	"net/http"

	"github.com/Real-Dev-Squad/wisee-backend/models"
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

func userRoutes(rg *gin.RouterGroup, db *bun.DB) {
	users := rg.Group("/users")

	users.GET("", func(ctx *gin.Context) {

		var users []models.User
		err := db.NewSelect().Model(&users).OrderExpr("id ASC").Limit(10).Scan(ctx)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "error",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": users,
		})
	})
}
