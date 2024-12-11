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
func RandomEmail(number int) string {
	return fmt.Sprintf("%s@email.com", RandomString(number))
}

func RandomString(number int) string {
	const alphabet = "abcdefghijklmnopqrstuvwxyz"
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < number; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func TestRandomEmail(t *testing.T) {

	tests := []struct {
		name      string
		input     int
		localLen  int
		expectErr bool
	}{
		{
			name:      "Generate Email with 5-Character Local Part",
			input:     5,
			localLen:  5,
			expectErr: false,
		},
		{
			name:      "Generate Email with 10-Character Local Part",
			input:     10,
			localLen:  10,
			expectErr: false,
		},
		{
			name:      "Generate Email with 0-Character Local Part",
			input:     0,
			localLen:  0,
			expectErr: false,
		},
		{
			name:      "Generate Email with 1-Character Local Part",
			input:     1,
			localLen:  1,
			expectErr: false,
		},
		{
			name:      "Generate Email with Maximum Reasonable Length Local Part",
			input:     100,
			localLen:  100,
			expectErr: false,
		},
		{
			name:      "Check Email Format Consistency",
			input:     8,
			localLen:  8,
			expectErr: false,
		},
		{
			name:      "Check Randomness of Generated Local Part",
			input:     8,
			localLen:  8,
			expectErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			email := RandomEmail(tt.input)

			if !strings.HasSuffix(email, "@email.com") {
				t.Errorf("Expected email to end with '@email.com', got %s", email)
			}

			localPart := strings.TrimSuffix(email, "@email.com")
			if len(localPart) != tt.localLen {
				t.Errorf("Expected local part length to be %d, got %d", tt.localLen, len(localPart))
			}

			if tt.name == "Check Randomness of Generated Local Part" {
				email2 := RandomEmail(tt.input)
				if email == email2 {
					t.Errorf("Expected different emails on multiple invocations, got same email: %s", email)
				}
			}

			t.Logf("Test case %s passed for input %d", tt.name, tt.input)
		})
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

/*
ROOST_METHOD_HASH=RandomString_d7e3599ac4
ROOST_METHOD_SIG_HASH=RandomString_c6fe4ad19a


 */
func TestRandomString(t *testing.T) {

	testCases := []struct {
		name     string
		input    int
		validate func(result string) bool
		expected string
	}{
		{
			name:  "Generate a Random String of Specified Length",
			input: 10,
			validate: func(result string) bool {
				return len(result) == 10
			},
			expected: "a string of length 10",
		},
		{
			name:  "Generate a Random String with Zero Length",
			input: 0,
			validate: func(result string) bool {
				return result == ""
			},
			expected: "an empty string",
		},
		{
			name:  "Generate Multiple Random Strings and Ensure Uniqueness",
			input: 10,
			validate: func(result string) bool {

				generatedStrings := map[string]bool{}
				for i := 0; i < 5; i++ {
					str := RandomString(10)
					if generatedStrings[str] {
						return false
					}
					generatedStrings[str] = true
				}
				return true
			},
			expected: "unique strings",
		},
		{
			name:  "Generate a Random String with Maximum Length",
			input: 1000000,
			validate: func(result string) bool {
				return len(result) == 1000000
			},
			expected: "a string of length 1000000",
		},
		{
			name:  "Generate a Random String with Negative Length",
			input: -10,
			validate: func(result string) bool {
				return result == ""
			},
			expected: "an empty string",
		},
		{
			name:  "Generate a Random String and Ensure Characters are from Alphabet",
			input: 10,
			validate: func(result string) bool {
				for _, char := range result {
					if !strings.ContainsRune(alphabet, char) {
						return false
					}
				}
				return true
			},
			expected: "a string with only alphabet characters",
		},
		{
			name:  "Generate Random String with Different Seeds",
			input: 10,
			validate: func(result string) bool {

				rand.Seed(time.Now().UnixNano())
				str1 := RandomString(10)
				rand.Seed(time.Now().UnixNano() + 1)
				str2 := RandomString(10)
				return str1 != str2
			},
			expected: "different strings with different seeds",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			result := RandomString(tc.input)

			if !tc.validate(result) {
				t.Errorf("Test %s failed. Expected %s, but got %s", tc.name, tc.expected, result)
			} else {
				t.Logf("Test %s passed. Got expected result: %s", tc.name, result)
			}
		})
	}
}

