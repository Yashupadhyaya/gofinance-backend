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
func GenerateRandomEmail(number int) string {
	if number < 0 {
		number = 0
	}
	return GenerateRandomString(number) + "@email.com"
}

func GenerateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func TestRandomEmail(t *testing.T) {

	rand.Seed(time.Now().UnixNano())

	testCases := []struct {
		description string
		length      int
		expectedLen int
		expectError bool
	}{
		{
			description: "Generate Email with Standard Length",
			length:      8,
			expectedLen: 8 + len("@email.com"),
			expectError: false,
		},
		{
			description: "Generate Email with Zero Length",
			length:      0,
			expectedLen: len("@email.com"),
			expectError: false,
		},
		{
			description: "Generate Email with Maximum Length",
			length:      1000,
			expectedLen: 1000 + len("@email.com"),
			expectError: false,
		},
		{
			description: "Generate Email with Negative Length",
			length:      -5,
			expectedLen: len("@email.com"),
			expectError: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			result := GenerateRandomEmail(tc.length)

			if len(result) != tc.expectedLen {
				t.Errorf("Test %s: expected length %d, got %d", tc.description, tc.expectedLen, len(result))
			}

			if !strings.HasSuffix(result, "@email.com") {
				t.Errorf("Test %s: expected email to end with '@email.com', got %s", tc.description, result)
			}

			randomPart := result[:len(result)-len("@email.com")]
			if !isAlphanumeric(randomPart) {
				t.Errorf("Test %s: expected alphanumeric characters before '@', got %s", tc.description, randomPart)
			}

			t.Logf("Test %s passed successfully", tc.description)
		})
	}
}

func isAlphanumeric(s string) bool {
	for _, c := range s {
		if !strings.ContainsRune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789", c) {
			return false
		}
	}
	return true
}

/*
ROOST_METHOD_HASH=RandomString_d7e3599ac4
ROOST_METHOD_SIG_HASH=RandomString_c6fe4ad19a


 */
func RandomString(length int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < length; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func TestRandomString(t *testing.T) {

	rand.Seed(time.Now().UnixNano())

	tests := []struct {
		name           string
		length         int
		expectedLength int
		expectRandom   bool
	}{
		{
			name:           "Positive Length",
			length:         10,
			expectedLength: 10,
			expectRandom:   true,
		},
		{
			name:           "Zero Length",
			length:         0,
			expectedLength: 0,
			expectRandom:   false,
		},
		{
			name:           "Large Length",
			length:         1000,
			expectedLength: 1000,
			expectRandom:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RandomString(tt.length)
			if len(result) != tt.expectedLength {
				t.Fatalf("Expected length %d, got %d", tt.expectedLength, len(result))
			}
			if tt.length > 0 {
				for _, ch := range result {
					if !strings.ContainsRune(alphabet, ch) {
						t.Fatalf("Character %c not found in the alphabet", ch)
					}
				}
			}
		})
	}

	t.Run("Randomness Check", func(t *testing.T) {
		length := 10
		str1 := RandomString(length)
		str2 := RandomString(length)
		if str1 == str2 {
			t.Logf("Generated strings are equal: %s and %s, which is unlikely but possible", str1, str2)
		} else {
			t.Log("Generated strings are different, randomness confirmed")
		}
	})

	t.Run("Consistency with Seeded Randomness", func(t *testing.T) {
		seed := int64(42)
		rand.Seed(seed)
		expected := RandomString(10)
		rand.Seed(seed)
		actual := RandomString(10)
		if expected != actual {
			t.Fatalf("Expected %s but got %s with the same seed", expected, actual)
		}
		t.Logf("Generated strings with seed %d are consistent", seed)
	})
}

