package util_test

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"net/http/httptest"
	"testing"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"util"
)

/*
ROOST_METHOD_HASH=ValidateToken_7440899dfa
ROOST_METHOD_SIG_HASH=ValidateToken_ff3cc8ee02


 */
func Testvalidatetoken(t *testing.T) {
	testCases := []struct {
		name		string
		token		string
		wantErr		bool
		httpCode	int
	}{{name: "valid token", token: "your_valid_token", wantErr: false, httpCode: http.StatusOK}, {name: "invalid token", token: "your_invalid_token", wantErr: true, httpCode: http.StatusUnauthorized}, {name: "empty token", token: "", wantErr: true, httpCode: http.StatusBadRequest}, {name: "malformed token", token: "your_malformed_token", wantErr: true, httpCode: http.StatusBadRequest}}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			err := ValidateToken(c, tc.token)
			if (err != nil) != tc.wantErr {
				t.Errorf("ValidateToken() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if w.Code != tc.httpCode {
				t.Errorf("HTTP status code = %v, want %v", w.Code, tc.httpCode)
			}
		})
	}
}

func ValidateToken(ctx *gin.Context, token string) error {
	claims := &Claims{}
	var jwtSignedKey = []byte("secret_key")
	tokenParse, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtSignedKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			ctx.JSON(http.StatusUnauthorized, err)
			return err
		}
		ctx.JSON(http.StatusBadRequest, err)
		return err
	}
	if !tokenParse.Valid {
		ctx.JSON(http.StatusUnauthorized, "Token is invalid")
		return nil
	}
	ctx.Next()
	return nil
}

func CreateTestContext(w http.ResponseWriter) (c *Context, r *Engine) {
	r = New()
	c = r.allocateContext(0)
	c.reset()
	c.writermem.reset(w)
	return
}

/*
ROOST_METHOD_HASH=GetTokenInHeaderAndVerify_c6fc249681
ROOST_METHOD_SIG_HASH=GetTokenInHeaderAndVerify_4459fbc010


 */
func TestgetTokenInHeaderAndVerify(t *testing.T) {
	testCases := []struct {
		name			string
		token			string
		mockValidateTokenErr	error
		expectedErr		error
	}{{name: "Valid Token", token: "Bearer validToken", mockValidateTokenErr: nil, expectedErr: nil}, {name: "Invalid Token", token: "Bearer invalidToken", mockValidateTokenErr: errors.New("invalid token"), expectedErr: errors.New("invalid token")}, {name: "No Token in Header", token: "", mockValidateTokenErr: errors.New("no token in header"), expectedErr: errors.New("no token in header")}}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockValidateToken := new(MockValidateToken)
			mockValidateToken.On("ValidateToken", mock.AnythingOfType("*gin.context"), mock.AnythingOfType("string")).Return(tc.mockValidateTokenErr)
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request, _ = http.NewRequest(http.MethodGet, "/", nil)
			c.Request.Header.Set("Authorization", tc.token)
			err := util.GetTokenInHeaderAndVerify(c)
			assert.Equal(t, tc.expectedErr, err)
			mockValidateToken.AssertExpectations(t)
		})
	}
}

func (m *MockValidateToken) ValidateToken(ctx *gin.Context, token string) error {
	args := m.Called(ctx, token)
	return args.Error(0)
}

func (a *Assertions) Equal(expected interface{}, actual interface{}, msgAndArgs ...interface{}) bool {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	return Equal(a.t, expected, actual, msgAndArgs...)
}

func CreateTestContext(w http.ResponseWriter) (c *Context, r *Engine) {
	r = New()
	c = r.allocateContext(0)
	c.reset()
	c.writermem.reset(w)
	return
}

func AnythingOfType(t string) AnythingOfTypeArgument {
	return AnythingOfTypeArgument(t)
}

