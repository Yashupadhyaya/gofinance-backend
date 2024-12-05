package util

import (
	"fmt"
	"testing"
	"time"
)

func TestRandomEmail(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		validate func(email string) error
	}{
		{
			name:  "Standard Length",
			input: 10,
			validate: func(email string) error {
				if !strings.Contains(email, "@") || !strings.HasSuffix(email, "@email.com") {
					return fmt.Errorf("invalid email format: %s", email)
				}
				return nil
			},
		},
		{
			name:  "Zero Length",
			input: 0,
			validate: func(email string) error {
				expected := "@email.com"
				if email != expected {
					return fmt.Errorf("expected %s, got %s", expected, email)
				}
				return nil
			},
		},
		{
			name:  "Large Length",
			input: 1000,
			validate: func(email string) error {
				expectedLength := 1000 + len("@email.com")
				if len(email) != expectedLength {
					return fmt.Errorf("expected length %d, got %d", expectedLength, len(email))
				}
				if !strings.HasSuffix(email, "@email.com") {
					return fmt.Errorf("invalid email format: %s", email)
				}
				return nil
			},
		},
		{
			name:  "Negative Length",
			input: -5,
			validate: func(email string) error {
				expected := "@email.com"
				if email != expected {
					return fmt.Errorf("expected %s, got %s", expected, email)
				}
				return nil
			},
		},
		{
			name:  "Special Characters",
			input: 15,
			validate: func(email string) error {
				if !strings.HasSuffix(email, "@email.com") {
					return fmt.Errorf("invalid email format: %s", email)
				}
				// Check for invalid characters in the local part
				localPart := strings.Split(email, "@")[0]
				for _, char := range localPart {
					if !strings.ContainsRune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789", char) {
						return fmt.Errorf("invalid character in email: %c", char)
					}
				}
				return nil
			},
		},
		{
			name:  "Consistency Across Multiple Calls",
			input: 10,
			validate: func(email string) error {
				// This test will run multiple times to check for randomness
				return nil
			},
		},
		{
			name:  "Email Generation Within a Specific Time Frame",
			input: 10,
			validate: func(email string) error {
				start := time.Now()
				_ = RandomEmail(10)
				duration := time.Since(start)
				if duration > time.Millisecond {
					return fmt.Errorf("email generation took too long: %v", duration)
				}
				return nil
			},
		},
		{
			name:  "Upper Bound Character Validation",
			input: 10,
			validate: func(email string) error {
				localPart := strings.Split(email, "@")[0]
				for _, char := range localPart {
					if !strings.ContainsRune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789", char) {
						return fmt.Errorf("invalid character in email: %c", char)
					}
				}
				return nil
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			email := RandomEmail(tt.input)
			if err := tt.validate(email); err != nil {
				t.Errorf("Test %s failed: %v", tt.name, err)
			} else {
				t.Logf("Test %s passed: %s", tt.name, email)
			}
		})
	}

	t.Run("Consistency Across Multiple Calls", func(t *testing.T) {
		emails := make(map[string]struct{})
		for i := 0; i < 100; i++ {
			email := RandomEmail(10)
			if _, exists := emails[email]; exists {
				t.Errorf("Duplicate email found: %s", email)
			}
			emails[email] = struct{}{}
		}
		t.Log("Test Consistency Across Multiple Calls passed")
	})
}
