package util_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"your_project/util"
)

/*
ROOST_METHOD_HASH=RandomString_d7e3599ac4
ROOST_METHOD_SIG_HASH=RandomString_c6fe4ad19a


 */
func TestRandomString(t *testing.T) {
	var tests = []struct {
		name		string
		input		int
		expectedError	error
	}{{name: "Test Scenario 1: Testing RandomString with a positive number", input: 5, expectedError: nil}, {name: "Test Scenario 2: Testing RandomString with zero", input: 0, expectedError: nil}, {name: "Test Scenario 3: Testing RandomString with a negative number", input: -5, expectedError: nil}}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := util.RandomString(test.input)
			if test.expectedError != nil {
				assert.Error(t, test.expectedError, "An error was expected")
			} else {
				if test.input > 0 {
					assert.Len(t, output, test.input, "The length of the output should be equal to the input")
				} else {
					assert.Empty(t, output, "The output should be an empty string when input is zero or less")
				}
			}
		})
	}
}

func (a *Assertions) Empty(object interface{}, msgAndArgs ...interface{}) bool {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	return Empty(a.t, object, msgAndArgs...)
}

func (a *Assertions) Error(err error, msgAndArgs ...interface{}) bool {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	return Error(a.t, err, msgAndArgs...)
}

func (a *Assertions) Len(object interface{}, length int, msgAndArgs ...interface{}) bool {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	return Len(a.t, object, length, msgAndArgs...)
}

