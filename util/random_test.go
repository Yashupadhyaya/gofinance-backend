package util

import (
	"fmt"
	"strings"
	"testing"
	"unicode"
	"math/rand"
	"time"
)

/*
ROOST_METHOD_HASH=RandomEmail_1905439733
ROOST_METHOD_SIG_HASH=RandomEmail_7a04f189fd


 */
func TestRandomEmail(t *testing.T) {
	type testCase struct {
		description string
		input       int
		expectedLen int
		assertFunc  func(t *testing.T, email string)
	}

	testCases := []testCase{
		{
			description: "Generate Email with 5-Character Local Part",
			input:       5,
			expectedLen: 16,
			assertFunc: func(t *testing.T, email string) {
				localPart := strings.Split(email, "@")[0]
				if len(localPart) != 5 {
					t.Errorf("Expected local part length of 5, got %d", len(localPart))
				}
			},
		},
		{
			description: "Generate Email with 0-Character Local Part",
			input:       0,
			expectedLen: 11,
			assertFunc: func(t *testing.T, email string) {
				localPart := strings.Split(email, "@")[0]
				if localPart != "" {
					t.Errorf("Expected empty local part, got %s", localPart)
				}
			},
		},
		{
			description: "Generate Email with Max Integer Local Part",
			input:       1000000,
			expectedLen: 1000011,
			assertFunc: func(t *testing.T, email string) {
				localPart := strings.Split(email, "@")[0]
				if len(localPart) != 1000000 {
					t.Errorf("Expected local part length of 1000000, got %d", len(localPart))
				}
			},
		},
		{
			description: "Generate Email with Negative Number",
			input:       -1,
			expectedLen: 11,
			assertFunc: func(t *testing.T, email string) {
				localPart := strings.Split(email, "@")[0]
				if localPart != "" {
					t.Errorf("Expected empty local part for negative input, got %s", localPart)
				}
			},
		},
		{
			description: "Generate Email with 1-Character Local Part",
			input:       1,
			expectedLen: 12,
			assertFunc: func(t *testing.T, email string) {
				localPart := strings.Split(email, "@")[0]
				if len(localPart) != 1 {
					t.Errorf("Expected local part length of 1, got %d", len(localPart))
				}
			},
		},
		{
			description: "Generate Email with Special Characters in Local Part",
			input:       10,
			expectedLen: 21,
			assertFunc: func(t *testing.T, email string) {
				localPart := strings.Split(email, "@")[0]
				for _, c := range localPart {
					if !unicode.IsLetter(c) {
						t.Errorf("Expected only alphabetic characters, got %c", c)
					}
				}
			},
		},
		{
			description: "Generate Email with Upper Bound of Alphabet",
			input:       26,
			expectedLen: 37,
			assertFunc: func(t *testing.T, email string) {
				localPart := strings.Split(email, "@")[0]
				usedChars := make(map[rune]bool)
				for _, c := range localPart {
					usedChars[c] = true
				}
				if len(usedChars) < 26 {
					t.Errorf("Expected all alphabet characters to be used, got %d unique characters", len(usedChars))
				}
			},
		},
		{
			description: "Generate Email with Consistent Output Length",
			input:       10,
			expectedLen: 21,
			assertFunc: func(t *testing.T, email string) {
				if len(email) != 21 {
					t.Errorf("Expected email length of 21, got %d", len(email))
				}
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			email := RandomEmail(tc.input)
			t.Logf("Generated email: %s", email)

			if len(email) != tc.expectedLen {
				t.Errorf("Expected total email length of %d, got %d", tc.expectedLen, len(email))
			}

			tc.assertFunc(t, email)
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
		name        string
		input       int
		validate    func(t *testing.T, result string)
		description string
	}{
		{
			name:  "Zero Length",
			input: 0,
			validate: func(t *testing.T, result string) {
				if result != "" {
					t.Errorf("Expected empty string, got %v", result)
				} else {
					t.Logf("Success: Received expected empty string")
				}
			},
			description: "This test checks if the function correctly handles the case where the requested string length is zero.",
		},
		{
			name:  "Positive Length",
			input: 10,
			validate: func(t *testing.T, result string) {
				if len(result) != 10 {
					t.Errorf("Expected string length of 10, got %v", len(result))
				} else {
					t.Logf("Success: Received string of correct length 10")
				}
			},
			description: "This test checks if the function generates a string of the correct length when given a positive integer.",
		},
		{
			name:  "Consistent Characters",
			input: 20,
			validate: func(t *testing.T, result string) {
				for _, c := range result {
					if !strings.ContainsRune(alphabet, c) {
						t.Errorf("Unexpected character %v in result", c)
					}
				}
				t.Logf("Success: Received string with consistent characters")
			},
			description: "This test checks if the generated string only contains characters from the defined alphabet.",
		},
		{
			name:  "Seeded Randomness",
			input: 10,
			validate: func(t *testing.T, result string) {
				rand.Seed(time.Now().UnixNano())
				result2 := RandomString(10)
				if result == result2 {
					t.Errorf("Expected different strings, got identical strings %v", result)
				} else {
					t.Logf("Success: Received different strings on different runs")
				}
			},
			description: "This test checks if the function generates different strings on different runs if the random seed is not set.",
		},
		{
			name:  "Fixed Seed",
			input: 10,
			validate: func(t *testing.T, result string) {
				rand.Seed(1)
				result2 := RandomString(10)
				rand.Seed(1)
				result3 := RandomString(10)
				if result2 != result3 {
					t.Errorf("Expected identical strings with fixed seed, got %v and %v", result2, result3)
				} else {
					t.Logf("Success: Received identical strings with fixed seed")
				}
			},
			description: "This test checks if the function generates the same string when the random seed is fixed.",
		},
		{
			name:  "Performance Large Input",
			input: 1000000,
			validate: func(t *testing.T, result string) {
				if len(result) != 1000000 {
					t.Errorf("Expected string length of 1000000, got %v", len(result))
				} else {
					t.Logf("Success: Function handled large input efficiently")
				}
			},
			description: "This test checks the performance of the function when generating a very large string.",
		},
		{
			name:  "Non-ASCII Characters",
			input: 15,
			validate: func(t *testing.T, result string) {
				for _, c := range result {
					if c < 0 || c > 127 {
						t.Errorf("Non-ASCII character %v in result", c)
					}
				}
				t.Logf("Success: Received string with only ASCII characters")
			},
			description: "This test checks that the function does not include non-ASCII characters.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(tt.description)
			result := RandomString(tt.input)
			tt.validate(t, result)
		})
	}
}

