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
	tests := []struct {
		name      string
		input     int
		expected  int
		validator func(string) bool
	}{
		{
			name:     "Generate Email with Minimum Length",
			input:    1,
			expected: 12,
			validator: func(email string) bool {
				return len(email) == 12
			},
		},
		{
			name:     "Generate Email with Typical Length",
			input:    10,
			expected: 21,
			validator: func(email string) bool {
				return len(email) == 21
			},
		},
		{
			name:     "Generate Email with Zero Length",
			input:    0,
			expected: 11,
			validator: func(email string) bool {
				return len(email) == 11
			},
		},
		{
			name:     "Generate Email with Large Length",
			input:    1000,
			expected: 1011,
			validator: func(email string) bool {
				return len(email) == 1011
			},
		},
		{
			name:     "Generate Email with Negative Length",
			input:    -5,
			expected: 11,
			validator: func(email string) bool {
				return len(email) == 11
			},
		},
		{
			name:     "Generate Email with Special Characters in Domain",
			input:    10,
			expected: 21,
			validator: func(email string) bool {
				for _, char := range email {
					if !strings.Contains(alphabet, string(char)) && char != '@' && char != '.' {
						return false
					}
				}
				return len(email) == 21
			},
		},
		{
			name:     "Consistency of Email Format",
			input:    10,
			expected: 21,
			validator: func(email string) bool {
				return strings.HasSuffix(email, "@email.com")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("Running test: %s", tt.name)
			email := RandomEmail(tt.input)
			if !tt.validator(email) {
				t.Errorf("Test %s failed: expected length %d, got %d", tt.name, tt.expected, len(email))
			} else {
				t.Logf("Test %s succeeded: email %s", tt.name, email)
			}
		})
	}
}

/*
ROOST_METHOD_HASH=RandomString_d7e3599ac4
ROOST_METHOD_SIG_HASH=RandomString_c6fe4ad19a


 */
func TestRandomString(t *testing.T) {

	tests := []struct {
		name     string
		input    int
		validate func(string) error
	}{
		{
			name:  "Generating Random String of Zero Length",
			input: 0,
			validate: func(result string) error {
				if len(result) != 0 {
					return fmt.Errorf("expected empty string, got %s", result)
				}
				return nil
			},
		},
		{
			name:  "Generating Random String of Positive Length",
			input: 10,
			validate: func(result string) error {
				if len(result) != 10 {
					return fmt.Errorf("expected string length 10, got %d", len(result))
				}
				return nil
			},
		},
		{
			name:  "Randomness of Generated String",
			input: 10,
			validate: func(result string) error {
				anotherResult := RandomString(10)
				if result == anotherResult {
					return fmt.Errorf("expected different strings, got same strings %s", result)
				}
				return nil
			},
		},
		{
			name:  "Valid Characters in Generated String",
			input: 15,
			validate: func(result string) error {
				for _, c := range result {
					if !strings.ContainsRune(alphabet, c) {
						return fmt.Errorf("invalid character %c in result %s", c, result)
					}
				}
				return nil
			},
		},
		{
			name:  "Large Input Length",
			input: 10000,
			validate: func(result string) error {
				if len(result) != 10000 {
					return fmt.Errorf("expected string length 10000, got %d", len(result))
				}
				return nil
			},
		},
		{
			name:  "Consistency with Fixed Seed",
			input: 10,
			validate: func(result string) error {
				rand.Seed(1)
				firstResult := RandomString(10)
				rand.Seed(1)
				secondResult := RandomString(10)
				if firstResult != secondResult {
					return fmt.Errorf("expected same strings with fixed seed, got different strings %s and %s", firstResult, secondResult)
				}
				return nil
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			rand.Seed(time.Now().UnixNano())

			result := RandomString(tt.input)

			if err := tt.validate(result); err != nil {
				t.Errorf("Test %s failed: %v", tt.name, err)
			} else {
				t.Logf("Test %s passed", tt.name)
			}
		})
	}
}

