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

func RandomString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	sb := strings.Builder{}
	for i := 0; i < n; i++ {
		sb.WriteByte(letters[rand.Intn(len(letters))])
	}
	return sb.String()
}

func TestRandomEmail(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	tests := []struct {
		name		string
		number		int
		expected	string
	}{{name: "Standard Length", number: 10, expected: "10 characters before @email.com"}, {name: "Minimum Length", number: 1, expected: "1 character before @email.com"}, {name: "Zero Length", number: 0, expected: "@email.com"}, {name: "Large Length", number: 1000, expected: "1000 characters before @email.com"}, {name: "Randomness Verification", number: 10, expected: "randomness check"}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			email := RandomEmail(tt.number)
			localPart := strings.Split(email, "@")[0]
			switch tt.name {
			case "Standard Length", "Minimum Length", "Large Length":
				if len(localPart) != tt.number {
					t.Errorf("Failed %s: expected local part length %d, got %d", tt.name, tt.number, len(localPart))
				} else {
					t.Logf("Success %s: local part length is %d", tt.name, len(localPart))
				}
			case "Zero Length":
				if email != tt.expected {
					t.Errorf("Failed %s: expected %s, got %s", tt.name, tt.expected, email)
				} else {
					t.Logf("Success %s: email is %s", tt.name, email)
				}
			case "Randomness Verification":
				emails := map[string]bool{}
				for i := 0; i < 5; i++ {
					email := RandomEmail(tt.number)
					if emails[email] {
						t.Errorf("Failed %s: duplicate email generated %s", tt.name, email)
					}
					emails[email] = true
				}
				t.Logf("Success %s: generated emails are unique", tt.name)
			}
		})
	}
}

/*
ROOST_METHOD_HASH=RandomString_d7e3599ac4
ROOST_METHOD_SIG_HASH=RandomString_c6fe4ad19a


 */
func RandomString(length int) string {
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < length; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func TestRandomString(t *testing.T) {
	testCases := []struct {
		description	string
		length		int
		validate	func(string) bool
	}{{description: "Generate a Random String of Positive Length", length: 10, validate: func(result string) bool {
		return len(result) == 10
	}}, {description: "Generate a Random String of Zero Length", length: 0, validate: func(result string) bool {
		return result == ""
	}}, {description: "Generate a Random String with Negative Length", length: -5, validate: func(result string) bool {
		return result == ""
	}}, {description: "Verify Randomness of Generated Strings", length: 10, validate: func(result string) bool {
		anotherResult := RandomString(10)
		return result != anotherResult
	}}, {description: "Test Upper Bound of String Length", length: 100000, validate: func(result string) bool {
		return len(result) == 100000
	}}, {description: "Validate Characters Used in Random String", length: 10, validate: func(result string) bool {
		for _, char := range result {
			if !strings.ContainsRune(alphabet, char) {
				return false
			}
		}
		return true
	}}}
	rand.Seed(time.Now().UnixNano())
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			result := RandomString(tc.length)
			if !tc.validate(result) {
				t.Errorf("Failed %s: got %v", tc.description, result)
			} else {
				t.Logf("Passed %s: got %v", tc.description, result)
			}
		})
	}
}

