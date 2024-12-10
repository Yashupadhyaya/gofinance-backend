package util

import (
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

	tests := []struct {
		name          string
		length        int
		expectedError bool
	}{
		{
			name:          "Normal Operation with Standard Length",
			length:        10,
			expectedError: false,
		},
		{
			name:          "Edge Case with Zero Length",
			length:        0,
			expectedError: false,
		},
		{
			name:          "Typical Length of Local Part",
			length:        15,
			expectedError: false,
		},
		{
			name:          "Maximum Length for Local Part",
			length:        64,
			expectedError: false,
		},
		{
			name:          "Negative Length Input",
			length:        -5,
			expectedError: true,
		},
		{
			name:          "Randomness Validation",
			length:        10,
			expectedError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			email := RandomEmail(tt.length)

			if tt.expectedError {
				if email != "" {
					t.Errorf("Expected error but got email: %s", email)
				}
				return
			}

			localPart := strings.Split(email, "@")[0]
			if tt.name == "Randomness Validation" {
				email1 := RandomEmail(tt.length)
				email2 := RandomEmail(tt.length)
				if email1 == email2 {
					t.Errorf("Expected different emails but got same: %s and %s", email1, email2)
				}
				return
			}

			if len(localPart) != tt.length {
				t.Errorf("Expected local part length %d but got %d", tt.length, len(localPart))
			}

			if !strings.HasSuffix(email, "@email.com") {
				t.Errorf("Expected suffix @email.com but got %s", email)
			}

			t.Logf("Test %s passed with email: %s", tt.name, email)
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
		name          string
		inputLength   int
		expectedEmpty bool
		expectedLen   int
		checkLower    bool
		checkRandom   bool
		seed          int64
		expectedStr   string
	}{
		{
			name:          "Zero Length",
			inputLength:   0,
			expectedEmpty: true,
		},
		{
			name:        "Specific Length",
			inputLength: 10,
			expectedLen: 10,
		},
		{
			name:        "All Lowercase Letters",
			inputLength: 15,
			checkLower:  true,
		},
		{
			name:        "Maximum Length",
			inputLength: 1000000,
			expectedLen: 1000000,
		},
		{
			name:        "Consistency of Randomness",
			inputLength: 10,
			checkRandom: true,
		},
		{
			name:        "Seeded Random Number Generator",
			inputLength: 10,
			seed:        42,
			expectedStr: "expected_string_for_seed_42",
		},
		{
			name:        "Performance Test for Large Inputs",
			inputLength: 1000000,
			expectedLen: 1000000,
		},
		{
			name:          "Non-Positive Length",
			inputLength:   -5,
			expectedEmpty: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.seed != 0 {
				rand.Seed(tt.seed)
			}

			result := RandomString(tt.inputLength)

			if tt.expectedEmpty && result != "" {
				t.Errorf("expected empty string, got %v", result)
				t.Logf("Test %s failed: Expected empty string for input length %d", tt.name, tt.inputLength)
			}

			if tt.expectedLen > 0 && len(result) != tt.expectedLen {
				t.Errorf("expected string length %d, got %d", tt.expectedLen, len(result))
				t.Logf("Test %s failed: Expected length %d, got %d", tt.name, tt.expectedLen, len(result))
			}

			if tt.checkLower {
				for _, char := range result {
					if !strings.ContainsRune(alphabet, char) {
						t.Errorf("expected only lowercase letters, got %v", result)
						t.Logf("Test %s failed: Expected only lowercase letters, got %v", tt.name, result)
						break
					}
				}
			}

			if tt.checkRandom {
				result2 := RandomString(tt.inputLength)
				if result == result2 {
					t.Errorf("expected different strings, got %v and %v", result, result2)
					t.Logf("Test %s failed: Expected different strings for consecutive calls", tt.name)
				}
			}

			if tt.seed != 0 && result != tt.expectedStr {
				t.Errorf("expected string %v, got %v", tt.expectedStr, result)
				t.Logf("Test %s failed: Expected string %v for seed %d, got %v", tt.name, tt.expectedStr, tt.seed, result)
			}
		})
	}
}

