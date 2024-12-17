package util

import (
	"math/rand"
	"strings"
	"testing"
	"time"
	"fmt"
)

/*
ROOST_METHOD_HASH=RandomEmail_1905439733
ROOST_METHOD_SIG_HASH=RandomEmail_7a04f189fd


 */
func TestRandomEmail(t *testing.T) {

	testCases := []struct {
		name             string
		input            int
		expectedLocalLen int
		expectedEmail    string
		expectError      bool
	}{
		{
			name:             "Generate Email with 5-Character Local Part",
			input:            5,
			expectedLocalLen: 5,
			expectedEmail:    "@email.com",
			expectError:      false,
		},
		{
			name:             "Generate Email with 0-Character Local Part",
			input:            0,
			expectedLocalLen: 0,
			expectedEmail:    "@email.com",
			expectError:      false,
		},
		{
			name:             "Generate Email with Maximum Reasonable Length",
			input:            100,
			expectedLocalLen: 100,
			expectedEmail:    "@email.com",
			expectError:      false,
		},
		{
			name:             "Generate Email with Negative Number",
			input:            -5,
			expectedLocalLen: 0,
			expectedEmail:    "@email.com",
			expectError:      false,
		},
		{
			name:             "Generate Email with Very Large Number",
			input:            1000,
			expectedLocalLen: 1000,
			expectedEmail:    "@email.com",
			expectError:      false,
		},
	}

	rand.Seed(time.Now().UnixNano())

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := RandomEmail(tc.input)
			localPart := strings.Split(result, "@")[0]

			if len(localPart) != tc.expectedLocalLen {
				t.Errorf("Test %s failed: expected local part length %d, got %d", tc.name, tc.expectedLocalLen, len(localPart))
			}

			if !strings.HasSuffix(result, tc.expectedEmail) {
				t.Errorf("Test %s failed: expected email to end with %s, got %s", tc.name, tc.expectedEmail, result)
			}

			t.Logf("Test %s passed: generated email %s", tc.name, result)
		})
	}

	t.Run("Consistency of Generated Email Format", func(t *testing.T) {
		email := RandomEmail(15)
		if !strings.HasSuffix(email, "@email.com") {
			t.Errorf("Test failed: expected email to end with @email.com, got %s", email)
		}
		t.Logf("Test passed: generated email %s", email)
	})

	t.Run("Verify Randomness of Generated Emails", func(t *testing.T) {
		email1 := RandomEmail(10)
		email2 := RandomEmail(10)
		if email1 == email2 {
			t.Errorf("Test failed: expected different emails, got %s and %s", email1, email2)
		}
		t.Logf("Test passed: generated emails %s and %s", email1, email2)
	})
}

/*
ROOST_METHOD_HASH=RandomString_d7e3599ac4
ROOST_METHOD_SIG_HASH=RandomString_c6fe4ad19a


 */
func TestRandomString(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	tests := []struct {
		name       string
		input      int
		assertFunc func(t *testing.T, result string)
	}{
		{
			name:  "Generate a Random String of Given Length",
			input: 10,
			assertFunc: func(t *testing.T, result string) {
				if len(result) != 10 {
					t.Errorf("Expected length 10, got %d", len(result))
				} else {
					t.Logf("Success: Generated string of length %d", len(result))
				}
			},
		},
		{
			name:  "Generate an Empty String When Length is Zero",
			input: 0,
			assertFunc: func(t *testing.T, result string) {
				if result != "" {
					t.Errorf("Expected empty string, got '%s'", result)
				} else {
					t.Logf("Success: Generated an empty string")
				}
			},
		},
		{
			name:  "Generate a Random String of Maximum Length",
			input: 10000,
			assertFunc: func(t *testing.T, result string) {
				if len(result) != 10000 {
					t.Errorf("Expected length 10000, got %d", len(result))
				} else {
					t.Logf("Success: Generated string of length %d", len(result))
				}
			},
		},
		{
			name:  "Ensure Randomness of Generated Strings",
			input: 10,
			assertFunc: func(t *testing.T, result string) {
				result2 := RandomString(10)
				if result == result2 {
					t.Errorf("Expected different strings, got '%s' and '%s'", result, result2)
				} else {
					t.Logf("Success: Generated different strings '%s' and '%s'", result, result2)
				}
			},
		},
		{
			name:  "Validate Characters in Generated String",
			input: 10,
			assertFunc: func(t *testing.T, result string) {
				for _, ch := range result {
					if !strings.ContainsRune(alphabet, ch) {
						t.Errorf("Unexpected character '%c' in generated string", ch)
					}
				}
				t.Logf("Success: All characters in the generated string are valid")
			},
		},
		{
			name:  "Performance Test for Large Input",
			input: 1000000,
			assertFunc: func(t *testing.T, result string) {
				start := time.Now()
				RandomString(1000000)
				duration := time.Since(start)
				if duration.Seconds() > 2 {
					t.Errorf("Function took too long: %v seconds", duration.Seconds())
				} else {
					t.Logf("Success: Generated string of length 1000000 in %v seconds", duration.Seconds())
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RandomString(tt.input)
			tt.assertFunc(t, result)
		})
	}
}

