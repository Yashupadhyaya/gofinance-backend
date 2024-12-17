package util

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"
	"unicode"
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
					t.Errorf("expected local part length of 5, got %d", len(localPart))
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
					t.Errorf("expected local part length of 1000, got %d", len(localPart))
				}
			},
		},
		{
			name:  "Validate Email Format",
			input: 10,
			validate: func(t *testing.T, email string) {
				if !strings.HasSuffix(email, "@email.com") {
					t.Errorf("expected email to end with '@email.com', got %s", email)
				}
				if strings.Count(email, "@") != 1 {
					t.Errorf("expected exactly one '@' character, got %s", email)
				}
			},
		},
		{
			name:  "Randomness of Generated Email",
			input: 10,
			validate: func(t *testing.T, email string) {
				email1 := RandomEmail(10)
				email2 := RandomEmail(10)
				if email1 == email2 {
					t.Errorf("expected different emails, got %s and %s", email1, email2)
				}
			},
		},
		{
			name:  "Consistency with Seeded Randomness",
			input: 10,
			validate: func(t *testing.T, email string) {
				rand.Seed(42)
				email1 := RandomEmail(10)
				rand.Seed(42)
				email2 := RandomEmail(10)
				if email1 != email2 {
					t.Errorf("expected same emails, got %s and %s", email1, email2)
				}
			},
		},
		{
			name:  "Check Email Domain Consistency",
			input: 15,
			validate: func(t *testing.T, email string) {
				if !strings.HasSuffix(email, "@email.com") {
					t.Errorf("expected email to end with '@email.com', got %s", email)
				}
			},
		},
		{
			name:  "Handling Negative Input",
			input: -5,
			validate: func(t *testing.T, email string) {
				if email != "@email.com" {
					t.Errorf("expected '@email.com', got %s", email)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
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
		name     string
		input    int
		expected int
	}{
		{
			name:     "Generate a Random String of Specified Length",
			input:    10,
			expected: 10,
		},
		{
			name:     "Generate a Random String with Zero Length",
			input:    0,
			expected: 0,
		},
		{
			name:     "Generate a Random String with Negative Length",
			input:    -1,
			expected: 0,
		},
		{
			name:     "Generate a Random String of Large Length",
			input:    10000,
			expected: 10000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RandomString(tt.input)
			if len(got) != tt.expected {
				t.Errorf("RandomString(%d) = %d; want %d", tt.input, len(got), tt.expected)
			} else {
				t.Logf("RandomString(%d) returned string of length %d", tt.input, len(got))
			}
		})
	}

	t.Run("Generate a Random String Multiple Times and Verify Uniqueness", func(t *testing.T) {
		input := 10
		str1 := RandomString(input)
		str2 := RandomString(input)
		if str1 == str2 {
			t.Error("RandomString generated the same string on multiple calls with the same length")
		} else {
			t.Log("RandomString generated unique strings on multiple calls")
		}
	})

	t.Run("Generate a Random String and Verify Content", func(t *testing.T) {
		input := 10
		got := RandomString(input)
		for _, c := range got {
			if !unicode.IsLetter(c) {
				t.Errorf("RandomString(%d) generated string with invalid character %c", input, c)
			}
		}
		t.Log("RandomString generated string with valid characters")
	})
}

