package main

import (
    "testing"
)

func TestEvaluateAverageGrade(t *testing.T) {
    tests := []struct {
        name     string
        grades   []int
        expected int
    }{
        {"average of empty slice", []int{}, 0},
        {"average of one element", []int{100}, 100},
        {"average of multiple elements", []int{50, 75, 100}, 75},
        {"average with zero grades", []int{0, 0, 0}, 0},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := evaluateAverageGrade(tt.grades)
            if result != tt.expected {
                t.Errorf("got %d, want %d", result, tt.expected)
            }
        })
    }
}

func TestIsValidGrade(t *testing.T) {
    tests := []struct {
        name     string
        grade    int
        expected bool
    }{
        {"valid grade", 85, true},
        {"invalid grade negative", -5, false},
        {"invalid grade over 100", 105, false},
        {"boundary grade 0", 0, true},
        {"boundary grade 100", 100, true},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := isValidGrade(tt.grade)
            if result != tt.expected {
                t.Errorf("got %v, want %v", result, tt.expected)
            }
        })
    }
}