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
	type testCase struct {
		description string
		input       int
		expected    string
		validate    func(t *testing.T, result string)
	}

	testCases := []testCase{
		{
			description: "Generate Email with 5-Character Local Part",
			input:       5,
			expected:    "5-character local part followed by @email.com",
			validate: func(t *testing.T, result string) {
				localPart := strings.Split(result, "@")[0]
				if len(localPart) != 5 || !strings.HasSuffix(result, "@email.com") {
					t.Errorf("expected local part of 5 characters followed by @email.com, got %s", result)
				}
			},
		},
		{
			description: "Generate Email with 0-Character Local Part",
			input:       0,
			expected:    "@email.com",
			validate: func(t *testing.T, result string) {
				if result != "@email.com" {
					t.Errorf("expected @email.com, got %s", result)
				}
			},
		},
		{
			description: "Generate Email with Maximum Character Local Part",
			input:       100,
			expected:    "100-character local part followed by @email.com",
			validate: func(t *testing.T, result string) {
				localPart := strings.Split(result, "@")[0]
				if len(localPart) != 100 || !strings.HasSuffix(result, "@email.com") {
					t.Errorf("expected local part of 100 characters followed by @email.com, got %s", result)
				}
			},
		},
		{
			description: "Generate Email with Negative Character Local Part",
			input:       -5,
			expected:    "@email.com",
			validate: func(t *testing.T, result string) {
				if result != "@email.com" {
					t.Errorf("expected @email.com for negative input, got %s", result)
				}
			},
		},
		{
			description: "Generate Email Multiple Times with Same Input",
			input:       10,
			expected:    "different email addresses each time",
			validate: func(t *testing.T, result string) {
				email1 := RandomEmail(10)
				email2 := RandomEmail(10)
				if email1 == email2 {
					t.Errorf("expected different email addresses, got %s and %s", email1, email2)
				}
			},
		},
		{
			description: "Generate Email with Special Characters in Local Part",
			input:       10,
			expected:    "valid characters in local part",
			validate: func(t *testing.T, result string) {
				localPart := strings.Split(result, "@")[0]
				for _, char := range localPart {
					if !strings.ContainsRune(alphabet, char) {
						t.Errorf("invalid character %c in local part", char)
					}
				}
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			t.Logf("Running test: %s", tc.description)
			rand.Seed(time.Now().UnixNano())
			result := RandomEmail(tc.input)
			tc.validate(t, result)
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
		expected func(string) bool
	}{
		{
			name:  "Generating Random String of Zero Length",
			input: 0,
			expected: func(s string) bool {
				return len(s) == 0
			},
		},
		{
			name:  "Generating Random String of Positive Length",
			input: 10,
			expected: func(s string) bool {
				return len(s) == 10
			},
		},
		{
			name:  "Consistency in Length of Generated Strings",
			input: 15,
			expected: func(s string) bool {
				return len(s) == 15
			},
		},
		{
			name:  "Randomness of Generated Strings",
			input: 8,
			expected: func(s string) bool {

				anotherString := RandomString(8)
				return s != anotherString
			},
		},
		{
			name:  "Handling Large String Lengths",
			input: 10000,
			expected: func(s string) bool {
				return len(s) == 10000
			},
		},
		{
			name:  "Valid Characters in Generated String",
			input: 20,
			expected: func(s string) bool {
				for _, char := range s {
					if !strings.Contains(alphabet, string(char)) {
						return false
					}
				}
				return true
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RandomString(tt.input)
			if !tt.expected(result) {
				t.Errorf("Test %s failed: expected criteria not met", tt.name)
			} else {
				t.Logf("Test %s succeeded: result = %s", tt.name, result)
			}
		})
	}
}

