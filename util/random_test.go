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

	testCases := []struct {
		description string
		input       int
		expectedLen int
		expectedFmt string
	}{
		{
			description: "Generate Random Email with Valid Length",
			input:       10,
			expectedLen: 10,
			expectedFmt: "@email.com",
		},
		{
			description: "Generate Random Email with Zero Length",
			input:       0,
			expectedLen: 0,
			expectedFmt: "@email.com",
		},
		{
			description: "Generate Random Email with Negative Length",
			input:       -5,
			expectedLen: 0,
			expectedFmt: "@email.com",
		},
		{
			description: "Generate Random Email with Large Length",
			input:       1000,
			expectedLen: 1000,
			expectedFmt: "@email.com",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			email := RandomEmail(tc.input)

			localPart := strings.TrimSuffix(email, "@email.com")
			if len(localPart) != tc.expectedLen {
				t.Errorf("Expected local part length %d, got %d", tc.expectedLen, len(localPart))
			}

			if !strings.HasSuffix(email, tc.expectedFmt) {
				t.Errorf("Expected email format to end with '%s', got '%s'", tc.expectedFmt, email)
			}

			t.Logf("Test passed for input: %d", tc.input)
		})
	}

	t.Run("Consistency Test for Repeated Calls", func(t *testing.T) {
		input := 8
		email1 := RandomEmail(input)
		email2 := RandomEmail(input)

		if email1 == email2 {
			t.Errorf("Expected different emails for repeated calls, got same: %s", email1)
		} else {
			t.Logf("Generated different emails: %s and %s", email1, email2)
		}
	})

	t.Run("Validate Email Format", func(t *testing.T) {
		for _, length := range []int{5, 10, 15} {
			email := RandomEmail(length)

			if !strings.Contains(email, "@") || !strings.Contains(email, ".") {
				t.Errorf("Invalid email format: %s", email)
			} else {
				t.Logf("Valid email format for length %d: %s", length, email)
			}
		}
	})
}

/*
ROOST_METHOD_HASH=RandomString_d7e3599ac4
ROOST_METHOD_SIG_HASH=RandomString_c6fe4ad19a


 */
func TestRandomString(t *testing.T) {
	tests := []struct {
		name        string
		length      int
		expectedLen int
		checkRandom bool
		checkChars  bool
		seed        int64
	}{
		{
			name:        "Generate a Random String of Specified Length",
			length:      10,
			expectedLen: 10,
		},
		{
			name:        "Generate an Empty String",
			length:      0,
			expectedLen: 0,
		},
		{
			name:        "Generate a String with Maximum Length",
			length:      10000,
			expectedLen: 10000,
		},
		{
			name:        "Verify Randomness of Generated String",
			length:      10,
			checkRandom: true,
		},
		{
			name:       "Check for Character Set Validity",
			length:     10,
			checkChars: true,
		},
		{
			name:        "Consistent Results with Seeded Randomness",
			length:      10,
			seed:        42,
			checkRandom: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.seed != 0 {
				rand.Seed(tt.seed)
			} else {
				rand.Seed(time.Now().UnixNano())
			}

			result := RandomString(tt.length)

			if len(result) != tt.expectedLen {
				t.Errorf("Expected length %d, but got %d", tt.expectedLen, len(result))
			} else {
				t.Logf("Success: Correct length of %d", tt.expectedLen)
			}

			if tt.checkRandom {
				otherResult := RandomString(tt.length)
				if result == otherResult {
					t.Errorf("Expected different strings, but got the same: %s", result)
				} else {
					t.Logf("Success: Randomness verified with different strings %s and %s", result, otherResult)
				}
			}

			if tt.checkChars {
				for _, char := range result {
					if !strings.ContainsRune(alphabet, char) {
						t.Errorf("Character %c not in alphabet", char)
					}
				}
				t.Logf("Success: All characters in result are valid")
			}

			if tt.seed != 0 {
				otherResult := RandomString(tt.length)
				if result != otherResult {
					t.Errorf("Expected same strings with seed %d, but got different strings: %s and %s", tt.seed, result, otherResult)
				} else {
					t.Logf("Success: Seeded randomness produced consistent results")
				}
			}
		})
	}
}

