package util

import (
	"math/rand"
	"strings"
	"testing"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"
const alphabet = "abcdefghijklmnopqrstuvwxyz"
/*
ROOST_METHOD_HASH=RandomEmail_1905439733
ROOST_METHOD_SIG_HASH=RandomEmail_7a04f189fd


 */
func RandomEmail(number int) string {
	return fmt.Sprintf("%s@email.com", RandomString(number))
}

func RandomString(n int) string {
	sb := strings.Builder{}
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func TestRandomEmail(t *testing.T) {
	type testScenario struct {
		description string
		input       int
		expected    string
		validate    func(string, int) bool
	}

	validateEmailFormat := func(email string, length int) bool {
		parts := strings.Split(email, "@")
		return len(parts) == 2 && len(parts[0]) == length && parts[1] == "email.com"
	}

	validateUniqueEmails := func(emails []string) bool {
		emailSet := make(map[string]struct{})
		for _, email := range emails {
			if _, exists := emailSet[email]; exists {
				return false
			}
			emailSet[email] = struct{}{}
		}
		return true
	}

	testScenarios := []testScenario{
		{
			description: "Generating a Random Email with a Valid Length",
			input:       10,
			expected:    "",
			validate:    validateEmailFormat,
		},
		{
			description: "Generating a Random Email with Zero Length",
			input:       0,
			expected:    "@email.com",
			validate: func(email string, _ int) bool {
				return email == "@email.com"
			},
		},
		{
			description: "Generating a Random Email with Maximum Integer Length",
			input:       int(^uint(0) >> 1),
			expected:    "",
			validate: func(email string, _ int) bool {
				return email != ""
			},
		},
		{
			description: "Generating a Random Email with Negative Length",
			input:       -5,
			expected:    "@email.com",
			validate: func(email string, _ int) bool {
				return email == "@email.com"
			},
		},
		{
			description: "Generating Multiple Random Emails with the Same Length",
			input:       8,
			expected:    "",
			validate: func(_ string, length int) bool {
				emails := make([]string, 5)
				for i := 0; i < 5; i++ {
					emails[i] = RandomEmail(length)
				}
				return validateUniqueEmails(emails)
			},
		},
	}

	rand.Seed(time.Now().UnixNano())

	for _, scenario := range testScenarios {
		t.Run(scenario.description, func(t *testing.T) {
			t.Logf("Running scenario: %s", scenario.description)

			result := RandomEmail(scenario.input)
			if !scenario.validate(result, scenario.input) {
				t.Errorf("Test failed for scenario: %s\nExpected: %v\nGot: %v", scenario.description, scenario.expected, result)
			} else {
				t.Logf("Success: %s", scenario.description)
			}
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

	rand.Seed(time.Now().UTC().UnixNano())

	tests := []struct {
		name        string
		length      int
		expectedLen int
	}{
		{
			name:        "Generate a Random String of Specified Length",
			length:      10,
			expectedLen: 10,
		},
		{
			name:        "Zero Length Input",
			length:      0,
			expectedLen: 0,
		},
		{
			name:        "Large Length Input",
			length:      10000,
			expectedLen: 10000,
		},
		{
			name:        "Consistency with Same Seed",
			length:      5,
			expectedLen: 5,
		},
		{
			name:        "Unique Characters in Short Strings",
			length:      2,
			expectedLen: 2,
		},
		{
			name:        "Character Set Validation",
			length:      50,
			expectedLen: 50,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "Consistency with Same Seed" {

				rand.Seed(42)
				first := RandomString(tt.length)
				second := RandomString(tt.length)
				if first != second {
					t.Errorf("expected consistent output with fixed seed, got %s and %s", first, second)
				}
				t.Log("Consistency with Same Seed passed")
				return
			}

			result := RandomString(tt.length)
			if len(result) != tt.expectedLen {
				t.Errorf("expected length %d, got %d", tt.expectedLen, len(result))
			}

			if tt.name == "Large Length Input" {

				for _, c := range result {
					if !strings.ContainsRune(alphabet, c) {
						t.Errorf("unexpected character %c in result", c)
					}
				}
				t.Log("Large Length Input passed")
			}

			if tt.name == "Unique Characters in Short Strings" {
				uniqueChars := make(map[rune]bool)
				for _, c := range result {
					if _, exists := uniqueChars[c]; !exists {
						uniqueChars[c] = true
					}
				}
				if len(uniqueChars) < 2 && tt.length > 1 {
					t.Errorf("expected more unique characters, got %d", len(uniqueChars))
				}
				t.Log("Unique Characters in Short Strings passed")
			}

			if tt.name == "Character Set Validation" {
				for _, c := range result {
					if !strings.ContainsRune(alphabet, c) {
						t.Errorf("unexpected character %c in result", c)
					}
				}
				t.Log("Character Set Validation passed")
			}

			t.Logf("%s passed", tt.name)
		})
	}
}

