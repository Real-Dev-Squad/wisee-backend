package routes

import "github.com/gin-gonic/gin"

var router = gin.Default()

func setupV1Routes(){
	v1:= router.Group("v1/")
	userRoutes(v1)
}

func Listen(listenAddress string) {
	setupV1Routes()
	router.Run(listenAddress) 
}
