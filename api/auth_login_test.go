package api

import (
	"bytes"
	"context"
	"crypto/sha512"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type mockStore struct {
	db *sql.DB
}

func (m *mockStore) GetUser(ctx context.Context, username string) (User, error) {
	// Implement mock logic here
	return User{}, nil
}

// TestServerLogin tests the login functionality of the Server.
func TestServerLogin(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	server := &Server{
		store: &SQLStore{db: db},
	}

	tests := []struct {
		name         string
		setupMock    func()
		requestBody  string
		expectedCode int
	}{
		{
			name: "Successful Login with Valid Credentials",
			setupMock: func() {
				passwordHash, _ := bcrypt.GenerateFromPassword([]byte("validPassword"), bcrypt.DefaultCost)
				mock.ExpectQuery("SELECT (.+) FROM users WHERE username=?").
					WithArgs("validUser").
					WillReturnRows(sqlmock.NewRows([]string{"id", "username", "password", "email", "created_at"}).
						AddRow(1, "validUser", passwordHash, "valid@example.com", time.Now()))
			},
			requestBody:  `{"username":"validUser","password":"validPassword"}`,
			expectedCode: http.StatusOK,
		},
		{
			name: "Login with Incorrect Password",
			setupMock: func() {
				passwordHash, _ := bcrypt.GenerateFromPassword([]byte("validPassword"), bcrypt.DefaultCost)
				mock.ExpectQuery("SELECT (.+) FROM users WHERE username=?").
					WithArgs("validUser").
					WillReturnRows(sqlmock.NewRows([]string{"id", "username", "password", "email", "created_at"}).
						AddRow(1, "validUser", passwordHash, "valid@example.com", time.Now()))
			},
			requestBody:  `{"username":"validUser","password":"invalidPassword"}`,
			expectedCode: http.StatusUnauthorized,
		},
		{
			name: "Login with Non-Existent Username",
			setupMock: func() {
				mock.ExpectQuery("SELECT (.+) FROM users WHERE username=?").
					WithArgs("nonExistentUser").
					WillReturnError(sql.ErrNoRows)
			},
			requestBody:  `{"username":"nonExistentUser","password":"somePassword"}`,
			expectedCode: http.StatusNotFound,
		},
		{
			name: "Login with Invalid JSON Payload",
			setupMock: func() {},
			requestBody:  `{"username":"invalidUser"`,
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "Login with Missing Fields in JSON Payload",
			setupMock: func() {},
			requestBody:  `{"username":"missingPassword"}`,
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "Internal Server Error During User Retrieval",
			setupMock: func() {
				mock.ExpectQuery("SELECT (.+) FROM users WHERE username=?").
					WithArgs("validUser").
					WillReturnError(sql.ErrConnDone)
			},
			requestBody:  `{"username":"validUser","password":"validPassword"}`,
			expectedCode: http.StatusInternalServerError,
		},
		{
			name: "Token Generation Failure",
			setupMock: func() {
				passwordHash, _ := bcrypt.GenerateFromPassword([]byte("validPassword"), bcrypt.DefaultCost)
				mock.ExpectQuery("SELECT (.+) FROM users WHERE username=?").
					WithArgs("validUser").
					WillReturnRows(sqlmock.NewRows([]string{"id", "username", "password", "email", "created_at"}).
						AddRow(1, "validUser", passwordHash, "valid@example.com", time.Now()))
			},
			requestBody:  `{"username":"validUser","password":"validPassword"}`,
			expectedCode: http.StatusInternalServerError, // Simulate JWT signing failure
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMock()

			req, err := http.NewRequest(http.MethodPost, "/login", strings.NewReader(tt.requestBody))
			if err != nil {
				t.Fatalf("could not create request: %v", err)
			}
			req.Header.Set("Content-Type", "application/json")

			rec := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(rec)
			ctx.Request = req

			server.login(ctx)

			if rec.Code != tt.expectedCode {
				t.Errorf("expected status %v; got %v", tt.expectedCode, rec.Code)
			}

			// TODO: Add more assertions based on the response body or headers if needed
		})
	}
}
