package util

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"bytes"
	"time"
	"github.com/stretchr/testify/assert"
)








/*
ROOST_METHOD_HASH=GetTokenInHeaderAndVerify_c6fc249681
ROOST_METHOD_SIG_HASH=GetTokenInHeaderAndVerify_4459fbc010

FUNCTION_DEF=func GetTokenInHeaderAndVerify(ctx *gin.Context) error 

*/
func TestGetTokenInHeaderAndVerify(t *testing.T) {

	gin.SetMode(gin.TestMode)

	type testCase struct {
		description    string
		authHeader     string
		expectedError  bool
		expectedStatus int
	}

	validToken := createValidToken()
	invalidToken := "invalid.token.parts"

	testCases := []testCase{
		{
			description:    "Valid Token in Authorization Header",
			authHeader:     "Bearer " + validToken,
			expectedError:  false,
			expectedStatus: http.StatusOK,
		},
		{
			description:    "Missing Authorization Header",
			authHeader:     "",
			expectedError:  true,
			expectedStatus: http.StatusBadRequest,
		},
		{
			description:    "Invalid Token Format in Authorization Header",
			authHeader:     "InvalidFormat",
			expectedError:  true,
			expectedStatus: http.StatusBadRequest,
		},
		{
			description:    "Invalid Token in Authorization Header",
			authHeader:     "Bearer " + invalidToken,
			expectedError:  true,
			expectedStatus: http.StatusUnauthorized,
		},
		{
			description:    "Authorization Header Without Token",
			authHeader:     "Bearer ",
			expectedError:  true,
			expectedStatus: http.StatusBadRequest,
		},
		{
			description:    "Non-'Bearer' Authorization Header",
			authHeader:     "Basic " + validToken,
			expectedError:  true,
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = &http.Request{
				Header: http.Header{
					"Authorization": []string{tc.authHeader},
				},
			}

			err := GetTokenInHeaderAndVerify(c)

			if tc.expectedError && err == nil {
				t.Errorf("Expected an error but got none. %s", tc.description)
			} else if !tc.expectedError && err != nil {
				t.Errorf("Did not expect an error but got one: %v. %s", err, tc.description)
			}

			if w.Code != tc.expectedStatus {
				t.Errorf("Expected status %d but got %d. %s", tc.expectedStatus, w.Code, tc.description)
			}

			t.Logf("Test '%s' completed with status: %d", tc.description, w.Code)
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
		name           string
		tokenGenerator func() string
		expectedStatus int
		expectedError  error
		nextCalled     bool
	}{
		{
			name: "Valid Token",
			tokenGenerator: func() string {
				claims := Claims{
					Username: "testuser",
					RegisteredClaims: jwt.RegisteredClaims{
						ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
					},
				}
				token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
				tokenString, _ := token.SignedString([]byte("secret_key"))
				return tokenString
			},
			expectedStatus: http.StatusOK,
			expectedError:  nil,
			nextCalled:     true,
		},
		{
			name: "Invalid Signature",
			tokenGenerator: func() string {
				claims := Claims{
					Username: "testuser",
					RegisteredClaims: jwt.RegisteredClaims{
						ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
					},
				}
				token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
				tokenString, _ := token.SignedString([]byte("wrong_key"))
				return tokenString
			},
			expectedStatus: http.StatusUnauthorized,
			expectedError:  jwt.ErrSignatureInvalid,
			nextCalled:     false,
		},
		{
			name: "Malformed Token",
			tokenGenerator: func() string {
				return "malformed.token.string"
			},
			expectedStatus: http.StatusBadRequest,
			expectedError:  jwt.NewValidationError("token contains an invalid number of segments", jwt.ValidationErrorMalformed),
			nextCalled:     false,
		},
		{
			name: "Expired Token",
			tokenGenerator: func() string {
				claims := Claims{
					Username: "testuser",
					RegisteredClaims: jwt.RegisteredClaims{
						ExpiresAt: jwt.NewNumericDate(time.Now().Add(-time.Hour)),
					},
				}
				token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
				tokenString, _ := token.SignedString([]byte("secret_key"))
				return tokenString
			},
			expectedStatus: http.StatusUnauthorized,
			expectedError:  nil,
			nextCalled:     false,
		},
		{
			name: "Token with Invalid Claims",
			tokenGenerator: func() string {
				type InvalidClaims struct {
					InvalidField string `json:"invalid"`
					jwt.RegisteredClaims
				}

				claims := InvalidClaims{
					InvalidField: "invalid",
					RegisteredClaims: jwt.RegisteredClaims{
						ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
					},
				}
				token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
				tokenString, _ := token.SignedString([]byte("secret_key"))
				return tokenString
			},
			expectedStatus: http.StatusUnauthorized,
			expectedError:  nil,
			nextCalled:     false,
		},
		{
			name: "Token Parsing Error",
			tokenGenerator: func() string {
				return "header.payload.signature"
			},
			expectedStatus: http.StatusBadRequest,
			expectedError:  jwt.NewValidationError("token contains an invalid number of segments", jwt.ValidationErrorMalformed),
			nextCalled:     false,
		},
		{
			name: "Missing Token",
			tokenGenerator: func() string {
				return ""
			},
			expectedStatus: http.StatusBadRequest,
			expectedError:  jwt.NewValidationError("token contains an invalid number of segments", jwt.ValidationErrorMalformed),
			nextCalled:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)

			token := tt.tokenGenerator()
			ctx.Request, _ = http.NewRequest(http.MethodGet, "/", nil)
			ctx.Request.Header.Set("Authorization", "Bearer "+token)

			nextCalled := false
			ctx.Next = func() {
				nextCalled = true
			}

			err := ValidateToken(ctx, token)

			assert.Equal(t, tt.expectedStatus, w.Code)
			if tt.expectedError != nil {
				assert.Equal(t, tt.expectedError.Error(), err.Error())
			} else {
				assert.Nil(t, err)
			}
			assert.Equal(t, tt.nextCalled, nextCalled)

			t.Logf("Test %s: Status=%d, NextCalled=%v", tt.name, w.Code, nextCalled)
		})
	}
}

