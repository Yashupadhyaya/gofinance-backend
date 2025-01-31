
// ********RoostGPT********
/*
Test generated by RoostGPT for test improve-test-golang using AI Type Open AI and AI Model gpt-4o

ROOST_METHOD_HASH=login_b45c9ba5a0
ROOST_METHOD_SIG_HASH=login_5743050a86

FUNCTION_DEF=func (server *Server) login(ctx *gin.Context)
Below are the test scenarios for the `login` function in the Go application:

### Scenario 1: Valid Login with Correct Credentials

**Details:**
- **Description:** This test checks that a valid login request with correct username and password results in a successful authentication and returns a JWT token.
- **Execution:**
  - **Arrange:** Mock the `GetUser` method to return a user with a known password hash. Use the bcrypt library to hash a known password and store it in the mock user data.
  - **Act:** Call the `login` function with the correct username and password.
  - **Assert:** Verify that the response status is `http.StatusOK` and a valid JWT token is returned.
- **Validation:**
  - The assertion checks if the token is present and valid, ensuring the login logic is correctly allowing access with valid credentials.
  - This test is crucial for verifying that users can successfully log in and receive a token for accessing protected resources.

### Scenario 2: Invalid Login with Incorrect Password

**Details:**
- **Description:** This test ensures that an incorrect password results in an unauthorized error, preventing access.
- **Execution:**
  - **Arrange:** Mock the `GetUser` method to return a user with a known password hash. Use a different password for the test input.
  - **Act:** Call the `login` function with the correct username but incorrect password.
  - **Assert:** Check that the response status is `http.StatusUnauthorized`.
- **Validation:**
  - The assertion confirms that incorrect credentials do not allow access, maintaining security.
  - This test is important to ensure that the application does not grant access with an incorrect password.

### Scenario 3: User Not Found

**Details:**
- **Description:** This test checks the behavior when a login attempt is made with a username that does not exist in the database.
- **Execution:**
  - **Arrange:** Mock the `GetUser` method to return `sql.ErrNoRows` for a non-existent username.
  - **Act:** Call the `login` function with a non-existent username.
  - **Assert:** Verify that the response status is `http.StatusNotFound`.
- **Validation:**
  - The assertion ensures the application responds appropriately when the user is not found, providing clarity to the user.
  - This test is significant for user experience, indicating when a username does not exist.

### Scenario 4: Malformed JSON Request

**Details:**
- **Description:** This test ensures that a malformed JSON request body results in a bad request error.
- **Execution:**
  - **Arrange:** Prepare a request with a malformed JSON body.
  - **Act:** Call the `login` function with the malformed JSON.
  - **Assert:** Check that the response status is `http.StatusBadRequest`.
- **Validation:**
  - The assertion checks that the function properly handles malformed input, maintaining application robustness.
  - This test is crucial for input validation and error handling, ensuring the application can gracefully handle invalid input formats.

### Scenario 5: Database Error on User Retrieval

**Details:**
- **Description:** This test checks the behavior when there is a database error while retrieving the user.
- **Execution:**
  - **Arrange:** Mock the `GetUser` method to return a generic SQL error.
  - **Act:** Call the `login` function with a valid username.
  - **Assert:** Verify that the response status is `http.StatusInternalServerError`.
- **Validation:**
  - The assertion ensures that the application can gracefully handle database errors without crashing.
  - This test is important for ensuring system reliability and proper error reporting to the user or logs.

### Scenario 6: JWT Token Generation Error

**Details:**
- **Description:** This test verifies that an error during JWT token generation results in an internal server error.
- **Execution:**
  - **Arrange:** Mock the JWT signing method to return an error.
  - **Act:** Call the `login` function with valid credentials.
  - **Assert:** Check that the response status is `http.StatusInternalServerError`.
- **Validation:**
  - The assertion ensures that the application handles unexpected errors during token generation, maintaining stability.
  - This test is important for error handling in token generation, ensuring users are informed of issues without exposing sensitive information.

These scenarios collectively cover the critical paths, edge cases, and error handling for the `login` function, ensuring comprehensive testing of its functionality.

roost_feedback [1/31/2025, 10:30:21 AM]:add some negative scenarios as well
*/

// ********RoostGPT********

package api

import (
	"bytes"
	"crypto/sha512"
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type mockStore struct {
	user sqlmock.Sqlmock
}

func (ms *mockStore) GetUser(ctx *gin.Context, username string) (User, error) {
	rows := sqlmock.NewRows([]string{"id", "username", "password", "email", "created_at"})
	if username == "validUser" {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("validPassword"), bcrypt.DefaultCost)
		rows.AddRow(1, "validUser", string(hashedPassword), "user@example.com", time.Now())
		ms.user.ExpectQuery("^SELECT (.+) FROM users WHERE username=?$").WithArgs(username).WillReturnRows(rows)
		return User{ID: 1, Username: "validUser", Password: string(hashedPassword), Email: "user@example.com", CreatedAt: time.Now()}, nil
	} else if username == "dbErrorUser" {
		ms.user.ExpectQuery("^SELECT (.+) FROM users WHERE username=?$").WithArgs(username).WillReturnError(sql.ErrConnDone)
		return User{}, sql.ErrConnDone
	} else {
		ms.user.ExpectQuery("^SELECT (.+) FROM users WHERE username=?$").WithArgs(username).WillReturnError(sql.ErrNoRows)
		return User{}, sql.ErrNoRows
	}
}

type Server struct {
	store  *mockStore
	router *gin.Engine
}

func (server *Server) login(ctx *gin.Context) {
	var req loginRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUser(ctx, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	hashedInput := sha512.Sum512_256([]byte(req.Password))
	trimmedHash := bytes.Trim(hashedInput[:], "\x00")
	preparedPassword := string(trimmedHash)
	plainTextInBytes := []byte(preparedPassword)
	hashTextInBytes := []byte(user.Password)
	err = bcrypt.CompareHashAndPassword(hashTextInBytes, plainTextInBytes)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	expirationTime := time.Now().Add(100 * time.Minute)

	claims := &Claims{
		Username: req.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	generatedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	var jwtSignedKey = []byte("secret_key")
	generatedTokenToString, err := generatedToken.SignedString(jwtSignedKey)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, generatedTokenToString)
}

func TestServerLogin(t *testing.T) {
	gin.SetMode(gin.TestMode)

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	store := &mockStore{user: mock}
	server := &Server{store: store}

	tests := []struct {
		name           string
		requestBody    string
		expectedStatus int
		expectedToken  bool
	}{
		{
			name:           "Valid Login with Correct Credentials",
			requestBody:    `{"username":"validUser","password":"validPassword"}`,
			expectedStatus: http.StatusOK,
			expectedToken:  true,
		},
		{
			name:           "Invalid Login with Incorrect Password",
			requestBody:    `{"username":"validUser","password":"wrongPassword"}`,
			expectedStatus: http.StatusUnauthorized,
			expectedToken:  false,
		},
		{
			name:           "User Not Found",
			requestBody:    `{"username":"nonExistentUser","password":"anyPassword"}`,
			expectedStatus: http.StatusNotFound,
			expectedToken:  false,
		},
		{
			name:           "Malformed JSON Request",
			requestBody:    `{"username":"malformedUser",`,
			expectedStatus: http.StatusBadRequest,
			expectedToken:  false,
		},
		{
			name:           "Database Error on User Retrieval",
			requestBody:    `{"username":"dbErrorUser","password":"anyPassword"}`,
			expectedStatus: http.StatusInternalServerError,
			expectedToken:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBufferString(tt.requestBody))
			rec := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(rec)
			ctx.Request = req

			server.login(ctx)

			if rec.Code != tt.expectedStatus {
				t.Errorf("expected status %v, got %v", tt.expectedStatus, rec.Code)
			}

			if tt.expectedToken {
				token := rec.Body.String()
				parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
					return []byte("secret_key"), nil
				})
				if err != nil || !parsedToken.Valid {
					t.Errorf("expected a valid token, but got an invalid one")
				}
			} else {
				if rec.Body.String() != "" {
					t.Errorf("expected no token, but got one")
				}
			}
		})
	}
}
