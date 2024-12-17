package util

import (
	"fmt"
	"math"
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

	rand.Seed(time.Now().UnixNano())

	type testCase struct {
		description string
		input       int
		expectedLen int
		expectedEnd string
		validate    func(string) bool
	}

	tests := []testCase{
		{
			description: "Generate Email with 5-Character Local Part",
			input:       5,
			expectedLen: 5 + 10,
			expectedEnd: "@email.com",
			validate: func(email string) bool {
				return len(email) == 15 && strings.HasSuffix(email, "@email.com")
			},
		},
		{
			description: "Generate Email with 0-Character Local Part",
			input:       0,
			expectedLen: 10,
			expectedEnd: "@email.com",
			validate: func(email string) bool {
				return len(email) == 10 && email == "@email.com"
			},
		},
		{
			description: "Generate Email with Large Local Part",
			input:       1000,
			expectedLen: 1000 + 10,
			expectedEnd: "@email.com",
			validate: func(email string) bool {
				return len(email) == 1010 && strings.HasSuffix(email, "@email.com")
			},
		},
		{
			description: "Generate Email with Negative Number",
			input:       -5,
			expectedLen: 10,
			expectedEnd: "@email.com",
			validate: func(email string) bool {
				return len(email) == 10 && email == "@email.com"
			},
		},
		{
			description: "Generate Email with Special Characters in Local Part",
			input:       10,
			expectedLen: 10 + 10,
			expectedEnd: "@email.com",
			validate: func(email string) bool {
				localPart := strings.Split(email, "@")[0]
				for _, char := range localPart {
					if !strings.ContainsRune(alphabet, char) {
						return false
					}
				}
				return true && strings.HasSuffix(email, "@email.com")
			},
		},
		{
			description: "Generate Email and Verify Length of Output",
			input:       15,
			expectedLen: 15 + 10,
			expectedEnd: "@email.com",
			validate: func(email string) bool {
				return len(email) == 25 && strings.HasSuffix(email, "@email.com")
			},
		},
		{
			description: "Generate Multiple Emails and Verify Uniqueness",
			input:       8,
			expectedLen: 8 + 10,
			expectedEnd: "@email.com",
			validate: func(email string) bool {
				emails := map[string]bool{}
				for i := 0; i < 100; i++ {
					email := RandomEmail(8)
					if emails[email] {
						return false
					}
					emails[email] = true
				}
				return true
			},
		},
		{
			description: "Generate Email with Maximum Integer Value",
			input:       math.MaxInt32,
			expectedLen: math.MaxInt32 + 10,
			expectedEnd: "@email.com",
			validate: func(email string) bool {

				return strings.HasSuffix(email, "@email.com")
			},
		},
		{
			description: "Performance Check for Large Input",
			input:       1000000,
			expectedLen: 1000000 + 10,
			expectedEnd: "@email.com",
			validate: func(email string) bool {
				start := time.Now()
				email := RandomEmail(1000000)
				elapsed := time.Since(start)
				return elapsed.Seconds() < 5 && len(email) == 1000010 && strings.HasSuffix(email, "@email.com")
			},
		},
		{
			description: "Generate Email and Check for Correct Domain",
			input:       10,
			expectedLen: 10 + 10,
			expectedEnd: "@email.com",
			validate: func(email string) bool {
				return strings.HasSuffix(email, "@email.com")
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			email := RandomEmail(tc.input)
			if !tc.validate(email) {
				t.Errorf("Test failed for %s: expected length %d and suffix %s, got %s", tc.description, tc.expectedLen, tc.expectedEnd, email)
			} else {
				t.Logf("Test passed for %s: %s", tc.description, email)
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

	tests := []struct {
		name     string
		length   int
		validate func(string, int) error
	}{
		{
			name:   "Generating Random String of Zero Length",
			length: 0,
			validate: func(result string, length int) error {
				if result != "" {
					return fmt.Errorf("expected empty string, got %s", result)
				}
				return nil
			},
		},
		{
			name:   "Generating Random String of Positive Length",
			length: 10,
			validate: func(result string, length int) error {
				if len(result) != length {
					return fmt.Errorf("expected length %d, got %d", length, len(result))
				}
				for _, c := range result {
					if !strings.ContainsRune(alphabet, c) {
						return fmt.Errorf("character %c not in alphabet", c)
					}
				}
				return nil
			},
		},
		{
			name:   "Consistency of Randomness",
			length: 10,
			validate: func(result string, length int) error {
				anotherResult := RandomString(length)
				if result == anotherResult {
					return fmt.Errorf("expected different strings, got identical strings %s and %s", result, anotherResult)
				}
				return nil
			},
		},
		{
			name:   "Handling Large Lengths",
			length: 10000,
			validate: func(result string, length int) error {
				if len(result) != length {
					return fmt.Errorf("expected length %d, got %d", length, len(result))
				}
				return nil
			},
		},
		{
			name:   "Consistent Alphabet Usage",
			length: 20,
			validate: func(result string, length int) error {
				for _, c := range result {
					if !strings.ContainsRune(alphabet, c) {
						return fmt.Errorf("character %c not in alphabet", c)
					}
				}
				return nil
			},
		},
		{
			name:   "Seed Consistency",
			length: 15,
			validate: func(result string, length int) error {
				rand.Seed(42)
				expectedResult := RandomString(length)
				if result != expectedResult {
					return fmt.Errorf("expected %s, got %s", expectedResult, result)
				}
				return nil
			},
		},
		{
			name:   "Boundary Condition for Single Character",
			length: 1,
			validate: func(result string, length int) error {
				if len(result) != length {
					return fmt.Errorf("expected length %d, got %d", length, len(result))
				}
				if !strings.ContainsRune(alphabet, rune(result[0])) {
					return fmt.Errorf("character %c not in alphabet", result[0])
				}
				return nil
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("Running test: %s", tt.name)
			result := RandomString(tt.length)
			if err := tt.validate(result, tt.length); err != nil {
				t.Fatalf("Test %s failed: %v", tt.name, err)
			} else {
				t.Logf("Test %s passed", tt.name)
			}
		})
	}
}

