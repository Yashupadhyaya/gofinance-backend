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
	type testCase struct {
		name       string
		input      int
		expectFunc func(email string) bool
		expectDesc string
	}

	tests := []testCase{
		{
			name:  "Generate Email with 5-Character Local Part",
			input: 5,
			expectFunc: func(email string) bool {
				localPart := strings.Split(email, "@")[0]
				return len(localPart) == 5
			},
			expectDesc: "Expected local part of the email to be exactly 5 characters long",
		},
		{
			name:  "Generate Email with 0-Character Local Part",
			input: 0,
			expectFunc: func(email string) bool {
				localPart := strings.Split(email, "@")[0]
				return localPart == ""
			},
			expectDesc: "Expected local part of the email to be empty",
		},
		{
			name:  "Generate Email with Maximum Integer Local Part",
			input: math.MaxInt32,
			expectFunc: func(email string) bool {
				localPart := strings.Split(email, "@")[0]
				return len(localPart) == math.MaxInt32
			},
			expectDesc: "Expected local part of the email to be of length math.MaxInt32",
		},
		{
			name:  "Generate Email with Negative Length",
			input: -1,
			expectFunc: func(email string) bool {
				return RandomEmail(-1) == ""
			},
			expectDesc: "Expected function to handle negative input gracefully",
		},
		{
			name:  "Generate Email with Special Characters in Local Part",
			input: 10,
			expectFunc: func(email string) bool {
				localPart := strings.Split(email, "@")[0]
				for _, char := range localPart {
					if !strings.ContainsRune(alphabet, char) {
						return false
					}
				}
				return true
			},
			expectDesc: "Expected local part of the email to contain only alphabetic characters",
		},
		{
			name:  "Consistency of Email Format",
			input: 10,
			expectFunc: func(email string) bool {
				return strings.HasSuffix(email, "@email.com")
			},
			expectDesc: "Expected email to consistently have the format <local_part>@email.com",
		},
		{
			name:  "Randomness of Generated Local Part",
			input: 10,
			expectFunc: func(email string) bool {
				email1 := RandomEmail(10)
				email2 := RandomEmail(10)
				return email1 != email2
			},
			expectDesc: "Expected local parts of generated email addresses to be different",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			email := RandomEmail(tt.input)
			if !tt.expectFunc(email) {
				t.Errorf("Test failed: %s. Got email: %s", tt.expectDesc, email)
			} else {
				t.Logf("Test passed: %s. Got email: %s", tt.expectDesc, email)
			}
		})
	}
}

/*
ROOST_METHOD_HASH=RandomString_d7e3599ac4
ROOST_METHOD_SIG_HASH=RandomString_c6fe4ad19a


 */
func TestRandomString(t *testing.T) {
	tests := []struct {
		name           string
		input          int
		expectedLength int
		expectedOutput string
	}{
		{
			name:           "Generate a Random String of Specified Length",
			input:          10,
			expectedLength: 10,
		},
		{
			name:           "Generate a Random String with Zero Length",
			input:          0,
			expectedLength: 0,
		},
		{
			name:           "Generate a Random String with Negative Length",
			input:          -5,
			expectedLength: 0,
		},
		{
			name:           "Generate a Random String with Maximum Length of the Alphabet",
			input:          26,
			expectedLength: 26,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RandomString(tt.input)
			if len(got) != tt.expectedLength {
				t.Errorf("RandomString() = %v, want length %v", got, tt.expectedLength)
			}

			if tt.expectedLength > 0 {
				for _, char := range got {
					if !strings.ContainsRune(alphabet, char) {
						t.Errorf("RandomString() contains invalid character %v", char)
					}
				}
			}
		})
	}

	t.Run("Generate a Random String Multiple Times and Verify Uniqueness", func(t *testing.T) {
		const input = 10
		str1 := RandomString(input)
		str2 := RandomString(input)
		if str1 == str2 {
			t.Errorf("RandomString() generated same strings %v and %v", str1, str2)
		}
	})
}

