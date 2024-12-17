package util

import (
	"math/rand"
	"strings"
	"testing"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"/*
ROOST_METHOD_HASH=RandomEmail_1905439733
ROOST_METHOD_SIG_HASH=RandomEmail_7a04f189fd


 */
func TestRandomEmail(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	type testCase struct {
		desc     string
		input    int
		validate func(string) bool
	}

	testCases := []testCase{
		{
			desc:  "Valid Email with Standard Length",
			input: 10,
			validate: func(result string) bool {
				localPart := strings.Split(result, "@")[0]
				return len(localPart) == 10 && strings.HasSuffix(result, "@email.com")
			},
		},
		{
			desc:  "Minimum Length Email",
			input: 1,
			validate: func(result string) bool {
				localPart := strings.Split(result, "@")[0]
				return len(localPart) == 1 && strings.HasSuffix(result, "@email.com")
			},
		},
		{
			desc:  "Maximum Length Email",
			input: 64,
			validate: func(result string) bool {
				localPart := strings.Split(result, "@")[0]
				return len(localPart) == 64 && strings.HasSuffix(result, "@email.com")
			},
		},
		{
			desc:  "Zero Length Email",
			input: 0,
			validate: func(result string) bool {
				localPart := strings.Split(result, "@")[0]
				return len(localPart) == 0 && strings.HasSuffix(result, "@email.com")
			},
		},
		{
			desc:  "Negative Length Email",
			input: -5,
			validate: func(result string) bool {
				localPart := strings.Split(result, "@")[0]
				return len(localPart) == 0 && strings.HasSuffix(result, "@email.com")
			},
		},
		{
			desc:  "Randomness of Generated Emails",
			input: 10,
			validate: func(result string) bool {

				anotherResult := RandomEmail(10)
				return result != anotherResult
			},
		},
		{
			desc:  "Valid Characters in Local Part",
			input: 10,
			validate: func(result string) bool {
				localPart := strings.Split(result, "@")[0]
				for _, char := range localPart {
					if !strings.ContainsRune(alphabet, char) {
						return false
					}
				}
				return true && strings.HasSuffix(result, "@email.com")
			},
		},
		{
			desc:  "Consistent Domain Part",
			input: 10,
			validate: func(result string) bool {
				return strings.HasSuffix(result, "@email.com")
			},
		},
		{
			desc:  "Email Length Validation",
			input: 10,
			validate: func(result string) bool {
				expectedLength := 10 + len("@email.com")
				return len(result) == expectedLength
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			result := RandomEmail(tc.input)
			if !tc.validate(result) {
				t.Errorf("Test case %s failed. Got %s", tc.desc, result)
			} else {
				t.Logf("Test case %s succeeded. Got %s", tc.desc, result)
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

	rand.Seed(time.Now().UnixNano())

	tests := []struct {
		name     string
		length   int
		validate func(result string) bool
	}{
		{
			name:   "Generating Random String of Zero Length",
			length: 0,
			validate: func(result string) bool {
				return result == ""
			},
		},
		{
			name:   "Generating Random String of Positive Length",
			length: 10,
			validate: func(result string) bool {
				return len(result) == 10
			},
		},
		{
			name:   "Randomness Check for Generated Strings",
			length: 10,
			validate: func(result string) bool {
				first := RandomString(10)
				second := RandomString(10)
				return first != second
			},
		},
		{
			name:   "Boundary Test for Maximum Length String",
			length: 1000000,
			validate: func(result string) bool {
				return len(result) == 1000000
			},
		},
		{
			name:   "Consistency Check for Seeded Randomness",
			length: 10,
			validate: func(result string) bool {
				rand.Seed(42)
				first := RandomString(10)
				rand.Seed(42)
				second := RandomString(10)
				return first == second
			},
		},
		{
			name:   "Character Validation in Generated String",
			length: 50,
			validate: func(result string) bool {
				for _, char := range result {
					if !strings.ContainsRune(alphabet, char) {
						return false
					}
				}
				return true
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RandomString(tt.length)
			if !tt.validate(result) {
				t.Errorf("Test %s failed: got %v", tt.name, result)
			} else {
				t.Logf("Test %s passed", tt.name)
			}
		})
	}
}

