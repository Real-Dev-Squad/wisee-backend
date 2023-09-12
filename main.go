package main

import (
	"database/sql"
	"flag"
	"log"
	"os"

	"github.com/Real-Dev-Squad/wisee-backend/routes"
	"github.com/joho/godotenv"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func setupDBConnection() *bun.DB {
	dsn := os.Getenv("DB_URL")
	pgDB := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	pgDB.SetMaxOpenConns(1)

	db := bun.NewDB(pgDB, pgdialect.New())

	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))

	return db
}

func main() {
	loadEnv()
	db := setupDBConnection()

	port := flag.String("port", ":8080", "server address to listen on")
	flag.Parse()

	routes.Listen("127.0.0.1"+*port, db)
}
