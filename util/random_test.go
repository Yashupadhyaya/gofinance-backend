package util

import (
	"fmt"
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
		name          string
		input         int
		expectedLocal int
		expectedEmail string
	}{
		{
			name:          "Generate Email with 5-Character Local Part",
			input:         5,
			expectedLocal: 5,
			expectedEmail: "",
		},
		{
			name:          "Generate Email with 0-Character Local Part",
			input:         0,
			expectedLocal: 0,
			expectedEmail: "@email.com",
		},
		{
			name:          "Generate Email with Maximum Expected Length",
			input:         100,
			expectedLocal: 100,
			expectedEmail: "",
		},
		{
			name:          "Check for Valid Characters in Local Part",
			input:         10,
			expectedLocal: 10,
			expectedEmail: "",
		},
		{
			name:          "Generate Email and Check Domain Consistency",
			input:         10,
			expectedLocal: 10,
			expectedEmail: "",
		},
		{
			name:          "Generate Multiple Emails and Ensure Uniqueness",
			input:         10,
			expectedLocal: 10,
			expectedEmail: "",
		},
		{
			name:          "Check Performance for Large Number of Calls",
			input:         10,
			expectedLocal: 10,
			expectedEmail: "",
		},
		{
			name:          "Check for Non-Empty Local Part",
			input:         1,
			expectedLocal: 1,
			expectedEmail: "",
		},
		{
			name:          "Check for Correct Formatting",
			input:         10,
			expectedLocal: 10,
			expectedEmail: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			email := RandomEmail(tt.input)
			localPart := strings.Split(email, "@")[0]

			switch tt.name {
			case "Generate Email with 5-Character Local Part":
				if len(localPart) != tt.expectedLocal {
					t.Errorf("Expected local part length %d, got %d", tt.expectedLocal, len(localPart))
				}
			case "Generate Email with 0-Character Local Part":
				if email != tt.expectedEmail {
					t.Errorf("Expected email %s, got %s", tt.expectedEmail, email)
				}
			case "Generate Email with Maximum Expected Length":
				if len(localPart) != tt.expectedLocal {
					t.Errorf("Expected local part length %d, got %d", tt.expectedLocal, len(localPart))
				}
			case "Check for Valid Characters in Local Part":
				for _, char := range localPart {
					if !strings.ContainsRune(alphabet, char) {
						t.Errorf("Invalid character %c in local part", char)
					}
				}
			case "Generate Email and Check Domain Consistency":
				if !strings.HasSuffix(email, "@email.com") {
					t.Errorf("Expected email to end with @email.com, got %s", email)
				}
			case "Generate Multiple Emails and Ensure Uniqueness":
				emails := make(map[string]bool)
				for i := 0; i < 100; i++ {
					email := RandomEmail(tt.input)
					localPart := strings.Split(email, "@")[0]
					if emails[localPart] {
						t.Errorf("Duplicate email found: %s", email)
					}
					emails[localPart] = true
				}
			case "Check Performance for Large Number of Calls":
				start := time.Now()
				for i := 0; i < 10000; i++ {
					RandomEmail(tt.input)
				}
				duration := time.Since(start)
				t.Logf("Generated 10000 emails in %v", duration)
			case "Check for Non-Empty Local Part":
				if len(localPart) == 0 {
					t.Error("Expected non-empty local part")
				}
			case "Check for Correct Formatting":
				expectedFormat := fmt.Sprintf("%s@email.com", localPart)
				if email != expectedFormat {
					t.Errorf("Expected email format %s, got %s", expectedFormat, email)
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
			name:  "Generating Random String of Zero Length",
			input: 0,
			validate: func(t *testing.T, result string) {
				if result != "" {
					t.Errorf("expected empty string, got %s", result)
				}
			},
		},
		{
			name:  "Generating Random String of Length One",
			input: 1,
			validate: func(t *testing.T, result string) {
				if len(result) != 1 {
					t.Errorf("expected string of length 1, got %d", len(result))
				}
			},
		},
		{
			name:  "Generating Random String of Specific Length",
			input: 10,
			validate: func(t *testing.T, result string) {
				if len(result) != 10 {
					t.Errorf("expected string of length 10, got %d", len(result))
				}
			},
		},
		{
			name:  "Randomness of Generated String",
			input: 10,
			validate: func(t *testing.T, result string) {

				anotherResult := RandomString(10)
				if result == anotherResult {
					t.Error("expected different strings on consecutive calls but got the same")
				}
			},
		},
		{
			name:  "Consistency of Length for Large Input",
			input: 10000,
			validate: func(t *testing.T, result string) {
				if len(result) != 10000 {
					t.Errorf("expected string of length 10000, got %d", len(result))
				}
			},
		},
		{
			name:  "Characters in Generated String",
			input: 50,
			validate: func(t *testing.T, result string) {
				for _, char := range result {
					if !strings.Contains(alphabet, string(char)) {
						t.Errorf("character %c not in alphabet", char)
					}
				}
			},
		},
		{
			name:  "Performance Under Stress",
			input: 100000,
			validate: func(t *testing.T, result string) {
				start := time.Now()
				for i := 0; i < 10; i++ {
					RandomString(100000)
				}
				elapsed := time.Since(start)
				if elapsed.Seconds() > 1 {
					t.Errorf("expected performance within 1 second, but took %s", elapsed)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RandomString(tt.input)
			tt.validate(t, result)
		})
	}
}

