package util

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"
/*
ROOST_METHOD_HASH=RandomEmail_1905439733
ROOST_METHOD_SIG_HASH=RandomEmail_7a04f189fd


 */
func RandomEmail(number int) string {
	return fmt.Sprintf("%s@email.com", RandomString(number))
}

func RandomString(number int) string {
	const alphabet = "abcdefghijklmnopqrstuvwxyz"
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < number; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func TestRandomEmail(t *testing.T) {

	tests := []struct {
		name           string
		input          int
		expectedLength int
		expectedDomain string
	}{
		{
			name:           "Generating Email with Standard Length",
			input:          10,
			expectedLength: 10,
			expectedDomain: "@email.com",
		},
		{
			name:           "Generating Email with Minimum Length",
			input:          1,
			expectedLength: 1,
			expectedDomain: "@email.com",
		},
		{
			name:           "Generating Email with Maximum Length",
			input:          100,
			expectedLength: 100,
			expectedDomain: "@email.com",
		},
		{
			name:           "Generating Email with Zero Length",
			input:          0,
			expectedLength: 0,
			expectedDomain: "@email.com",
		},
		{
			name:           "Generating Email with Upper Bound Length",
			input:          64,
			expectedLength: 64,
			expectedDomain: "@email.com",
		},
	}

	rand.Seed(time.Now().UnixNano())

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RandomEmail(tt.input)
			localPart := strings.Split(result, "@")[0]
			domainPart := strings.Split(result, "@")[1]

			if len(localPart) != tt.expectedLength {
				t.Errorf("expected local part length %d, got %d", tt.expectedLength, len(localPart))
			}

			if domainPart != tt.expectedDomain[1:] {
				t.Errorf("expected domain %s, got %s", tt.expectedDomain, domainPart)
			}

			t.Logf("Test case '%s' passed with result: %s", tt.name, result)
		})
	}

	t.Run("Generating Multiple Emails and Ensuring Uniqueness", func(t *testing.T) {
		emailSet := make(map[string]struct{})
		for i := 0; i < 100; i++ {
			email := RandomEmail(10)
			if _, exists := emailSet[email]; exists {
				t.Errorf("duplicate email generated: %s", email)
			}
			emailSet[email] = struct{}{}
		}
		t.Log("All generated emails are unique")
	})
}

/*
ROOST_METHOD_HASH=RandomString_d7e3599ac4
ROOST_METHOD_SIG_HASH=RandomString_c6fe4ad19a


 */
func RandomString(number int) string {
	if number <= 0 {
		return ""
	}
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

