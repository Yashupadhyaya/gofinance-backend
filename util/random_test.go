package util

import (
	"fmt"
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

	rand.Seed(1)

	type testCase struct {
		description string
		length      int
		seed        int64
		expectedLen int
		expectEmpty bool
	}

	testCases := []testCase{
		{
			description: "Generate a random string of positive length",
			length:      10,
			expectedLen: 10,
		},
		{
			description: "Generate a random string of zero length",
			length:      0,
			expectedLen: 0,
			expectEmpty: true,
		},
		{
			description: "Generate a random string of negative length",
			length:      -5,
			expectedLen: 0,
			expectEmpty: true,
		},
		{
			description: "Consistency of randomness",
			length:      10,
			expectedLen: 10,
		},
		{
			description: "Consistency of randomness with seed",
			length:      10,
			expectedLen: 10,
			seed:        1,
		},

		{
			description: "Maximum length string",
			length:      10000,
			expectedLen: 10000,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			if tc.seed != 0 {
				rand.Seed(tc.seed)
			}
			result := RandomString(tc.length)

			if tc.expectEmpty {
				if result != "" {
					t.Errorf("Expected empty string, got %s", result)
				}
			} else {
				if len(result) != tc.expectedLen {
					t.Errorf("Expected length %d, got %d", tc.expectedLen, len(result))
				}

				for _, char := range result {
					if !strings.ContainsRune(alphabet, char) {
						t.Errorf("Generated string contains invalid character: %c", char)
					}
				}

				if tc.description == "Consistency of randomness" {
					anotherResult := RandomString(tc.length)
					if result == anotherResult {
						t.Errorf("Expected different strings, got the same: %s", result)
					}
				}

				if tc.description == "Consistency of randomness with seed" {
					anotherResult := RandomString(tc.length)
					if result != anotherResult {
						t.Errorf("Expected same strings with seed, got different: %s and %s", result, anotherResult)
					}
				}
			}
			t.Logf("Test '%s' passed with generated string: %s", tc.description, result)
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

	tests := []struct {
		name        string
		length      int
		expectLocal string
	}{
		{
			name:        "Generate Email with Specified Length",
			length:      10,
			expectLocal: "10 characters expected",
		},
		{
			name:        "Generate Email with Zero Length",
			length:      0,
			expectLocal: "empty local part expected",
		},
		{
			name:        "Generate Email with Negative Length",
			length:      -5,
			expectLocal: "empty local part or default handling expected",
		},
		{
			name:        "Generate Email with Large Length",
			length:      1000,
			expectLocal: "1000 characters expected",
		},
		{
			name:        "Generate Multiple Emails and Ensure Uniqueness",
			length:      10,
			expectLocal: "uniqueness check",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.name {
			case "Generate Email with Specified Length":
				email := RandomEmail(tt.length)
				localPart := strings.Split(email, "@")[0]
				if len(localPart) != tt.length {
					t.Errorf("Failed %s: got %d characters, want %d", tt.name, len(localPart), tt.length)
				} else {
					t.Logf("Success %s: %s", tt.name, tt.expectLocal)
				}

			case "Generate Email with Zero Length":
				email := RandomEmail(tt.length)
				if email != "@email.com" {
					t.Errorf("Failed %s: got %s, want @email.com", tt.name, email)
				} else {
					t.Logf("Success %s: %s", tt.name, tt.expectLocal)
				}

			case "Generate Email with Negative Length":
				email := RandomEmail(tt.length)
				localPart := strings.Split(email, "@")[0]
				if len(localPart) != 0 {
					t.Errorf("Failed %s: unexpected local part length %d", tt.name, len(localPart))
				} else {
					t.Logf("Success %s: %s", tt.name, tt.expectLocal)
				}

			case "Generate Email with Large Length":
				start := time.Now()
				email := RandomEmail(tt.length)
				localPart := strings.Split(email, "@")[0]
				duration := time.Since(start)

				if len(localPart) != tt.length {
					t.Errorf("Failed %s: got %d characters, want %d", tt.name, len(localPart), tt.length)
				} else if duration.Seconds() > 1 {
					t.Errorf("Failed %s: execution took too long", tt.name)
				} else {
					t.Logf("Success %s: %s", tt.name, tt.expectLocal)
				}

			case "Generate Multiple Emails and Ensure Uniqueness":
				emailSet := make(map[string]struct{})
				for i := 0; i < 100; i++ {
					email := RandomEmail(tt.length)
					if _, exists := emailSet[email]; exists {
						t.Errorf("Failed %s: duplicate email found %s", tt.name, email)
						return
					}
					emailSet[email] = struct{}{}
				}
				t.Logf("Success %s: %s", tt.name, tt.expectLocal)
			}
		})
	}
}

