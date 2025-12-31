package calc

import (
	"math"
	"testing"
)

func TestCalculator(t *testing.T) {
	type testCase struct {
		name       string
		input      string
		expected   float64
		shouldErr  bool
		useDegrees bool
	}

	tests := []testCase{
		{"Addition", "2 + 2", 4, false, false},
		{"Priority", "2 + 3 * 4", 14, false, false},
		{"Parantheses", "(2 + 3) * 4", 20, false, false},
		{"Decimal", "1.5 + 2.5", 4, false, false},

		{"Sqrt", "sqrt(16)", 4, false, false},
		{"Binary Root", "root(27, 3)", 3, false, false},
		{"Negative Root Odd Number", "root(-27, 3)", -3, false, false},

		{"Sin Rad", "sin(0)", 0, false, false},
		{"Tan Rad", "tan(0)", 0, false, false},

		{"Sin 30 Deg", "sin(30)", 0.5, false, true},
		{"Cos 60 Deg", "cos(60)", 0.5, false, true},
		{"Tan 45 Deg", "tan(45)", 1, false, true},
		{"Cos 90 Deg Fix", "cos(90)", 0, false, true},

		{"Simple assigment", "x = 10", 10, false, false},
		{"Variable usage", "x * 2", 20, false, false},

		{"Div Zero", "10 / 0", 0, true, false},
		{"Negative Sqrt", "sqrt(-2)", 0, true, false},
		{"Tan 90 Deg", "tan(90)", 0, true, true}, // Asimptota
		{"Negative Root Even Number", "root(-4, 2)", 0, true, false},
	}

	c := New()
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			c.useDegrees = tc.useDegrees
			got, err := c.Calculate(tc.input)
			if tc.shouldErr {
				if err == nil {
					t.Errorf("Expected error at '%s', but actual result: %.4f", tc.input, got)
				}
				return
			}
			if err != nil {
				t.Errorf("Unexpected error at '%s': %v", tc.input, err)
				return
			}
			if math.Abs(got-tc.expected) > 1e-9 {
				t.Errorf("To '%s': expected %.4f, actual result %.4f", tc.input, tc.expected, got)
			}
		})
	}
}
