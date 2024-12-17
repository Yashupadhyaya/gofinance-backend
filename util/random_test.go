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
	type testCase struct {
		description string
		input       int
		expected    string
		validate    func(email string) (bool, string)
	}

	tests := []testCase{
		{
			description: "Generate Email with 5-Character Local Part",
			input:       5,
			validate: func(email string) (bool, string) {
				expectedSuffix := "@email.com"
				if !strings.HasSuffix(email, expectedSuffix) {
					return false, fmt.Sprintf("Expected suffix %s, got %s", expectedSuffix, email)
				}
				localPart := strings.TrimSuffix(email, expectedSuffix)
				if len(localPart) != 5 {
					return false, fmt.Sprintf("Expected local part length 5, got %d", len(localPart))
				}
				return true, ""
			},
		},
		{
			description: "Generate Email with Zero-Character Local Part",
			input:       0,
			validate: func(email string) (bool, string) {
				expectedEmail := "@email.com"
				if email != expectedEmail {
					return false, fmt.Sprintf("Expected email %s, got %s", expectedEmail, email)
				}
				return true, ""
			},
		},
		{
			description: "Generate Email with Maximum Character Local Part",
			input:       256,
			validate: func(email string) (bool, string) {
				expectedSuffix := "@email.com"
				if !strings.HasSuffix(email, expectedSuffix) {
					return false, fmt.Sprintf("Expected suffix %s, got %s", expectedSuffix, email)
				}
				localPart := strings.TrimSuffix(email, expectedSuffix)
				if len(localPart) != 256 {
					return false, fmt.Sprintf("Expected local part length 256, got %d", len(localPart))
				}
				return true, ""
			},
		},
		{
			description: "Generate Email with Negative Character Local Part",
			input:       -1,
			validate: func(email string) (bool, string) {
				expectedEmail := "@email.com"
				if email != expectedEmail {
					return false, fmt.Sprintf("Expected email %s, got %s", expectedEmail, email)
				}
				return true, ""
			},
		},
		{
			description: "Generate Email with Non-Alphabet Characters",
			input:       10,
			validate: func(email string) (bool, string) {
				expectedSuffix := "@email.com"
				if !strings.HasSuffix(email, expectedSuffix) {
					return false, fmt.Sprintf("Expected suffix %s, got %s", expectedSuffix, email)
				}
				localPart := strings.TrimSuffix(email, expectedSuffix)
				for _, char := range localPart {
					if !strings.ContainsRune(alphabet, char) {
						return false, fmt.Sprintf("Local part contains non-alphabet character: %c", char)
					}
				}
				return true, ""
			},
		},
		{
			description: "Generate Email with Consistent Domain",
			input:       10,
			validate: func(email string) (bool, string) {
				expectedSuffix := "@email.com"
				if !strings.HasSuffix(email, expectedSuffix) {
					return false, fmt.Sprintf("Expected suffix %s, got %s", expectedSuffix, email)
				}
				return true, ""
			},
		},
		{
			description: "Generate Email with Randomness Verification",
			input:       10,
			validate: func(email string) (bool, string) {
				email1 := RandomEmail(10)
				email2 := RandomEmail(10)
				if email1 == email2 {
					return false, "Expected different emails on multiple invocations"
				}
				return true, ""
			},
		},
		{
			description: "Generate Email with Minimum Valid Character Local Part",
			input:       1,
			validate: func(email string) (bool, string) {
				expectedSuffix := "@email.com"
				if !strings.HasSuffix(email, expectedSuffix) {
					return false, fmt.Sprintf("Expected suffix %s, got %s", expectedSuffix, email)
				}
				localPart := strings.TrimSuffix(email, expectedSuffix)
				if len(localPart) != 1 {
					return false, fmt.Sprintf("Expected local part length 1, got %d", len(localPart))
				}
				return true, ""
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			email := RandomEmail(tc.input)
			if valid, msg := tc.validate(email); !valid {
				t.Errorf("Test failed: %s\nReason: %s", tc.description, msg)
			} else {
				t.Logf("Test passed: %s", tc.description)
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
		name     string
		input    int
		validate func(t *testing.T, result string)
	}{
		{
			name:  "Generate a Random String of Specified Length",
			input: 10,
			validate: func(t *testing.T, result string) {
				if len(result) != 10 {
					t.Errorf("Expected string length of 10, but got %d", len(result))
				}
			},
		},
		{
			name:  "Generate a Random String of Zero Length",
			input: 0,
			validate: func(t *testing.T, result string) {
				if result != "" {
					t.Errorf("Expected empty string, but got %s", result)
				}
			},
		},
		{
			name:  "Generate a Random String Contains Only Alphabet Characters",
			input: 15,
			validate: func(t *testing.T, result string) {
				for _, char := range result {
					if !strings.ContainsRune(alphabet, char) {
						t.Errorf("Character %c not in alphabet", char)
					}
				}
			},
		},
		{
			name:  "Generate a Random String Multiple Times and Verify Uniqueness",
			input: 10,
			validate: func(t *testing.T, result string) {
				uniqueStrings := make(map[string]struct{})
				for i := 0; i < 5; i++ {
					str := RandomString(10)
					if _, exists := uniqueStrings[str]; exists {
						t.Errorf("Duplicate string found: %s", str)
					}
					uniqueStrings[str] = struct{}{}
				}
			},
		},
		{
			name:  "Generate a Random String with Maximum Length",
			input: 1000000,
			validate: func(t *testing.T, result string) {
				if len(result) != 1000000 {
					t.Errorf("Expected string length of 1000000, but got %d", len(result))
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("Running test scenario: %s", tt.name)
			result := RandomString(tt.input)
			tt.validate(t, result)
		})
	}
}

