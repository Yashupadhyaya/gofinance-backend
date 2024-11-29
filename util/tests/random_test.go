package util

import (
	"strings"
	"testing"
	"math/rand"
	"time"
)

/*
ROOST_METHOD_HASH=RandomEmail_1905439733
ROOST_METHOD_SIG_HASH=RandomEmail_7a04f189fd


 */
func TestRandomEmail(t *testing.T) {
	type testCase struct {
		name           string
		input          int
		expectedFormat string
		expectedLength int
	}

	testCases := []testCase{
		{
			name:           "Generate Email with Standard Length",
			input:          10,
			expectedFormat: "@email.com",
			expectedLength: 10,
		},
		{
			name:           "Generate Email with Minimum Length",
			input:          1,
			expectedFormat: "@email.com",
			expectedLength: 1,
		},
		{
			name:           "Generate Email with Zero Length",
			input:          0,
			expectedFormat: "@email.com",
			expectedLength: 0,
		},
		{
			name:           "Generate Email with Maximum Length",
			input:          1000,
			expectedFormat: "@email.com",
			expectedLength: 1000,
		},
		{
			name:           "Validate Email Format",
			input:          5,
			expectedFormat: "@email.com",
			expectedLength: 5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			email := RandomEmail(tc.input)
			if !strings.HasSuffix(email, tc.expectedFormat) {
				t.Errorf("Failed %s: expected email to end with %s, got %s", tc.name, tc.expectedFormat, email)
			}

			randomPart := strings.TrimSuffix(email, tc.expectedFormat)
			if len(randomPart) != tc.expectedLength {
				t.Errorf("Failed %s: expected random part length to be %d, got %d", tc.name, tc.expectedLength, len(randomPart))
			}

			if !strings.Contains(email, "@") || strings.Count(email, "@") != 1 {
				t.Errorf("Failed %s: email does not contain exactly one @ symbol", tc.name)
			}

			t.Logf("Success %s: email format and length are as expected", tc.name)
		})
	}

	t.Run("Consistency Across Multiple Calls", func(t *testing.T) {
		email1 := RandomEmail(8)
		email2 := RandomEmail(8)

		if email1 == email2 {
			t.Errorf("Failed Consistency Across Multiple Calls: expected different email addresses, got %s and %s", email1, email2)
		} else {
			t.Log("Success Consistency Across Multiple Calls: different email addresses generated")
		}
	})
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

func Testrandomstring(t *testing.T) {

	rand.Seed(time.Now().UnixNano())

	type testCase struct {
		name      string
		length    int
		expected  string
		checkFunc func(string, string) bool
	}

	testCases := []testCase{
		{
			name:     "Generate a Random String of Positive Length",
			length:   10,
			expected: "length 10",
			checkFunc: func(got, _ string) bool {
				return len(got) == 10
			},
		},
		{
			name:     "Generate an Empty String for Zero Length",
			length:   0,
			expected: "",
			checkFunc: func(got, expected string) bool {
				return got == expected
			},
		},
		{
			name:     "Consistent Randomness with Seeded Random Generator",
			length:   10,
			expected: "consistent output",
			checkFunc: func(got, _ string) bool {
				rand.Seed(42)
				firstOutput := RandomString(10)
				rand.Seed(42)
				secondOutput := RandomString(10)
				return firstOutput == secondOutput
			},
		},
		{
			name:     "Generate String with Maximum Length",
			length:   1000000,
			expected: "length 1000000",
			checkFunc: func(got, _ string) bool {
				return len(got) == 1000000
			},
		},
		{
			name:     "Validate Character Set Usage",
			length:   50,
			expected: "valid characters",
			checkFunc: func(got, _ string) bool {
				for _, char := range got {
					if !strings.ContainsRune(alphabet, char) {
						return false
					}
				}
				return true
			},
		},
		{
			name:     "Handle Negative Length Input Gracefully",
			length:   -5,
			expected: "",
			checkFunc: func(got, expected string) bool {
				return got == expected
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := RandomString(tc.length)
			if !tc.checkFunc(result, tc.expected) {
				t.Errorf("Test %s failed: expected %s, got %s", tc.name, tc.expected, result)
			} else {
				t.Logf("Test %s passed", tc.name)
			}
		})
	}
}

