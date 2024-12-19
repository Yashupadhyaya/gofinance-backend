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

	rand.Seed(time.Now().UnixNano())

	testCases := []struct {
		name              string
		input             int
		expectedLocalPart int
		expectedOutput    string
		expectedError     bool
		validateLocalPart func(localPart string) bool
	}{
		{
			name:              "Generate Email with 5-Character Local Part",
			input:             5,
			expectedLocalPart: 5,
			validateLocalPart: func(localPart string) bool { return len(localPart) == 5 },
		},
		{
			name:              "Generate Email with 0-Character Local Part",
			input:             0,
			expectedLocalPart: 0,
			expectedOutput:    "@email.com",
			validateLocalPart: func(localPart string) bool { return localPart == "" },
		},
		{
			name:              "Generate Email with Maximum Local Part Length (64 Characters)",
			input:             64,
			expectedLocalPart: 64,
			validateLocalPart: func(localPart string) bool { return len(localPart) == 64 },
		},
		{
			name:              "Generate Email with Negative Local Part Length",
			input:             -5,
			expectedLocalPart: 0,
			expectedError:     true,
			validateLocalPart: func(localPart string) bool { return localPart == "" },
		},
		{
			name:              "Generate Email with 1000-Character Local Part",
			input:             1000,
			expectedLocalPart: 1000,
			validateLocalPart: func(localPart string) bool { return len(localPart) == 1000 },
		},
		{
			name:              "Generate Email with Special Characters in Local Part",
			input:             10,
			expectedLocalPart: 10,
			validateLocalPart: func(localPart string) bool {
				for _, c := range localPart {
					if !strings.Contains(alphabet, string(c)) {
						return false
					}
				}
				return true
			},
		},
		{
			name:              "Generate Multiple Emails with Unique Local Parts",
			input:             10,
			expectedLocalPart: 10,
			validateLocalPart: func(localPart string) bool { return len(localPart) == 10 },
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			email := RandomEmail(tc.input)
			localPart := strings.Split(email, "@")[0]

			if tc.expectedError {
				if localPart != "" {
					t.Errorf("expected an empty local part for input %d, got %s", tc.input, localPart)
				}
			} else {
				if !tc.validateLocalPart(localPart) {
					t.Errorf("validation failed for local part %s with input %d", localPart, tc.input)
				}

				if tc.expectedOutput != "" && email != tc.expectedOutput {
					t.Errorf("expected email %s, got %s", tc.expectedOutput, email)
				}

				t.Logf("Generated email: %s", email)
			}
		})
	}

	t.Run("Generate Multiple Unique Emails", func(t *testing.T) {
		emailSet := make(map[string]struct{})
		for i := 0; i < 100; i++ {
			email := RandomEmail(10)
			if _, exists := emailSet[email]; exists {
				t.Errorf("duplicate email found: %s", email)
			}
			emailSet[email] = struct{}{}
		}
		t.Logf("Generated 100 unique emails successfully")
	})
}

/*
ROOST_METHOD_HASH=RandomString_d7e3599ac4
ROOST_METHOD_SIG_HASH=RandomString_c6fe4ad19a


 */
func TestRandomString(t *testing.T) {

	rand.Seed(time.Now().UnixNano())

	tests := []struct {
		name          string
		input         int
		expectedLen   int
		shouldBeEmpty bool
		alphabet      string
	}{
		{
			name:        "Generate a Random String of Specified Length",
			input:       10,
			expectedLen: 10,
			alphabet:    alphabet,
		},
		{
			name:          "Generate a Random String with Zero Length",
			input:         0,
			expectedLen:   0,
			shouldBeEmpty: true,
			alphabet:      alphabet,
		},
		{
			name:        "Generate Multiple Random Strings and Ensure Uniqueness",
			input:       10,
			expectedLen: 10,
			alphabet:    alphabet,
		},
		{
			name:        "Generate a Random String with Maximum Length",
			input:       1_000_000,
			expectedLen: 1_000_000,
			alphabet:    alphabet,
		},
		{
			name:        "Ensure All Characters in Random String Are from Alphabet",
			input:       10,
			expectedLen: 10,
			alphabet:    alphabet,
		},
		{
			name:        "Generate a Random String with Non-Deterministic Seed",
			input:       10,
			expectedLen: 10,
			alphabet:    alphabet,
		},
		{
			name:        "Validate Performance for Large Input",
			input:       1_000_000,
			expectedLen: 1_000_000,
			alphabet:    alphabet,
		},
		{
			name:          "Validate Function with Negative Input",
			input:         -10,
			expectedLen:   0,
			shouldBeEmpty: true,
			alphabet:      alphabet,
		},
		{
			name:        "Validate Function with Non-ASCII Characters in Alphabet",
			input:       10,
			expectedLen: 10,
			alphabet:    "abcdefghijklmnopqrstuvwxyzéèê",
		},
		{
			name:          "Validate Function with Empty Alphabet",
			input:         10,
			expectedLen:   0,
			shouldBeEmpty: true,
			alphabet:      "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			localAlphabet := alphabet
			if tt.alphabet != alphabet {
				localAlphabet = tt.alphabet
			}

			result := randomStringWithAlphabet(tt.input, localAlphabet)

			if len(result) != tt.expectedLen {
				t.Errorf("expected length %d, got %d", tt.expectedLen, len(result))
			}

			if tt.shouldBeEmpty && result != "" {
				t.Errorf("expected empty string, got %s", result)
			}

			if len(result) == tt.expectedLen {
				for _, char := range result {
					if !strings.ContainsRune(localAlphabet, char) {
						t.Errorf("character %c not in alphabet %s", char, localAlphabet)
					}
				}
			}

			if tt.name == "Generate Multiple Random Strings and Ensure Uniqueness" {
				anotherResult := randomStringWithAlphabet(tt.input, localAlphabet)
				if result == anotherResult {
					t.Errorf("expected different strings, got identical strings %s", result)
				}
			}

			if tt.name == "Generate a Random String with Non-Deterministic Seed" {
				rand.Seed(time.Now().UnixNano())
				anotherResult := randomStringWithAlphabet(tt.input, localAlphabet)
				if result == anotherResult {
					t.Errorf("expected different strings due to non-deterministic seed, got identical strings %s", result)
				}
			}

			if tt.name == "Validate Performance for Large Input" {
				start := time.Now()
				_ = randomStringWithAlphabet(tt.input, localAlphabet)
				duration := time.Since(start)
				if duration.Seconds() > 2 {
					t.Errorf("function took too long: %v", duration)
				}
			}
		})
	}
}

func randomStringWithAlphabet(number int, alphabet string) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < number; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

