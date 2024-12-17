package util

import (
	"math/rand"
	"strings"
	"testing"
	"time"
	"unicode"
	"github.com/stretchr/testify/assert"
)

/*
ROOST_METHOD_HASH=RandomEmail_1905439733
ROOST_METHOD_SIG_HASH=RandomEmail_7a04f189fd


 */
func TestRandomEmail(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	tests := []struct {
		name          string
		input         int
		expectedLen   int
		expectedError bool
	}{
		{
			name:          "Generate Email with 5-Character Local Part",
			input:         5,
			expectedLen:   5,
			expectedError: false,
		},
		{
			name:          "Generate Email with 0-Character Local Part",
			input:         0,
			expectedLen:   0,
			expectedError: false,
		},
		{
			name:          "Generate Email with 1000-Character Local Part",
			input:         1000,
			expectedLen:   1000,
			expectedError: false,
		},
		{
			name:          "Generate Email with Random Lengths",
			input:         rand.Intn(100),
			expectedLen:   -1,
			expectedError: false,
		},
		{
			name:          "Generate Email with Negative Input",
			input:         -5,
			expectedLen:   0,
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			email := RandomEmail(tt.input)
			localPart := strings.Split(email, "@")[0]

			if tt.expectedError {
				if localPart != "" {
					t.Errorf("Expected an empty local part for negative input, got %s", localPart)
				}
			} else {
				if tt.expectedLen != -1 && len(localPart) != tt.expectedLen {
					t.Errorf("Expected local part length of %d, got %d", tt.expectedLen, len(localPart))
				}

				if !strings.HasSuffix(email, "@email.com") {
					t.Errorf("Expected email to end with '@email.com', got %s", email)
				}
			}

			t.Logf("Generated email: %s", email)

			for _, char := range localPart {
				if !strings.ContainsRune(alphabet, char) {
					t.Errorf("Character %c in local part is not in the allowed alphabet", char)
				}
			}
		})
	}

	t.Run("Generate Email and Check Uniqueness", func(t *testing.T) {
		emailSet := make(map[string]bool)
		for i := 0; i < 100; i++ {
			email := RandomEmail(10)
			if emailSet[email] {
				t.Errorf("Duplicate email generated: %s", email)
			}
			emailSet[email] = true
		}
	})
}

/*
ROOST_METHOD_HASH=RandomString_d7e3599ac4
ROOST_METHOD_SIG_HASH=RandomString_c6fe4ad19a


 */
func TestRandomString(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	tests := []struct {
		name           string
		inputLength    int
		expectedLength int
		expectedError  error
	}{
		{
			name:           "Generate a Random String of Specified Length",
			inputLength:    10,
			expectedLength: 10,
		},
		{
			name:           "Generate a Random String with Zero Length",
			inputLength:    0,
			expectedLength: 0,
		},
		{
			name:           "Generate a Random String with Negative Length",
			inputLength:    -1,
			expectedLength: 0,
		},
		{
			name:           "Generate a Random String with Maximum Length",
			inputLength:    100000,
			expectedLength: 100000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RandomString(tt.inputLength)

			assert.Equal(t, tt.expectedLength, len(result), "Expected length does not match")

			for _, char := range result {
				assert.True(t, unicode.IsLetter(char), "String contains non-alphabet characters")
			}
		})
	}

	t.Run("Verify Randomness of Generated Strings", func(t *testing.T) {
		str1 := RandomString(10)
		str2 := RandomString(10)
		assert.NotEqual(t, str1, str2, "Random strings should not be equal")
	})

	t.Run("Ensure Alphabet Characters in Generated String", func(t *testing.T) {
		result := RandomString(10)
		for _, char := range result {
			assert.Contains(t, alphabet, string(char), "String contains characters outside the defined alphabet")
		}
	})
}

