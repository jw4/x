package contract_test

import (
	"testing"

	"github.com/jw4/x/contract"
)

func TestAssertNotNil(t *testing.T) {
	testCases := []struct {
		given       []interface{}
		expectPanic bool
	}{
		{},
		{given: []interface{}{nil}, expectPanic: true},
		{given: []interface{}{1, 3, ""}, expectPanic: false},
	}

	for _, testCase := range testCases {
		if msg, pass := checkForPanic(t, testCase.given, testCase.expectPanic); !pass {
			expected := "to panic"
			if !testCase.expectPanic {
				expected = "NOT " + expected
			}

			t.Errorf("given %v, expected %s", testCase.given, msg)
		}
	}
}

func checkForPanic(t *testing.T, given []interface{}, expectPanic bool) (string, bool) {
	if didPanic(t, func() { contract.AssertNotNil(given...) }) != expectPanic {
		expected := "to panic"
		if !expectPanic {
			expected = "NOT " + expected
		}

		return expected, false
	}

	return "", true
}

func didPanic(t *testing.T, fn func()) (ok bool) {
	defer func() {
		if r := recover(); r != nil {
			ok = true
		}
	}()

	fn()

	return
}
