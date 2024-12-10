package util

import (
	"fmt"
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
	tests := []struct {
		name           string
		input          int
		expectedLength int
		expectedOutput string
		expectError    bool
	}{
		{
			name:           "Generate Email with 5-Character Local Part",
			input:          5,
			expectedLength: 5,
			expectedOutput: "",
			expectError:    false,
		},
		{
			name:           "Generate Email with 0-Character Local Part",
			input:          0,
			expectedLength: 0,
			expectedOutput: "@email.com",
			expectError:    false,
		},
		{
			name:           "Generate Email with Maximum Integer Value",
			input:          int(^uint(0) >> 1),
			expectedLength: int(^uint(0) >> 1),
			expectedOutput: "",
			expectError:    false,
		},
		{
			name:           "Generate Email with Negative Number",
			input:          -1,
			expectedLength: 0,
			expectedOutput: "@email.com",
			expectError:    true,
		},
		{
			name:           "Generate Email with 1-Character Local Part",
			input:          1,
			expectedLength: 1,
			expectedOutput: "",
			expectError:    false,
		},
		{
			name:           "Generate Email with 10-Character Local Part",
			input:          10,
			expectedLength: 10,
			expectedOutput: "",
			expectError:    false,
		},
		{
			name:           "Generate Email with Consistent Length",
			input:          8,
			expectedLength: 8,
			expectedOutput: "",
			expectError:    false,
		},
		{
			name:           "Generate Email with Varying Lengths",
			input:          3,
			expectedLength: 3,
			expectedOutput: "",
			expectError:    false,
		},
		{
			name:           "Generate Email with Varying Lengths",
			input:          7,
			expectedLength: 7,
			expectedOutput: "",
			expectError:    false,
		},
		{
			name:           "Generate Email with Varying Lengths",
			input:          12,
			expectedLength: 12,
			expectedOutput: "",
			expectError:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := RandomEmail(tt.input)

			if tt.expectError {
				if output != tt.expectedOutput {
					t.Errorf("Expected error output %v, but got %v", tt.expectedOutput, output)
				}
			} else {
				localPart := strings.Split(output, "@")[0]
				if len(localPart) != tt.expectedLength {
					t.Errorf("Expected local part length %d, but got %d", tt.expectedLength, len(localPart))
				}

				if tt.input == 10 {
					for _, char := range localPart {
						if !strings.Contains(alphabet, string(char)) {
							t.Errorf("Expected local part to contain only alphabet characters, but got %v", localPart)
							break
						}
					}
				}

				if tt.input == int(^uint(0)>>1) {
					if len(output) == 0 {
						t.Errorf("Expected a non-empty email address for maximum integer value input")
					}
				}
			}
		})
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

