package util

import (
	"fmt"
	"strings"
	"testing"
	"time"
	"math/rand"
)

/*
ROOST_METHOD_HASH=RandomEmail_1905439733
ROOST_METHOD_SIG_HASH=RandomEmail_7a04f189fd


 */
func RandomEmail(number int) string {
	return fmt.Sprintf("%s@email.com", RandomString(number))
}

func TestRandomEmail(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	tests := []struct {
		name		string
		input		int
		expectedPrefix	string
		expectedSuffix	string
		expectedLength	int
		expectError	bool
	}{{name: "Generate Email with a Standard Length", input: 10, expectedPrefix: "xxxxxxxxxx", expectedSuffix: "@email.com", expectedLength: 10 + len("@email.com"), expectError: false}, {name: "Generate Email with Zero Length", input: 0, expectedPrefix: "", expectedSuffix: "@email.com", expectedLength: len("@email.com"), expectError: false}, {name: "Generate Email with Negative Length", input: -5, expectedPrefix: "", expectedSuffix: "@email.com", expectedLength: len("@email.com"), expectError: false}, {name: "Generate Email with Maximum Length", input: 1000, expectedPrefix: strings.Repeat("x", 1000), expectedSuffix: "@email.com", expectedLength: 1000 + len("@email.com"), expectError: false}, {name: "Validate Email Format", input: 5, expectedPrefix: "xxxxx", expectedSuffix: "@email.com", expectedLength: 5 + len("@email.com"), expectError: false}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RandomEmail(tt.input)
			t.Logf("Generated Email: %s", result)
			if len(result) != tt.expectedLength {
				t.Errorf("expected length %d, got %d", tt.expectedLength, len(result))
			}
			if !strings.HasSuffix(result, tt.expectedSuffix) {
				t.Errorf("expected suffix %s, got %s", tt.expectedSuffix, result)
			}
			if tt.input > 0 {
				firstResult := RandomEmail(tt.input)
				secondResult := RandomEmail(tt.input)
				if firstResult == secondResult {
					t.Error("expected different results for the same input due to randomness, got identical results")
				}
			}
		})
	}
	t.Run("Consistent Randomness for Same Input", func(t *testing.T) {
		input := 10
		firstResult := RandomEmail(input)
		secondResult := RandomEmail(input)
		if firstResult == secondResult {
			t.Error("expected different results for the same input due to randomness, got identical results")
		}
	})
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
		name	string
		length	int
		check	func(string) bool
		message	string
	}{{name: "Generate a Random String of Specified Length", length: 10, check: func(s string) bool {
		return len(s) == 10
	}, message: "Expected string of length 10"}, {name: "Generate an Empty String When Length is Zero", length: 0, check: func(s string) bool {
		return len(s) == 0
	}, message: "Expected empty string"}, {name: "Generate a String with Upper Bound Length", length: 10000, check: func(s string) bool {
		return len(s) == 10000
	}, message: "Expected string of length 10000"}, {name: "Ensure Randomness of Generated String", length: 100, check: func(s string) bool {
		otherString := RandomString(100)
		return s != otherString
	}, message: "Expected different strings on multiple runs"}, {name: "Handle Negative Length Gracefully", length: -5, check: func(s string) bool {
		return len(s) == 0
	}, message: "Expected empty string for negative length"}, {name: "Consistent Character Set Usage", length: 50, check: func(s string) bool {
		for _, c := range s {
			if !strings.ContainsRune(alphabet, c) {
				return false
			}
		}
		return true
	}, message: "Expected all characters to be part of the known alphabet"}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RandomString(tt.length)
			if !tt.check(result) {
				t.Errorf("%s: got %v", tt.message, result)
			} else {
				t.Logf("Success: %s", tt.message)
			}
		})
	}
	t.Run("Performance Under Repeated Calls", func(t *testing.T) {
		start := time.Now()
		for i := 0; i < 1000; i++ {
			_ = RandomString(50)
		}
		duration := time.Since(start)
		if duration.Seconds() > 1 {
			t.Errorf("Performance test exceeded acceptable duration: %v", duration)
		} else {
			t.Logf("Performance test completed in acceptable duration: %v", duration)
		}
	})
}

