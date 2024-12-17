package util

import (
	"math/rand"
	"strings"
	"testing"
	"time"
)

/*
ROOST_METHOD_HASH=RandomEmail_1905439733
ROOST_METHOD_SIG_HASH=RandomEmail_7a04f189fd


 */
func TestRandomEmail(t *testing.T) {

	rand.Seed(time.Now().UnixNano())

	tests := []struct {
		name        string
		input       int
		expectedLen int
		expectError bool
	}{
		{
			name:        "Generating Email with 5 Characters",
			input:       5,
			expectedLen: 5 + len("@email.com"),
			expectError: false,
		},
		{
			name:        "Generating Email with 0 Characters",
			input:       0,
			expectedLen: len("@email.com"),
			expectError: false,
		},
		{
			name:        "Generating Email with a Large Number of Characters",
			input:       1000,
			expectedLen: 1000 + len("@email.com"),
			expectError: false,
		},
		{
			name:        "Generating Email with Special Characters in Local Part",
			input:       10,
			expectedLen: 10 + len("@email.com"),
			expectError: false,
		},
		{
			name:        "Consistency of Output Length",
			input:       15,
			expectedLen: 15 + len("@email.com"),
			expectError: false,
		},
		{
			name:        "Multiple Invocations Yield Different Results",
			input:       8,
			expectedLen: 8 + len("@email.com"),
			expectError: false,
		},
		{
			name:        "Handling Negative Number of Characters",
			input:       -5,
			expectedLen: len("@email.com"),
			expectError: true,
		},
		{
			name:        "Boundary Test with 1 Character",
			input:       1,
			expectedLen: 1 + len("@email.com"),
			expectError: false,
		},
		{
			name:        "Verifying Alphabetic Characters Only",
			input:       10,
			expectedLen: 10 + len("@email.com"),
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			email := RandomEmail(tt.input)

			if len(email) != tt.expectedLen {
				t.Errorf("expected length %d, got %d", tt.expectedLen, len(email))
			}

			if tt.input >= 0 {
				localPart := strings.Split(email, "@")[0]
				if len(localPart) != tt.input {
					t.Errorf("expected local part length %d, got %d", tt.input, len(localPart))
				}

				for _, char := range localPart {
					if !strings.ContainsRune(alphabet, char) {
						t.Errorf("local part contains invalid character: %c", char)
					}
				}
			} else {
				if email != "@email.com" {
					t.Errorf("expected '@email.com', got %s", email)
				}
			}

			if tt.name == "Multiple Invocations Yield Different Results" {
				email2 := RandomEmail(tt.input)
				if email == email2 {
					t.Errorf("expected different emails on multiple invocations, got same: %s", email)
				}
			}
		})
	}
}

/*
ROOST_METHOD_HASH=RandomString_d7e3599ac4
ROOST_METHOD_SIG_HASH=RandomString_c6fe4ad19a


 */
func TestRandomString(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	tests := []struct {
		name     string
		input    int
		validate func(t *testing.T, result string)
	}{
		{
			name:  "Generate a Random String of Given Length",
			input: 10,
			validate: func(t *testing.T, result string) {
				if len(result) != 10 {
					t.Errorf("Expected string length of 10, but got %d", len(result))
				} else {
					t.Log("String length matches the expected value")
				}
			},
		},
		{
			name:  "Generate a Random String of Length Zero",
			input: 0,
			validate: func(t *testing.T, result string) {
				if len(result) != 0 {
					t.Errorf("Expected empty string, but got string of length %d", len(result))
				} else {
					t.Log("String length matches the expected value for zero input")
				}
			},
		},
		{
			name:  "Generate a Random String with Large Length",
			input: 10000,
			validate: func(t *testing.T, result string) {
				if len(result) != 10000 {
					t.Errorf("Expected string length of 10000, but got %d", len(result))
				} else {
					t.Log("String length matches the expected value for large input")
				}
			},
		},
		{
			name:  "Ensure Randomness of Generated Strings",
			input: 10,
			validate: func(t *testing.T, result string) {
				first := RandomString(10)
				second := RandomString(10)
				if first == second {
					t.Errorf("Expected different strings, but got identical strings: %s", first)
				} else {
					t.Log("Strings are different, ensuring randomness")
				}
			},
		},
		{
			name:  "Verify Characters in Generated String",
			input: 10,
			validate: func(t *testing.T, result string) {
				for _, char := range result {
					if !strings.ContainsRune(alphabet, char) {
						t.Errorf("String contains invalid character: %c", char)
					}
				}
				t.Log("All characters in the string are valid")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("Running test: %s", tt.name)
			result := RandomString(tt.input)
			tt.validate(t, result)
		})
	}
}

