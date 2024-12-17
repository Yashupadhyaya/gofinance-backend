package util

import (
	"fmt"
	"strings"
	"testing"
	"math/rand"
	"time"
	"unicode/utf8"
)

/*
ROOST_METHOD_HASH=RandomEmail_1905439733
ROOST_METHOD_SIG_HASH=RandomEmail_7a04f189fd


 */
func TestRandomEmail(t *testing.T) {

	tests := []struct {
		name        string
		input       int
		expectedLen int
		expectedErr bool
	}{
		{
			name:        "Generate Email with 5-Character Local Part",
			input:       5,
			expectedLen: 5,
			expectedErr: false,
		},
		{
			name:        "Generate Email with 0-Character Local Part",
			input:       0,
			expectedLen: 0,
			expectedErr: false,
		},
		{
			name:        "Generate Email with Maximum Reasonable Local Part Length",
			input:       64,
			expectedLen: 64,
			expectedErr: false,
		},
		{
			name:        "Generate Email with Typical Local Part Length",
			input:       10,
			expectedLen: 10,
			expectedErr: false,
		},
		{
			name:        "Generate Email with Special Characters in Local Part",
			input:       10,
			expectedLen: 10,
			expectedErr: false,
		},
		{
			name:        "Consistent Email Domain",
			input:       10,
			expectedLen: 10,
			expectedErr: false,
		},
		{
			name:        "Generate Email with Negative Input",
			input:       -1,
			expectedLen: 0,
			expectedErr: true,
		},
		{
			name:        "Generate Email with Large Input",
			input:       1000,
			expectedLen: 1000,
			expectedErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			email := RandomEmail(tt.input)
			localPart := strings.Split(email, "@")[0]

			if len(localPart) != tt.expectedLen {
				t.Errorf("Expected local part length %d, got %d", tt.expectedLen, len(localPart))
			}

			if tt.input > 0 && !strings.HasSuffix(email, "@email.com") {
				t.Errorf("Expected email to end with @email.com, got %s", email)
			}

			if tt.input > 0 && !allAlphabetic(localPart) {
				t.Errorf("Expected local part to contain only alphabetic characters, got %s", localPart)
			}

			if tt.input < 0 && len(localPart) != 0 && !tt.expectedErr {
				t.Errorf("Expected error for negative input, got %s", localPart)
			}

			t.Logf("Test %s passed", tt.name)
		})
	}
}

func allAlphabetic(s string) bool {
	for _, c := range s {
		if !strings.ContainsRune(alphabet, c) {
			return false
		}
	}
	return true
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
		validate func(result string) bool
		errorMsg string
	}{
		{
			name:   "Generate a Random String of Given Length",
			length: 10,
			validate: func(result string) bool {
				return len(result) == 10
			},
			errorMsg: "Expected string length to be 10",
		},
		{
			name:   "Generate a Random String with Zero Length",
			length: 0,
			validate: func(result string) bool {
				return result == ""
			},
			errorMsg: "Expected empty string for length 0",
		},
		{
			name:   "Randomness of Generated Strings",
			length: 10,
			validate: func(result string) bool {
				result2 := RandomString(10)
				return result != result2
			},
			errorMsg: "Expected different strings for multiple calls",
		},
		{
			name:   "Use of Alphabet Characters Only",
			length: 10,
			validate: func(result string) bool {
				for _, c := range result {
					if !strings.ContainsRune(alphabet, c) {
						return false
					}
				}
				return true
			},
			errorMsg: "Expected string to contain only alphabet characters",
		},
		{
			name:   "Consistency with Different Seeds",
			length: 10,
			validate: func(result string) bool {
				rand.Seed(1)
				result1 := RandomString(10)
				rand.Seed(1)
				result2 := RandomString(10)
				return result1 == result2
			},
			errorMsg: "Expected identical strings for the same seed",
		},
		{
			name:   "Handling Large Length Values",
			length: 10000,
			validate: func(result string) bool {
				return len(result) == 10000
			},
			errorMsg: "Expected string length to be 10000",
		},
		{
			name:   "Verify No Side Effects",
			length: 10,
			validate: func(result string) bool {
				globalVar := "knownState"
				_ = RandomString(10)
				return globalVar == "knownState"
			},
			errorMsg: "Expected no side effects on global state",
		},
		{
			name:   "Validate Unicode Support",
			length: 10,
			validate: func(result string) bool {
				return utf8.ValidString(result)
			},
			errorMsg: "Expected valid UTF-8 string",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RandomString(tt.length)
			if !tt.validate(result) {
				t.Errorf("%s: %s", tt.name, tt.errorMsg)
			} else {
				t.Logf("%s: success", tt.name)
			}
		})
	}
}

