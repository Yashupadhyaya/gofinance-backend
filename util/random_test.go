package util

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"
	"unicode"
)

/*
ROOST_METHOD_HASH=RandomEmail_1905439733
ROOST_METHOD_SIG_HASH=RandomEmail_7a04f189fd


 */
func RandomEmail(number int) string {
	return fmt.Sprintf("%s@email.com", RandomString(number))
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

	tests := []struct {
		name        string
		inputLength int
		expectedLen int
		expectedFmt string
	}{
		{
			name:        "Standard Length",
			inputLength: 10,
			expectedLen: 10,
			expectedFmt: "@email.com",
		},
		{
			name:        "Zero Length",
			inputLength: 0,
			expectedLen: 0,
			expectedFmt: "@email.com",
		},
		{
			name:        "Maximum Length",
			inputLength: 64,
			expectedLen: 64,
			expectedFmt: "@email.com",
		},
		{
			name:        "Special Characters Check",
			inputLength: 12,
			expectedLen: 12,
			expectedFmt: "@email.com",
		},
		{
			name:        "Consistent Domain",
			inputLength: 8,
			expectedLen: 8,
			expectedFmt: "@email.com",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			email := RandomEmail(tt.inputLength)
			parts := strings.Split(email, "@")

			if len(parts) != 2 {
				t.Errorf("Email format invalid: %s", email)
			}

			localPart := parts[0]
			domainPart := "@" + parts[1]

			if len(localPart) != tt.expectedLen {
				t.Errorf("Local part length = %d; want %d", len(localPart), tt.expectedLen)
			}

			if domainPart != tt.expectedFmt {
				t.Errorf("Domain part = %s; want %s", domainPart, tt.expectedFmt)
			}

			for _, r := range localPart {
				if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
					t.Errorf("Local part contains special character: %c", r)
					break
				}
			}

			t.Logf("Test %s passed with email: %s", tt.name, email)
		})
	}
}

/*
ROOST_METHOD_HASH=RandomString_d7e3599ac4
ROOST_METHOD_SIG_HASH=RandomString_c6fe4ad19a


 */
func RandomString(length int) string {
	var sb strings.Builder
	k := len(testAlphabet)

	for i := 0; i < length; i++ {
		c := testAlphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func TestRandomString(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	tests := []struct {
		name     string
		length   int
		validate func(result string) error
	}{
		{
			name:   "Generate a Random String of Positive Length",
			length: 10,
			validate: func(result string) error {
				if len(result) != 10 {
					return fmt.Errorf("expected length 10, got %d", len(result))
				}
				return nil
			},
		},
		{
			name:   "Generate a Random String of Length Zero",
			length: 0,
			validate: func(result string) error {
				if result != "" {
					return fmt.Errorf("expected empty string, got %s", result)
				}
				return nil
			},
		},
		{
			name:   "Handle Negative Length Input Gracefully",
			length: -5,
			validate: func(result string) error {
				if result != "" {
					return fmt.Errorf("expected empty string for negative length, got %s", result)
				}
				return nil
			},
		},
		{
			name:   "Consistent Output Length with Multiple Calls",
			length: 15,
			validate: func(result string) error {
				if len(result) != 15 {
					return fmt.Errorf("expected length 15, got %d", len(result))
				}
				return nil
			},
		},
		{
			name:   "Correct Character Set Utilization",
			length: 20,
			validate: func(result string) error {
				for _, char := range result {
					if !strings.ContainsRune(testAlphabet, char) {
						return fmt.Errorf("unexpected character %c in result", char)
					}
				}
				return nil
			},
		},
		{
			name:   "Randomness of Output",
			length: 10,
			validate: func(result string) error {
				otherResult := RandomString(10)
				if result == otherResult {
					return fmt.Errorf("expected different strings, got identical: %s", result)
				}
				return nil
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RandomString(tt.length)
			if err := tt.validate(result); err != nil {
				t.Errorf("Test %s failed: %v", tt.name, err)
			} else {
				t.Logf("Test %s passed successfully", tt.name)
			}
		})
	}
}

