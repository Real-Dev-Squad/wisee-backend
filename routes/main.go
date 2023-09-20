package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

func SetupV1Routes(db *bun.DB) *gin.Engine {
	var router = gin.Default()

	v1 := router.Group("v1/")
	userRoutes(v1, db)
	authRoutes(v1, db)

	return router
}

func Listen(listenAddress string, db *bun.DB) {
	router := SetupV1Routes(db)
	router.Run(listenAddress)
}
