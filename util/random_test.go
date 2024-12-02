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
func RandomEmailTest(number int) string {
	return strings.ToLower(fmt.Sprintf("%s@email.com", RandomStringTest(number)))
}

func RandomStringTest(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	sb := strings.Builder{}
	for i := 0; i < n; i++ {
		sb.WriteByte(letters[rand.Intn(len(letters))])
	}
	return sb.String()
}

func Testrandom_randomemail(t *testing.T) {

	rand.Seed(time.Now().UnixNano())

	type testCase struct {
		name     string
		number   int
		expected string
		assert   func(t *testing.T, result string, expected string)
	}

	testCases := []testCase{
		{
			name:   "Generate Email with a Standard Length",
			number: 10,
			assert: func(t *testing.T, result string, expected string) {
				if !strings.Contains(result, "@") || !strings.HasSuffix(result, "@email.com") {
					t.Errorf("Expected valid email format, got: %s", result)
				}
				t.Log("Passed: Generate Email with a Standard Length")
			},
		},
		{
			name:   "Generate Email with Zero Length",
			number: 0,
			assert: func(t *testing.T, result string, expected string) {
				if result != "@email.com" {
					t.Errorf("Expected '@email.com', got: %s", result)
				}
				t.Log("Passed: Generate Email with Zero Length")
			},
		},
		{
			name:   "Generate Email with Maximum Length",
			number: 256,
			assert: func(t *testing.T, result string, expected string) {
				if len(result) <= 256 {
					t.Errorf("Expected email length greater than 256, got: %d", len(result))
				}
				t.Log("Passed: Generate Email with Maximum Length")
			},
		},
		{
			name:   "Generate Email with Negative Length",
			number: -5,
			assert: func(t *testing.T, result string, expected string) {
				if result != "@email.com" {
					t.Errorf("Expected '@email.com' for negative input, got: %s", result)
				}
				t.Log("Passed: Generate Email with Negative Length")
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := RandomEmailTest(tc.number)
			tc.assert(t, result, tc.expected)
		})
	}
}

/*
ROOST_METHOD_HASH=RandomString_d7e3599ac4
ROOST_METHOD_SIG_HASH=RandomString_c6fe4ad19a


 */
func RandomStringTest(number int) string {
	var sb strings.Builder
	k := len(testAlphabet)

	for i := 0; i < number; i++ {
		c := testAlphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func Testrandom_randomstring(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	type testCase struct {
		name     string
		input    int
		expected int
	}

	testCases := []testCase{
		{
			name:     "Generate a Random String of Specified Length",
			input:    10,
			expected: 10,
		},
		{
			name:     "Generate an Empty String",
			input:    0,
			expected: 0,
		},
		{
			name:     "Handle Negative Input Gracefully",
			input:    -5,
			expected: 0,
		},
		{
			name:     "High Volume String Generation",
			input:    100000,
			expected: 100000,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := RandomStringTest(tc.input)
			if len(result) != tc.expected {
				t.Errorf("Expected string length %d, got %d", tc.expected, len(result))
			}

			for _, char := range result {
				if !strings.ContainsRune(testAlphabet, char) {
					t.Errorf("Character %c not found in alphabet", char)
				}
			}
		})
	}

	t.Run("Consistent Use of Characters from Alphabet", func(t *testing.T) {
		result := RandomStringTest(10)
		for _, char := range result {
			if !strings.ContainsRune(testAlphabet, char) {
				t.Errorf("Character %c not found in alphabet", char)
			}
		}
	})

	t.Run("Randomness Verification", func(t *testing.T) {
		results := make(map[string]bool)
		for i := 0; i < 100; i++ {
			result := RandomStringTest(10)
			if results[result] {
				t.Errorf("Duplicate string found: %s", result)
			}
			results[result] = true
		}
	})

}

