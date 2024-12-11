package util

import (
	"math/rand"
	"strings"
	"testing"
	"time"
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

func TestRandomEmail(t *testing.T) {
	type testCase struct {
		description string
		input       int
		assert      func(t *testing.T, result string)
	}

	testCases := []testCase{
		{
			description: "Generate Email with 5-Character Local Part",
			input:       5,
			assert: func(t *testing.T, result string) {
				localPart := strings.Split(result, "@")[0]
				if len(localPart) != 5 {
					t.Errorf("Expected local part length of 5, got %d", len(localPart))
				}
				if !strings.HasSuffix(result, "@email.com") {
					t.Errorf("Expected suffix '@email.com', got %s", result)
				}
			},
		},
		{
			description: "Generate Email with 0-Character Local Part",
			input:       0,
			assert: func(t *testing.T, result string) {
				if result != "@email.com" {
					t.Errorf("Expected '@email.com', got %s", result)
				}
			},
		},
		{
			description: "Generate Email with Maximum Reasonable Length",
			input:       1000,
			assert: func(t *testing.T, result string) {
				localPart := strings.Split(result, "@")[0]
				if len(localPart) != 1000 {
					t.Errorf("Expected local part length of 1000, got %d", len(localPart))
				}
				if !strings.HasSuffix(result, "@email.com") {
					t.Errorf("Expected suffix '@email.com', got %s", result)
				}
			},
		},
		{
			description: "Generate Email with Special Characters in Local Part",
			input:       10,
			assert: func(t *testing.T, result string) {
				localPart := strings.Split(result, "@")[0]
				for _, char := range localPart {
					if !strings.ContainsRune(alphabet, char) {
						t.Errorf("Expected only alphabetic characters, found '%c'", char)
					}
				}
			},
		},
		{
			description: "Consistency of Local Part Length",
			input:       8,
			assert: func(t *testing.T, result string) {
				for i := 0; i < 5; i++ {
					result := RandomEmail(8)
					localPart := strings.Split(result, "@")[0]
					if len(localPart) != 8 {
						t.Errorf("Expected local part length of 8, got %d", len(localPart))
					}
					if !strings.HasSuffix(result, "@email.com") {
						t.Errorf("Expected suffix '@email.com', got %s", result)
					}
				}
			},
		},
		{
			description: "Randomness of Local Part",
			input:       10,
			assert: func(t *testing.T, result string) {
				results := make(map[string]bool)
				for i := 0; i < 5; i++ {
					result := RandomEmail(10)
					if results[result] {
						t.Errorf("Duplicate email found: %s", result)
					}
					results[result] = true
				}
			},
		},
		{
			description: "Handling Negative Input",
			input:       -5,
			assert: func(t *testing.T, result string) {

				if result != "@email.com" {
					t.Errorf("Expected '@email.com', got %s", result)
				}
			},
		},
		{
			description: "Performance with Large Input",
			input:       10000,
			assert: func(t *testing.T, result string) {
				start := time.Now()
				result := RandomEmail(10000)
				elapsed := time.Since(start)
				localPart := strings.Split(result, "@")[0]
				if len(localPart) != 10000 {
					t.Errorf("Expected local part length of 10000, got %d", len(localPart))
				}
				if !strings.HasSuffix(result, "@email.com") {
					t.Errorf("Expected suffix '@email.com', got %s", result)
				}
				t.Logf("Time taken: %s", elapsed)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			t.Log(tc.description)
			result := RandomEmail(tc.input)
			tc.assert(t, result)
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
		assert func(t *testing.T, result string)
	}{
		{
			name:   "Generating Random String of Positive Length",
			number: 10,
			assert: func(t *testing.T, result string) {
				if len(result) != 10 {
					t.Errorf("Expected length 10, but got %d", len(result))
				}
			},
		},
		{
			name:   "Generating Random String of Zero Length",
			number: 0,
			assert: func(t *testing.T, result string) {
				if result != "" {
					t.Errorf("Expected empty string, but got %s", result)
				}
			},
		},
		{
			name:   "Generating Random String with Negative Length",
			number: -1,
			assert: func(t *testing.T, result string) {
				if result != "" {
					t.Errorf("Expected empty string for negative length, but got %s", result)
				}
			},
		},
		{
			name:   "Consistency of String Length for Multiple Calls",
			number: 15,
			assert: func(t *testing.T, result string) {
				for i := 0; i < 5; i++ {
					if len(result) != 15 {
						t.Errorf("Expected length 15, but got %d on iteration %d", len(result), i)
					}
				}
			},
		},
		{
			name:   "Randomness of Generated Strings",
			number: 20,
			assert: func(t *testing.T, result string) {
				results := make(map[string]bool)
				for i := 0; i < 10; i++ {
					newResult := RandomString(20)
					if results[newResult] {
						t.Errorf("Duplicate string found: %s", newResult)
					}
					results[newResult] = true
				}
			},
		},
		{
			name:   "Character Set Validation",
			number: 30,
			assert: func(t *testing.T, result string) {
				for _, char := range result {
					if !strings.ContainsRune(alphabet, char) {
						t.Errorf("Invalid character found: %c", char)
					}
				}
			},
		},
		{
			name:   "Performance with Large Length",
			number: 1000000,
			assert: func(t *testing.T, result string) {
				start := time.Now()
				if len(result) != 1000000 {
					t.Errorf("Expected length 1000000, but got %d", len(result))
				}
				duration := time.Since(start)
				t.Logf("Generated large string in %v", duration)
			},
		},
		{
			name:   "Seed Consistency Check",
			number: 25,
			assert: func(t *testing.T, result string) {
				rand.Seed(42)
				expected := RandomString(25)
				for i := 0; i < 5; i++ {
					rand.Seed(42)
					if result != RandomString(25) {
						t.Errorf("Expected consistent string, but got different result on iteration %d", i)
					}
				}
			},
		},
		{
			name:   "Boundary Test with Length One",
			number: 1,
			assert: func(t *testing.T, result string) {
				if len(result) != 1 {
					t.Errorf("Expected length 1, but got %d", len(result))
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RandomString(tt.number)
			tt.assert(t, result)
		})
	}
}

