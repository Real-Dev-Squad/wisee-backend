package main

import (
	"flag"

	"github.com/Real-Dev-Squad/wisee-backend/routes"
)

func main() {
	port := flag.String("port", ":8080", "server address to listen on")
	flag.Parse()

	routes.Listen(*port)
}
