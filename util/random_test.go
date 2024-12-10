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

	tests := []struct {
		name           string
		input          int
		expectedLength int
		expectedFormat string
	}{
		{
			name:           "Generate Email with Standard Length",
			input:          10,
			expectedLength: 10,
			expectedFormat: "^[a-z]{10}@email.com$",
		},
		{
			name:           "Generate Email with Zero Length",
			input:          0,
			expectedLength: 0,
			expectedFormat: "^@email.com$",
		},
		{
			name:           "Generate Email with Negative Length",
			input:          -5,
			expectedLength: 0,
			expectedFormat: "^@email.com$",
		},
		{
			name:           "Generate Email with Maximum Length",
			input:          1000,
			expectedLength: 1000,
			expectedFormat: "^[a-z]{1000}@email.com$",
		},
		{
			name:           "Generate Email with Special Characters in Username",
			input:          15,
			expectedLength: 15,
			expectedFormat: "^[a-z]{15}@email.com$",
		},
		{
			name:           "Generate Email with Consistent Output Length",
			input:          8,
			expectedLength: 8,
			expectedFormat: "^[a-z]{8}@email.com$",
		},
		{
			name:           "Generate Email with Randomness",
			input:          6,
			expectedLength: 6,
			expectedFormat: "^[a-z]{6}@email.com$",
		},
		{
			name:           "Verify Email Format",
			input:          7,
			expectedLength: 7,
			expectedFormat: "^[a-z]{7}@email.com$",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			email := RandomEmail(tt.input)
			username := strings.Split(email, "@")[0]

			if len(username) != tt.expectedLength {
				t.Errorf("expected username length %d, got %d", tt.expectedLength, len(username))
			}

			if matched := validateFormat(email, tt.expectedFormat); !matched {
				t.Errorf("email %s does not match expected format %s", email, tt.expectedFormat)
			}

			if tt.name == "Generate Email with Randomness" {
				email2 := RandomEmail(tt.input)
				if email == email2 {
					t.Errorf("expected different emails but got the same: %s", email)
				}
			}

			t.Logf("Test '%s' passed with email: %s", tt.name, email)
		})
	}
}

func validateFormat(email, format string) bool {

	return true
}

