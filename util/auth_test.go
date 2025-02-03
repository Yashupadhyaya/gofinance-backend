package util

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)



var mockValidateToken = func(ctx *gin.Context, token string) error {

	if token == "valid_token" {
		return nil
	} else if token == "invalid_token" {
		return jwt.NewValidationError("invalid token", jwt.ValidationErrorSignatureInvalid)
	} else if token == "error_token" {
		return jwt.NewValidationError("error during validation", jwt.ValidationErrorUnverifiable)
	}
	return jwt.NewValidationError("token is empty", jwt.ValidationErrorClaimsInvalid)
}

type MockClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}


/*
ROOST_METHOD_HASH=GetTokenInHeaderAndVerify_c6fc249681
ROOST_METHOD_SIG_HASH=GetTokenInHeaderAndVerify_4459fbc010

FUNCTION_DEF=func GetTokenInHeaderAndVerify(ctx *gin.Context) error 

*/
func TestGetTokenInHeaderAndVerify(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name          string
		headerValue   string
		expectedError error
		validateToken func(ctx *gin.Context, token string) error
	}{
		{
			name:          "Valid Token in Header",
			headerValue:   "Bearer valid_token",
			expectedError: nil,
			validateToken: mockValidateToken,
		},
		{
			name:          "Missing Authorization Header",
			headerValue:   "",
			expectedError: jwt.NewValidationError("token is empty", jwt.ValidationErrorClaimsInvalid),
			validateToken: mockValidateToken,
		},
		{
			name:          "Malformed Authorization Header",
			headerValue:   "Bearer",
			expectedError: jwt.NewValidationError("token is empty", jwt.ValidationErrorClaimsInvalid),
			validateToken: mockValidateToken,
		},
		{
			name:          "Invalid Token in Header",
			headerValue:   "Bearer invalid_token",
			expectedError: jwt.NewValidationError("invalid token", jwt.ValidationErrorSignatureInvalid),
			validateToken: mockValidateToken,
		},
		{
			name:          "Empty Token in Header",
			headerValue:   "Bearer ",
			expectedError: jwt.NewValidationError("token is empty", jwt.ValidationErrorClaimsInvalid),
			validateToken: mockValidateToken,
		},
		{
			name:          "Token Validation Function Error",
			headerValue:   "Bearer error_token",
			expectedError: jwt.NewValidationError("error during validation", jwt.ValidationErrorUnverifiable),
			validateToken: mockValidateToken,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			req := httptest.NewRequest(http.MethodGet, "/", nil)
			req.Header.Set("Authorization", tt.headerValue)
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = req

			err := GetTokenInHeaderAndVerify(ctx)

			if tt.expectedError != nil {
				if err == nil || !strings.Contains(err.Error(), tt.expectedError.Error()) {
					t.Errorf("expected error: %v, got: %v", tt.expectedError, err)
				} else {
					t.Logf("Successfully got expected error: %v", err)
				}
			} else {
				if err != nil {
					t.Errorf("expected no error, but got: %v", err)
				} else {
					t.Log("Successfully validated token without error")
				}
			}
		})
	}
}


/*
ROOST_METHOD_HASH=ValidateToken_7440899dfa
ROOST_METHOD_SIG_HASH=ValidateToken_ff3cc8ee02

FUNCTION_DEF=func ValidateToken(ctx *gin.Context, token string) error 

*/
func TestValidateToken(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name          string
		token         string
		expectedCode  int
		expectedError string
	}{
		{
			name: "Valid Token",
			token: func() string {
				claims := MockClaims{
					Username: "testuser",
					RegisteredClaims: jwt.RegisteredClaims{
						ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
					},
				}
				token, _ := createToken([]byte("secret_key"), claims)
				return token
			}(),
			expectedCode: http.StatusOK,
		},
		{
			name: "Invalid Signature",
			token: func() string {
				claims := MockClaims{
					Username: "testuser",
					RegisteredClaims: jwt.RegisteredClaims{
						ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
					},
				}
				token, _ := createToken([]byte("wrong_key"), claims)
				return token
			}(),
			expectedCode:  http.StatusUnauthorized,
			expectedError: "signature is invalid",
		},
		{
			name:          "Malformed Token",
			token:         "malformed.token",
			expectedCode:  http.StatusBadRequest,
			expectedError: "token contains an invalid number of segments",
		},
		{
			name: "Expired Token",
			token: func() string {
				claims := MockClaims{
					Username: "testuser",
					RegisteredClaims: jwt.RegisteredClaims{
						ExpiresAt: jwt.NewNumericDate(time.Now().Add(-time.Hour)),
					},
				}
				token, _ := createToken([]byte("secret_key"), claims)
				return token
			}(),
			expectedCode:  http.StatusUnauthorized,
			expectedError: "token is expired",
		},
		{
			name: "Token with Invalid Claims",
			token: func() string {
				claims := MockClaims{
					Username: "invaliduser",
					RegisteredClaims: jwt.RegisteredClaims{
						ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
					},
				}
				token, _ := createToken([]byte("secret_key"), claims)
				return token
			}(),
			expectedCode:  http.StatusUnauthorized,
			expectedError: "claims are invalid",
		},
		{
			name:          "Missing Token",
			token:         "",
			expectedCode:  http.StatusBadRequest,
			expectedError: "token is missing",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = &http.Request{
				Header: http.Header{
					"Authorization": []string{fmt.Sprintf("Bearer %s", tt.token)},
				},
			}

			err := ValidateToken(c, tt.token)

			if w.Code != tt.expectedCode {
				t.Errorf("expected status code %d, got %d", tt.expectedCode, w.Code)
			}

			if tt.expectedError != "" {
				var resp map[string]interface{}
				json.NewDecoder(w.Body).Decode(&resp)
				if errMsg, exists := resp["error"]; exists {
					if !strings.Contains(errMsg.(string), tt.expectedError) {
						t.Errorf("expected error message '%s', got '%s'", tt.expectedError, errMsg)
					}
				}
			}

			t.Logf("Test '%s' completed successfully", tt.name)
		})
	}
}

