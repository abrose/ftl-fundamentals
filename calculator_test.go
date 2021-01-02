package calculator_test

import (
	"calculator"
	"math/rand"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Parallel()

	type testCase struct {
		a, b float64
		want float64
	}

	testCases := []testCase{
		{a: 2, b: 2, want: 4},
		{a: 1, b: 1, want: 2},
		{a: 5, b: 0, want: 5},
	}

	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
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
		a, b float64
		want float64
	}

	testCases := []testCase{
		{a: 5, b: 3, want: 2},
		{a: 5.4, b: 1.2, want: 4.2},
		{a: 1.5, b: 3, want: -1.5},
	}

	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)
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
}
