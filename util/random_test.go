package util

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"testing"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"
const defaultAlphabet = "abcdefghijklmnopqrstuvwxyz"/*
ROOST_METHOD_HASH=RandomEmail_1905439733
ROOST_METHOD_SIG_HASH=RandomEmail_7a04f189fd


 */
func RandomEmail(number int) string {
	if number < 0 {
		return "@email.com"
	}
	return fmt.Sprintf("%s@email.com", RandomString(number))
}

func RandomString(number int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < number; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func TestRandomEmail(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	tests := []struct {
		name         string
		localPartLen int
		expectedLen  int
		expectErr    bool
	}{
		{
			name:         "Generate Email with 5-Character Local Part",
			localPartLen: 5,
			expectedLen:  5,
			expectErr:    false,
		},
		{
			name:         "Generate Email with 0-Character Local Part",
			localPartLen: 0,
			expectedLen:  0,
			expectErr:    false,
		},
		{
			name:         "Generate Email with Maximum Character Local Part",
			localPartLen: 1000,
			expectedLen:  1000,
			expectErr:    false,
		},
		{
			name:         "Generate Email with Negative Length",
			localPartLen: -1,
			expectedLen:  0,
			expectErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			email := RandomEmail(tt.localPartLen)

			if !tt.expectErr {
				matched, err := regexp.MatchString(`^[a-zA-Z]*@email\.com$`, email)
				if err != nil || !matched {
					t.Errorf("RandomEmail() = %v, want match for email format", email)
				}

				localPart := strings.Split(email, "@")[0]
				if len(localPart) != tt.expectedLen {
					t.Errorf("Local part length = %d, want %d", len(localPart), tt.expectedLen)
				}
			} else {
				if email != "@email.com" {
					t.Errorf("Expected default or error behavior, got %v", email)
				}
			}
		})
	}

	t.Run("Validate Email Format with Special Characters", func(t *testing.T) {
		email := RandomEmail(rand.Intn(100))
		matched, err := regexp.MatchString(`^[a-zA-Z]*@email\.com$`, email)
		if err != nil || !matched {
			t.Errorf("RandomEmail() = %v, want match for email format", email)
		}
	})

	t.Run("Generate Multiple Emails with Same Length", func(t *testing.T) {
		emailSet := make(map[string]struct{})
		for i := 0; i < 100; i++ {
			email := RandomEmail(10)
			if _, exists := emailSet[email]; exists {
				t.Errorf("Duplicate email found: %v", email)
			}
			emailSet[email] = struct{}{}
		}
	})

	t.Run("Generate Email with Non-Alphabetic Characters in Local Part", func(t *testing.T) {
		email := RandomEmail(10)
		localPart := strings.Split(email, "@")[0]
		for _, char := range localPart {
			if !strings.ContainsRune(alphabet, char) {
				t.Errorf("Local part contains non-alphabetic character: %c", char)
			}
		}
	})
}

/*
ROOST_METHOD_HASH=RandomString_d7e3599ac4
ROOST_METHOD_SIG_HASH=RandomString_c6fe4ad19a


 */
func RandomString(number int, alphabet string) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < number; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func TestRandomString(t *testing.T) {
	type testCase struct {
		name           string
		input          int
		expectedLength int
		seed           int64
		shouldSeed     bool
		shouldBeUnique bool
		alphabet       string
	}

	nonASCIIAlphabet := "abcdefghijklmnopqrstuvwxyzñáéíóúü"

	testCases := []testCase{
		{
			name:           "Zero Length",
			input:          0,
			expectedLength: 0,
		},
		{
			name:           "Length One",
			input:          1,
			expectedLength: 1,
		},
		{
			name:           "Specific Length",
			input:          10,
			expectedLength: 10,
		},
		{
			name:           "Upper Bound Length",
			input:          1000,
			expectedLength: 1000,
		},
		{
			name:           "Randomness Check",
			input:          10,
			expectedLength: 10,
			shouldBeUnique: true,
		},
		{
			name:           "Seeded Random Generator",
			input:          10,
			expectedLength: 10,
			seed:           42,
			shouldSeed:     true,
		},
		{
			name:           "Non-ASCII Alphabet",
			input:          10,
			expectedLength: 10,
			alphabet:       nonASCIIAlphabet,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.shouldSeed {
				rand.Seed(tc.seed)
			} else {
				rand.Seed(time.Now().UnixNano())
			}

			alphabet := defaultAlphabet
			if tc.alphabet != "" {
				alphabet = tc.alphabet
			}

			result := RandomString(tc.input, alphabet)
			if len(result) != tc.expectedLength {
				t.Errorf("Test %s failed: expected length %d, got %d", tc.name, tc.expectedLength, len(result))
			}

			if tc.shouldBeUnique {

				anotherResult := RandomString(tc.input, alphabet)
				if anotherResult == result {
					t.Errorf("Test %s failed: expected unique strings, got %s and %s", tc.name, result, anotherResult)
				}
			}

			if tc.shouldSeed {
				expectedResult := "jvgqfytthr"
				if result != expectedResult {
					t.Errorf("Test %s failed: expected %s, got %s", tc.name, expectedResult, result)
				}
			}

			if tc.alphabet != "" {
				for _, char := range result {
					if !strings.ContainsRune(tc.alphabet, char) {
						t.Errorf("Test %s failed: character %c not in custom alphabet", tc.name, char)
					}
				}
			}

			t.Logf("Test %s passed: generated string %s", tc.name, result)
		})
	}
}

