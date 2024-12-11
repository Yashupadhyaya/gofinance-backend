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
		name        string
		input       int
		expectedLen int
		verifyFunc  func(t *testing.T, email string)
	}{
		{
			name:        "Generating Email with Standard Length",
			input:       10,
			expectedLen: 10,
			verifyFunc: func(t *testing.T, email string) {
				localPart := strings.Split(email, "@")[0]
				if len(localPart) != 10 {
					t.Errorf("Expected local part length 10, got %d", len(localPart))
				}
			},
		},
		{
			name:        "Generating Email with Zero Length",
			input:       0,
			expectedLen: 0,
			verifyFunc: func(t *testing.T, email string) {
				localPart := strings.Split(email, "@")[0]
				if len(localPart) != 0 {
					t.Errorf("Expected local part length 0, got %d", len(localPart))
				}
			},
		},
		{
			name:        "Generating Email with Maximum Length",
			input:       1000,
			expectedLen: 1000,
			verifyFunc: func(t *testing.T, email string) {
				localPart := strings.Split(email, "@")[0]
				if len(localPart) != 1000 {
					t.Errorf("Expected local part length 1000, got %d", len(localPart))
				}
			},
		},
		{
			name:        "Generating Email with Special Characters",
			input:       10,
			expectedLen: 10,
			verifyFunc: func(t *testing.T, email string) {
				localPart := strings.Split(email, "@")[0]
				for _, c := range localPart {
					if !strings.ContainsRune(alphabet, c) {
						t.Errorf("Expected only lowercase alphabetic characters, got %c", c)
					}
				}
			},
		},
		{
			name:        "Consistency Across Multiple Calls",
			input:       10,
			expectedLen: 10,
			verifyFunc: func(t *testing.T, email string) {
				emails := make(map[string]bool)
				for i := 0; i < 100; i++ {
					email := RandomEmail(10)
					if emails[email] {
						t.Errorf("Duplicate email found: %s", email)
					}
					emails[email] = true
				}
			},
		},
		{
			name:        "Boundary Testing with Negative Input",
			input:       -1,
			expectedLen: 0,
			verifyFunc: func(t *testing.T, email string) {
				localPart := strings.Split(email, "@")[0]
				if len(localPart) != 0 {
					t.Errorf("Expected local part length 0 for negative input, got %d", len(localPart))
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			email := RandomEmail(tt.input)
			t.Logf("Generated email: %s", email)
			tt.verifyFunc(t, email)
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
		expected int
	}{
		{
			name:     "Generate a Random String of Given Length",
			length:   10,
			expected: 10,
		},
		{
			name:     "Generate a Random String with Zero Length",
			length:   0,
			expected: 0,
		},
		{
			name:     "Generate a Random String with Negative Length",
			length:   -5,
			expected: 0,
		},
		{
			name:     "Generate a Random String of Maximum Possible Length",
			length:   100000,
			expected: 100000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("Running test: %s", tt.name)
			result := RandomString(tt.length)
			if len(result) != tt.expected {
				t.Errorf("Expected string length %d, but got %d", tt.expected, len(result))
			}
		})
	}

	t.Run("Ensure Randomness of Generated String", func(t *testing.T) {
		t.Log("Running test: Ensure Randomness of Generated String")
		length := 10
		first := RandomString(length)
		second := RandomString(length)
		if first == second {
			t.Errorf("Expected different strings, but got the same: %s", first)
		}
	})

	t.Run("Verify Characters in Generated String", func(t *testing.T) {
		t.Log("Running test: Verify Characters in Generated String")
		length := 10
		result := RandomString(length)
		for _, char := range result {
			if !strings.ContainsRune(alphabet, char) {
				t.Errorf("Character %c is not in the defined alphabet", char)
			}
		}
	})

	t.Run("Consistency with Seeded Random Number Generator", func(t *testing.T) {
		t.Log("Running test: Consistency with Seeded Random Number Generator")
		rand.Seed(42)
		length := 10
		expected := RandomString(length)
		rand.Seed(42)
		result := RandomString(length)
		if result != expected {
			t.Errorf("Expected %s, but got %s with the same seed", expected, result)
		}
	})
}

