package utils

import (
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateNewAccessToken() (string, error) {

	secret := os.Getenv("jwt_secret")

	minutesCount, _ := strconv.Atoi(os.Getenv("JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT"))

	claims := jwt.MapClaims{}

	claims["exp"] = time.Now().Add(time.Minute * time.Duration(minutesCount)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", err
	}
	return t, nil
}
