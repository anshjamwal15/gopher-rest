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

func TestCreateOrg(t *testing.T) {

	token, tokenErr := utils.GenerateNewAccessToken()

	if tokenErr != nil {
		t.Fatal(tokenErr)
	}

	expectedData := `{"Name": "testorg", "Created_By": 1}`
	unexpectedData := `{"Name": "testorg", "Created_By": 2}`

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
			description:   "Create Org",
			route:         "/api/v1/create",
			method:        "POST",
			body:          strings.NewReader(expectedData),
			expectedError: false,
			expectedCode:  200,
			testType:      "Success",
		},
		{
			description:   "Create Org with invalid user.",
			route:         "/api/v1/create",
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

		assert.Equalf(t, test.expectedError, err != nil, test.description)

		if test.expectedError {
			continue
		}

		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}

}
