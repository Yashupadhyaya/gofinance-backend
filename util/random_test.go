package util

import (
	"fmt"
	"strings"
	"testing"
	"time"
	"math/rand"
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
		expected    string
	}{
		{
			name:        "Standard Length",
			input:       10,
			expectedLen: 10,
			expected:    "@email.com",
		},
		{
			name:        "Zero Length",
			input:       0,
			expectedLen: 0,
			expected:    "@email.com",
		},
		{
			name:        "Large Length",
			input:       1000,
			expectedLen: 1000,
			expected:    "@email.com",
		},
		{
			name:        "Negative Length",
			input:       -5,
			expectedLen: 0,
			expected:    "@email.com",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			result := RandomEmail(tt.input)

			if len(result) != tt.expectedLen+len(tt.expected) {
				t.Errorf("Expected email length %d, but got %d", tt.expectedLen+len(tt.expected), len(result))
			}

			if !strings.HasSuffix(result, tt.expected) {
				t.Errorf("Expected email to end with %s, but got %s", tt.expected, result)
			}

			t.Logf("Test %s passed with input %d", tt.name, tt.input)
		})
	}

	t.Run("Special Characters", func(t *testing.T) {

		originalRandomString := RandomString
		RandomString = func(number int) string {
			return "!@#$%"
		}
		defer func() { RandomString = originalRandomString }()

		result := RandomEmail(5)

		if !strings.HasSuffix(result, "@email.com") || len(result) != 10 {
			t.Errorf("Expected email with special characters, but got %s", result)
		}

		t.Log("Test Special Characters passed")
	})

	t.Run("Consistency of Format", func(t *testing.T) {
		const length = 5
		emails := make(map[string]bool)

		for i := 0; i < 100; i++ {
			email := RandomEmail(length)
			if !strings.HasSuffix(email, "@email.com") || len(email) != length+len("@email.com") {
				t.Errorf("Inconsistent email format: %s", email)
			}
			emails[email] = true
		}

		if len(emails) != 100 {
			t.Errorf("Expected 100 unique emails, but got %d", len(emails))
		}

		t.Log("Test Consistency of Format passed")
	})
}

/*
ROOST_METHOD_HASH=RandomString_d7e3599ac4
ROOST_METHOD_SIG_HASH=RandomString_c6fe4ad19a


 */
func TestRandomString(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	type testCase struct {
		name     string
		input    int
		validate func(result string) bool
	}

	alphabet := "abcdefghijklmnopqrstuvwxyz"

	tests := []testCase{
		{
			name:  "Positive Length",
			input: 10,
			validate: func(result string) bool {
				if len(result) != 10 {
					t.Logf("Expected length 10, got %d", len(result))
					return false
				}
				return true
			},
		},
		{
			name:  "Zero Length",
			input: 0,
			validate: func(result string) bool {
				if result != "" {
					t.Logf("Expected empty string, got %s", result)
					return false
				}
				return true
			},
		},
		{
			name:  "Negative Length",
			input: -5,
			validate: func(result string) bool {
				if result != "" {
					t.Logf("Expected empty string for negative input, got %s", result)
					return false
				}
				return true
			},
		},
		{
			name:  "Character Set Consistency",
			input: 15,
			validate: func(result string) bool {
				for _, char := range result {
					if !strings.ContainsRune(alphabet, char) {
						t.Logf("Character %c not in alphabet", char)
						return false
					}
				}
				return true
			},
		},
		{
			name:  "Randomness Verification",
			input: 10,
			validate: func(result string) bool {

				variant := RandomString(10)
				if result == variant {
					t.Logf("Generated strings are identical: %s and %s", result, variant)
					return false
				}
				return true
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := RandomString(tc.input)
			if !tc.validate(result) {
				t.Errorf("Test case %s failed", tc.name)
			} else {
				t.Logf("Test case %s passed with result: %s", tc.name, result)
			}
		})
	}

	t.Run("Concurrency Test", func(t *testing.T) {
		done := make(chan bool)
		for i := 0; i < 5; i++ {
			go func() {
				result := RandomString(10)
				if len(result) != 10 {
					t.Errorf("Concurrent test failed, expected length 10, got %d", len(result))
				}
				done <- true
			}()
		}
		for i := 0; i < 5; i++ {
			<-done
		}
		t.Log("Concurrency test passed")
	})
}

