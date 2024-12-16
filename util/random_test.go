package util

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"testing"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"/*
ROOST_METHOD_HASH=RandomEmail_1905439733
ROOST_METHOD_SIG_HASH=RandomEmail_7a04f189fd


 */
func RandomEmailWithAlphabet(number int, alphabet string) string {
	return fmt.Sprintf("%s@email.com", RandomStringWithAlphabet(number, alphabet))
}

func RandomStringWithAlphabet(number int, alphabet string) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < number; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func TestRandomEmail(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	type testCase struct {
		name         string
		input        int
		expectLength int
		expectFormat string
	}

	testCases := []testCase{
		{
			name:         "Generate Email with 5-Character Local Part",
			input:        5,
			expectLength: 5,
			expectFormat: `^[a-z]{5}@email\.com$`,
		},
		{
			name:         "Generate Email with 0-Character Local Part",
			input:        0,
			expectLength: 0,
			expectFormat: `^@email\.com$`,
		},
		{
			name:         "Generate Email with Maximum Length Local Part",
			input:        100,
			expectLength: 100,
			expectFormat: `^[a-z]{100}@email\.com$`,
		},
		{
			name:         "Check Randomness of Generated Emails",
			input:        10,
			expectLength: 10,
			expectFormat: `^[a-z]{10}@email\.com$`,
		},
		{
			name:         "Validate Email Format",
			input:        10,
			expectLength: 10,
			expectFormat: `^[a-z]{10}@email\.com$`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			email1 := RandomEmail(tc.input)
			localPart := strings.Split(email1, "@")[0]

			if len(localPart) != tc.expectLength {
				t.Errorf("expected local part length %d, got %d", tc.expectLength, len(localPart))
			}

			matched, err := regexp.MatchString(tc.expectFormat, email1)
			if err != nil {
				t.Fatalf("failed to compile regex: %v", err)
			}
			if !matched {
				t.Errorf("email %s does not match format %s", email1, tc.expectFormat)
			}

			if tc.name == "Check Randomness of Generated Emails" {
				email2 := RandomEmail(tc.input)
				if email1 == email2 {
					t.Errorf("expected different emails, got same: %s and %s", email1, email2)
				}
			}

			t.Logf("success: %s", email1)
		})
	}

	t.Run("Generate Email with Special Characters in Local Part", func(t *testing.T) {
		originalAlphabet := alphabet
		defer func() { alphabet = originalAlphabet }()

		alphabetWithSpecialChars := "abc123!@#"
		email := RandomEmailWithAlphabet(5, alphabetWithSpecialChars)
		localPart := strings.Split(email, "@")[0]

		if len(localPart) != 5 {
			t.Errorf("expected local part length 5, got %d", len(localPart))
		}

		matched, err := regexp.MatchString(`^[abc123!@#]{5}@email\.com$`, email)
		if err != nil {
			t.Fatalf("failed to compile regex: %v", err)
		}
		if !matched {
			t.Errorf("email %s does not match format with special characters", email)
		}

		t.Logf("success: %s", email)
	})
}

/*
ROOST_METHOD_HASH=RandomString_d7e3599ac4
ROOST_METHOD_SIG_HASH=RandomString_c6fe4ad19a


 */
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

func TestRandomString(t *testing.T) {

	rand.Seed(time.Now().UnixNano())

	tests := []struct {
		name                  string
		input                 int
		expectedLen           int
		shouldContainAlphabet bool
	}{
		{
			name:                  "Generate a Random String of Given Length",
			input:                 10,
			expectedLen:           10,
			shouldContainAlphabet: true,
		},
		{
			name:                  "Generate a Random String with Zero Length",
			input:                 0,
			expectedLen:           0,
			shouldContainAlphabet: true,
		},
		{
			name:                  "Generate a Random String with Negative Length",
			input:                 -1,
			expectedLen:           0,
			shouldContainAlphabet: true,
		},
		{
			name:                  "Verify Characters in the Generated String",
			input:                 100,
			expectedLen:           100,
			shouldContainAlphabet: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RandomString(tt.input)

			if len(result) != tt.expectedLen {
				t.Errorf("expected length %d, got %d", tt.expectedLen, len(result))
			}

			if tt.shouldContainAlphabet {
				const alphabet = "abcdefghijklmnopqrstuvwxyz"
				for _, char := range result {
					if !strings.ContainsRune(alphabet, char) {
						t.Errorf("string contains invalid character: %c", char)
					}
				}
			}
		})
	}

	t.Run("Consistency in Randomness", func(t *testing.T) {
		result1 := RandomString(10)
		result2 := RandomString(10)
		if result1 == result2 {
			t.Errorf("expected different strings, got identical strings: %s", result1)
		}
	})

	t.Run("Performance for Large Input Lengths", func(t *testing.T) {
		start := time.Now()
		_ = RandomString(1000000)
		duration := time.Since(start)
		if duration.Seconds() > 1 {
			t.Errorf("function took too long: %v", duration)
		}
	})

	t.Run("Verify Seed Initialization", func(t *testing.T) {
		result1 := RandomString(10)
		time.Sleep(1 * time.Second)
		result2 := RandomString(10)
		if result1 == result2 {
			t.Errorf("expected different strings for different times, got identical strings: %s", result1)
		}
	})
}

