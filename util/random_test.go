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
func TestRandomEmail(t *testing.T) {

	rand.Seed(time.Now().UnixNano())

	tests := []struct {
		name     string
		input    int
		expected string
	}{
		{
			name:     "Standard Length",
			input:    10,
			expected: "10 characters + @email.com",
		},
		{
			name:     "Minimum Length",
			input:    1,
			expected: "1 character + @email.com",
		},
		{
			name:     "Maximum Length",
			input:    1000,
			expected: "1000 characters + @email.com",
		},
		{
			name:     "Zero Length",
			input:    0,
			expected: "@email.com",
		},
		{
			name:     "Negative Length",
			input:    -5,
			expected: "handle negative input gracefully",
		},
		{
			name:     "Multiple Emails for Uniqueness",
			input:    10,
			expected: "unique emails",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.name {
			case "Standard Length", "Minimum Length", "Maximum Length", "Zero Length":
				email := RandomEmail(tt.input)
				localPart := strings.Split(email, "@")[0]
				expectedLength := tt.input

				if tt.name == "Zero Length" {
					expectedLength = 0
				}

				if len(localPart) != expectedLength {
					t.Errorf("expected local part length %d, got %d", expectedLength, len(localPart))
				}
				if !strings.HasSuffix(email, "@email.com") {
					t.Errorf("expected email to end with @email.com, got %s", email)
				}
				t.Logf("Generated email: %s", email)

			case "Negative Length":
				defer func() {
					if r := recover(); r != nil {
						t.Logf("Recovered from panic: %v", r)
					}
				}()
				email := RandomEmail(tt.input)
				if email != "@email.com" {
					t.Errorf("expected '@email.com' for negative input, got %s", email)
				}

			case "Multiple Emails for Uniqueness":
				emails := make(map[string]bool)
				for i := 0; i < 100; i++ {
					email := RandomEmail(tt.input)
					if emails[email] {
						t.Errorf("duplicate email found: %s", email)
					}
					emails[email] = true
				}
				t.Logf("Generated %d unique emails", len(emails))
			}
		})
	}
}

