package util

import (
	"math"
	"math/rand"
	"strings"
	"testing"
	"time"
	"fmt"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"
const alphabet = "abcdefghijklmnopqrstuvwxyz"
/*
ROOST_METHOD_HASH=RandomEmail_1905439733
ROOST_METHOD_SIG_HASH=RandomEmail_7a04f189fd


 */
func RandomEmail(number int) string {
	return fmt.Sprintf("%s@email.com", RandomString(number))
}

func RandomString(number int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < number; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func TestRandomEmail(t *testing.T) {

	tests := []struct {
		description string
		input       int
		expected    func(string) bool
	}{
		{
			description: "Generate Email with 5-Character Local Part",
			input:       5,
			expected: func(email string) bool {
				parts := strings.Split(email, "@")
				return len(parts[0]) == 5 && parts[1] == "email.com"
			},
		},
		{
			description: "Generate Email with 0-Character Local Part",
			input:       0,
			expected: func(email string) bool {
				return email == "@email.com"
			},
		},
		{
			description: "Generate Email with Negative Number",
			input:       -5,
			expected: func(email string) bool {
				return email == "@email.com"
			},
		},
		{
			description: "Generate Email with Large Number",
			input:       1000,
			expected: func(email string) bool {
				parts := strings.Split(email, "@")
				return len(parts[0]) == 1000 && parts[1] == "email.com"
			},
		},
		{
			description: "Generate Email with Special Characters in Local Part",
			input:       10,
			expected: func(email string) bool {
				parts := strings.Split(email, "@")
				for _, char := range parts[0] {
					if !strings.ContainsRune(alphabet, char) {
						return false
					}
				}
				return parts[1] == "email.com"
			},
		},
		{
			description: "Consistent Domain Part",
			input:       10,
			expected: func(email string) bool {
				return strings.HasSuffix(email, "@email.com")
			},
		},
		{
			description: "Randomness of Local Part",
			input:       10,
			expected: func(email string) bool {
				email1 := RandomEmail(10)
				email2 := RandomEmail(10)
				return email1 != email2
			},
		},
		{
			description: "Performance with Large Inputs",
			input:       10000,
			expected: func(email string) bool {
				start := time.Now()
				_ = RandomEmail(10000)
				duration := time.Since(start)
				return duration < time.Second
			},
		},
		{
			description: "Edge Case with Maximum Integer Input",
			input:       math.MaxInt32,
			expected: func(email string) bool {
				defer func() {
					if r := recover(); r != nil {
						t.Errorf("Function panicked with input %d", math.MaxInt32)
					}
				}()
				email := RandomEmail(math.MaxInt32)
				parts := strings.Split(email, "@")
				return len(parts[0]) == math.MaxInt32 && parts[1] == "email.com"
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			email := RandomEmail(tc.input)
			if !tc.expected(email) {
				t.Errorf("Test failed for %s. Got: %s", tc.description, email)
			} else {
				t.Logf("Test passed for %s", tc.description)
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
		name   string
		number int
		verify func(t *testing.T, result string)
	}{
		{
			name:   "Generating Random String of Zero Length",
			number: 0,
			verify: func(t *testing.T, result string) {
				if result != "" {
					t.Errorf("Expected empty string, got %s", result)
				} else {
					t.Log("Zero length test passed")
				}
			},
		},
		{
			name:   "Generating Random String of Length One",
			number: 1,
			verify: func(t *testing.T, result string) {
				if len(result) != 1 {
					t.Errorf("Expected string length 1, got %d", len(result))
				} else if !strings.Contains(alphabet, result) {
					t.Errorf("Expected character from alphabet, got %s", result)
				} else {
					t.Log("Length one test passed")
				}
			},
		},
		{
			name:   "Generating Random String of Typical Length",
			number: 10,
			verify: func(t *testing.T, result string) {
				if len(result) != 10 {
					t.Errorf("Expected string length 10, got %d", len(result))
				} else {
					for _, c := range result {
						if !strings.Contains(alphabet, string(c)) {
							t.Errorf("Unexpected character %c in result", c)
						}
					}
					t.Log("Typical length test passed")
				}
			},
		},
		{
			name:   "Consistency of Random String Length",
			number: 15,
			verify: func(t *testing.T, result string) {
				if len(result) != 15 {
					t.Errorf("Expected string length 15, got %d", len(result))
				} else {
					t.Log("Consistency of length test passed")
				}
			},
		},
		{
			name:   "Valid Characters in Generated String",
			number: 20,
			verify: func(t *testing.T, result string) {
				for _, c := range result {
					if !strings.Contains(alphabet, string(c)) {
						t.Errorf("Unexpected character %c in result", c)
					}
				}
				t.Log("Valid characters test passed")
			},
		},
		{
			name:   "Large Length Input",
			number: 1000,
			verify: func(t *testing.T, result string) {
				if len(result) != 1000 {
					t.Errorf("Expected string length 1000, got %d", len(result))
				} else {
					for _, c := range result {
						if !strings.Contains(alphabet, string(c)) {
							t.Errorf("Unexpected character %c in result", c)
						}
					}
					t.Log("Large length test passed")
				}
			},
		},
		{
			name:   "Consistency with Seed",
			number: 10,
			verify: func(t *testing.T, result string) {
				seed := int64(12345)
				rand.Seed(seed)
				expected := RandomString(10)
				rand.Seed(seed)
				actual := RandomString(10)
				if expected != actual {
					t.Errorf("Expected %s, got %s", expected, actual)
				} else {
					t.Log("Consistency with seed test passed")
				}
			},
		},
		{
			name:   "Handling Negative Length Input",
			number: -5,
			verify: func(t *testing.T, result string) {
				if result != "" {
					t.Errorf("Expected empty string for negative input, got %s", result)
				} else {
					t.Log("Negative length input test passed")
				}
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := RandomString(test.number)
			test.verify(t, result)
		})
	}
}

