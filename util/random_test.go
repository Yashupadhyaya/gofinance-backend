package util

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"
	"unicode"
)

/*
ROOST_METHOD_HASH=RandomEmail_1905439733
ROOST_METHOD_SIG_HASH=RandomEmail_7a04f189fd


 */
func RandomEmail(number int) string {
	return fmt.Sprintf("%s@email.com", RandomString(number))
}

func RandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	if length <= 0 {
		return ""
	}
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func TestRandomEmail(t *testing.T) {

	rand.Seed(time.Now().UnixNano())

	testCases := []struct {
		name        string
		input       int
		expectedLen int
		description string
	}{
		{
			name:        "Standard Length",
			input:       10,
			expectedLen: 10,
			description: "Generate Email with a Standard Length",
		},
		{
			name:        "Minimum Length",
			input:       1,
			expectedLen: 1,
			description: "Generate Email with Minimum Length",
		},
		{
			name:        "Large Length",
			input:       1000,
			expectedLen: 1000,
			description: "Generate Email with Large Length",
		},
		{
			name:        "Zero Length",
			input:       0,
			expectedLen: 0,
			description: "Generate Email with Zero Length",
		},
		{
			name:        "Negative Length",
			input:       -5,
			expectedLen: 0,
			description: "Generate Email with Negative Length",
		},
		{
			name:        "Special Characters in Random String",
			input:       10,
			expectedLen: 10,
			description: "Ensure no special characters in Random String",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Log(tc.description)

			email := RandomEmail(tc.input)

			localPart := strings.Split(email, "@")[0]
			if len(localPart) != tc.expectedLen {
				t.Errorf("Expected local part length %d, got %d", tc.expectedLen, len(localPart))
			}

			if tc.name == "Special Characters in Random String" {
				for _, char := range localPart {
					if !unicode.IsLetter(char) && !unicode.IsDigit(char) {
						t.Errorf("Invalid character in email local part: %c", char)
					}
				}
			}

			t.Logf("Generated email: %s", email)
		})
	}
}

/*
ROOST_METHOD_HASH=RandomString_d7e3599ac4
ROOST_METHOD_SIG_HASH=RandomString_c6fe4ad19a


 */
func RandomString(number int) string {
	if number < 0 {
		return ""
	}

	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < number; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func TestRandomString(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	type testCase struct {
		name     string
		input    int
		expected func(string) bool
	}

	tests := []testCase{
		{
			name:  "Positive Length",
			input: 10,
			expected: func(result string) bool {
				return len(result) == 10 && containsOnlyAlphabet(result)
			},
		},
		{
			name:  "Zero Length",
			input: 0,
			expected: func(result string) bool {
				return result == ""
			},
		},
		{
			name:  "Large Length",
			input: 100000,
			expected: func(result string) bool {
				return len(result) == 100000 && containsOnlyAlphabet(result)
			},
		},
		{
			name:  "Negative Length",
			input: -5,
			expected: func(result string) bool {
				return result == ""
			},
		},
		{
			name:  "Consistency of Randomness",
			input: 10,
			expected: func(result string) bool {
				anotherResult := RandomString(10)
				return result != anotherResult
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := RandomString(tc.input)
			if !tc.expected(result) {
				t.Errorf("Test %s failed: unexpected result %q", tc.name, result)
			} else {
				t.Logf("Test %s succeeded: result %q", tc.name, result)
			}
		})
	}
}

func containsOnlyAlphabet(s string) bool {
	for _, c := range s {
		if !strings.ContainsRune(alphabet, c) {
			return false
		}
	}
	return true
}

