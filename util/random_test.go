package util_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/yourorg/util"
)

/*
ROOST_METHOD_HASH=RandomEmail_1905439733
ROOST_METHOD_SIG_HASH=RandomEmail_7a04f189fd


 */
func TestRandomEmail(t *testing.T) {

	testCases := []struct {
		name   string
		length int
		want   int
	}{
		{
			name:   "Valid Input with Single Character Length",
			length: 1,
			want:   11,
		},
		{
			name:   "Valid Input with Multiple Characters Length",
			length: 5,
			want:   15,
		},
		{
			name:   "Zero Character Length",
			length: 0,
			want:   10,
		},
		{
			name:   "Negative Character Length",
			length: -1,
			want:   10,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			got := RandomEmail(tc.length)

			if len(got) != tc.want {
				t.Errorf("RandomEmail() = %v; want %v", len(got), tc.want)
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

