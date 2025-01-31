package util

import (
	"math/rand"
	"strings"
	"testing"
	"time"
	"fmt"
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
		number   int
		validate func(result string) bool
	}{
		{
			name:   "Generate a Random String of Positive Length",
			number: 10,
			validate: func(result string) bool {
				if len(result) != 10 {
					t.Errorf("Expected string of length 10, got %d", len(result))
					return false
				}
				for _, char := range result {
					if !strings.ContainsRune(alphabet, char) {
						t.Errorf("Character %c is not in the alphabet", char)
						return false
					}
				}
				return true
			},
		},
		{
			name:   "Generate a Random String of Zero Length",
			number: 0,
			validate: func(result string) bool {
				if len(result) != 0 {
					t.Errorf("Expected empty string, got %d", len(result))
					return false
				}
				return true
			},
		},
		{
			name:   "Generate a Random String of Maximum Reasonable Length",
			number: 100000,
			validate: func(result string) bool {
				if len(result) != 100000 {
					t.Errorf("Expected string of length 100000, got %d", len(result))
					return false
				}
				for _, char := range result {
					if !strings.ContainsRune(alphabet, char) {
						t.Errorf("Character %c is not in the alphabet", char)
						return false
					}
				}
				return true
			},
		},
		{
			name:   "Consistent Randomness Across Multiple Calls",
			number: 15,
			validate: func(result string) bool {
				anotherResult := RandomString(15)
				if result == anotherResult {
					t.Errorf("Expected different strings but got identical: %s", result)
					return false
				}
				return true
			},
		},
		{
			name:   "Validate Characters Used in Generated String",
			number: 20,
			validate: func(result string) bool {
				for _, char := range result {
					if !strings.ContainsRune(alphabet, char) {
						t.Errorf("Character %c is not in the alphabet", char)
						return false
					}
				}
				return true
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RandomString(tt.number)
			if !tt.validate(result) {
				t.Logf("Test %s failed with result: %s", tt.name, result)
			} else {
				t.Logf("Test %s succeeded with result: %s", tt.name, result)
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
	type testCase struct {
		description string
		number      int
		expectedLen int
		expectedFmt string
	}

	testCases := []testCase{
		{
			description: "Generate Email with Typical Length",
			number:      10,
			expectedLen: 10,
			expectedFmt: "%s@email.com",
		},
		{
			description: "Generate Email with Minimum Length",
			number:      1,
			expectedLen: 1,
			expectedFmt: "%s@email.com",
		},
		{
			description: "Generate Email with Maximum Length",
			number:      64,
			expectedLen: 64,
			expectedFmt: "%s@email.com",
		},
		{
			description: "Generate Email with Zero Length",
			number:      0,
			expectedLen: 0,
			expectedFmt: "@email.com",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			email := RandomEmail(tc.number)
			localPart := strings.Split(email, "@")[0]

			if tc.number == 0 {
				if email != tc.expectedFmt {
					t.Errorf("expected email format '%s', got '%s'", tc.expectedFmt, email)
				}
			} else {
				if len(localPart) != tc.expectedLen {
					t.Errorf("expected local part length %d, got %d", tc.expectedLen, len(localPart))
				}
				if !strings.HasSuffix(email, "@email.com") {
					t.Errorf("expected email to end with '@email.com', got '%s'", email)
				}
			}

			t.Logf("Test '%s' passed successfully with generated email: %s", tc.description, email)
		})
	}

	t.Run("Randomness and Uniqueness of Generated Emails", func(t *testing.T) {
		uniqueEmails := make(map[string]bool)
		const iterations = 100

		for i := 0; i < iterations; i++ {
			email := RandomEmail(10)
			if _, exists := uniqueEmails[email]; exists {
				t.Errorf("email '%s' was generated more than once", email)
				return
			}
			uniqueEmails[email] = true
		}

		t.Log("All generated emails are unique over 100 iterations.")
	})

}

