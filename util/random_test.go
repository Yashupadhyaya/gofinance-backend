package util

import (
	"fmt"
	"math"
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
		name           string
		input          int
		expectedLength int
		expectedSuffix string
	}{
		{
			name:           "Generate Email with 5-Character Local Part",
			input:          5,
			expectedLength: 5,
			expectedSuffix: "@email.com",
		},
		{
			name:           "Generate Email with 10-Character Local Part",
			input:          10,
			expectedLength: 10,
			expectedSuffix: "@email.com",
		},
		{
			name:           "Generate Email with 0-Character Local Part",
			input:          0,
			expectedLength: 0,
			expectedSuffix: "@email.com",
		},
		{
			name:           "Generate Email with Large Local Part",
			input:          100,
			expectedLength: 100,
			expectedSuffix: "@email.com",
		},
		{
			name:           "Generate Email with Negative Number",
			input:          -5,
			expectedLength: 0,
			expectedSuffix: "@email.com",
		},
		{
			name:           "Multiple Calls Consistency",
			input:          5,
			expectedLength: 5,
			expectedSuffix: "@email.com",
		},
		{
			name:           "Consistent Domain Part",
			input:          5,
			expectedLength: 5,
			expectedSuffix: "@email.com",
		},
		{
			name:           "Input as Boundary Value (1)",
			input:          1,
			expectedLength: 1,
			expectedSuffix: "@email.com",
		},
		{
			name:           "Input as Boundary Value (Max Int)",
			input:          math.MaxInt32,
			expectedLength: math.MaxInt32,
			expectedSuffix: "@email.com",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			email := RandomEmail(tt.input)
			localPart := strings.TrimSuffix(email, tt.expectedSuffix)

			if len(localPart) != tt.expectedLength {
				t.Errorf("expected local part length %d, got %d", tt.expectedLength, len(localPart))
			}

			if !strings.HasSuffix(email, tt.expectedSuffix) {
				t.Errorf("expected email suffix %s, got %s", tt.expectedSuffix, email)
			}

			if tt.name == "Multiple Calls Consistency" {
				email2 := RandomEmail(tt.input)
				if email == email2 {
					t.Errorf("expected different emails on multiple calls, got the same: %s", email)
				}
			}
		})
	}
}

