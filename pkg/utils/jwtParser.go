package utils

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

type TokenMetadata struct {
	Expires int64
}

func ExtractTokenMetaData(c *fiber.Ctx) (*TokenMetadata, error) {
	token, err := verifyToken(c)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {

		expires := int64(claims["exp"].(float64))

		return &TokenMetadata{
			Expires: expires,
		}, nil
	}
	return nil, err
}

func extractToken(c *fiber.Ctx) string {
	bearToken := c.Get("Authorization")

	onlyToken := strings.Split(bearToken, " ")

	if len(onlyToken) == 2 {
		return onlyToken[1]
	}
	return ""
}

func verifyToken(c *fiber.Ctx) (*jwt.Token, error) {
	tokenString := extractToken(c)

	token, err := jwt.Parse(tokenString, jwtKeyFunc)

	if err != nil {
		return nil, err
	}
	return token, nil
}

func jwtKeyFunc(token *jwt.Token) (interface{}, error) {
	return []byte(os.Getenv("jwt_secret")), nil
}
