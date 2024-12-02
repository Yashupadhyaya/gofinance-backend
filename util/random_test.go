package util

import (
	"strings"
	"testing"
	"time"
	"math/rand"
)

/*
ROOST_METHOD_HASH=RandomEmail_1905439733
ROOST_METHOD_SIG_HASH=RandomEmail_7a04f189fd


 */
func RandomEmail(number int) string {
	return fmt.Sprintf("%s@email.com", RandomString(number))
}

func TestRandomEmail(t *testing.T) {

	testCases := []struct {
		name          string
		number        int
		expectedLen   int
		expectedEmail string
	}{
		{
			name:        "Standard Length",
			number:      10,
			expectedLen: 10 + len("@email.com"),
		},
		{
			name:          "Zero Length",
			number:        0,
			expectedEmail: "@email.com",
		},
		{
			name:        "Very Large Length",
			number:      1000,
			expectedLen: 1000 + len("@email.com"),
		},
		{
			name:          "Consistent Domain",
			number:        5,
			expectedEmail: "xxxxx@email.com",
		},
		{
			name:          "Negative Length",
			number:        -1,
			expectedEmail: "@email.com",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			email := RandomEmail(tc.number)
			t.Logf("Generated email: %s", email)

			if tc.expectedLen > 0 && len(email) != tc.expectedLen {
				t.Errorf("Expected email length %d, got %d", tc.expectedLen, len(email))
			}

			if tc.expectedEmail != "" && !strings.HasSuffix(email, "@email.com") {
				t.Errorf("Expected email to end with '@email.com', got %s", email)
			}

			if tc.name == "Consistent Domain" && !strings.HasSuffix(email, "@email.com") {
				t.Errorf("Email does not have consistent domain, got %s", email)
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

	type testCase struct {
		name     string
		input    int
		expected int
		validate func(string) bool
	}

	testCases := []testCase{
		{
			name:     "Generate a Random String of Positive Length",
			input:    5,
			expected: 5,
			validate: func(output string) bool {
				return len(output) == 5
			},
		},
		{
			name:     "Generate a Random String of Zero Length",
			input:    0,
			expected: 0,
			validate: func(output string) bool {
				return output == ""
			},
		},
		{
			name:     "Handle Negative Length Input Gracefully",
			input:    -1,
			expected: 0,
			validate: func(output string) bool {
				return output == ""
			},
		},
		{
			name:     "Generate a Random String with Maximum Length",
			input:    10000,
			expected: 10000,
			validate: func(output string) bool {
				return len(output) == 10000
			},
		},
		{
			name:     "Verify Randomness of Generated Strings",
			input:    10,
			expected: 10,
			validate: func(output string) bool {
				output2 := RandomString(10)
				return output != output2
			},
		},
		{
			name:     "Ensure Characters Are from the Alphabet",
			input:    8,
			expected: 8,
			validate: func(output string) bool {
				for _, char := range output {
					if !strings.ContainsRune(testAlphabet, char) {
						return false
					}
				}
				return true
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := RandomString(tc.input)
			if len(output) != tc.expected {
				t.Errorf("Test %s failed: expected length %d, got %d", tc.name, tc.expected, len(output))
			}
			if !tc.validate(output) {
				t.Errorf("Test %s failed validation", tc.name)
			}
			t.Logf("Test %s succeeded", tc.name)
		})
	}
}

