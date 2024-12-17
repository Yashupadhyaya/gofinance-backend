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

	tests := []struct {
		name           string
		input          int
		expectedLength int
		expectedRegex  string
		expectedError  bool
	}{
		{
			name:           "Normal Operation with Valid Input",
			input:          10,
			expectedLength: 10,
			expectedRegex:  `^[a-z]{10}@email\.com$`,
			expectedError:  false,
		},
		{
			name:           "Edge Case with Zero Length",
			input:          0,
			expectedLength: 0,
			expectedRegex:  `^@email\.com$`,
			expectedError:  false,
		},
		{
			name:           "Edge Case with Maximum Length",
			input:          1024,
			expectedLength: 1024,
			expectedRegex:  `^[a-z]{1024}@email\.com$`,
			expectedError:  false,
		},
		{
			name:           "Normal Operation with Minimum Length",
			input:          1,
			expectedLength: 1,
			expectedRegex:  `^[a-z]{1}@email\.com$`,
			expectedError:  false,
		},
		{
			name:           "Consistent Email Format",
			input:          5,
			expectedLength: 5,
			expectedRegex:  `^[a-z]{5}@email\.com$`,
			expectedError:  false,
		},
		{
			name:           "Randomness Check",
			input:          8,
			expectedLength: 8,
			expectedRegex:  `^[a-z]{8}@email\.com$`,
			expectedError:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("Running test: %s", tt.name)
			result := RandomEmail(tt.input)

			if tt.expectedError {
				if result != "" {
					t.Errorf("expected an error but got result: %s", result)
				}
				return
			}

			localPart := strings.Split(result, "@")[0]
			if len(localPart) != tt.expectedLength {
				t.Errorf("expected local part length: %d, got: %d", tt.expectedLength, len(localPart))
			}

			matched, err := regexp.MatchString(tt.expectedRegex, result)
			if err != nil {
				t.Fatalf("regex match error: %v", err)
			}
			if !matched {
				t.Errorf("email format does not match regex: %s, got: %s", tt.expectedRegex, result)
			}

			if tt.name == "Randomness Check" {
				otherResult := RandomEmail(tt.input)
				if result == otherResult {
					t.Errorf("expected different results for successive calls, got same: %s", result)
				}
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
		validate func(t *testing.T, result string)
	}{
		{
			name:   "Generating Random String of Zero Length",
			length: 0,
			validate: func(t *testing.T, result string) {
				if result != "" {
					t.Errorf("Expected empty string, got %s", result)
				}
			},
		},
		{
			name:   "Generating Random String of Positive Length",
			length: 10,
			validate: func(t *testing.T, result string) {
				if len(result) != 10 {
					t.Errorf("Expected string of length 10, got %d", len(result))
				}
			},
		},
		{
			name:   "Consistency of Random String Length",
			length: 5,
			validate: func(t *testing.T, result string) {
				if len(result) != 5 {
					t.Errorf("Expected string of length 5, got %d", len(result))
				}
			},
		},
		{
			name:   "Character Inclusion in Alphabet",
			length: 100,
			validate: func(t *testing.T, result string) {
				for _, char := range result {
					if !strings.Contains(alphabet, string(char)) {
						t.Errorf("Character %c not in alphabet", char)
					}
				}
			},
		},
		{
			name:   "Handling Large Length Requests",
			length: 1000000,
			validate: func(t *testing.T, result string) {
				if len(result) != 1000000 {
					t.Errorf("Expected string of length 1000000, got %d", len(result))
				}
			},
		},
		{
			name:   "Randomness of Output",
			length: 10,
			validate: func(t *testing.T, result string) {

				anotherResult := RandomString(10)
				if result == anotherResult {
					t.Errorf("Expected different strings, got identical strings: %s and %s", result, anotherResult)
				}
			},
		},
		{
			name:   "Seeding the Random Number Generator",
			length: 10,
			validate: func(t *testing.T, result string) {
				rand.Seed(1)
				expectedResult := RandomString(10)
				rand.Seed(1)
				actualResult := RandomString(10)
				if expectedResult != actualResult {
					t.Errorf("Expected %s, got %s", expectedResult, actualResult)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("Running test: %s", tt.name)
			result := RandomString(tt.length)
			tt.validate(t, result)
		})
	}
}

