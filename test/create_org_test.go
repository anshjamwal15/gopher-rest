package test

import (
	"fmt"
	"gopher-rest/pkg/routes"
	"gopher-rest/pkg/utils"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestCreateOrg(t *testing.T) {

	app := fiber.New()
	routes.PrivateRoutes(app)

	token, tokenErr := utils.GenerateNewAccessToken()

	if tokenErr != nil {
		t.Fatal(tokenErr)
	}

	jsonStr := `{"Name": "orgname", "Created_By": 1}`

	req := httptest.NewRequest("POST", "/api/v1/create", strings.NewReader(jsonStr))

	req.Header.Set("Authorization", token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)

	fmt.Printf("Type is %T\n", resp.Body)

	assert.Equalf(t, false, err != nil, "Create Org")

	assert.Equalf(t, 200, resp.StatusCode, "Create Org")

}
