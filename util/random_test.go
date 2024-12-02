package util

import (
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
	type testScenario struct {
		description string
		length      int
		expected    string
	}

	scenarios := []testScenario{
		{
			description: "Generate Email with a Standard Length",
			length:      10,
			expected:    "10 characters local part",
		},
		{
			description: "Generate Email with Zero Length",
			length:      0,
			expected:    "@email.com",
		},
		{
			description: "Generate Email with Maximum Length",
			length:      1000,
			expected:    "1000 characters local part",
		},
		{
			description: "Generate Email with Negative Length",
			length:      -5,
			expected:    "empty local part or error",
		},
		{
			description: "Consistent Domain Name",
			length:      10,
			expected:    "@email.com",
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.description, func(t *testing.T) {
			email := RandomEmail(scenario.length)

			if scenario.length >= 0 {
				parts := strings.Split(email, "@")
				localPartLength := len(parts[0])

				if scenario.length == 0 {
					if email != scenario.expected {
						t.Errorf("Expected email '%s', but got '%s'", scenario.expected, email)
					} else {
						t.Logf("Success: %s", scenario.description)
					}
				} else if scenario.length > 0 {
					if localPartLength != scenario.length {
						t.Errorf("Expected local part length of %d, but got %d", scenario.length, localPartLength)
					} else {
						t.Logf("Success: %s", scenario.description)
					}
				}
			} else {

				if len(email) != len("@email.com") {
					t.Errorf("Expected email to handle negative input with length of %d, but got %d", len("@email.com"), len(email))
				} else {
					t.Logf("Success: %s", scenario.description)
				}
			}

			if !strings.HasSuffix(email, "@email.com") {
				t.Errorf("Expected domain '@email.com', but got '%s'", email)
			} else {
				t.Logf("Domain part is consistent for %s", scenario.description)
			}
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

	testCases := []struct {
		description string
		length      int
		expectedLen int
		expectedErr bool
	}{
		{"Generate a Random String of Specified Length", 10, 10, false},
		{"Generate an Empty String When Length is Zero", 0, 0, false},
		{"Generate a String with Upper Bound Length", 10000, 10000, false},
		{"Verify Function for Negative Input", -5, 0, true},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			result := RandomString(tc.length)
			if len(result) != tc.expectedLen {
				t.Errorf("Expected length %d, but got %d", tc.expectedLen, len(result))
			}

			if tc.expectedErr && result != "" {
				t.Errorf("Expected an empty string for negative input, but got '%s'", result)
			}

			for _, char := range result {
				if !strings.ContainsRune(alphabet, char) {
					t.Errorf("Character '%c' not in predefined alphabet", char)
				}
			}

			t.Logf("Test '%s' passed with result: '%s'", tc.description, result)
		})
	}

	t.Run("Confirm Consistent Results with Seeded Randomness", func(t *testing.T) {
		rand.Seed(1)
		expected := RandomString(10)

		rand.Seed(1)
		actual := RandomString(10)

		if expected != actual {
			t.Errorf("Expected seeded result '%s', but got '%s'", expected, actual)
		}
		t.Logf("Test 'Confirm Consistent Results with Seeded Randomness' passed with result: '%s'", actual)
	})

	t.Run("Ensure Unique Characters Are Randomly Distributed", func(t *testing.T) {
		length := 100
		iterations := 1000
		charCount := make(map[rune]int)

		for i := 0; i < iterations; i++ {
			result := RandomString(length)
			for _, char := range result {
				charCount[char]++
			}
		}

		average := float64(length*iterations) / float64(len(alphabet))
		threshold := 0.1 * average

		for char, count := range charCount {
			if float64(count) < average-threshold || float64(count) > average+threshold {
				t.Errorf("Character '%c' is disproportionally represented: %d occurrences", char, count)
			}
		}
		t.Logf("Test 'Ensure Unique Characters Are Randomly Distributed' passed with valid distribution")
	})
}

