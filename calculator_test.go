package calculator_test

import (
	"calculator"
	"math/rand"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Parallel()

	type testCase struct {
		inputs []float64
		want   float64
	}

	testCases := []testCase{
		{inputs: []float64{2, 4, 6}, want: 12},
		{inputs: []float64{-2, -4, -6}, want: -12},
		{inputs: []float64{10, 5, 3, 7, 5, 20}, want: 50},
		{inputs: []float64{10, 5, 3, 7, 5, 20, 7, 5}, want: 62},
	}

	for _, tc := range testCases {
		got := calculator.Add(tc.inputs...)
		if tc.want != got {
			t.Errorf("want %f, got %f", tc.want, got)
		}
	}
}

func TestAddRandom(t *testing.T) {
	for i := 0; i < 1000; i++ {
		a := rand.Float64() * 10000
		b := rand.Float64() * 10000
		want := a + b

		got := calculator.Add(a, b)

		if got != want {
			t.Errorf("want %f, got %f", want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()

	type testCase struct {
		inputs []float64
		want   float64
	}

	testCases := []testCase{
		{inputs: []float64{10, 3, 5, 5}, want: -3},
		{inputs: []float64{10, 3, 5, 5, -10, 13}, want: -6},
		{inputs: []float64{10, 3, 5, 5, -10}, want: 7},
	}

	for _, tc := range testCases {
		got := calculator.Subtract(tc.inputs...)
		if tc.want != got {
			t.Errorf("want %f, got %f", tc.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()

	type testCase struct {
		a, b float64
		want float64
		test string
	}

	testCases := []testCase{
		{a: 5, b: 5, want: 25, test: "Multiplying two positive numbers results in an other positive number"},
		{a: 6, b: 0, want: 0, test: "Multiplying a number with zero results in zero"},
		{a: 40, b: -5, want: -200, test: "Multiplying a positive number with a negative number results in an other positive number"},
		{a: -40, b: -3, want: 120, test: "Multiplying two negative numbers results in a positive number"},
	}

	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)

		if tc.want != got {
			t.Errorf("%s - want %f, got %f", tc.test, tc.want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()

	type testCase struct {
		a, b        float64
		want        float64
		errExpected bool
	}

	testCases := []testCase{
		{a: 6, b: 3, want: 2, errExpected: false},
		{a: 6, b: 0, want: 0, errExpected: true},
		{a: 0, b: 6, want: 0, errExpected: false},
	}

	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)

		errReceived := err != nil

		if tc.errExpected != errReceived {
			t.Fatalf("Divide (%f, %f): unexpected error status: %v", tc.a, tc.b, errReceived)
		}

		if !tc.errExpected && tc.want != got {
			t.Errorf("want %f, got %f", tc.a, tc.b)
		}
	}
}

func TestSqrt(t *testing.T) {
	t.Parallel()

	type testCase struct {
		a           float64
		want        float64
		errExpected bool
	}

	testCases := []testCase{
		{a: 16, want: 4, errExpected: false},
		{a: -16, want: 0, errExpected: true},
		{a: 0, want: 0, errExpected: false},
	}

	for _, tc := range testCases {
		got, err := calculator.Sqrt(tc.a)

		if tc.errExpected != (err != nil) {
			t.Fatalf("Sqrt (%f): Unexpected error status: %v", tc.a, err)
		}

		if got != tc.want {
			t.Errorf("want %f, got %f", tc.want, got)
		}
	}
}

func TestEvaluate(t *testing.T) {
	t.Parallel()

	type testCase struct {
		expr        string
		want        float64
		errExpected bool
	}

	testCases := []testCase{
		{expr: "2 + 4", want: 6, errExpected: false},
		{expr: "-2 + 4", want: 2, errExpected: false},
		{expr: "2 / 0", want: 0, errExpected: true},
		{expr: "20 / 5", want: 4, errExpected: false},
		{expr: "20/ 5", want: 0, errExpected: true},
		{expr: "20 /5", want: 0, errExpected: true},
		{expr: "20 * 5", want: 100, errExpected: false},
		{expr: "20 - 5", want: 15, errExpected: false},
	}

	for _, tc := range testCases {
		got, err := calculator.Evaluate(tc.expr)

		if tc.errExpected != (err != nil) {
			t.Fatalf("Evaluate %s: Unexpected error status: %v", tc.expr, err)
		}

		if got != tc.want {
			t.Errorf("Evaluate: want %f, got %f", tc.want, got)
		}
	}
}
