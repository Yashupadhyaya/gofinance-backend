package util

import (
	"fmt"
	"math/rand"
	"regexp"
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
		input          int
		expectedLength int
		expectedOutput string
		description    string
	}

	testCases := []testCase{
		{
			input:          10,
			expectedLength: 10,
			description:    "Generate Email with Positive Number of Characters",
		},
		{
			input:          0,
			expectedLength: 0,
			expectedOutput: "@email.com",
			description:    "Generate Email with Zero Characters",
		},
		{
			input:          -5,
			expectedLength: 0,
			expectedOutput: "@email.com",
			description:    "Generate Email with Negative Number of Characters",
		},
		{
			input:          1000,
			expectedLength: 1000,
			description:    "Generate Email with Large Number of Characters",
		},
		{
			input:          15,
			expectedLength: 15,
			description:    "Generate Email with Special Characters in Local Part",
		},
		{
			input:          8,
			expectedLength: 8,
			description:    "Generate Email with Consistent Output Format",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			t.Logf("Running test case: %v", tc.description)

			result := RandomEmail(tc.input)

			if tc.expectedOutput != "" {
				if result != tc.expectedOutput {
					t.Errorf("Expected output: %v, but got: %v", tc.expectedOutput, result)
				}
			} else {
				localPart := strings.Split(result, "@")[0]

				if len(localPart) != tc.expectedLength {
					t.Errorf("Expected length of local part: %v, but got: %v", tc.expectedLength, len(localPart))
				}

				if !strings.HasSuffix(result, "@email.com") {
					t.Errorf("Expected email to end with '@email.com', but got: %v", result)
				}

				if tc.input > 0 {
					matched, err := regexp.MatchString(fmt.Sprintf("^[a-z]{%d}$", tc.input), localPart)
					if err != nil {
						t.Fatalf("Error matching regex: %v", err)
					}
					if !matched {
						t.Errorf("Local part contains invalid characters or incorrect length: %v", localPart)
					}
				}
			}
		})
	}

	t.Run("Consistent Output Format", func(t *testing.T) {
		input := 8
		expectedLength := 8
		for i := 0; i < 10; i++ {
			result := RandomEmail(input)
			localPart := strings.Split(result, "@")[0]

			if len(localPart) != expectedLength {
				t.Errorf("Expected length of local part: %v, but got: %v", expectedLength, len(localPart))
			}

			if !strings.HasSuffix(result, "@email.com") {
				t.Errorf("Expected email to end with '@email.com', but got: %v", result)
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
		name     string
		input    int
		validate func(t *testing.T, result string)
	}{
		{
			name:  "Generating Random String of Zero Length",
			input: 0,
			validate: func(t *testing.T, result string) {
				if result != "" {
					t.Errorf("Expected empty string, got %s", result)
				} else {
					t.Log("Success: Zero length string generated correctly.")
				}
			},
		},
		{
			name:  "Generating Random String of Length One",
			input: 1,
			validate: func(t *testing.T, result string) {
				if len(result) != 1 {
					t.Errorf("Expected string of length 1, got %d", len(result))
				} else {
					t.Log("Success: String of length 1 generated correctly.")
				}
			},
		},
		{
			name:  "Generating Random String of Arbitrary Length",
			input: 10,
			validate: func(t *testing.T, result string) {
				if len(result) != 10 {
					t.Errorf("Expected string of length 10, got %d", len(result))
				} else {
					t.Log("Success: String of length 10 generated correctly.")
				}
			},
		},
		{
			name:  "Randomness of the Generated Strings",
			input: 5,
			validate: func(t *testing.T, result string) {
				anotherResult := RandomString(5)
				if result == anotherResult {
					t.Errorf("Expected different strings, got identical strings: %s", result)
				} else {
					t.Log("Success: Different strings generated, confirming randomness.")
				}
			},
		},
		{
			name:  "Handling Large Length Requests",
			input: 1000,
			validate: func(t *testing.T, result string) {
				if len(result) != 1000 {
					t.Errorf("Expected string of length 1000, got %d", len(result))
				} else {
					t.Log("Success: String of length 1000 generated correctly.")
				}
			},
		},
		{
			name:  "Ensuring All Characters are from the Alphabet",
			input: 50,
			validate: func(t *testing.T, result string) {
				for _, c := range result {
					if !strings.ContainsRune(alphabet, c) {
						t.Errorf("Expected characters from alphabet, got %c", c)
						return
					}
				}
				t.Log("Success: All characters are from the alphabet.")
			},
		},
		{
			name:  "Consistency with Seeded Random Number Generator",
			input: 10,
			validate: func(t *testing.T, result string) {
				rand.Seed(42)
				expected := RandomString(10)
				if result != expected {
					t.Errorf("Expected %s, got %s", expected, result)
				} else {
					t.Log("Success: Consistent result with seeded random number generator.")
				}
			},
		},
		{
			name:  "Handling Negative Length Requests",
			input: -5,
			validate: func(t *testing.T, result string) {
				if result != "" {
					t.Errorf("Expected empty string for negative length, got %s", result)
				} else {
					t.Log("Success: Handled negative length request gracefully.")
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RandomString(tt.input)
			tt.validate(t, result)
		})
	}
}

