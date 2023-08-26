package routes

import "github.com/gin-gonic/gin"

var router = gin.Default()

func setupRoutes(){
	v1:= router.Group("v1/")
	userRoutes(v1)
}

func Listen(listenAddress string) {
	setupRoutes()
	router.Run(listenAddress) 
}
