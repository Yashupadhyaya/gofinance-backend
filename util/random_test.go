package util

import (
	"fmt"
	"math/rand"
	"regexp"
	"testing"
	"time"
	"strings"
)

/*
ROOST_METHOD_HASH=RandomEmail_1905439733
ROOST_METHOD_SIG_HASH=RandomEmail_7a04f189fd


 */
func DummyRandomEmail(number int) string {
	return fmt.Sprintf("%s@email.com", DummyRandomString(number))
}

func DummyRandomString(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func TestRandomRandomEmail(t *testing.T) {
	t.Parallel()

	type testCase struct {
		description string
		number      int
		expected    string
		shouldMatch bool
	}

	testCases := []testCase{
		{
			description: "Generate Email with a Standard Length",
			number:      10,
			expected:    "^.{10}@email\\.com$",
			shouldMatch: true,
		},
		{
			description: "Generate Email with Minimum Length",
			number:      1,
			expected:    "^.{1}@email\\.com$",
			shouldMatch: true,
		},
		{
			description: "Generate Email with Zero Length",
			number:      0,
			expected:    "^@email\\.com$",
			shouldMatch: true,
		},
		{
			description: "Generate Email with Maximum Length",
			number:      64,
			expected:    "^.{64}@email\\.com$",
			shouldMatch: true,
		},
		{
			description: "Consistent Output Format",
			number:      5,
			expected:    "^[a-zA-Z0-9._%+-]+@email\\.com$",
			shouldMatch: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()

			rand.Seed(time.Now().UnixNano())

			email := DummyRandomEmail(tc.number)

			if tc.shouldMatch {
				matched, err := regexp.MatchString(tc.expected, email)
				if err != nil {
					t.Fatalf("Error matching regex: %v", err)
				}
				if !matched {
					t.Errorf("Expected email to match %v, but got %v", tc.expected, email)
				}
			} else {
				if email != tc.expected {
					t.Errorf("Expected email to be %v, but got %v", tc.expected, email)
				}
			}

			t.Logf("Test '%s' passed with email: %s", tc.description, email)
		})
	}

}

/*
ROOST_METHOD_HASH=RandomString_d7e3599ac4
ROOST_METHOD_SIG_HASH=RandomString_c6fe4ad19a


 */
func TestRandomRandomString(t *testing.T) {

	type testCase struct {
		description string
		input       int
		expectedLen int
		expectEmpty bool
	}

	rand.Seed(1)

	testCases := []testCase{
		{
			description: "Generate a Random String of Specified Length",
			input:       10,
			expectedLen: 10,
		},
		{
			description: "Generate an Empty String",
			input:       0,
			expectEmpty: true,
		},
		{
			description: "Generate a String with Maximum Length",
			input:       10000,
			expectedLen: 10000,
		},
		{
			description: "Handle Negative Length Input",
			input:       -5,
			expectEmpty: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			result := RandomString(tc.input)
			if tc.expectEmpty && result != "" {
				t.Errorf("Expected empty string, got %s", result)
			} else if !tc.expectEmpty && len(result) != tc.expectedLen {
				t.Errorf("Expected string length %d, got %d", tc.expectedLen, len(result))
			}
			t.Logf("Test passed for scenario: %s", tc.description)
		})
	}

	t.Run("Randomness of Generated String", func(t *testing.T) {
		n := 100
		str1 := RandomString(n)
		str2 := RandomString(n)
		if str1 == str2 {
			t.Error("Expected different strings for the same length, got identical strings")
		}
		t.Log("Successfully verified randomness of generated strings")
	})

	t.Run("Consistent Output with Seeded Randomness", func(t *testing.T) {
		n := 10
		rand.Seed(42)
		expected := RandomString(n)
		rand.Seed(42)
		actual := RandomString(n)
		if expected != actual {
			t.Errorf("Expected consistent output with seeded randomness, got different outputs: %s and %s", expected, actual)
		}
		t.Log("Successfully verified consistency with seeded randomness")
	})
}

