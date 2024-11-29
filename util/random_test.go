package util

import (
	"fmt"
	"regexp"
	"testing"
	"math/rand"
	"strings"
	"time"
)

/*
ROOST_METHOD_HASH=RandomEmail_1905439733
ROOST_METHOD_SIG_HASH=RandomEmail_7a04f189fd


 */
func TestRandomEmail(t *testing.T) {

	validateEmailFormat := func(email string, expectedLength int) bool {
		re := regexp.MustCompile(fmt.Sprintf(`^[a-zA-Z0-9]{%d}@email\.com`, expectedLength))
		return re.MatchString(email)
	}

	type testCase struct {
		description    string
		number         int
		expectedLength int
		expectedEmail  string
	}

	testCases := []testCase{
		{
			description:    "Valid Email with Standard Length",
			number:         10,
			expectedLength: 10,
		},
		{
			description:    "Minimum Length Local Part",
			number:         1,
			expectedLength: 1,
		},
		{
			description:    "Zero Length Local Part",
			number:         0,
			expectedLength: 0,
			expectedEmail:  "@email.com",
		},
		{
			description:    "Large Length Local Part",
			number:         1000,
			expectedLength: 1000,
		},
		{
			description:    "Randomness of Email Addresses",
			number:         10,
			expectedLength: 10,
		},
		{
			description:    "Medium Length Local Part",
			number:         50,
			expectedLength: 50,
		},
		{
			description:    "Empty Email Generation",
			number:         0,
			expectedLength: 0,
			expectedEmail:  "@email.com",
		},
		{
			description:    "Single Character Email",
			number:         1,
			expectedLength: 1,
		},
		{
			description:    "Short Length Local Part",
			number:         5,
			expectedLength: 5,
		},
		{
			description:    "Long Length Local Part",
			number:         100,
			expectedLength: 100,
		},
		{
			description:    "Edge Case for Length Just Above Zero",
			number:         2,
			expectedLength: 2,
		},
		{
			description:    "Edge Case for Length Just Below Standard",
			number:         9,
			expectedLength: 9,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {

			email := RandomEmail(tc.number)

			if tc.expectedEmail != "" {
				if email != tc.expectedEmail {
					t.Errorf("Expected email: %s, but got: %s", tc.expectedEmail, email)
				}
			} else {
				if !validateEmailFormat(email, tc.expectedLength) {
					t.Errorf("Email format validation failed for email: %s", email)
				}
			}

			if tc.description == "Randomness of Email Addresses" {
				email2 := RandomEmail(tc.number)
				if email == email2 {
					t.Errorf("Expected different emails on repeated calls, but got the same: %s", email)
				}
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

	const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	tests := []struct {
		name   string
		number int
		want   int
	}{
		{"Generate a Random String of Specified Length", 10, 10},
		{"Generate a Random String of Length Zero", 0, 0},
		{"Generate a Random String with Negative Length", -5, 0},
		{"Generate Strings with Maximum Length", 10000, 10000},
		{"Generate Strings with Different Lengths 1", 1, 1},
		{"Generate Strings with Different Lengths 5", 5, 5},
		{"Generate Strings with Different Lengths 50", 50, 50},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RandomString(tt.number)
			if len(got) != tt.want {
				t.Errorf("RandomString(%d) = %v, want %v", tt.number, len(got), tt.want)
			}
		})
	}

	t.Run("Generate Multiple Random Strings and Verify Uniqueness", func(t *testing.T) {
		length := 10
		generatedStrings := map[string]bool{}
		for i := 0; i < 100; i++ {
			str := RandomString(length)
			if _, exists := generatedStrings[str]; exists {
				t.Errorf("Duplicate string found: %v", str)
			}
			generatedStrings[str] = true
		}
	})

	t.Run("Consistency of Alphabet Used", func(t *testing.T) {
		length := 50
		result := RandomString(length)
		for _, char := range result {
			if !strings.ContainsRune(alphabet, char) {
				t.Errorf("Character %c is not in the predefined alphabet", char)
			}
		}
	})
}

