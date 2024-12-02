package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5

	if result != expected {
		t.Errorf("Add(2, 3) = %d; expected %d", result, expected)
	}
}

func TestAdd_TableDriven(t *testing.T) {
	testCases := []struct {
		name     string
		a, b     int
		expected int
	}{
		// {"Add two positive integers": 12, 13, 25},
		{"Add two negative integers:", -2, -3, -5},
		{"Add a positive and a negative integer:", 2, -3, -1},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Add(tc.a, tc.b)
			if result != tc.expected {
				t.Errorf("For %s, expected %d but got %d", tc.name, tc.expected, result)
			}
		})
	}
}
