package main

import (
	"reflect"
	"testing"
)

func TestWordFrequency(t *testing.T) {
	tests := []struct {
		input    string
		expected map[string]int
	}{
		{
			input: "Hello, World! Hello, Universe.",
			expected: map[string]int{
				"hello":    2,
				"world":    1,
				"universe": 1,
			},
		},
		{
			input: "Go is great. Go is fun.",
			expected: map[string]int{
				"go":   2,
				"is":   2,
				"great": 1,
				"fun":  1,
			},
		},
		{
			input: "Test-driven development, test-driven design.",
			expected: map[string]int{
				"test": 2,
                "driven" : 2,
				"development": 1,
				"design":      1,
			},
		},
		{
			input: "A.B.C. A B C A-B-C",
			expected: map[string]int{
				"a": 3,
				"b": 3,
				"c": 3,
			},
		},
	}

	for _, test := range tests {
		result := wordFrequency(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For input %q, expected %v, but got %v", test.input, test.expected, result)
		}
	}
}
