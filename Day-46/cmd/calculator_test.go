package main

import (
	"errors"
	"testing"
)

func TestDivide(t *testing.T) {
	testCases := []struct {
		name      string
		a, b      int
		expected  int
		expectErr error
	}{
		{"Divide positives", 10, 2, 5, nil},
		{"Divide by zero", 10, 0, 0, errors.New("cannot divide by zero")},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := Divide(tc.a, tc.b)
			if result != tc.expected || (err != nil && err.Error() != tc.expectErr.Error()) {
				t.Errorf("Expected result %d and error %v, got result %d and error %v",
					tc.expected, tc.expectErr, result, err)
			}
		})
	}
}
