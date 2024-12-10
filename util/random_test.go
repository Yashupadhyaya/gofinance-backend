package util

import (
	"math/rand"
	"strings"
	"testing"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"
const alphabet = "abcdefghijklmnopqrstuvwxyz"
var sb strings.Builder
/*
ROOST_METHOD_HASH=RandomEmail_1905439733
ROOST_METHOD_SIG_HASH=RandomEmail_7a04f189fd


 */
func RandomEmail(number int) string {
	return fmt.Sprintf("%s@email.com", RandomString(number))
}

func RandomString(n int) string {
	sb.Reset()
	for i := 0; i < n; i++ {
		index := rand.Intn(len(alphabet))
		sb.WriteByte(alphabet[index])
	}
	return sb.String()
}

func TestRandomEmail(t *testing.T) {
	t.Parallel()

	rand.Seed(time.Now().UnixNano())

	type testCase struct {
		name           string
		input          int
		expectedPrefix string
		expectedLength int
		shouldBeUnique bool
		expectedOutput string
	}

	testCases := []testCase{
		{
			name:           "Positive Number",
			input:          10,
			expectedLength: 10,
			shouldBeUnique: false,
		},
		{
			name:           "Zero Length",
			input:          0,
			expectedOutput: "@email.com",
		},
		{
			name:           "Negative Number",
			input:          -1,
			expectedOutput: "@email.com",
		},
		{
			name:           "Large Number",
			input:          1000,
			expectedLength: 1000,
			shouldBeUnique: false,
		},
		{
			name:           "Uniqueness Test",
			input:          10,
			shouldBeUnique: true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			if tc.shouldBeUnique {
				emails := make(map[string]struct{})
				for i := 0; i < 100; i++ {
					email := RandomEmail(tc.input)
					if _, exists := emails[email]; exists {
						t.Errorf("Duplicate email found: %s", email)
					}
					emails[email] = struct{}{}
				}
				t.Log("All generated emails are unique.")
			} else {
				email := RandomEmail(tc.input)
				localPart := strings.Split(email, "@")[0]

				if tc.expectedLength > 0 {
					if len(localPart) != tc.expectedLength {
						t.Errorf("Expected local part of length %d, got %d", tc.expectedLength, len(localPart))
					} else {
						t.Logf("Email local part length is correct: %d", tc.expectedLength)
					}
				}

				if tc.expectedOutput != "" {
					if email != tc.expectedOutput {
						t.Errorf("Expected email %s, got %s", tc.expectedOutput, email)
					} else {
						t.Logf("Email output is correct: %s", tc.expectedOutput)
					}
				}
			}
		})
	}
}

/*
ROOST_METHOD_HASH=RandomString_d7e3599ac4
ROOST_METHOD_SIG_HASH=RandomString_c6fe4ad19a


 */
func RandomString(number int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < number; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func TestRandomString(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	tests := []struct {
		name           string
		input          int
		expectedLength int
		verifyRandom   bool
	}{
		{
			name:           "Generate a Random String of Specified Length",
			input:          10,
			expectedLength: 10,
			verifyRandom:   false,
		},
		{
			name:           "Generate a Random String with Zero Length",
			input:          0,
			expectedLength: 0,
			verifyRandom:   false,
		},
		{
			name:           "Generate a Random String with Negative Length",
			input:          -5,
			expectedLength: 0,
			verifyRandom:   false,
		},
		{
			name:           "Generate a Random String with Large Length",
			input:          10000,
			expectedLength: 10000,
			verifyRandom:   false,
		},
		{
			name:           "Verify Randomness of Generated Strings",
			input:          10,
			expectedLength: 10,
			verifyRandom:   true,
		},
		{
			name:           "Verify Characters Are from Defined Alphabet",
			input:          50,
			expectedLength: 50,
			verifyRandom:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RandomString(tt.input)

			if len(result) != tt.expectedLength {
				t.Errorf("Expected length %d, got %d", tt.expectedLength, len(result))
			}

			if tt.verifyRandom {
				anotherResult := RandomString(tt.input)
				if result == anotherResult {
					t.Errorf("Expected different strings on consecutive calls, got identical strings")
				}
			}

			if tt.name == "Verify Characters Are from Defined Alphabet" {
				for _, char := range result {
					if !strings.Contains(alphabet, string(char)) {
						t.Errorf("Character %c is not in the defined alphabet", char)
					}
				}
			}

			t.Logf("Test '%s' passed.", tt.name)
		})
	}

}

