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

func TestRandomEmail(t *testing.T) {

	rand.Seed(time.Now().UnixNano())

	testCases := []struct {
		name        string
		input       int
		expectedLen int
		expectedEnd string
	}{
		{
			name:        "Standard Length",
			input:       10,
			expectedLen: 10,
			expectedEnd: "@email.com",
		},
		{
			name:        "Minimum Length",
			input:       1,
			expectedLen: 1,
			expectedEnd: "@email.com",
		},
		{
			name:        "Maximum Length",
			input:       100,
			expectedLen: 100,
			expectedEnd: "@email.com",
		},
		{
			name:        "Zero Length",
			input:       0,
			expectedLen: 0,
			expectedEnd: "@email.com",
		},
		{
			name:        "Upper Bound Length",
			input:       64,
			expectedLen: 64,
			expectedEnd: "@email.com",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			email := RandomEmail(tc.input)

			localPart := strings.Split(email, "@")[0]
			if len(localPart) != tc.expectedLen {
				t.Errorf("expected local part length %d, got %d", tc.expectedLen, len(localPart))
			}

			if !strings.HasSuffix(email, tc.expectedEnd) {
				t.Errorf("expected email to end with %s, got %s", tc.expectedEnd, email)
			}

			t.Logf("Email generated: %s", email)
		})
	}

	t.Run("Multiple Emails Uniqueness", func(t *testing.T) {
		emailSet := make(map[string]struct{})
		for i := 0; i < 100; i++ {
			email := RandomEmail(10)
			if _, exists := emailSet[email]; exists {
				t.Errorf("duplicate email found: %s", email)
			}
			emailSet[email] = struct{}{}
		}
		t.Log("All generated emails are unique")
	})
}

