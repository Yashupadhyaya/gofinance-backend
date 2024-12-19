package util

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
)

/*
ROOST_METHOD_HASH=RandomEmail_1905439733
ROOST_METHOD_SIG_HASH=RandomEmail_7a04f189fd


 */
func TestRandomEmail(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	tests := []struct {
		name     string
		number   int
		expected string
	}{
		{
			name:     "Generate Email with 5-Character Local Part",
			number:   5,
			expected: "@email.com",
		},
		{
			name:     "Generate Email with 0-Character Local Part",
			number:   0,
			expected: "@email.com",
		},
		{
			name:     "Generate Email with Maximum Length Local Part",
			number:   100,
			expected: "@email.com",
		},
		{
			name:     "Generate Email with Single Character Local Part",
			number:   1,
			expected: "@email.com",
		},
		{
			name:     "Invalid Input Handling (Negative Number)",
			number:   -5,
			expected: "@email.com",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			email := RandomEmail(tt.number)
			t.Logf("Generated Email: %s", email)

			if !strings.HasSuffix(email, tt.expected) {
				t.Errorf("expected suffix %s, got %s", tt.expected, email)
			}

			localPart := strings.TrimSuffix(email, "@email.com")
			if len(localPart) != tt.number {
				t.Errorf("expected local part length %d, got %d", tt.number, len(localPart))
			}
		})
	}

	t.Run("Consistency of Email Format", func(t *testing.T) {
		for i := 0; i < 10; i++ {
			email := RandomEmail(10)
			t.Logf("Generated Email: %s", email)

			if !strings.HasSuffix(email, "@email.com") {
				t.Errorf("expected suffix @email.com, got %s", email)
			}

			localPart := strings.TrimSuffix(email, "@email.com")
			if len(localPart) != 10 {
				t.Errorf("expected local part length 10, got %d", len(localPart))
			}
		}
	})

	t.Run("Validate Characters in Local Part", func(t *testing.T) {
		email := RandomEmail(15)
		t.Logf("Generated Email: %s", email)

		localPart := strings.TrimSuffix(email, "@email.com")
		for _, char := range localPart {
			if !strings.ContainsRune(alphabet, char) {
				t.Errorf("unexpected character %c in local part", char)
			}
		}
	})

	t.Run("Generate Email with Non-ASCII Characters in Local Part", func(t *testing.T) {
		for i := 0; i < 10; i++ {
			email := RandomEmail(20)
			t.Logf("Generated Email: %s", email)

			localPart := strings.TrimSuffix(email, "@email.com")
			for _, char := range localPart {
				if char > 127 {
					t.Errorf("unexpected non-ASCII character %c in local part", char)
				}
			}
		}
	})

	t.Run("Performance Under High Volume", func(t *testing.T) {
		start := time.Now()
		for i := 0; i < 10000; i++ {
			RandomEmail(10)
		}
		duration := time.Since(start)
		t.Logf("Duration for 10000 invocations: %v", duration)

		if duration > time.Second {
			t.Errorf("performance issue: took %v", duration)
		}
	})

}

/*
ROOST_METHOD_HASH=RandomString_d7e3599ac4
ROOST_METHOD_SIG_HASH=RandomString_c6fe4ad19a


 */
func TestRandomString(t *testing.T) {

	tests := []struct {
		name      string
		input     int
		expected  string
		expectErr bool
	}{
		{
			name:      "Generate a Random String of Specified Length",
			input:     10,
			expected:  "",
			expectErr: false,
		},
		{
			name:      "Generate a Random String with Zero Length",
			input:     0,
			expected:  "",
			expectErr: false,
		},
		{
			name:      "Generate a Random String with a Negative Length",
			input:     -1,
			expected:  "",
			expectErr: false,
		},
		{
			name:      "Performance Testing for Large String Generation",
			input:     1000000,
			expected:  "",
			expectErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			rand.Seed(time.Now().UnixNano())

			result := RandomString(tc.input)

			if tc.expectErr {

			} else {
				if tc.input >= 0 {
					assert.Equal(t, len(result), tc.input, "The length of the generated string should be equal to the input number.")
				} else {
					assert.Equal(t, result, "", "The function should return an empty string for negative length.")
				}
			}

			t.Logf("Test '%s' passed with input %d and generated string: %s", tc.name, tc.input, result)
		})
	}

	t.Run("Verify Randomness of Generated Strings", func(t *testing.T) {

		input := 10
		rand.Seed(time.Now().UnixNano())

		result1 := RandomString(input)
		result2 := RandomString(input)

		assert.NotEqual(t, result1, result2, "The generated strings should be different for the same input length.")
		t.Logf("Generated strings: '%s' and '%s'", result1, result2)
	})

	t.Run("Consistency Across Multiple Runs with Same Seed", func(t *testing.T) {

		input := 10
		seed := int64(42)
		rand.Seed(seed)

		result1 := RandomString(input)
		rand.Seed(seed)
		result2 := RandomString(input)

		assert.Equal(t, result1, result2, "The generated strings should be identical for the same seed and input length.")
		t.Logf("Generated strings with fixed seed: '%s' and '%s'", result1, result2)
	})
}

