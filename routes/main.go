package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

var router = gin.Default()

func setupV1Routes(db *bun.DB) {
	v1 := router.Group("v1/")
	userRoutes(v1, db)
}

func Listen(listenAddress string, db *bun.DB) {
	setupV1Routes(db)
	router.Run(listenAddress)
}
