package utils

import (
	"errors"
	"time"

	"github.com/Real-Dev-Squad/wisee-backend/src/config"
	"github.com/Real-Dev-Squad/wisee-backend/src/models"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = config.JwtSecret
var jwtValidityInDays = config.JwtValidityInDays

/*
 * GenerateToken generates a JWT token for the user
 */
func GenerateToken(user *models.User) (string, error) {
	issuer := config.JwtIssuer
	key := []byte(jwtSecret)

	t := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"iss":   issuer,
		"email": user.Email,
		"iat":   jwt.NewNumericDate(time.Now()),
		"exp":   jwt.NewNumericDate(time.Now().AddDate(0, 0, jwtValidityInDays)),
	})

	token, error := t.SignedString(key)

	return token, error
}

/*
 * VerifyToken verifies the token and returns the email of the user
 */
func VerifyToken(tokenString string) (string, error) {
	token, tokenErr := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if token.Method.Alg() != jwt.SigningMethodHS512.Alg() {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(jwtSecret), nil
	})

	if tokenErr != nil || !token.Valid {
		return "", tokenErr
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("token claims are not in expected format")
	}

	exp, err := claims.GetExpirationTime()
	if err != nil {
		return "", err
	}

	if exp.Before(time.Now().UTC()) {
		return "", errors.New("token has expired")
	}

	email, ok := claims["email"].(string)
	if !ok {
		return "", errors.New("email not found in token")
	}

	return email, nil
}
