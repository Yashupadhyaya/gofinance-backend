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

	type testCase struct {
		name          string
		input         int
		expectedError bool
		validate      func(t *testing.T, result string)
	}

	rand.Seed(time.Now().UnixNano())

	testCases := []testCase{
		{
			name:  "Generating Email with Standard Length",
			input: 10,
			validate: func(t *testing.T, result string) {
				localPart := strings.Split(result, "@")[0]
				if len(localPart) != 10 {
					t.Errorf("Expected local part length 10, but got %d", len(localPart))
				}
			},
		},
		{
			name:  "Generating Email with Minimum Length",
			input: 1,
			validate: func(t *testing.T, result string) {
				localPart := strings.Split(result, "@")[0]
				if len(localPart) != 1 {
					t.Errorf("Expected local part length 1, but got %d", len(localPart))
				}
			},
		},
		{
			name:  "Generating Email with Zero Length",
			input: 0,
			validate: func(t *testing.T, result string) {
				localPart := strings.Split(result, "@")[0]
				if len(localPart) != 0 {
					t.Errorf("Expected local part length 0, but got %d", len(localPart))
				}
			},
		},
		{
			name:  "Generating Email with Maximum Length",
			input: 100,
			validate: func(t *testing.T, result string) {
				localPart := strings.Split(result, "@")[0]
				if len(localPart) != 100 {
					t.Errorf("Expected local part length 100, but got %d", len(localPart))
				}
			},
		},
		{
			name:  "Consistency Check",
			input: 10,
			validate: func(t *testing.T, result string) {
				emails := make(map[string]struct{})
				for i := 0; i < 10; i++ {
					email := RandomEmail(10)
					if _, exists := emails[email]; exists {
						t.Errorf("Duplicate email found: %s", email)
					}
					emails[email] = struct{}{}
				}
			},
		},
		{
			name:  "Special Characters in Domain",
			input: 10,
			validate: func(t *testing.T, result string) {
				domain := strings.Split(result, "@")[1]
				if domain != "email.com" {
					t.Errorf("Expected domain 'email.com', but got %s", domain)
				}
			},
		},
		{
			name:  "Valid Characters in Local Part",
			input: 10,
			validate: func(t *testing.T, result string) {
				localPart := strings.Split(result, "@")[0]
				for _, char := range localPart {
					if !strings.ContainsRune(alphabet, char) {
						t.Errorf("Invalid character found in local part: %c", char)
					}
				}
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := RandomEmail(tc.input)
			tc.validate(t, result)
			t.Logf("Test '%s' passed", tc.name)
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
		description string
		input       int
		assert      func(t *testing.T, result string)
	}

	testCases := []testCase{
		{
			description: "Generating Random String of Positive Length",
			input:       10,
			assert: func(t *testing.T, result string) {
				if len(result) != 10 {
					t.Errorf("expected length 10, got %d", len(result))
				} else {
					t.Log("Successfully generated random string of length 10")
				}
			},
		},
		{
			description: "Generating Empty String for Zero Length",
			input:       0,
			assert: func(t *testing.T, result string) {
				if len(result) != 0 {
					t.Errorf("expected length 0, got %d", len(result))
				} else {
					t.Log("Successfully generated empty string for zero length")
				}
			},
		},
		{
			description: "Generating Random String with Maximum Length",
			input:       10000,
			assert: func(t *testing.T, result string) {
				if len(result) != 10000 {
					t.Errorf("expected length 10000, got %d", len(result))
				} else {
					t.Log("Successfully generated random string of length 10000")
				}
			},
		},
		{
			description: "Checking Randomness of Generated String",
			input:       10,
			assert: func(t *testing.T, result string) {
				anotherResult := RandomString(10)
				if result == anotherResult {
					t.Errorf("expected different strings, got identical strings")
				} else {
					t.Log("Successfully generated different random strings")
				}
			},
		},
		{
			description: "Verifying Characters in Generated String",
			input:       10,
			assert: func(t *testing.T, result string) {
				for _, char := range result {
					if !strings.ContainsRune(alphabet, char) {
						t.Errorf("unexpected character %c in result string", char)
						return
					}
				}
				t.Log("Successfully verified all characters in the generated string")
			},
		},
		{
			description: "Generating Random String for Negative Length",
			input:       -5,
			assert: func(t *testing.T, result string) {
				if len(result) != 0 {
					t.Errorf("expected length 0 for negative input, got %d", len(result))
				} else {
					t.Log("Successfully handled negative length by returning empty string")
				}
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			t.Logf("Running scenario: %s", tc.description)
			result := RandomString(tc.input)
			tc.assert(t, result)
		})
	}
}

