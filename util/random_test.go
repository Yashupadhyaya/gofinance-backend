package util

import (
	"fmt"
	"math/rand"
	"regexp"
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
		name      string
		input     int
		validate  func(got string) error
		wantError bool
	}{
		{
			name:  "Generate Email with 5-Character Local Part",
			input: 5,
			validate: func(got string) error {
				expectedLength := 5
				if len(got) != expectedLength+10 {
					return fmt.Errorf("expected length %d, got %d", expectedLength+10, len(got))
				}
				return nil
			},
			wantError: false,
		},
		{
			name:  "Generate Email with 0-Character Local Part",
			input: 0,
			validate: func(got string) error {
				expected := "@email.com"
				if got != expected {
					return fmt.Errorf("expected %s, got %s", expected, got)
				}
				return nil
			},
			wantError: false,
		},
		{
			name:  "Generate Email with Maximum Character Local Part",
			input: 1000,
			validate: func(got string) error {
				expectedLength := 1000
				if len(got) != expectedLength+10 {
					return fmt.Errorf("expected length %d, got %d", expectedLength+10, len(got))
				}
				return nil
			},
			wantError: false,
		},
		{
			name:  "Validate Email Format",
			input: 10,
			validate: func(got string) error {
				match, _ := regexp.MatchString("^[a-z]{10}@email.com$", got)
				if !match {
					return fmt.Errorf("email format is incorrect: %s", got)
				}
				return nil
			},
			wantError: false,
		},
		{
			name:  "Generate Email with Consistent Length",
			input: 10,
			validate: func(got string) error {
				expectedLength := 10
				if len(got) != expectedLength+10 {
					return fmt.Errorf("expected length %d, got %d", expectedLength+10, len(got))
				}
				return nil
			},
			wantError: false,
		},
		{
			name:  "Performance Test for Large Input",
			input: 10000,
			validate: func(got string) error {
				start := time.Now()
				_ = RandomEmail(10000)
				duration := time.Since(start)
				if duration > time.Second {
					return fmt.Errorf("function took too long: %v", duration)
				}
				return nil
			},
			wantError: false,
		},
		{
			name:  "Generate Unique Emails for Sequential Calls",
			input: 10,
			validate: func(got string) error {
				emails := make(map[string]bool)
				for i := 0; i < 100; i++ {
					email := RandomEmail(10)
					if emails[email] {
						return fmt.Errorf("duplicate email found: %s", email)
					}
					emails[email] = true
				}
				return nil
			},
			wantError: false,
		},
		{
			name:  "Stress Test with Multiple Concurrent Calls",
			input: 10,
			validate: func(got string) error {
				const numGoroutines = 100
				results := make(chan string, numGoroutines)

				for i := 0; i < numGoroutines; i++ {
					go func() {
						results <- RandomEmail(10)
					}()
				}

				emails := make(map[string]bool)
				for i := 0; i < numGoroutines; i++ {
					email := <-results
					if emails[email] {
						return fmt.Errorf("duplicate email found: %s", email)
					}
					emails[email] = true
				}
				return nil
			},
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RandomEmail(tt.input)
			err := tt.validate(got)
			if (err != nil) != tt.wantError {
				t.Errorf("TestRandomEmail() error = %v, wantError %v", err, tt.wantError)
			}
			if err != nil {
				t.Logf("Failure reason: %v", err)
			} else {
				t.Logf("Success: %s", tt.name)
			}
		})
	}
}

/*
ROOST_METHOD_HASH=RandomString_d7e3599ac4
ROOST_METHOD_SIG_HASH=RandomString_c6fe4ad19a


 */
func TestRandomString(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	tests := []struct {
		name        string
		input       int
		expectedLen int
		expectEmpty bool
		expectError bool
	}{
		{
			name:        "Generate a Random String of Specified Length",
			input:       10,
			expectedLen: 10,
			expectEmpty: false,
			expectError: false,
		},
		{
			name:        "Generate a Random String with Zero Length",
			input:       0,
			expectedLen: 0,
			expectEmpty: true,
			expectError: false,
		},
		{
			name:        "Generate a Random String with a Large Length",
			input:       10000,
			expectedLen: 10000,
			expectEmpty: false,
			expectError: false,
		},
		{
			name:        "Generate a Random String with Negative Length",
			input:       -1,
			expectedLen: 0,
			expectEmpty: true,
			expectError: false,
		},
		{
			name:        "Check Randomness of Generated Strings",
			input:       10,
			expectedLen: 10,
			expectEmpty: false,
			expectError: false,
		},
		{
			name:        "Validate Characters in Generated String",
			input:       10,
			expectedLen: 10,
			expectEmpty: false,
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RandomString(tt.input)

			if len(result) != tt.expectedLen {
				t.Errorf("Expected string length %d, but got %d", tt.expectedLen, len(result))
			}

			if tt.expectEmpty && result != "" {
				t.Errorf("Expected an empty string, but got %s", result)
			}

			if tt.name == "Check Randomness of Generated Strings" {
				anotherResult := RandomString(tt.input)
				if result == anotherResult {
					t.Errorf("Expected different strings, but got the same: %s", result)
				}
			}

			if tt.name == "Validate Characters in Generated String" {
				for _, c := range result {
					if !strings.ContainsRune(alphabet, c) {
						t.Errorf("Unexpected character %c in result string", c)
					}
				}
			}
		})
	}

	t.Run("Performance Test for Random String Generation", func(t *testing.T) {
		lengths := []int{100, 1000, 10000}
		for _, length := range lengths {
			start := time.Now()
			RandomString(length)
			elapsed := time.Since(start)
			t.Logf("Time taken to generate string of length %d: %v", length, elapsed)

		}
	})
}

