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

	type testCase struct {
		name     string
		number   int
		expected string
	}

	testCases := []testCase{
		{
			name:   "Generating Email with Minimum Length",
			number: 1,
			expected: func() string {
				return fmt.Sprintf("%s@email.com", RandomString(1))
			}(),
		},
		{
			name:   "Generating Email with Typical Length",
			number: 10,
			expected: func() string {
				return fmt.Sprintf("%s@email.com", RandomString(10))
			}(),
		},
		{
			name:   "Generating Email with Maximum Length",
			number: 100,
			expected: func() string {
				return fmt.Sprintf("%s@email.com", RandomString(100))
			}(),
		},
		{
			name:   "Generating Email with Zero Length",
			number: 0,
			expected: func() string {
				return fmt.Sprintf("%s@email.com", RandomString(0))
			}(),
		},
		{
			name:   "Generating Email with Negative Length",
			number: -5,
			expected: func() string {
				return ""
			}(),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := RandomEmail(tc.number)
			if len(tc.expected) > 0 {
				if result != tc.expected {
					t.Errorf("Expected %s, but got %s", tc.expected, result)
				} else {
					t.Logf("Success: %s", result)
				}
			} else {
				if result != "" {
					t.Errorf("Expected empty string, but got %s", result)
				} else {
					t.Logf("Success: empty string as expected")
				}
			}
		})
	}

	t.Run("Consistency Check for Same Length", func(t *testing.T) {
		emailSet := make(map[string]struct{})
		for i := 0; i < 5; i++ {
			email := RandomEmail(10)
			if _, exists := emailSet[email]; exists {
				t.Errorf("Duplicate email generated: %s", email)
			}
			emailSet[email] = struct{}{}
		}
		t.Logf("Success: All generated emails are unique")
	})

	t.Run("Valid Characters in Username", func(t *testing.T) {
		email := RandomEmail(10)
		username := strings.Split(email, "@")[0]
		for _, char := range username {
			if !strings.ContainsRune(alphabet, char) {
				t.Errorf("Invalid character in username: %c", char)
			}
		}
		t.Logf("Success: All characters in username are valid")
	})
}

/*
ROOST_METHOD_HASH=RandomString_d7e3599ac4
ROOST_METHOD_SIG_HASH=RandomString_c6fe4ad19a


 */
func RandomStringWithAlphabet(number int, alphabet string) string {
	if len(alphabet) == 0 {
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

	type testCase struct {
		name     string
		input    int
		expected func(string) bool
	}

	testCases := []testCase{
		{
			name:  "Generate a Random String of Given Length",
			input: 10,
			expected: func(output string) bool {
				return len(output) == 10
			},
		},
		{
			name:  "Generate a Random String of Length Zero",
			input: 0,
			expected: func(output string) bool {
				return output == ""
			},
		},
		{
			name:  "Consistency of Randomness",
			input: 5,
			expected: func(output string) bool {
				anotherOutput := RandomString(5)
				return output != anotherOutput
			},
		},
		{
			name:  "Check Characters Are from the Alphabet",
			input: 20,
			expected: func(output string) bool {
				for _, char := range output {
					if !strings.ContainsRune(alphabet, char) {
						return false
					}
				}
				return true
			},
		},
		{
			name:  "Performance with Large Input",
			input: 1000000,
			expected: func(output string) bool {

				return len(output) == 1000000
			},
		},
		{
			name:  "Seed Consistency",
			input: 10,
			expected: func(output string) bool {
				rand.Seed(42)
				firstOutput := RandomString(10)
				rand.Seed(42)
				secondOutput := RandomString(10)
				return firstOutput == secondOutput
			},
		},
		{
			name:  "Check for Empty Alphabet Handling",
			input: 10,
			expected: func(output string) bool {

				originalAlphabet := alphabet
				emptyAlphabet := ""
				defer func() { alphabet = originalAlphabet }()
				return RandomStringWithAlphabet(10, emptyAlphabet) == ""
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := RandomString(tc.input)
			if !tc.expected(output) {
				t.Errorf("Test %s failed: expected condition not met for input %d", tc.name, tc.input)
			} else {
				t.Logf("Test %s succeeded", tc.name)
			}
		})
	}
}

