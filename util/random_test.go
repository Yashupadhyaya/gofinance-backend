package util

import (
	"math/rand"
	"strings"
	"testing"
	"time"
	"fmt"
)

/*
ROOST_METHOD_HASH=RandomEmail_1905439733
ROOST_METHOD_SIG_HASH=RandomEmail_7a04f189fd


 */
func GenerateRandomEmail(number int) string {
	return RandomString(number) + "@email.com"
}

func RandomString(n int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func TestGenerateRandomEmail(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	type testCase struct {
		length       int
		expectedLen  int
		expectedPart string
		description  string
	}

	testCases := []testCase{
		{
			length:       8,
			expectedLen:  8 + len("@email.com"),
			expectedPart: "@email.com",
			description:  "Generate Email with Standard Length",
		},
		{
			length:       0,
			expectedLen:  len("@email.com"),
			expectedPart: "@email.com",
			description:  "Generate Email with Zero Length",
		},
		{
			length:       1000,
			expectedLen:  1000 + len("@email.com"),
			expectedPart: "@email.com",
			description:  "Generate Email with Maximum Length",
		},
		{
			length:       -5,
			expectedLen:  len("@email.com"),
			expectedPart: "@email.com",
			description:  "Generate Email with Negative Length",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			email := GenerateRandomEmail(tc.length)
			if len(email) != tc.expectedLen {
				t.Errorf("Failed %s: Expected length %d, got %d", tc.description, tc.expectedLen, len(email))
			}
			if !strings.HasSuffix(email, tc.expectedPart) {
				t.Errorf("Failed %s: Expected suffix %s, got %s", tc.description, tc.expectedPart, email)
			}
			t.Logf("Passed %s: Email generated successfully with expected format and length", tc.description)
		})
	}

	t.Run("Consistent Email Domain", func(t *testing.T) {
		email := GenerateRandomEmail(10)
		expectedDomain := "@email.com"
		if !strings.HasSuffix(email, expectedDomain) {
			t.Errorf("Failed Consistent Email Domain: Expected domain %s, got %s", expectedDomain, email)
		} else {
			t.Log("Passed Consistent Email Domain: Email domain is consistent")
		}
	})

	t.Run("Randomness of Generated Emails", func(t *testing.T) {
		length := 10
		email1 := GenerateRandomEmail(length)
		email2 := GenerateRandomEmail(length)

		if email1 == email2 {
			t.Errorf("Failed Randomness of Generated Emails: Expected different emails, got %s and %s", email1, email2)
		} else {
			t.Log("Passed Randomness of Generated Emails: Generated emails are random")
		}
	})
}

/*
ROOST_METHOD_HASH=RandomString_d7e3599ac4
ROOST_METHOD_SIG_HASH=RandomString_c6fe4ad19a


 */
func TestRandomString(number int) string {
	var sb strings.Builder
	k := len(testAlphabet)

	for i := 0; i < number; i++ {
		c := testAlphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func TestRandomStringFunction(t *testing.T) {

	tests := []struct {
		name     string
		input    int
		expected int
		validate func(string) error
	}{
		{
			name:     "Positive Length",
			input:    10,
			expected: 10,
			validate: func(s string) error {
				if len(s) != 10 {
					return fmt.Errorf("expected length 10, got %d", len(s))
				}
				return nil
			},
		},
		{
			name:     "Zero Length",
			input:    0,
			expected: 0,
			validate: func(s string) error {
				if s != "" {
					return fmt.Errorf("expected empty string, got %q", s)
				}
				return nil
			},
		},
		{
			name:     "Negative Length",
			input:    -5,
			expected: 0,
			validate: func(s string) error {
				if s != "" {
					return fmt.Errorf("expected empty string for negative input, got %q", s)
				}
				return nil
			},
		},
		{
			name:     "Consistent Randomness with Seed",
			input:    10,
			expected: 10,
			validate: func(s string) error {
				rand.Seed(12345)
				first := TestRandomString(10)
				rand.Seed(12345)
				second := TestRandomString(10)
				if first != second {
					return fmt.Errorf("expected consistent result with fixed seed, got %q and %q", first, second)
				}
				return nil
			},
		},
		{
			name:     "Verify Character Set Usage",
			input:    50,
			expected: 50,
			validate: func(s string) error {
				for _, c := range s {
					if !strings.ContainsRune(testAlphabet, c) {
						return fmt.Errorf("character %q not in allowed alphabet", c)
					}
				}
				return nil
			},
		},
		{
			name:     "Large Length Input",
			input:    1000000,
			expected: 1000000,
			validate: func(s string) error {
				if len(s) != 1000000 {
					return fmt.Errorf("expected length 1000000, got %d", len(s))
				}
				return nil
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			rand.Seed(time.Now().UnixNano())
			result := TestRandomString(tc.input)
			if err := tc.validate(result); err != nil {
				t.Errorf("Test %s failed: %v", tc.name, err)
			} else {
				t.Logf("Test %s succeeded", tc.name)
			}
		})
	}
}

