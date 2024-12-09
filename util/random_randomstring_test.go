package util

import (
	"math/rand"
	"strings"
	"testing"
	"time"
)

// Assume 'alphabet' is defined elsewhere and imported
// var alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func TestRandomString(t *testing.T) {
	// Seed the random number generator for consistency in testing
	rand.Seed(time.Now().UnixNano())

	tests := []struct {
		name        string
		length      int
		alphabet    string
		expectError bool
	}{
		{
			name:        "Zero Length String",
			length:      0,
			alphabet:    "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
			expectError: false,
		},
		{
			name:        "Positive Length String",
			length:      10,
			alphabet:    "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
			expectError: false,
		},
		{
			name:        "Large Length String",
			length:      10000,
			alphabet:    "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
			expectError: false,
		},
		{
			name:        "Undefined Alphabet",
			length:      10,
			alphabet:    "",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			alphabet = tt.alphabet // Assuming 'alphabet' is a global variable
			// Act
			result := RandomString(tt.length)

			// Assert
			if tt.expectError {
				if result != "" {
					t.Errorf("Expected an error or empty string, got: %s", result)
				}
				t.Logf("Handled undefined alphabet gracefully")
			} else {
				if len(result) != tt.length {
					t.Errorf("Expected length %d, got %d", tt.length, len(result))
				}
				for _, ch := range result {
					if !strings.ContainsRune(tt.alphabet, ch) {
						t.Errorf("Character %c is not in the alphabet", ch)
					}
				}
				t.Logf("Generated string: %s", result)
			}
		})
	}

	t.Run("Consistency of Randomness", func(t *testing.T) {
		alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
		str1 := RandomString(10)
		str2 := RandomString(10)
		if str1 == str2 {
			t.Errorf("Expected different strings, got identical: %s and %s", str1, str2)
		}
		t.Logf("Generated different strings: %s and %s", str1, str2)
	})
}

// TODO: Modify the alphabet variable import as per actual implementation details
// Note: This test assumes the 'alphabet' variable is defined and modifiable for testing purposes.
