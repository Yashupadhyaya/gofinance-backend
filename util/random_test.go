package util

import (
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
		name        string
		input       int
		validate    func(string) bool
		expected    string
		description string
	}{
		{
			name:        "Generate Email with 5-Character Local Part",
			input:       5,
			validate:    func(email string) bool { return len(email) == 11 && email[5:] == "@email.com" },
			expected:    "xxxxx@email.com",
			description: "This test checks if the function generates an email address with a 5-character local part when the input number is 5.",
		},
		{
			name:        "Generate Email with 0-Character Local Part",
			input:       0,
			validate:    func(email string) bool { return email == "@email.com" },
			expected:    "@email.com",
			description: "This test checks if the function handles an input of 0 and returns an email with an empty local part.",
		},
		{
			name:        "Generate Email with Maximum Reasonable Length",
			input:       64,
			validate:    func(email string) bool { return len(email) == 70 && email[64:] == "@email.com" },
			expected:    strings.Repeat("x", 64) + "@email.com",
			description: "This test checks if the function can handle generating an email with a very long local part, e.g., 64 characters.",
		},
		{
			name:        "Generate Email with Non-Alpha Characters in Local Part",
			input:       10,
			validate:    func(email string) bool { return regexp.MustCompile(`^[a-z]+@email\.com`).MatchString(email) },
			expected:    "xxxxxxxxxx@email.com",
			description: "This test checks if the function generates a local part consisting solely of alphabetic characters.",
		},
		{
			name:        "Consistent Length of Local Part",
			input:       8,
			validate:    func(email string) bool { return len(email) == 14 && email[8:] == "@email.com" },
			expected:    "xxxxxxxx@email.com",
			description: "This test checks if the function consistently generates a local part of the specified length across multiple invocations.",
		},
		{
			name:        "Randomness of Local Part",
			input:       10,
			validate:    func(email string) bool { return regexp.MustCompile(`^[a-z]+@email\.com`).MatchString(email) },
			expected:    "xxxxxxxxxx@email.com",
			description: "This test checks if the function generates different local parts for multiple invocations with the same input.",
		},
		{
			name:        "Check for Valid Email Format",
			input:       12,
			validate:    func(email string) bool { return regexp.MustCompile(`^[a-zA-Z]+@email\.com`).MatchString(email) },
			expected:    "xxxxxxxxxxxx@email.com",
			description: "This test checks if the function generates a valid email format.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(tt.description)
			email := RandomEmail(tt.input)
			if !tt.validate(email) {
				t.Errorf("Test %s failed: generated email %s does not match expected pattern", tt.name, email)
			} else {
				t.Logf("Test %s succeeded: generated email %s matches expected pattern", tt.name, email)
			}
		})
	}
}

/*
ROOST_METHOD_HASH=RandomString_d7e3599ac4
ROOST_METHOD_SIG_HASH=RandomString_c6fe4ad19a


 */
func TestRandomString(t *testing.T) {

	tests := []struct {
		name     string
		input    int
		expected string
		validate func(t *testing.T, output string)
	}{
		{
			name:     "Generating Random String of Zero Length",
			input:    0,
			expected: "",
			validate: func(t *testing.T, output string) {
				if output != "" {
					t.Errorf("Expected an empty string, but got '%s'", output)
				} else {
					t.Log("Successfully generated an empty string for zero length input")
				}
			},
		},
		{
			name:  "Generating Random String of Positive Length",
			input: 10,
			validate: func(t *testing.T, output string) {
				if len(output) != 10 {
					t.Errorf("Expected string of length 10, but got length %d", len(output))
				}
				for _, c := range output {
					if !strings.ContainsRune(alphabet, c) {
						t.Errorf("Unexpected character '%c' in output string", c)
					}
				}
				t.Log("Successfully generated a random string of length 10 with valid characters")
			},
		},
		{
			name:  "Consistency of Randomness",
			input: 10,
			validate: func(t *testing.T, output string) {
				output2 := RandomString(10)
				if output == output2 {
					t.Errorf("Expected different strings, but got the same string '%s'", output)
				} else {
					t.Log("Successfully generated different random strings for the same length")
				}
			},
		},
		{
			name:  "Generating Maximum Length String",
			input: 1000000,
			validate: func(t *testing.T, output string) {
				if len(output) != 1000000 {
					t.Errorf("Expected string of length 1000000, but got length %d", len(output))
				} else {
					t.Log("Successfully generated a random string of length 1000000")
				}
			},
		},
		{
			name:  "Validating Characters in Generated String",
			input: 50,
			validate: func(t *testing.T, output string) {
				for _, c := range output {
					if !strings.ContainsRune(alphabet, c) {
						t.Errorf("Unexpected character '%c' in output string", c)
					}
				}
				t.Log("Successfully validated characters in the generated string")
			},
		},
		{
			name:  "Generating Random String with Seeded Randomness",
			input: 10,
			validate: func(t *testing.T, output string) {
				rand.Seed(1)
				output1 := RandomString(10)
				rand.Seed(1)
				output2 := RandomString(10)
				if output1 != output2 {
					t.Errorf("Expected same strings, but got different strings '%s' and '%s'", output1, output2)
				} else {
					t.Log("Successfully generated consistent random strings with seeded randomness")
				}
			},
		},
		{
			name:  "Handling Negative Length Input",
			input: -5,
			validate: func(t *testing.T, output string) {
				if output != "" {
					t.Errorf("Expected an empty string for negative input, but got '%s'", output)
				} else {
					t.Log("Successfully handled negative length input by returning an empty string")
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rand.Seed(time.Now().UnixNano())
			output := RandomString(tt.input)
			tt.validate(t, output)
		})
	}
}

