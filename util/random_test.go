package util

import (
	"fmt"
	"strings"
	"testing"
	"math/rand"
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
			name:           "Generate Email with a Standard Length",
			input:          10,
			expectedLength: 10,
			expectedFormat: "xxxxxxxxxx@email.com",
		},
		{
			name:           "Generate Email with Minimum Length",
			input:          1,
			expectedLength: 1,
			expectedFormat: "x@email.com",
		},
		{
			name:           "Generate Email with Maximum Length",
			input:          64,
			expectedLength: 64,
			expectedFormat: strings.Repeat("x", 64) + "@email.com",
		},
		{
			name:           "Generate Email with Zero Length",
			input:          0,
			expectedLength: 0,
			expectedFormat: "@email.com",
		},
		{
			name:           "Generate Email with Negative Length",
			input:          -5,
			expectedLength: 0,
			expectedFormat: "@email.com",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			email := RandomEmail(tt.input)

			localPart := strings.Split(email, "@")[0]
			if len(localPart) != tt.expectedLength {
				t.Errorf("expected local part length %d, got %d", tt.expectedLength, len(localPart))
			}

			if tt.expectedFormat != "" && strings.Count(email, "@") != 1 {
				t.Errorf("expected format %s, got %s", tt.expectedFormat, email)
			}

			t.Logf("Test '%s' passed with generated email: %s", tt.name, email)
		})
	}
}

/*
ROOST_METHOD_HASH=RandomString_d7e3599ac4
ROOST_METHOD_SIG_HASH=RandomString_c6fe4ad19a


 */
func TestRandomString(t *testing.T) {

	alphabet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	rand.Seed(time.Now().UnixNano())

	tests := []struct {
		name           string
		inputLength    int
		expectedLength int
		expectEmpty    bool
	}{
		{
			name:           "Positive Length",
			inputLength:    10,
			expectedLength: 10,
			expectEmpty:    false,
		},
		{
			name:           "Zero Length",
			inputLength:    0,
			expectedLength: 0,
			expectEmpty:    true,
		},
		{
			name:           "Negative Length",
			inputLength:    -5,
			expectedLength: 0,
			expectEmpty:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			result := RandomString(tt.inputLength)

			if len(result) != tt.expectedLength {
				t.Errorf("expected length %d, got %d", tt.expectedLength, len(result))
			}

			if tt.expectEmpty && result != "" {
				t.Errorf("expected empty string, got %s", result)
			}
		})
	}

	t.Run("Consistent Character Set Usage", func(t *testing.T) {
		length := 100
		result := RandomString(length)

		for _, char := range result {
			if !strings.ContainsRune(alphabet, char) {
				t.Errorf("character %c not found in alphabet", char)
			}
		}
	})

	t.Run("Randomness Check", func(t *testing.T) {
		length := 10
		set := make(map[string]bool)

		for i := 0; i < 100; i++ {
			str := RandomString(length)
			if set[str] {
				t.Errorf("duplicate string found: %s", str)
			}
			set[str] = true
		}
	})
}

