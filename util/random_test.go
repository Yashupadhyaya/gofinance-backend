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

	isValidEmail := func(email string) bool {
		rx := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@email\.com$`)
		return rx.MatchString(email)
	}

	tests := []struct {
		name        string
		input       int
		expectError bool
		validate    func(string) bool
	}{
		{
			name:  "Generate Email with 5-Character Local Part",
			input: 5,
			validate: func(email string) bool {
				localPart := strings.Split(email, "@")[0]
				return len(localPart) == 5 && strings.HasSuffix(email, "@email.com")
			},
		},
		{
			name:  "Generate Email with 0-Character Local Part",
			input: 0,
			validate: func(email string) bool {
				return email == "@email.com"
			},
		},
		{
			name:  "Generate Email with Maximum Character Local Part",
			input: 1000,
			validate: func(email string) bool {
				localPart := strings.Split(email, "@")[0]
				return len(localPart) == 1000 && strings.HasSuffix(email, "@email.com")
			},
		},
		{
			name:  "Validate Email Format",
			input: 10,
			validate: func(email string) bool {
				return isValidEmail(email)
			},
		},
		{
			name:  "Generate Email with Special Characters in Local Part",
			input: 10,
			validate: func(email string) bool {

				const specialAlphabet = "abcdefghijklmnopqrstuvwxyz!#$%&'*+-/=?^_`{|}~"
				originalAlphabet := alphabet
				defer func() { alphabet = originalAlphabet }()
				alphabet = specialAlphabet

				localPart := strings.Split(email, "@")[0]
				for _, char := range localPart {
					if !strings.ContainsRune(specialAlphabet, char) {
						return false
					}
				}
				return strings.HasSuffix(email, "@email.com")
			},
		},
		{
			name:  "Generate Multiple Emails with Unique Local Parts",
			input: 10,
			validate: func(email string) bool {
				emails := make(map[string]struct{})
				for i := 0; i < 100; i++ {
					email := RandomEmail(10)
					localPart := strings.Split(email, "@")[0]
					if _, exists := emails[localPart]; exists {
						return false
					}
					emails[localPart] = struct{}{}
				}
				return true
			},
		},
		{
			name:  "Generate Email with Negative Input",
			input: -1,
			validate: func(email string) bool {
				return email == "@email.com"
			},
		},
		{
			name:  "Generate Email with Non-ASCII Characters in Alphabet",
			input: 10,
			validate: func(email string) bool {

				const nonASCIIAlphabet = "abcdefghijklmnopqrstuvwxyzäöüß"
				originalAlphabet := alphabet
				defer func() { alphabet = originalAlphabet }()
				alphabet = nonASCIIAlphabet

				localPart := strings.Split(email, "@")[0]
				for _, char := range localPart {
					if !strings.ContainsRune(nonASCIIAlphabet, char) {
						return false
					}
				}
				return strings.HasSuffix(email, "@email.com")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			email := RandomEmail(tt.input)
			if !tt.validate(email) {
				t.Errorf("Test failed for input %d: %s", tt.input, email)
			} else {
				t.Logf("Test passed for input %d: %s", tt.input, email)
			}
		})
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

/*
ROOST_METHOD_HASH=RandomString_d7e3599ac4
ROOST_METHOD_SIG_HASH=RandomString_c6fe4ad19a


 */
func TestRandomString(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	const alphabet = "abcdefghijklmnopqrstuvwxyz"

	tests := []struct {
		name         string
		input        int
		expectedLen  int
		validateFunc func(string) error
	}{
		{
			name:        "Generating Random String of Specified Length",
			input:       10,
			expectedLen: 10,
			validateFunc: func(s string) error {
				if len(s) != 10 {
					return fmt.Errorf("expected length 10, got %d", len(s))
				}
				return nil
			},
		},
		{
			name:        "Generating Random String with Zero Length",
			input:       0,
			expectedLen: 0,
			validateFunc: func(s string) error {
				if len(s) != 0 {
					return fmt.Errorf("expected length 0, got %d", len(s))
				}
				return nil
			},
		},
		{
			name:        "Randomness of Generated String",
			input:       5,
			expectedLen: 5,
			validateFunc: func(s string) error {
				s2 := RandomString(5)
				if s == s2 {
					return fmt.Errorf("expected different strings, got identical strings %s and %s", s, s2)
				}
				return nil
			},
		},
		{
			name:        "Valid Characters in Generated String",
			input:       50,
			expectedLen: 50,
			validateFunc: func(s string) error {
				for _, c := range s {
					if !strings.ContainsRune(alphabet, c) {
						return fmt.Errorf("invalid character %c in string", c)
					}
				}
				return nil
			},
		},
		{
			name:        "Generating Random String of Maximum Reasonable Length",
			input:       10000,
			expectedLen: 10000,
			validateFunc: func(s string) error {
				if len(s) != 10000 {
					return fmt.Errorf("expected length 10000, got %d", len(s))
				}
				return nil
			},
		},
		{
			name:        "Multiple Consecutive Calls",
			input:       20,
			expectedLen: 20,
			validateFunc: func(s string) error {
				for i := 0; i < 10; i++ {
					s := RandomString(20)
					if len(s) != 20 {
						return fmt.Errorf("expected length 20, got %d", len(s))
					}
					for _, c := range s {
						if !strings.ContainsRune(alphabet, c) {
							return fmt.Errorf("invalid character %c in string", c)
						}
					}
				}
				return nil
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RandomString(tt.input)
			if err := tt.validateFunc(result); err != nil {
				t.Errorf("test %s failed: %v", tt.name, err)
			} else {
				t.Logf("test %s passed", tt.name)
			}
		})
	}
}

