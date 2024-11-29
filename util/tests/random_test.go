package util_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/yourusername/yourproject/util"
)

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

