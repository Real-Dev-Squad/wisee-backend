package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

func SetupV1Routes(db *bun.DB) *gin.Engine {
	var router = gin.Default()

	// TODO: Configure CORS properly to allow only access from certain origins
	router.Use(cors.Default())

	v1 := router.Group("wisee/v1/")
	UserRoutes(v1, db)
	AuthRoutes(v1, db)
	FormRoutes(v1, db)
	HealthRoutes(v1, db)

	return router
}

func Listen(listenAddress string, db *bun.DB) {
	router := SetupV1Routes(db)

	router.Run(listenAddress)
}
