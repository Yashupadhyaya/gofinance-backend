package util_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/yourusername/yourproject/util"
)

/*
ROOST_METHOD_HASH=RandomEmail_1905439733
ROOST_METHOD_SIG_HASH=RandomEmail_7a04f189fd


 */
func Testrandomemail(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	testCases := []struct {
		name        string
		input       int
		expectError bool
	}{
		{"Valid Input with Single Character Length", 1, false},
		{"Valid Input with Multiple Characters Length", 5, false},
		{"Zero Character Length", 0, false},
		{"Negative Character Length", -5, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Log("Scenario:", tc.name)

			result := RandomEmail(tc.input)

			if tc.input < 0 {
				tc.input = 0
			}

			expectedLength := 10 + tc.input

			if len(result) != expectedLength {
				t.Errorf("Expected email of length %d, but got %d", expectedLength, len(result))
			} else {
				t.Logf("Success: Length of email is as expected: %d", len(result))
			}

			if !strings.Contains(result, "@email.com") {
				t.Errorf("Expected email to contain '@email.com', but it did not: %s", result)
			} else {
				t.Logf("Success: Email contains '@email.com'")
			}
		})
	}
}

/*
ROOST_METHOD_HASH=RandomString_d7e3599ac4
ROOST_METHOD_SIG_HASH=RandomString_c6fe4ad19a


 */
func TestRandomString(t *testing.T) {
	testCases := []struct {
		desc     string
		input    int
		expected int
	}{
		{
			desc:     "Testing RandomString with a positive number",
			input:    5,
			expected: 5,
		},
		{
			desc:     "Testing RandomString with zero",
			input:    0,
			expected: 0,
		},
		{
			desc:     "Testing RandomString with a negative number",
			input:    -3,
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			result := util.RandomString(tc.input)
			assert.Equal(t, tc.expected, len(result), "they should be equal")
		})
	}
}

