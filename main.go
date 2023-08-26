package main

import (
	"flag"

	"github.com/Real-Dev-Squad/wisee-backend/routes"
)

func main() {
	address := flag.String("address", ":8080", "server address to listen on")
	flag.Parse()

	routes.Listen(*address)
}
