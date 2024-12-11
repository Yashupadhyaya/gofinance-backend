package util

import (
	"fmt"
	"math"
	"math/rand"
	"regexp"
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
		expected string
		validate func(t *testing.T, email string)
	}{
		{
			name:  "Generate Email with 5-Character Local Part",
			input: 5,
			validate: func(t *testing.T, email string) {
				localPart := strings.Split(email, "@")[0]
				if len(localPart) != 5 {
					t.Errorf("Expected local part length of 5, got %d", len(localPart))
				}
			},
		},
		{
			name:  "Generate Email with 0-Character Local Part",
			input: 0,
			validate: func(t *testing.T, email string) {
				expected := "@email.com"
				if email != expected {
					t.Errorf("Expected email %s, got %s", expected, email)
				}
			},
		},
		{
			name:  "Generate Email with 100-Character Local Part",
			input: 100,
			validate: func(t *testing.T, email string) {
				localPart := strings.Split(email, "@")[0]
				if len(localPart) != 100 {
					t.Errorf("Expected local part length of 100, got %d", len(localPart))
				}
			},
		},
		{
			name:  "Generate Email with 1-Character Local Part",
			input: 1,
			validate: func(t *testing.T, email string) {
				localPart := strings.Split(email, "@")[0]
				if len(localPart) != 1 {
					t.Errorf("Expected local part length of 1, got %d", len(localPart))
				}
			},
		},
		{
			name:  "Generate Email with Negative Local Part Length",
			input: -5,
			validate: func(t *testing.T, email string) {
				expected := "@email.com"
				if email != expected {
					t.Errorf("Expected email %s for negative input, got %s", expected, email)
				}
			},
		},
		{
			name:  "Generate Email with Maximum Safe Integer Local Part Length",
			input: math.MaxInt32,
			validate: func(t *testing.T, email string) {

				t.Logf("Generated email with potentially very large local part: %s", email)
			},
		},
		{
			name:  "Generate Multiple Emails and Verify Uniqueness",
			input: 10,
			validate: func(t *testing.T, email string) {
				emails := make(map[string]bool)
				for i := 0; i < 100; i++ {
					email := RandomEmail(10)
					if emails[email] {
						t.Errorf("Duplicate email found: %s", email)
					}
					emails[email] = true
				}
			},
		},
		{
			name:  "Generate Email and Verify Format",
			input: 10,
			validate: func(t *testing.T, email string) {
				re := regexp.MustCompile(`^[a-zA-Z]{10}@email\.com$`)
				if !re.MatchString(email) {
					t.Errorf("Email format is incorrect: %s", email)
				}
			},
		},
		{
			name:  "Generate Email with Special Characters in Local Part",
			input: 10,
			validate: func(t *testing.T, email string) {
				localPart := strings.Split(email, "@")[0]
				re := regexp.MustCompile(`^[a-zA-Z]+$`)
				if !re.MatchString(localPart) {
					t.Errorf("Local part contains special characters: %s", localPart)
				}
			},
		},
		{
			name:  "Generate Email with Different Lengths and Verify Consistency",
			input: 3,
			validate: func(t *testing.T, email string) {
				lengths := []int{3, 7, 12}
				for _, length := range lengths {
					email := RandomEmail(length)
					localPart := strings.Split(email, "@")[0]
					if len(localPart) != length {
						t.Errorf("Expected local part length of %d, got %d", length, len(localPart))
					}
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("Running test: %s", tt.name)
			email := RandomEmail(tt.input)
			tt.validate(t, email)
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
		name   string
		number int
	}{
		{"Zero length", 0},
		{"Positive length", 10},
		{"All lowercase letters", 15},
		{"Consistency of random seed", 20},
		{"Handling large input size", 1000000},
		{"Same length multiple times", 8},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.name {
			case "Zero length":
				result := RandomString(tt.number)
				if len(result) != 0 {
					t.Errorf("expected empty string, got %v", result)
				} else {
					t.Log("Zero length: Passed")
				}

			case "Positive length":
				result := RandomString(tt.number)
				if len(result) != tt.number {
					t.Errorf("expected string length %d, got %d", tt.number, len(result))
				} else {
					t.Log("Positive length: Passed")
				}

			case "All lowercase letters":
				result := RandomString(tt.number)
				for _, c := range result {
					if c < 'a' || c > 'z' {
						t.Errorf("expected all lowercase letters, got %v", result)
						return
					}
				}
				t.Log("All lowercase letters: Passed")

			case "Consistency of random seed":
				result1 := RandomString(tt.number)
				result2 := RandomString(tt.number)
				if result1 == result2 {
					t.Errorf("expected different strings, got identical strings %v", result1)
				} else {
					t.Log("Consistency of random seed: Passed")
				}

			case "Handling large input size":
				result := RandomString(tt.number)
				if len(result) != tt.number {
					t.Errorf("expected string length %d, got %d", tt.number, len(result))
				}
				for _, c := range result {
					if c < 'a' || c > 'z' {
						t.Errorf("expected all lowercase letters, got %v", result)
						return
					}
				}
				t.Log("Handling large input size: Passed")

			case "Same length multiple times":
				result1 := RandomString(tt.number)
				result2 := RandomString(tt.number)
				if result1 == result2 {
					t.Errorf("expected different strings, got identical strings %v", result1)
				} else {
					t.Log("Same length multiple times: Passed")
				}
			}
		})
	}
}

