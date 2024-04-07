package main

import (
	"flag"

	"github.com/Real-Dev-Squad/wisee-backend/src/config"
	"github.com/Real-Dev-Squad/wisee-backend/src/routes"
	"github.com/Real-Dev-Squad/wisee-backend/src/utils"
)

func main() {
	dsn := config.DbUrl
	_, bunDbInstance := utils.SetupDBConnection(dsn)

	port := flag.String("port", ":8080", "server address to listen on")
	flag.Parse()

	routes.Listen("0.0.0.0"+*port, bunDbInstance)
}
