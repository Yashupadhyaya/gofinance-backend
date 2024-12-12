package util

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"testing"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"
const alphabet = "abcdefghijklmnopqrstuvwxyz"
var originalRandomString = RandomString
/*
ROOST_METHOD_HASH=RandomEmail_1905439733
ROOST_METHOD_SIG_HASH=RandomEmail_7a04f189fd


 */
func RandomEmail(number int) string {
	return fmt.Sprintf("%s@email.com", RandomString(number))
}

func TestRandomEmail(t *testing.T) {
	tests := []struct {
		name     string
		length   int
		expected string
		mockFunc func(int) string
	}{
		{
			name:     "Standard Length",
			length:   10,
			expected: "^[a-z]{10}@email.com$",
			mockFunc: originalRandomString,
		},
		{
			name:     "Zero Length",
			length:   0,
			expected: "^@email.com$",
			mockFunc: originalRandomString,
		},
		{
			name:     "Maximum Length",
			length:   1000,
			expected: "^[a-z]{1000}@email.com$",
			mockFunc: originalRandomString,
		},
		{
			name:     "Special Characters in Local Part",
			length:   10,
			expected: "^!@#$%^&*()@email.com$",
			mockFunc: mockRandomStringSpecialChars,
		},
		{
			name:     "Upper Bound of Characters",
			length:   64,
			expected: "^[a-z]{64}@email.com$",
			mockFunc: originalRandomString,
		},
		{
			name:     "Consistent Format",
			length:   10,
			expected: "^[a]{10}@email.com$",
			mockFunc: mockRandomStringFixed,
		},
		{
			name:     "Randomness of Generated Local Part",
			length:   10,
			expected: "^[a-z]{10}@email.com$",
			mockFunc: originalRandomString,
		},
		{
			name:     "Negative Length Input",
			length:   -5,
			expected: "^@email.com$",
			mockFunc: originalRandomString,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			currentRandomString := RandomString
			RandomString = tt.mockFunc
			defer func() { RandomString = currentRandomString }()

			got := RandomEmail(tt.length)

			matched, err := regexp.MatchString(tt.expected, got)
			if err != nil {
				t.Fatalf("Error compiling regex: %v", err)
			}

			if !matched {
				t.Errorf("Test: '%s' failed. Expected pattern: '%s', but got: '%s'", tt.name, tt.expected, got)
			} else {
				t.Logf("Test: '%s' passed. Generated email: '%s'", tt.name, got)
			}
		})
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func mockRandomStringFixed(number int) string {
	return strings.Repeat("a", number)
}

func mockRandomStringSpecialChars(number int) string {
	return "!@#$%^&*()"
}

/*
ROOST_METHOD_HASH=RandomString_d7e3599ac4
ROOST_METHOD_SIG_HASH=RandomString_c6fe4ad19a


 */
func RandomString(number int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < number; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func TestRandomString(t *testing.T) {

	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{
			name:     "Generate a Random String of Given Length",
			input:    10,
			expected: 10,
		},
		{
			name:     "Generate a Random String of Length Zero",
			input:    0,
			expected: 0,
		},
		{
			name:     "Generate a Random String with Negative Length",
			input:    -5,
			expected: 0,
		},
		{
			name:     "Generate a Random String of Maximum Length",
			input:    10000,
			expected: 10000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RandomString(tt.input)
			if len(result) != tt.expected {
				t.Errorf("Expected length %d, but got %d", tt.expected, len(result))
			}
			t.Logf("Test %s passed", tt.name)
		})
	}

	t.Run("Randomness of Generated String", func(t *testing.T) {
		length := 10
		result1 := RandomString(length)
		result2 := RandomString(length)
		if result1 == result2 {
			t.Errorf("Expected different strings, but got the same: %s", result1)
		}
		t.Logf("Test Randomness of Generated String passed")
	})

	t.Run("Verify Characters in the Generated String", func(t *testing.T) {
		length := 10
		result := RandomString(length)
		for _, char := range result {
			if !strings.ContainsRune(alphabet, char) {
				t.Errorf("Unexpected character %c in the result string", char)
			}
		}
		t.Logf("Test Verify Characters in the Generated String passed")
	})

	t.Run("Consistent Seed Behavior", func(t *testing.T) {
		rand.Seed(1)
		result1 := RandomString(10)
		rand.Seed(1)
		result2 := RandomString(10)
		if result1 != result2 {
			t.Errorf("Expected the same string for the same seed, but got different strings: %s and %s", result1, result2)
		}
		t.Logf("Test Consistent Seed Behavior passed")
	})

	t.Run("Performance Test for Large Input", func(t *testing.T) {
		start := time.Now()
		RandomString(1000000)
		duration := time.Since(start)
		if duration.Seconds() > 1 {
			t.Errorf("Function took too long: %v", duration)
		}
		t.Logf("Test Performance Test for Large Input passed")
	})
}

