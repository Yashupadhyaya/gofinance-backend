package util

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"net/http/httptest"
	"testing"
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

