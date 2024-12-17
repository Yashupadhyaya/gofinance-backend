package util

import (
	"math/rand"
	"strings"
	"testing"
	"time"
)

/*
ROOST_METHOD_HASH=RandomEmail_1905439733
ROOST_METHOD_SIG_HASH=RandomEmail_7a04f189fd


 */
func TestRandomEmail(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	tests := []struct {
		name     string
		input    int
		validate func(t *testing.T, email string)
	}{
		{
			name:  "Generate Email with 5-Character Local Part",
			input: 5,
			validate: func(t *testing.T, email string) {
				localPart := strings.Split(email, "@")[0]
				if len(localPart) != 5 {
					t.Errorf("expected local part length 5, got %d", len(localPart))
				}
			},
		},
		{
			name:  "Generate Email with 0-Character Local Part",
			input: 0,
			validate: func(t *testing.T, email string) {
				localPart := strings.Split(email, "@")[0]
				if localPart != "" {
					t.Errorf("expected empty local part, got %s", localPart)
				}
			},
		},
		{
			name:  "Generate Email with Maximum Character Local Part",
			input: 1000,
			validate: func(t *testing.T, email string) {
				localPart := strings.Split(email, "@")[0]
				if len(localPart) != 1000 {
					t.Errorf("expected local part length 1000, got %d", len(localPart))
				}
			},
		},
		{
			name:  "Generate Email with Special Characters in Local Part",
			input: 10,
			validate: func(t *testing.T, email string) {
				localPart := strings.Split(email, "@")[0]
				for _, c := range localPart {
					if !strings.ContainsRune(alphabet, c) {
						t.Errorf("expected only alphabetic characters, got %c", c)
					}
				}
			},
		},
		{
			name:  "Generate Email with Non-Positive Integer",
			input: -1,
			validate: func(t *testing.T, email string) {
				localPart := strings.Split(email, "@")[0]
				if localPart != "" {
					t.Errorf("expected empty local part for non-positive input, got %s", localPart)
				}
			},
		},
		{
			name:  "Consistency of Email Domain",
			input: 5,
			validate: func(t *testing.T, email string) {
				domain := strings.Split(email, "@")[1]
				if domain != "email.com" {
					t.Errorf("expected domain email.com, got %s", domain)
				}
			},
		},
		{
			name:  "Randomness of Generated Local Part",
			input: 10,
			validate: func(t *testing.T, email string) {
				email2 := RandomEmail(10)
				if email == email2 {
					t.Errorf("expected different local parts, got identical emails %s and %s", email, email2)
				}
			},
		},
		{
			name:  "Performance with Large Input",
			input: 10000,
			validate: func(t *testing.T, email string) {
				start := time.Now()
				_ = RandomEmail(10000)
				duration := time.Since(start)
				if duration.Seconds() > 1 {
					t.Errorf("expected function to complete in less than 1 second, took %v", duration)
				}
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Logf("Running test: %s", tc.name)
			email := RandomEmail(tc.input)
			tc.validate(t, email)
		})
	}
}

/*
ROOST_METHOD_HASH=RandomString_d7e3599ac4
ROOST_METHOD_SIG_HASH=RandomString_c6fe4ad19a


 */
func TestRandomString(t *testing.T) {

	rand.Seed(time.Now().UnixNano())

	tests := []struct {
		name      string
		input     int
		expected  int
		assertion func(got string, expected int) bool
	}{
		{
			name:     "Generate a Random String of Specified Length",
			input:    10,
			expected: 10,
			assertion: func(got string, expected int) bool {
				return len(got) == expected
			},
		},
		{
			name:     "Generate a Random String with Zero Length",
			input:    0,
			expected: 0,
			assertion: func(got string, expected int) bool {
				return len(got) == expected
			},
		},
		{
			name:     "Generate a Random String with a Negative Length",
			input:    -5,
			expected: 0,
			assertion: func(got string, expected int) bool {
				return len(got) == expected
			},
		},
		{
			name:     "Verify Randomness of Generated Strings",
			input:    15,
			expected: 15,
			assertion: func(got string, expected int) bool {
				anotherString := RandomString(expected)
				return got != anotherString
			},
		},
		{
			name:     "Ensure Usage of Only Alphabet Characters",
			input:    20,
			expected: 20,
			assertion: func(got string, expected int) bool {
				for _, char := range got {
					if !strings.ContainsRune(alphabet, char) {
						return false
					}
				}
				return true
			},
		},
		{
			name:     "Performance Test for Large Input Value",
			input:    1000000,
			expected: 1000000,
			assertion: func(got string, expected int) bool {
				return len(got) == expected
			},
		},
		{
			name:     "Consistent Results with Same Seed",
			input:    12,
			expected: 12,
			assertion: func(got string, expected int) bool {
				rand.Seed(42)
				firstString := RandomString(expected)
				rand.Seed(42)
				secondString := RandomString(expected)
				return firstString == secondString
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RandomString(tt.input)
			if !tt.assertion(got, tt.expected) {
				t.Errorf("Test %s failed: got %v, expected %v", tt.name, got, tt.expected)
			} else {
				t.Logf("Test %s passed: got %v, expected %v", tt.name, got, tt.expected)
			}
		})
	}
}

