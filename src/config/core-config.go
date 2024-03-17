package config

import (
	"os"
	"strconv"

	"github.com/Real-Dev-Squad/wisee-backend/src/utils/logger"
	"github.com/joho/godotenv"
)

var Env string

var JwtSecret string
var JwtValidityInHours int
var JwtIssuer string

var Domain string
var AuthRedirectUrl string

var DbUrl string
var TestDbUrl string
var DbMaxOpenConnections int

var GoogleClientId string
var GoogleClientSecret string
var GoogleRedirectUrl string

func loadEnv() {
	env := os.Getenv("ENV")

	if env != "production" {
		if err := godotenv.Load(".env"); err != nil {
			logger.Fatal("Error loading .env file")
		}
	}
}

func init() {
	loadEnv()

	env := os.Getenv("ENV")

	if env == "" {
		Env = "dev"
	} else {
		Env = env
	}

	JwtSecret = os.Getenv("JWT_SECRET")
	JwtValidityInHours, _ = strconv.Atoi(os.Getenv("JWT_VALIDITY_IN_HOURS"))
	JwtIssuer = os.Getenv("JWT_ISSUER")

	Domain = os.Getenv("DOMAIN")
	AuthRedirectUrl = os.Getenv("AUTH_REDIRECT_URL")

	DbUrl = os.Getenv("DB_URL")
	TestDbUrl = os.Getenv("TEST_DB_URL")
	DbMaxOpenConnections, _ = strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONNECTIONS"))

	GoogleClientId = os.Getenv("GOOGLE_CLIENT_ID")
	GoogleClientSecret = os.Getenv("GOOGLE_CLIENT_SECRET")
	GoogleRedirectUrl = os.Getenv("GOOGLE_REDIRECT_URL")

	logger.Info("Loaded environment variables")
}
