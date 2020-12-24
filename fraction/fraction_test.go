package fraction

import (
	"reflect"
	"testing"
)

func Test_signum(t *testing.T) {
	type testCase struct {
		name     string
		argument int
		expected int
	}
	tests := []testCase{
		{name: "Positive value", argument: 10, expected: +1},
		{name: "Negative value", argument: -12, expected: -1},
		{name: "Zero value", argument: 0, expected: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if result := signum(tt.argument); !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("signum() %v current:'%v' result:%v, expected: %v", tt.name, tt.argument, result, tt.expected)
			}
		})
	}

}

func Test_normalize(t *testing.T) {
	type testCase struct {
		name     string
		argument Fraction
		expected Fraction
	}
	tests := []testCase{
		{name: "Keep the fraction as is", argument: Fraction{1, 2}, expected: Fraction{1, 2}},
		{name: "Reduce into 1/1", argument: Fraction{2, 2}, expected: Fraction{1, 1}},
		{name: "Reduce into 2/1", argument: Fraction{4, 2}, expected: Fraction{2, 1}},
		{name: "Reduce into 3/1", argument: Fraction{9, 3}, expected: Fraction{3, 1}},
		{name: "Do not reduce 13/15", argument: Fraction{13, 15}, expected: Fraction{13, 15}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if result := tt.argument.normalize(); !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("normalize() current:'%v' result:%v, expected: %v", tt.name, result, tt.expected)
			}
		})
	}

}

func Test_Plus(t *testing.T) {
	type arguments struct {
		summand1 Fraction
		summand2 Fraction
	}
	type testCase struct {
		name     string
		args     arguments
		expected Fraction
	}

	tests := []testCase{
		{
			name: "Two proper fraction which produces a reduced result.",
			args: arguments{
				summand1: Fraction{1, 2},
				summand2: Fraction{1, 2},
			},
			expected: Fraction{1, 1},
		},
		{
			name: "A proper fraction and a non proper fraction with reduced result.",
			args: arguments{
				summand1: Fraction{1, 2},
				summand2: Fraction{3, 2},
			},
			expected: Fraction{2, 1},
		},
		{
			name: "Two proper fraction with non reduced result.",
			args: arguments{
				summand1: Fraction{1, 3},
				summand2: Fraction{2, 3},
			},
			expected: Fraction{1, 1},
		},
		{
			name: "Two proper fraction with result which can not being reduced.",
			args: arguments{
				summand1: Fraction{2, 3},
				summand2: Fraction{1, 5},
			},
			expected: Fraction{13, 15},
		},
		{
			name: "Two proper fractions where one has negative numerator with reduced result.",
			args: arguments{
				summand1: Fraction{-2, 3},
				summand2: Fraction{1, 5},
			},
			expected: Fraction{-7, 15},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if sum := tt.args.summand1.Plus(tt.args.summand2); !reflect.DeepEqual(sum, tt.expected) {
				t.Errorf("Plus() '%v' = %v, expected %v", tt.name, sum, tt.expected)
			}
		})
	}
}

func Test_Minus(t *testing.T) {
	type arguments struct {
		minuend    Fraction
		subtrahend Fraction
	}
	type testCase struct {
		name     string
		args     arguments
		expected Fraction
	}

	tests := []testCase{
		{
			name: "Subtract two proper fraction with reduced result.",
			args: arguments{
				minuend:    Fraction{1, 2},
				subtrahend: Fraction{1, 2},
			},
			expected: Fraction{0, 1},
		},
		{
			name: "Subtract proper fractions with negative result and reduced.",
			args: arguments{
				minuend:    Fraction{1, 2},
				subtrahend: Fraction{3, 2},
			},
			expected: Fraction{-1, 1},
		},
		{
			name: "Plus two proper fraction with result which can not being reduced.",
			args: arguments{
				minuend:    Fraction{2, 3},
				subtrahend: Fraction{1, 5},
			},
			expected: Fraction{7, 15},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if difference := tt.args.minuend.Minus(tt.args.subtrahend); !reflect.DeepEqual(difference, tt.expected) {
				t.Errorf("Minus(%v) %v-%v = %v, expected %v", tt.name, tt.args.minuend, tt.args.subtrahend, difference, tt.expected)
			}
		})
	}
}

func Test_Multiply(t *testing.T) {
	type arguments struct {
		minuend    Fraction
		subtrahend Fraction
	}
	type testCase struct {
		name     string
		args     arguments
		expected Fraction
	}

	tests := []testCase{
		{
			name: "Multiply two fraction with reduced result.",
			args: arguments{
				minuend:    Fraction{2, 3},
				subtrahend: Fraction{4, 5},
			},
			expected: Fraction{8, 15},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if difference := tt.args.minuend.Multiply(tt.args.subtrahend); !reflect.DeepEqual(difference, tt.expected) {
				t.Errorf("Minus(%v) %v-%v = %v, expected %v", tt.name, tt.args.minuend, tt.args.subtrahend, difference, tt.expected)
			}
		})
	}
}
