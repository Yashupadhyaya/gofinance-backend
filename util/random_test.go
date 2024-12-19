package util

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"
	"math"
)

/*
ROOST_METHOD_HASH=RandomEmail_1905439733
ROOST_METHOD_SIG_HASH=RandomEmail_7a04f189fd


 */
func TestRandomEmail(t *testing.T) {

	rand.Seed(time.Now().UnixNano())

	tests := []struct {
		name        string
		number      int
		expectError bool
	}{
		{
			name:        "Standard Length",
			number:      10,
			expectError: false,
		},
		{
			name:        "Zero Length",
			number:      0,
			expectError: false,
		},
		{
			name:        "Maximum Length",
			number:      1000,
			expectError: false,
		},
		{
			name:        "Negative Length",
			number:      -5,
			expectError: true,
		},
		{
			name:        "Special Characters",
			number:      10,
			expectError: false,
		},
		{
			name:        "Multiple Emails for Uniqueness",
			number:      10,
			expectError: false,
		},
		{
			name:        "Non-ASCII Characters",
			number:      10,
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("Starting test: %s", tt.name)

			email := RandomEmail(tt.number)
			t.Logf("Generated email: %s", email)

			atIndex := strings.Index(email, "@")
			if atIndex == -1 {
				t.Fatalf("Generated email does not contain '@': %s", email)
			}

			localPart := email[:atIndex]
			domainPart := email[atIndex:]

			t.Logf("Local part: %s, Domain part: %s", localPart, domainPart)

			switch tt.name {
			case "Standard Length":
				if len(localPart) != tt.number {
					t.Errorf("Expected local part length %d, got %d", tt.number, len(localPart))
				}

			case "Zero Length":
				if len(localPart) != 0 {
					t.Errorf("Expected local part length 0, got %d", len(localPart))
				}

			case "Maximum Length":
				if len(localPart) != tt.number {
					t.Errorf("Expected local part length %d, got %d", tt.number, len(localPart))
				}

			case "Negative Length":
				if len(localPart) != 0 {
					t.Errorf("Expected local part length 0 for negative input, got %d", len(localPart))
				}

			case "Special Characters":

				if strings.ContainsAny(localPart, "!@#$%^&*()") {
					t.Errorf("Local part contains special characters: %s", localPart)
				}

			case "Multiple Emails for Uniqueness":
				emails := make(map[string]bool)
				for i := 0; i < 100; i++ {
					email := RandomEmail(tt.number)
					if emails[email] {
						t.Errorf("Duplicate email generated: %s", email)
					}
					emails[email] = true
				}

			case "Non-ASCII Characters":

				for _, r := range localPart {
					if r > 127 {
						t.Errorf("Local part contains non-ASCII character: %s", localPart)
					}
				}
			}

			if domainPart != "@email.com" {
				t.Errorf("Expected domain part '@email.com', got %s", domainPart)
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
		name          string
		input         int
		expectedLen   int
		expectEmpty   bool
		expectDiff    bool
		expectValid   bool
		measurePerf   bool
		withinSeconds float64
	}{
		{
			name:        "Generate a Random String of Specified Length",
			input:       10,
			expectedLen: 10,
			expectEmpty: false,
			expectDiff:  false,
			expectValid: true,
		},
		{
			name:        "Generate a Random String with Zero Length",
			input:       0,
			expectedLen: 0,
			expectEmpty: true,
			expectDiff:  false,
			expectValid: true,
		},
		{
			name:        "Generate a Random String with Negative Length",
			input:       -5,
			expectedLen: 0,
			expectEmpty: true,
			expectDiff:  false,
			expectValid: true,
		},
		{
			name:          "Generate a Random String with Maximum Integer Length",
			input:         math.MaxInt32,
			expectedLen:   math.MaxInt32,
			expectEmpty:   false,
			expectDiff:    false,
			expectValid:   true,
			measurePerf:   true,
			withinSeconds: 10.0,
		},
		{
			name:        "Ensure Randomness of Generated Strings",
			input:       10,
			expectedLen: 10,
			expectEmpty: false,
			expectDiff:  true,
			expectValid: true,
		},
		{
			name:        "Validate Characters in Generated String",
			input:       10,
			expectedLen: 10,
			expectEmpty: false,
			expectDiff:  false,
			expectValid: true,
		},
		{
			name:          "Check Performance for Large String Generation",
			input:         1000000,
			expectedLen:   1000000,
			expectEmpty:   false,
			expectValid:   true,
			measurePerf:   true,
			withinSeconds: 2.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			startTime := time.Now()

			got := RandomString(tt.input)

			if tt.measurePerf {
				duration := time.Since(startTime).Seconds()
				if duration > tt.withinSeconds {
					t.Errorf("performance test failed: got duration %v, expected within %v seconds", duration, tt.withinSeconds)
				} else {
					t.Logf("performance test passed: got duration %v, expected within %v seconds", duration, tt.withinSeconds)
				}
			}

			if len(got) != tt.expectedLen {
				t.Errorf("unexpected string length: got %v, want %v", len(got), tt.expectedLen)
			} else {
				t.Logf("string length test passed: got %v, want %v", len(got), tt.expectedLen)
			}

			if tt.expectEmpty && got != "" {
				t.Errorf("expected empty string, got %v", got)
			} else if !tt.expectEmpty && got == "" {
				t.Errorf("expected non-empty string, got empty string")
			}

			if tt.expectDiff {
				got2 := RandomString(tt.input)
				if got == got2 {
					t.Errorf("expected different strings on multiple invocations, got same strings: %v and %v", got, got2)
				} else {
					t.Logf("randomness test passed: got different strings: %v and %v", got, got2)
				}
			}

			if tt.expectValid {
				for _, c := range got {
					if !strings.ContainsRune(alphabet, c) {
						t.Errorf("unexpected character in generated string: %v", c)
					}
				}
				t.Logf("character validation test passed")
			}
		})
	}
}

