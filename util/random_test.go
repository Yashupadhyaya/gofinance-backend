package util

import (
	"math/rand"
	"strings"
	"testing"
	"time"
)

/*
ROOST_METHOD_HASH=RandomEmail_1905439733
ROOST_METHOD_SIG_HASH=RandomEmail_7a04f189fd


 */
func RandomEmail(number int) string {
	if number < 0 {
		number = 0
	}
	return RandomString(number) + "@email.com"
}

func RandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func TestRandomEmail(t *testing.T) {

	rand.Seed(time.Now().UnixNano())

	type testCase struct {
		description string
		length      int
		expected    string
		assertFunc  func(t *testing.T, result string, expected string)
	}

	testCases := []testCase{
		{
			description: "Generate Email with Standard Length",
			length:      10,
			expected:    "@email.com",
			assertFunc: func(t *testing.T, result string, expected string) {
				if len(result) != 10+len(expected) || !strings.HasSuffix(result, expected) {
					t.Fatalf("Expected length %d+%d and suffix %s, got %s", 10, len(expected), expected, result)
				}
			},
		},
		{
			description: "Generate Email with Zero Length",
			length:      0,
			expected:    "@email.com",
			assertFunc: func(t *testing.T, result string, expected string) {
				if result != expected {
					t.Fatalf("Expected %s, got %s", expected, result)
				}
			},
		},
		{
			description: "Generate Email with Maximum Length",
			length:      1000,
			expected:    "@email.com",
			assertFunc: func(t *testing.T, result string, expected string) {
				if len(result) != 1000+len(expected) || !strings.HasSuffix(result, expected) {
					t.Fatalf("Expected length %d+%d and suffix %s, got %s", 1000, len(expected), expected, result)
				}
			},
		},
		{
			description: "Consistent Email Domain",
			length:      5,
			expected:    "@email.com",
			assertFunc: func(t *testing.T, result string, expected string) {
				if !strings.HasSuffix(result, expected) {
					t.Fatalf("Expected suffix %s, got %s", expected, result)
				}
			},
		},
		{
			description: "Non-Negative Number Length",
			length:      -5,
			expected:    "@email.com",
			assertFunc: func(t *testing.T, result string, expected string) {
				if result != expected {
					t.Fatalf("Expected %s, got %s", expected, result)
				}
			},
		},
		{
			description: "Randomness of Email Generation",
			length:      8,
			expected:    "@email.com",
			assertFunc: func(t *testing.T, result string, expected string) {
				anotherResult := RandomEmail(8)
				if result == anotherResult {
					t.Fatalf("Expected different results for subsequent calls, got %s and %s", result, anotherResult)
				}
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			t.Logf("Running test: %s", tc.description)
			result := RandomEmail(tc.length)
			tc.assertFunc(t, result, tc.expected)
		})
	}
}

/*
ROOST_METHOD_HASH=RandomString_d7e3599ac4
ROOST_METHOD_SIG_HASH=RandomString_c6fe4ad19a


 */
func TestRandomString(t *testing.T) {

	rand.Seed(time.Now().UnixNano())

	testCases := []struct {
		description string
		length      int
	}{
		{
			description: "Generate a Random String of Specified Length",
			length:      10,
		},
		{
			description: "Generate an Empty String",
			length:      0,
		},
		{
			description: "Generate a String with Maximum Length",
			length:      10000,
		},
		{
			description: "Randomness of Generated String",
			length:      10,
		},
		{
			description: "Character Set Validation",
			length:      10,
		},
		{
			description: "Consistent Behavior Across Calls",
			length:      10,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			result := RandomString(tc.length)

			switch tc.description {
			case "Generate a Random String of Specified Length":
				if len(result) != tc.length {
					t.Errorf("Expected length %d, but got %d", tc.length, len(result))
				} else {
					t.Logf("Success: Generated string of length %d", tc.length)
				}

			case "Generate an Empty String":
				if result != "" {
					t.Errorf("Expected empty string, but got '%s'", result)
				} else {
					t.Log("Success: Generated empty string")
				}

			case "Generate a String with Maximum Length":
				if len(result) != tc.length {
					t.Errorf("Expected length %d, but got %d", tc.length, len(result))
				} else {
					t.Logf("Success: Generated string of maximum length %d", tc.length)
				}

			case "Randomness of Generated String":
				otherResult := RandomString(tc.length)
				if result == otherResult {
					t.Errorf("Expected different strings, but got identical '%s'", result)
				} else {
					t.Log("Success: Generated different strings showing randomness")
				}

			case "Character Set Validation":
				for _, char := range result {
					if !strings.ContainsRune(testAlphabet, char) {
						t.Errorf("Unexpected character '%c' found in string", char)
						return
					}
				}
				t.Log("Success: All characters are from the predefined alphabet")

			case "Consistent Behavior Across Calls":
				for i := 0; i < 5; i++ {
					consistencyResult := RandomString(tc.length)
					if len(consistencyResult) != tc.length {
						t.Errorf("Inconsistent length. Expected %d, got %d", tc.length, len(consistencyResult))
					}
					for _, char := range consistencyResult {
						if !strings.ContainsRune(testAlphabet, char) {
							t.Errorf("Unexpected character '%c' found in string during consistency test", char)
							return
						}
					}
				}
				t.Log("Success: Consistent behavior across multiple calls")
			}
		})
	}
}

