// ********RoostGPT********
/*
Test generated by RoostGPT for test go-single-testfile using AI Type Claude AI and AI Model claude-3-5-sonnet-20240620

ROOST_METHOD_HASH=RandomEmail_1905439733
ROOST_METHOD_SIG_HASH=RandomEmail_7a04f189fd

Based on the provided function and context, here are several test scenarios for the RandomEmail function:

Scenario 1: Generate a Random Email with Positive Integer Input

Details:
  Description: This test checks if the RandomEmail function generates a valid email address when given a positive integer input.
Execution:
  Arrange: Set up a test case with a positive integer input (e.g., 8).
  Act: Call RandomEmail(8) and store the result.
  Assert: Check if the returned string is a valid email address format and has the correct length.
Validation:
  Use string manipulation and regular expressions to verify the email format (local-part@email.com) and ensure the local part length matches the input. This test is crucial to verify the basic functionality of the email generation.

Scenario 2: Generate a Random Email with Zero Input

Details:
  Description: This test verifies the behavior of RandomEmail when given a zero as input.
Execution:
  Arrange: Prepare a test case with input 0.
  Act: Call RandomEmail(0) and capture the output.
  Assert: Verify that the function returns a valid email address with an empty local part.
Validation:
  Check if the result is exactly "@email.com". This test is important to ensure the function handles edge cases gracefully.

Scenario 3: Generate a Random Email with Large Integer Input

Details:
  Description: This test examines how RandomEmail handles a large integer input.
Execution:
  Arrange: Set up a test with a large integer (e.g., 1000).
  Act: Execute RandomEmail(1000) and store the result.
  Assert: Confirm that the function returns a valid email address with the correct local part length.
Validation:
  Verify the email format and ensure the local part length is exactly 1000 characters. This test is vital to check the function's performance and behavior with large inputs.

Scenario 4: Generate Multiple Random Emails and Check Uniqueness

Details:
  Description: This test generates multiple random emails and checks for uniqueness among them.
Execution:
  Arrange: Prepare to generate multiple emails (e.g., 100) with a fixed length (e.g., 10).
  Act: Call RandomEmail(10) 100 times and store the results in a slice.
  Assert: Verify that all generated emails are unique.
Validation:
  Use a map to check for duplicates. This test is important to ensure the randomness and uniqueness of generated emails, which is crucial for many applications.

Scenario 5: Generate a Random Email with Negative Integer Input

Details:
  Description: This test checks the behavior of RandomEmail when given a negative integer input.
Execution:
  Arrange: Set up a test case with a negative integer input (e.g., -5).
  Act: Call RandomEmail(-5) and capture the output.
  Assert: Verify the function's behavior with negative input (e.g., returns an error or a default email).
Validation:
  Check if the function handles negative inputs gracefully, either by returning an error or using a sensible default. This test is crucial for robust error handling.

Scenario 6: Verify Consistency of Random Email Generation

Details:
  Description: This test checks if RandomEmail generates different emails on subsequent calls with the same input.
Execution:
  Arrange: Prepare to call RandomEmail multiple times with the same input.
  Act: Call RandomEmail(10) twice in succession and store the results.
  Assert: Verify that the two generated emails are different.
Validation:
  Compare the two results to ensure they're not identical. This test is important to confirm the randomness of the email generation process.

These scenarios cover various aspects of the RandomEmail function, including normal operation, edge cases, and potential error conditions. They help ensure the function works correctly across different inputs and maintains the expected randomness and format of the generated email addresses.
*/

// ********RoostGPT********
package util

import (
	"regexp"
	"strings"
	"testing"
)

func TestRandomEmail(t *testing.T) {
	tests := []struct {
		name      string
		input     int
		wantLen   int
		wantRegex string
	}{
		{
			name:      "Positive Integer Input",
			input:     8,
			wantLen:   8,
			wantRegex: `^[a-zA-Z0-9]{8}@email\.com$`,
		},
		{
			name:      "Zero Input",
			input:     0,
			wantLen:   0,
			wantRegex: `^@email\.com$`,
		},
		{
			name:      "Large Integer Input",
			input:     1000,
			wantLen:   1000,
			wantRegex: `^[a-zA-Z0-9]{1000}@email\.com$`,
		},
		{
			name:      "Negative Integer Input",
			input:     -5,
			wantLen:   0,
			wantRegex: `^@email\.com$`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RandomEmail(tt.input)
			
			// Check email format
			if !regexp.MustCompile(tt.wantRegex).MatchString(got) {
				t.Errorf("RandomEmail(%d) = %v, want match for %v", tt.input, got, tt.wantRegex)
			}

			// Check local part length
			localPart := strings.Split(got, "@")[0]
			if len(localPart) != tt.wantLen {
				t.Errorf("RandomEmail(%d) local part length = %d, want %d", tt.input, len(localPart), tt.wantLen)
			}
		})
	}
}

func TestRandomEmailUniqueness(t *testing.T) {
	const numEmails = 100
	const emailLength = 10
	emails := make(map[string]bool)

	for i := 0; i < numEmails; i++ {
		email := RandomEmail(emailLength)
		if emails[email] {
			t.Errorf("Duplicate email generated: %s", email)
		}
		emails[email] = true
	}

	if len(emails) != numEmails {
		t.Errorf("Expected %d unique emails, got %d", numEmails, len(emails))
	}
}

func TestRandomEmailConsistency(t *testing.T) {
	const emailLength = 10
	email1 := RandomEmail(emailLength)
	email2 := RandomEmail(emailLength)

	if email1 == email2 {
		t.Errorf("Two consecutive calls to RandomEmail(%d) returned the same email: %s", emailLength, email1)
	}
}
