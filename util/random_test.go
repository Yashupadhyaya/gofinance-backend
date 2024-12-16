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
		name          string
		number        int
		expectedLen   int
		expectedEmail string
		validate      func(email string) bool
	}{
		{
			name:        "Generating Email with Minimum Length",
			number:      1,
			expectedLen: 12,
			validate: func(email string) bool {
				return len(email) == 12 && strings.HasPrefix(email, "@") == false && strings.HasSuffix(email, "@email.com")
			},
		},
		{
			name:        "Generating Email with Typical Length",
			number:      10,
			expectedLen: 20,
			validate: func(email string) bool {
				return len(email) == 20 && strings.HasPrefix(email, "@") == false && strings.HasSuffix(email, "@email.com")
			},
		},
		{
			name:        "Generating Email with Maximum Length",
			number:      1000,
			expectedLen: 1010,
			validate: func(email string) bool {
				return len(email) == 1010 && strings.HasPrefix(email, "@") == false && strings.HasSuffix(email, "@email.com")
			},
		},
		{
			name:        "Generating Email with Zero Length",
			number:      0,
			expectedLen: 10,
			validate: func(email string) bool {
				return len(email) == 10 && email == "@email.com"
			},
		},
		{
			name:        "Consistency of Generated Email Length",
			number:      15,
			expectedLen: 25,
			validate: func(email string) bool {
				return len(email) == 25 && strings.HasPrefix(email, "@") == false && strings.HasSuffix(email, "@email.com")
			},
		},
		{
			name:        "Valid Characters in Generated Email",
			number:      10,
			expectedLen: 20,
			validate: func(email string) bool {
				for _, char := range email[:10] {
					if !strings.Contains(alphabet, string(char)) {
						return false
					}
				}
				return strings.HasSuffix(email, "@email.com")
			},
		},
		{
			name:        "Multiple Calls Produce Different Results",
			number:      10,
			expectedLen: 20,
			validate: func(email string) bool {
				emails := make(map[string]bool)
				for i := 0; i < 5; i++ {
					em := RandomEmail(10)
					if emails[em] {
						return false
					}
					emails[em] = true
				}
				return true
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			email := RandomEmail(tt.number)
			t.Logf("Generated email: %s", email)

			if len(email) != tt.expectedLen {
				t.Errorf("expected length %d, got %d", tt.expectedLen, len(email))
			}

			if !tt.validate(email) {
				t.Errorf("validation failed for email: %s", email)
			}
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
	rand.Seed(time.Now().UnixNano())

	tests := []struct {
		name       string
		input      int
		assertFunc func(t *testing.T, result string, input int)
	}{
		{
			name:  "Generate a Random String of Specified Length",
			input: 10,
			assertFunc: func(t *testing.T, result string, input int) {
				if len(result) != input {
					t.Errorf("Expected length %d, but got %d", input, len(result))
				} else {
					t.Logf("Successfully generated string of length %d", input)
				}
			},
		},
		{
			name:  "Generate a Random String with Zero Length",
			input: 0,
			assertFunc: func(t *testing.T, result string, input int) {
				if result != "" {
					t.Errorf("Expected empty string, but got %s", result)
				} else {
					t.Logf("Successfully handled zero length, got empty string")
				}
			},
		},
		{
			name:  "Generate a Random String with a Large Length",
			input: 10000,
			assertFunc: func(t *testing.T, result string, input int) {
				if len(result) != input {
					t.Errorf("Expected length %d, but got %d", input, len(result))
				} else {
					t.Logf("Successfully generated string of length %d", input)
				}
			},
		},
		{
			name:  "Verify Randomness of Generated Strings",
			input: 10,
			assertFunc: func(t *testing.T, result string, input int) {
				anotherResult := RandomString(input)
				if result == anotherResult {
					t.Errorf("Expected different strings, but got same: %s", result)
				} else {
					t.Logf("Successfully generated different strings: %s and %s", result, anotherResult)
				}
			},
		},
		{
			name:  "Generate a Random String Using Different Characters from the Alphabet",
			input: 10,
			assertFunc: func(t *testing.T, result string, input int) {
				for _, c := range result {
					if !strings.ContainsRune(alphabet, c) {
						t.Errorf("Character %c not in defined alphabet", c)
					}
				}
				t.Logf("All characters in generated string are part of the defined alphabet")
			},
		},
		{
			name:  "Generate Random Strings with Consistent Length Across Multiple Invocations",
			input: 10,
			assertFunc: func(t *testing.T, result string, input int) {
				for i := 0; i < 5; i++ {
					anotherResult := RandomString(input)
					if len(anotherResult) != input {
						t.Errorf("Expected length %d, but got %d on iteration %d", input, len(anotherResult), i)
					}
				}
				t.Logf("Successfully generated strings of consistent length %d across multiple invocations", input)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RandomString(tt.input)
			tt.assertFunc(t, result, tt.input)
		})
	}
}

