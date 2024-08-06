package main

import (
    "testing"
)

func TestIsPalindrome(t *testing.T) {
    tests := []struct {
        input    string
        expected bool
    }{
        {"A man, a plan, a canal, Panama", true},
        {"racecar", true},
        {"hello", false},
        {"", true},
        {"No 'x' in Nixon", true},
        {"Was it a car or a cat I saw?", true},
    }

    for _, test := range tests {
        t.Run(test.input, func(t *testing.T) {
            result := isPalindrome(test.input)
            if result != test.expected {
                t.Errorf("isPalindrome(%q) = %v; want %v", test.input, result, test.expected)
            }
        })
    }
}