package test

import (
	"gopher-rest/pkg/routes"
	"gopher-rest/pkg/utils"
	"io"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestOrgAddUser(t *testing.T) {

	token, tokenErr := utils.GenerateNewAccessToken()

	if tokenErr != nil {
		t.Fatal(tokenErr)
	}

	expectedData := `{"Username": "orgtestuser", "Password": "orgtest123", "OrgId": 3, "AdminId": 1}`
	unexpectedData := `{"Username": "orgtestuser", "Password": "orgtest123", "OrgId": 3, "AdminId": 2}`

	tests := []struct {
		description   string
		route         string
		method        string
		token         string
		body          io.Reader
		expectedError bool
		expectedCode  int
		testType      string
	}{
		{
			description:   "Add User in org",
			route:         "/api/v1/add",
			method:        "POST",
			body:          strings.NewReader(expectedData),
			expectedError: false,
			expectedCode:  200,
			testType:      "Success",
		},
		{
			description:   "Add User in org with wrong admin.",
			route:         "/api/v1/add",
			method:        "POST",
			body:          strings.NewReader(unexpectedData),
			expectedError: false,
			expectedCode:  400,
			testType:      "Failure",
		},
	}

	app := fiber.New()
	routes.PrivateRoutes(app)

	for _, test := range tests {

		req := httptest.NewRequest(test.method, test.route, test.body)

		req.Header.Set("Authorization", token)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req, -1)

		assert.Equalf(t, false, err != nil, test.description)

		if test.expectedError {
			continue
		}

		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}

}
