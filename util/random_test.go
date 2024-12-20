package util

import (
	"math/rand"
	"strings"
	"testing"
	"time"
)


/*
ROOST_METHOD_HASH=RandomString_d7e3599ac4
ROOST_METHOD_SIG_HASH=RandomString_c6fe4ad19a


 */
func TestRandomString(t *testing.T) {

	rand.Seed(time.Now().UnixNano())

	tests := []struct {
		name     string
		input    int
		expected int
		validate func(string) bool
	}{
		{
			name:     "Generate a Random String of Specified Length",
			input:    10,
			expected: 10,
			validate: func(s string) bool {
				return len(s) == 10
			},
		},
		{
			name:     "Generate a Random String with Zero Length",
			input:    0,
			expected: 0,
			validate: func(s string) bool {
				return s == ""
			},
		},
		{
			name:     "Generate a Random String with a Negative Length",
			input:    -1,
			expected: 0,
			validate: func(s string) bool {
				return s == ""
			},
		},
		{
			name:     "Verify Randomness of Generated Strings",
			input:    10,
			expected: 10,
			validate: func(s string) bool {
				s1 := RandomString(10)
				s2 := RandomString(10)
				return s1 != s2
			},
		},
		{
			name:     "Verify Characters in Generated String are from Alphabet",
			input:    10,
			expected: 10,
			validate: func(s string) bool {
				for _, c := range s {
					if !strings.ContainsRune(alphabet, c) {
						return false
					}
				}
				return true
			},
		},
		{
			name:     "Generate a Random String with a Large Length",
			input:    10000,
			expected: 10000,
			validate: func(s string) bool {
				return len(s) == 10000
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("Running test: %s", tt.name)
			result := RandomString(tt.input)
			if len(result) != tt.expected {
				t.Errorf("Expected length %d, but got %d", tt.expected, len(result))
			}
			if !tt.validate(result) {
				t.Errorf("Validation failed for test: %s", tt.name)
			}
		})
	}
}

