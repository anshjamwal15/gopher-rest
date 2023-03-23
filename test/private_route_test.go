package test

import (
	"gopher-rest/pkg/utils"
	"testing"

	"github.com/joho/godotenv"
)

func TestPrivateRoutes(t *testing.T) {

	if err := godotenv.Load("../.env"); err != nil {
		panic(err)
	}

	token, err := utils.GenerateNewAccessToken()
	if err != nil {
		panic(err)
		panic(token)
	}

}
