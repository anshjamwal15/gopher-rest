package test

import (
	"gopher-rest/pkg/routes"
	"io"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestPublicRoutes(t *testing.T) {

	expectedData := `{"username": "testuser", "password": "test1234"}`
	unexpectedData := `{"username": "testuser", "password": "123"}`

	tests := []struct {
		description   string
		route         string
		method        string
		body          io.Reader
		expectedError bool
		expectedCode  int
		testType      string
	}{
		{
			description:   "User Registration",
			route:         "/api/v1/user",
			method:        "POST",
			body:          strings.NewReader(expectedData),
			expectedError: false,
			expectedCode:  200,
			testType:      "Success",
		},
		{
			description:   "User registration with invalid password.",
			route:         "/api/v1/user",
			method:        "POST",
			body:          strings.NewReader(unexpectedData),
			expectedError: false,
			expectedCode:  400,
			testType:      "Failure",
		},
		{
			description:   "User Login",
			route:         "/api/v1/login",
			method:        "POST",
			body:          strings.NewReader(expectedData),
			expectedError: false,
			expectedCode:  200,
			testType:      "Success",
		},
		{
			description:   "User Login with invalid credentials.",
			route:         "/api/v1/login",
			method:        "POST",
			body:          strings.NewReader(unexpectedData),
			expectedError: false,
			expectedCode:  400,
			testType:      "Failure",
		},
	}

	app := fiber.New()

	routes.PublicRoutes(app)

	for _, test := range tests {

		req := httptest.NewRequest(test.method, test.route, test.body)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req, -1)

		assert.Equalf(t, test.expectedError, err != nil, test.description)

		if test.expectedError {
			continue
		}

		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}

}
