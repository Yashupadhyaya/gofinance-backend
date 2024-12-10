package util

import (
	"fmt"
	"strings"
	"testing"
)

/*
ROOST_METHOD_HASH=RandomEmail_1905439733
ROOST_METHOD_SIG_HASH=RandomEmail_7a04f189fd


 */
func TestRandomEmail(t *testing.T) {
	tests := []struct {
		name       string
		input      int
		wantLength int
		wantSuffix string
	}{
		{
			name:       "Generate Email with Standard Length",
			input:      10,
			wantLength: 10,
			wantSuffix: "@email.com",
		},
		{
			name:       "Generate Email with Zero Length",
			input:      0,
			wantLength: 0,
			wantSuffix: "@email.com",
		},
		{
			name:       "Generate Email with Maximum Length",
			input:      1000,
			wantLength: 1000,
			wantSuffix: "@email.com",
		},
		{
			name:       "Generate Email with Negative Length",
			input:      -5,
			wantLength: 0,
			wantSuffix: "@email.com",
		},
		{
			name:       "Generate Email with Special Characters in Domain",
			input:      10,
			wantLength: 10,
			wantSuffix: "@email.com",
		},
		{
			name:       "Generate Email with Different Lengths Multiple Times",
			input:      5,
			wantLength: 5,
			wantSuffix: "@email.com",
		},
		{
			name:       "Generate Email with Same Length Multiple Times",
			input:      8,
			wantLength: 8,
			wantSuffix: "@email.com",
		},
		{
			name:       "Generate Email and Check Alphabet Usage",
			input:      12,
			wantLength: 12,
			wantSuffix: "@email.com",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RandomEmail(tt.input)

			username := strings.TrimSuffix(got, tt.wantSuffix)
			if len(username) != tt.wantLength {
				t.Errorf("Expected username length: %d, but got: %d", tt.wantLength, len(username))
			}

			if !strings.HasSuffix(got, tt.wantSuffix) {
				t.Errorf("Email should end with %s, but got: %s", tt.wantSuffix, got)
			}

			if tt.name == "Generate Email and Check Alphabet Usage" {
				for _, char := range username {
					if !strings.ContainsRune(alphabet, char) {
						t.Errorf("Username contains invalid character: %c", char)
					}
				}
			}

			t.Logf("Generated email: %s", got)
		})
	}
}

