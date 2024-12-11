package util

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"strings"
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
		expectError bool
		validate    func(t *testing.T, result string)
	}{
		{
			name:  "Valid Email Generation with Standard Length",
			input: 10,
			validate: func(t *testing.T, result string) {
				localPart := strings.Split(result, "@")[0]
				assert.Equal(t, 10, len(localPart), "Expected local part to be 10 characters long")
				assert.True(t, strings.HasSuffix(result, "@email.com"), "Expected email to end with @email.com")
			},
		},
		{
			name:  "Minimum Length for Local Part",
			input: 1,
			validate: func(t *testing.T, result string) {
				localPart := strings.Split(result, "@")[0]
				assert.Equal(t, 1, len(localPart), "Expected local part to be 1 character long")
				assert.True(t, strings.HasSuffix(result, "@email.com"), "Expected email to end with @email.com")
			},
		},
		{
			name:  "Zero Length for Local Part",
			input: 0,
			validate: func(t *testing.T, result string) {
				localPart := strings.Split(result, "@")[0]
				assert.Equal(t, 0, len(localPart), "Expected local part to be 0 characters long")
				assert.True(t, strings.HasSuffix(result, "@email.com"), "Expected email to end with @email.com")
			},
		},
		{
			name:  "Large Length for Local Part",
			input: 1000,
			validate: func(t *testing.T, result string) {
				localPart := strings.Split(result, "@")[0]
				assert.Equal(t, 1000, len(localPart), "Expected local part to be 1000 characters long")
				assert.True(t, strings.HasSuffix(result, "@email.com"), "Expected email to end with @email.com")
			},
		},
		{
			name:  "Check for Unique Email Generation",
			input: 10,
			validate: func(t *testing.T, result string) {
				emailSet := make(map[string]struct{})
				for i := 0; i < 100; i++ {
					email := RandomEmail(10)
					if _, exists := emailSet[email]; exists {
						t.Fatalf("Duplicate email found: %s", email)
					}
					emailSet[email] = struct{}{}
				}
			},
		},
		{
			name:  "Valid Email Domain",
			input: 5,
			validate: func(t *testing.T, result string) {
				assert.True(t, strings.HasSuffix(result, "@email.com"), "Expected email to end with @email.com")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RandomEmail(tt.input)
			t.Logf("Generated email: %s", result)
			tt.validate(t, result)
		})
	}
}

/*
ROOST_METHOD_HASH=RandomString_d7e3599ac4
ROOST_METHOD_SIG_HASH=RandomString_c6fe4ad19a


 */
func TestRandomString(t *testing.T) {

	type testCase struct {
		name     string
		input    int
		expected int
		validate func(t *testing.T, result string)
	}

	testCases := []testCase{
		{
			name:     "Generate a Random String of Specified Length",
			input:    10,
			expected: 10,
			validate: func(t *testing.T, result string) {
				if len(result) != 10 {
					t.Errorf("Expected length 10, got %d", len(result))
				} else {
					t.Logf("Success: Generated string of length 10")
				}
			},
		},
		{
			name:     "Generate a Random String with Zero Length",
			input:    0,
			expected: 0,
			validate: func(t *testing.T, result string) {
				if len(result) != 0 {
					t.Errorf("Expected empty string, got length %d", len(result))
				} else {
					t.Logf("Success: Generated an empty string")
				}
			},
		},
		{
			name:     "Generate a Random String of Maximum Length",
			input:    10000,
			expected: 10000,
			validate: func(t *testing.T, result string) {
				if len(result) != 10000 {
					t.Errorf("Expected length 10000, got %d", len(result))
				} else {
					t.Logf("Success: Generated string of length 10000")
				}
			},
		},
		{
			name:  "Verify Randomness of Generated Strings",
			input: 10,
			validate: func(t *testing.T, result string) {
				result2 := RandomString(10)
				if result == result2 {
					t.Errorf("Expected different strings, got identical strings")
				} else {
					t.Logf("Success: Generated different strings")
				}
			},
		},
		{
			name:  "Consistent Output with Seeded Random Number Generator",
			input: 10,
			validate: func(t *testing.T, result string) {
				rand.Seed(42)
				result1 := RandomString(10)
				rand.Seed(42)
				result2 := RandomString(10)
				if result1 != result2 {
					t.Errorf("Expected identical strings, got different strings")
				} else {
					t.Logf("Success: Generated identical strings with seeded RNG")
				}
			},
		},
		{
			name:  "Generate a Random String with Non-Standard Alphabet",
			input: 10,
			validate: func(t *testing.T, result string) {
				originalAlphabet := alphabet
				defer func() { alphabet = originalAlphabet }()
				alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
				result := RandomString(10)
				for _, char := range result {
					if !strings.Contains(alphabet, string(char)) {
						t.Errorf("Generated string contains invalid character: %c", char)
						return
					}
				}
				t.Logf("Success: Generated string with non-standard alphabet")
			},
		},
		{
			name:  "Performance Test for Large Input Sizes",
			input: 1000000,
			validate: func(t *testing.T, result string) {
				start := time.Now()
				result := RandomString(1000000)
				duration := time.Since(start)
				t.Logf("Generated string of length 1000000 in %s", duration)
				if len(result) != 1000000 {
					t.Errorf("Expected length 1000000, got %d", len(result))
				}
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := RandomString(tc.input)
			tc.validate(t, result)
		})
	}
}

