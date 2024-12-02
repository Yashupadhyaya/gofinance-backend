package util

import (
	"fmt"
	"strings"
	"testing"
	"math/rand"
	"time"
)

/*
ROOST_METHOD_HASH=RandomEmail_1905439733
ROOST_METHOD_SIG_HASH=RandomEmail_7a04f189fd


 */
func TestRandomEmail(t *testing.T) {

	rand.Seed(time.Now().UnixNano())

	tests := []struct {
		name           string
		input          int
		expectedLength int
		expectedOutput string
		expectError    bool
	}{
		{
			name:           "Generate Email with a Standard Length",
			input:          10,
			expectedLength: 10,
			expectedOutput: "",
			expectError:    false,
		},
		{
			name:           "Generate Email with Zero Length",
			input:          0,
			expectedLength: 0,
			expectedOutput: "@email.com",
			expectError:    false,
		},
		{
			name:           "Generate Email with Maximum Length",
			input:          254,
			expectedLength: 254,
			expectedOutput: "",
			expectError:    false,
		},
		{
			name:           "Generate Email with Negative Length",
			input:          -5,
			expectedLength: 0,
			expectedOutput: "",
			expectError:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if !tt.expectError {
						t.Errorf("Unexpected error: %v", r)
					}
				}
			}()

			result := RandomEmail(tt.input)
			localPart := strings.Split(result, "@")[0]

			if tt.expectError {
				if len(localPart) != 0 {
					t.Errorf("Expected error, got valid email: %s", result)
				}
				return
			}

			if tt.expectedOutput != "" && result != tt.expectedOutput {
				t.Errorf("Expected %s, got %s", tt.expectedOutput, result)
			}

			if tt.expectedLength != 0 && len(localPart) != tt.expectedLength {
				t.Errorf("Expected local part length %d, got %d", tt.expectedLength, len(localPart))
			}

			t.Logf("Generated email: %s", result)
		})
	}

	t.Run("Generate Email with a Typical Random Seed", func(t *testing.T) {
		email1 := RandomEmail(10)
		email2 := RandomEmail(10)

		if email1 == email2 {
			t.Errorf("Expected different emails, got %s and %s", email1, email2)
		}

		t.Logf("Generated emails: %s, %s", email1, email2)
	})
}

/*
ROOST_METHOD_HASH=RandomString_d7e3599ac4
ROOST_METHOD_SIG_HASH=RandomString_c6fe4ad19a


 */
func TestRandomString(t *testing.T) {

	rand.Seed(time.Now().UnixNano())

	tests := []struct {
		name        string
		input       int
		expectEmpty bool
		expectError bool
	}{
		{
			name:        "Positive Length",
			input:       5,
			expectEmpty: false,
			expectError: false,
		},
		{
			name:        "Zero Length",
			input:       0,
			expectEmpty: true,
			expectError: false,
		},
		{
			name:        "Negative Length",
			input:       -5,
			expectEmpty: true,
			expectError: false,
		},
		{
			name:        "Large Length",
			input:       100000,
			expectEmpty: false,
			expectError: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := RandomString(tc.input)

			if tc.expectEmpty {
				if len(result) != 0 {
					t.Errorf("expected empty string, got '%s'", result)
				} else {
					t.Logf("success: expected empty string and got empty string")
				}
			} else if len(result) != tc.input {
				t.Errorf("expected string length %d, got %d", tc.input, len(result))
			} else {
				t.Logf("success: expected string length %d and got %d", tc.input, len(result))
			}

			if !tc.expectEmpty {
				for _, char := range result {
					if !strings.ContainsRune(alphabet, char) {
						t.Errorf("character '%c' is not in the alphabet", char)
					}
				}
			}

			if !tc.expectEmpty && tc.input > 0 {
				anotherResult := RandomString(tc.input)
				if result == anotherResult {
					t.Errorf("generated strings should be different: '%s' and '%s'", result, anotherResult)
				} else {
					t.Logf("success: generated strings are different: '%s' and '%s'", result, anotherResult)
				}
			}
		})
	}

	t.Run("Performance with Large Length", func(t *testing.T) {
		start := time.Now()
		result := RandomString(100000)
		duration := time.Since(start)
		if len(result) != 100000 {
			t.Errorf("expected string length 100000, got %d", len(result))
		}
		t.Logf("success: generated string of length %d in %s", len(result), duration)
	})

}

