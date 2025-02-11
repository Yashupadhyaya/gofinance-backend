package util

import (
	"math/rand"
	"strings"
	"testing"
	"time"
)








/*
ROOST_METHOD_HASH=RandomString_d7e3599ac4
ROOST_METHOD_SIG_HASH=RandomString_c6fe4ad19a

FUNCTION_DEF=func RandomString(number int) string 

*/
func TestRandomString(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	tests := []struct {
		name     string
		length   int
		expected func(string) bool
	}{
		{
			name:   "Generate a Random String of Specified Length",
			length: 10,
			expected: func(s string) bool {
				return len(s) == 10
			},
		},
		{
			name:   "Generate a Random String with Zero Length",
			length: 0,
			expected: func(s string) bool {
				return s == ""
			},
		},
		{
			name:   "Generate a Random String with Negative Length",
			length: -5,
			expected: func(s string) bool {
				return s == ""
			},
		},
		{
			name:   "Consistency of Randomness in Generated Strings",
			length: 15,
			expected: func(s string) bool {

				anotherString := RandomString(15)
				return s != anotherString
			},
		},
		{
			name:   "Performance Test for Large String Length",
			length: 1000000,
			expected: func(s string) bool {
				return len(s) == 1000000
			},
		},
		{
			name:   "Verify Character Set Usage in Generated String",
			length: 20,
			expected: func(s string) bool {
				for _, char := range s {
					if !strings.ContainsRune(alphabet, char) {
						return false
					}
				}
				return true
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RandomString(tt.length)
			if !tt.expected(result) {
				t.Errorf("Test %s failed: expected condition not met for result %s", tt.name, result)
			} else {
				t.Logf("Test %s passed: result %s", tt.name, result)
			}
		})
	}
}


/*
ROOST_METHOD_HASH=RandomEmail_1905439733
ROOST_METHOD_SIG_HASH=RandomEmail_7a04f189fd

FUNCTION_DEF=func RandomEmail(number int) string 

*/
func TestRandomEmail(t *testing.T) {

	rand.Seed(time.Now().UnixNano())

	testCases := []struct {
		name     string
		input    int
		expected string
	}{
		{
			name:     "Valid Length",
			input:    10,
			expected: "local part of length 10",
		},
		{
			name:     "Zero Length",
			input:    0,
			expected: "@email.com",
		},
		{
			name:     "Maximum Length",
			input:    1000,
			expected: "local part of length 1000",
		},
		{
			name:     "Negative Length",
			input:    -5,
			expected: "handle negative gracefully",
		},
		{
			name:     "Consistent Format",
			input:    5,
			expected: "ends with @email.com",
		},
		{
			name:     "Randomness Check",
			input:    10,
			expected: "unique local parts",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Logf("Running test case: %s", tc.name)
			result := RandomEmail(tc.input)

			switch tc.name {
			case "Valid Length":
				localPart := strings.TrimSuffix(result, "@email.com")
				if len(localPart) != tc.input {
					t.Errorf("Expected local part length %d, got %d", tc.input, len(localPart))
				}
				if !strings.HasSuffix(result, "@email.com") {
					t.Errorf("Expected email to end with '@email.com', got %s", result)
				}

			case "Zero Length":
				if result != tc.expected {
					t.Errorf("Expected %s, got %s", tc.expected, result)
				}

			case "Maximum Length":
				localPart := strings.TrimSuffix(result, "@email.com")
				if len(localPart) != tc.input {
					t.Errorf("Expected local part length %d, got %d", tc.input, len(localPart))
				}
				if !strings.HasSuffix(result, "@email.com") {
					t.Errorf("Expected email to end with '@email.com', got %s", result)
				}

			case "Negative Length":

				if result != "@email.com" {
					t.Errorf("Expected '@email.com' for negative input, got %s", result)
				}

			case "Consistent Format":
				if !strings.HasSuffix(result, "@email.com") {
					t.Errorf("Expected email to end with '@email.com', got %s", result)
				}

			case "Randomness Check":

				uniqueEmails := make(map[string]bool)
				for i := 0; i < 5; i++ {
					email := RandomEmail(tc.input)
					localPart := strings.TrimSuffix(email, "@email.com")
					if uniqueEmails[localPart] {
						t.Errorf("Duplicate local part found: %s", localPart)
					}
					uniqueEmails[localPart] = true
				}
			}
		})
	}
}

