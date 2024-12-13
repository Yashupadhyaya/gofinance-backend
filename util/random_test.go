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

	testCases := []struct {
		name        string
		input       int
		expectedLen int
		expectedEnd string
	}{
		{
			name:        "Standard Length",
			input:       10,
			expectedLen: 10,
			expectedEnd: "@email.com",
		},
		{
			name:        "Minimum Length",
			input:       1,
			expectedLen: 1,
			expectedEnd: "@email.com",
		},
		{
			name:        "Maximum Length",
			input:       100,
			expectedLen: 100,
			expectedEnd: "@email.com",
		},
		{
			name:        "Zero Length",
			input:       0,
			expectedLen: 0,
			expectedEnd: "@email.com",
		},
		{
			name:        "Upper Bound Length",
			input:       64,
			expectedLen: 64,
			expectedEnd: "@email.com",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			email := RandomEmail(tc.input)

			localPart := strings.Split(email, "@")[0]
			if len(localPart) != tc.expectedLen {
				t.Errorf("expected local part length %d, got %d", tc.expectedLen, len(localPart))
			}

			if !strings.HasSuffix(email, tc.expectedEnd) {
				t.Errorf("expected email to end with %s, got %s", tc.expectedEnd, email)
			}

			t.Logf("Email generated: %s", email)
		})
	}

	t.Run("Multiple Emails Uniqueness", func(t *testing.T) {
		emailSet := make(map[string]struct{})
		for i := 0; i < 100; i++ {
			email := RandomEmail(10)
			if _, exists := emailSet[email]; exists {
				t.Errorf("duplicate email found: %s", email)
			}
			emailSet[email] = struct{}{}
		}
		t.Log("All generated emails are unique")
	})
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
		expected int
		validate func(string) bool
	}{
		{
			name:     "Generate a Random String of Specified Length",
			input:    10,
			expected: 10,
			validate: func(s string) bool {
				return len(s) == 10
			},
		},
		{
			name:     "Generate a Random String with Zero Length",
			input:    0,
			expected: 0,
			validate: func(s string) bool {
				return s == ""
			},
		},
		{
			name:     "Generate a Random String with a Negative Length",
			input:    -1,
			expected: 0,
			validate: func(s string) bool {
				return s == ""
			},
		},
		{
			name:     "Verify Randomness of Generated Strings",
			input:    10,
			expected: 10,
			validate: func(s string) bool {
				s1 := RandomString(10)
				s2 := RandomString(10)
				return s1 != s2
			},
		},
		{
			name:     "Verify Characters in Generated String are from Alphabet",
			input:    10,
			expected: 10,
			validate: func(s string) bool {
				for _, c := range s {
					if !strings.ContainsRune(alphabet, c) {
						return false
					}
				}
				return true
			},
		},
		{
			name:     "Generate a Random String with a Large Length",
			input:    10000,
			expected: 10000,
			validate: func(s string) bool {
				return len(s) == 10000
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("Running test: %s", tt.name)
			result := RandomString(tt.input)
			if len(result) != tt.expected {
				t.Errorf("Expected length %d, but got %d", tt.expected, len(result))
			}
			if !tt.validate(result) {
				t.Errorf("Validation failed for test: %s", tt.name)
			}
		})
	}
}

