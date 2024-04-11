package tests

import (
	"github.com/RakhimovAns/FinalYandexTask/govaluate"
	"testing"
)

func TestCalculate(t *testing.T) {
	type testCase struct {
		name       string
		expression string
		expected   float64
	}
	tests := []testCase{
		{
			name:       "plus",
			expression: "1+2",
			expected:   3.0,
		},
		{
			name:       "minus",
			expression: "5-3",
			expected:   2.0,
		},
		{
			name:       "multiply",
			expression: "2*3",
			expected:   6.0,
		},
		{
			name:       "divide",
			expression: "10/2",
			expected:   5.0,
		},
		{
			name:       "modulo",
			expression: "10%3",
			expected:   1.0,
		},
		{
			name:       "mixed",
			expression: "(1+2)*3-4/2",
			expected:   7.0,
		},
	}
	for _, tc := range tests {
		expr, err := govaluate.NewEvaluableExpression(tc.expression)
		if err != nil {
			t.Errorf("error with parsing: %v", err)
			continue
		}
		result, err := expr.Evaluate(nil)
		if err != nil {
			t.Errorf("error evaluating expression: %v", err)
			continue
		}
		actual, ok := result.(float64)
		if !ok {
			t.Errorf("result is not a float64: %v", result)
			continue
		}
		if actual != tc.expected {
			t.Errorf("unexpected result for test '%s', got: %v, want: %v", tc.name, actual, tc.expected)
		}
	}
}
