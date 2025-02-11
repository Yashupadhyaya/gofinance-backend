package util

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)








/*
ROOST_METHOD_HASH=ValidateToken_7440899dfa
ROOST_METHOD_SIG_HASH=ValidateToken_ff3cc8ee02

FUNCTION_DEF=func ValidateToken(ctx *gin.Context, token string) error 

*/
func TestValidateToken(t *testing.T) {
	gin.SetMode(gin.TestMode)

	type testCase struct {
		name           string
		token          string
		expectedStatus int
	}

	createToken := func(secretKey string, claims jwt.Claims) string {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, _ := token.SignedString([]byte(secretKey))
		return tokenString
	}

	tests := []testCase{
		{
			name: "Valid Token with Correct Signature",
			token: createToken("secret_key", jwt.MapClaims{
				"username": "testuser",
				"exp":      time.Now().Add(time.Hour).Unix(),
			}),
			expectedStatus: http.StatusOK,
		},
		{
			name: "Invalid Token Signature",
			token: createToken("wrong_key", jwt.MapClaims{
				"username": "testuser",
				"exp":      time.Now().Add(time.Hour).Unix(),
			}),
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name: "Expired Token",
			token: createToken("secret_key", jwt.MapClaims{
				"username": "testuser",
				"exp":      time.Now().Add(-time.Hour).Unix(),
			}),
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name:           "Malformed Token",
			token:          "malformed.token.string",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "Token with Invalid Claims",
			token: createToken("secret_key", jwt.MapClaims{
				"username": "testuser",
				"aud":      "wrong_audience",
			}),
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name:           "Missing Token",
			token:          "",
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			c.Request = httptest.NewRequest(http.MethodGet, "/", nil)
			if tc.token != "" {
				c.Request.Header.Set("Authorization", "Bearer "+tc.token)
			}

			var buf bytes.Buffer
			gin.DefaultWriter = &buf

			err := ValidateToken(c, tc.token)

			if w.Code != tc.expectedStatus {
				t.Errorf("Expected status %d, but got %d", tc.expectedStatus, w.Code)
			}

			if err != nil {
				t.Logf("Test %s failed with error: %v", tc.name, err)
			} else {
				t.Logf("Test %s succeeded", tc.name)
			}

			buf.Reset()
		})
	}
}

