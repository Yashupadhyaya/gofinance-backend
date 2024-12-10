package util

import (
	"fmt"
	"strings"
	"testing"
)

/*
ROOST_METHOD_HASH=RandomEmail_1905439733
ROOST_METHOD_SIG_HASH=RandomEmail_7a04f189fd


 */
func RandomString(number int) string {
	var sb strings.Builder
	const alphabet = "abcdefghijklmnopqrstuvwxyz"
	k := len(alphabet)

	for i := 0; i < number; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func TestRandomEmail(t *testing.T) {
	type testCase struct {
		name      string
		input     int
		wantLen   int
		wantError bool
	}

	testCases := []testCase{
		{
			name:    "Generate Email with Standard Length",
			input:   10,
			wantLen: 10,
		},
		{
			name:    "Generate Email with Zero Length",
			input:   0,
			wantLen: 0,
		},
		{
			name:    "Generate Email with Large Length",
			input:   1000,
			wantLen: 1000,
		},
		{
			name:    "Generate Email with Special Characters in Domain Part",
			input:   5,
			wantLen: 5,
		},
		{
			name:      "Generate Email with Negative Length",
			input:     -5,
			wantLen:   0,
			wantError: true,
		},
		{
			name:    "Generate Multiple Emails to Ensure Uniqueness",
			input:   10,
			wantLen: 10,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Logf("Running test case: %s", tc.name)
			email := RandomEmail(tc.input)
			localPart := strings.Split(email, "@")[0]

			if !strings.HasSuffix(email, "@email.com") {
				t.Errorf("Email does not end with '@email.com': %s", email)
			}

			if len(localPart) != tc.wantLen {
				t.Errorf("Expected local part length %d, got %d", tc.wantLen, len(localPart))
			}

			if tc.name == "Generate Multiple Emails to Ensure Uniqueness" {
				emailSet := make(map[string]struct{})
				for i := 0; i < 100; i++ {
					email := RandomEmail(tc.input)
					if _, exists := emailSet[email]; exists {
						t.Errorf("Duplicate email found: %s", email)
					}
					emailSet[email] = struct{}{}
				}
			}

			if tc.wantError && len(localPart) != 0 {
				t.Errorf("Expected error for negative input, got local part: %s", localPart)
			}

			t.Logf("Successfully ran test case: %s", tc.name)
		})
	}
}

