package utils

import (
	"os"
	"strconv"
	"time"

	"github.com/Real-Dev-Squad/wisee-backend/models"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(user *models.User) (string, error) {
	issuer := os.Getenv("JWT_ISSUER")
	key := []byte(os.Getenv("JWT_SECRET"))

	tokenValidityInHours, err := strconv.ParseInt(os.Getenv("JWT_VALIDITY_IN_HOURS"), 10, 8)

	if err != nil {
		return "", err
	}

	tokenExpiryTime := time.Now().Add(time.Hour * time.Duration(tokenValidityInHours)).UTC()

	t := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"iss":   issuer,
		"exp":   tokenExpiryTime,
		"email": user.Email,
	})

	token, error := t.SignedString(key)

	return token, error
}
