package util

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"bytes"
	"time"
	"github.com/golang-jwt/jwt/v4"
)




var mockValidateToken = func(ctx *gin.Context, token string)




/*
ROOST_METHOD_HASH=GetTokenInHeaderAndVerify_c6fc249681
ROOST_METHOD_SIG_HASH=GetTokenInHeaderAndVerify_4459fbc010

FUNCTION_DEF=func GetTokenInHeaderAndVerify(ctx *gin.Context) error 

*/
func TestGetTokenInHeaderAndVerify(t *testing.T) {
	type testCase struct {
		name          string
		headerValue   string
		validateToken func(ctx *gin.Context, token string) error
		expectedError error
	}

	tests := []testCase{
		{
			name:        "Valid Token in Header",
			headerValue: "Bearer validToken",
			validateToken: func(ctx *gin.Context, token string) error {

				return nil
			},
			expectedError: nil,
		},
		{
			name:          "Missing Authorization Header",
			headerValue:   "",
			validateToken: mockValidateToken,
			expectedError: errors.New("missing authorization header"),
		},
		{
			name:          "Invalid Token Format in Header",
			headerValue:   "InvalidFormatToken",
			validateToken: mockValidateToken,
			expectedError: errors.New("invalid authorization header format"),
		},
		{
			name:        "Token Validation Failure",
			headerValue: "Bearer invalidToken",
			validateToken: func(ctx *gin.Context, token string) error {

				return errors.New("token validation failed")
			},
			expectedError: errors.New("token validation failed"),
		},
		{
			name:          "Empty Token in Header",
			headerValue:   "Bearer ",
			validateToken: mockValidateToken,
			expectedError: errors.New("empty token"),
		},
	}

	originalValidateToken := ValidateToken
	defer func() { ValidateToken = originalValidateToken }()
	ValidateToken = mockValidateToken

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			gin.SetMode(gin.TestMode)
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			req.Header.Set("Authorization", tc.headerValue)
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = req

			err := GetTokenInHeaderAndVerify(ctx)

			if tc.expectedError == nil {
				assert.NoError(t, err)
				t.Logf("Scenario '%s' passed: Expected no error, got no error", tc.name)
			} else {
				assert.EqualError(t, err, tc.expectedError.Error())
				t.Logf("Scenario '%s' passed: Expected error '%s', got error '%s'", tc.name, tc.expectedError.Error(), err.Error())
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
	type testCase struct {
		name       string
		token      string
		wantStatus int
		wantErr    error
	}

	tests := []testCase{
		{
			name:       "Valid Token",
			token:      generateValidToken(),
			wantStatus: http.StatusOK,
			wantErr:    nil,
		},
		{
			name:       "Invalid Signature",
			token:      generateInvalidSignatureToken(),
			wantStatus: http.StatusUnauthorized,
			wantErr:    jwt.ErrSignatureInvalid,
		},
		{
			name:       "Malformed Token",
			token:      "malformed.token.string",
			wantStatus: http.StatusBadRequest,
			wantErr:    jwt.NewValidationError("token contains an invalid number of segments", jwt.ValidationErrorMalformed),
		},
		{
			name:       "Expired Token",
			token:      generateExpiredToken(),
			wantStatus: http.StatusUnauthorized,
			wantErr:    jwt.NewValidationError("token is expired", jwt.ValidationErrorExpired),
		},
		{
			name:       "Missing Token",
			token:      "",
			wantStatus: http.StatusBadRequest,
			wantErr:    jwt.NewValidationError("token is missing", jwt.ValidationErrorMalformed),
		},
		{
			name:       "Token with Missing Claims",
			token:      generateTokenWithMissingClaims(),
			wantStatus: http.StatusUnauthorized,
			wantErr:    jwt.NewValidationError("token claims are invalid", jwt.ValidationErrorClaimsInvalid),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			err := ValidateToken(c, tc.token)

			if w.Code != tc.wantStatus {
				t.Errorf("got status %d, want %d", w.Code, tc.wantStatus)
			}

			if err != nil && !strings.Contains(err.Error(), tc.wantErr.Error()) {
				t.Errorf("got error %v, want %v", err, tc.wantErr)
			}
		})
	}
}

