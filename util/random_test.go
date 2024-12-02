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
func RandomEmail(number int) string {
	return fmt.Sprintf("%s@email.com", RandomString(number))
}

func TestRandomEmail(t *testing.T) {

	rand.Seed(time.Now().UnixNano())

	type testCase struct {
		name          string
		number        int
		expectedLocal int
		expectedEmail string
	}

	testCases := []testCase{
		{
			name:          "Standard Length",
			number:        10,
			expectedLocal: 10,
			expectedEmail: "@email.com",
		},
		{
			name:          "Minimum Length",
			number:        1,
			expectedLocal: 1,
			expectedEmail: "@email.com",
		},
		{
			name:          "Large Length",
			number:        1000,
			expectedLocal: 1000,
			expectedEmail: "@email.com",
		},
		{
			name:          "Zero Length",
			number:        0,
			expectedLocal: 0,
			expectedEmail: "@email.com",
		},
		{
			name:          "Negative Length",
			number:        -5,
			expectedLocal: 0,
			expectedEmail: "@email.com",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			email := RandomEmail(tc.number)

			parts := strings.Split(email, "@")
			if len(parts) != 2 {
				t.Fatalf("invalid email format: %s", email)
			}
			localPart := parts[0]
			domainPart := "@" + parts[1]

			if len(localPart) != tc.expectedLocal {
				t.Errorf("expected local part length %d, got %d", tc.expectedLocal, len(localPart))
			}

			if domainPart != tc.expectedEmail {
				t.Errorf("expected domain %s, got %s", tc.expectedEmail, domainPart)
			}

			t.Logf("Success for %s: Generated email = %s", tc.name, email)
		})
	}

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
	type testCase struct {
		name     string
		input    int
		expected int
		validate func(string) bool
	}

	tests := []testCase{
		{
			name:     "Positive Length",
			input:    10,
			expected: 10,
			validate: func(result string) bool {
				return len(result) == 10
			},
		},
		{
			name:     "Zero Length",
			input:    0,
			expected: 0,
			validate: func(result string) bool {
				return result == ""
			},
		},
		{
			name:     "Negative Length",
			input:    -5,
			expected: 0,
			validate: func(result string) bool {
				return result == ""
			},
		},
		{
			name:     "Randomness Check",
			input:    10,
			expected: 10,
			validate: func(result string) bool {

				another := RandomString(10)
				return result != another
			},
		},
		{
			name:     "Maximum Length",
			input:    100000,
			expected: 100000,
			validate: func(result string) bool {
				return len(result) == 100000
			},
		},
		{
			name:     "Consistent Output with Seed",
			input:    10,
			expected: 10,
			validate: func(result string) bool {
				rand.Seed(42)
				first := RandomString(10)
				rand.Seed(42)
				second := RandomString(10)
				return first == second
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			rand.Seed(time.Now().UnixNano())
			result := RandomString(tc.input)
			if !tc.validate(result) {
				t.Errorf("Test %s failed: expected length %d, got %d", tc.name, tc.expected, len(result))
			} else {
				t.Logf("Test %s passed", tc.name)
			}
		})
	}
}

