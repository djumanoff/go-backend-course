package calculator_test

import (
	"testing"

	"github.com/djumanoff/go-backend/lesson-08/calculator"
)

var calc = calculator.NewCalc()

var sumTests = []struct {
	args     []int64
	expected int64
}{
	{[]int64{1, 2, 3, 4}, 10},
	{[]int64{5, 6, 7, 8}, 26},
}

var multiplyTests = []struct {
	args     []int64
	expected int64
}{
	{[]int64{1, 2, 3, 4}, 24},
	{[]int64{5, 6, 7}, 210},
}

func TestSum(t *testing.T) {
	for _, test := range sumTests {
		result := calc.Sum(test.args...)
		if result != test.expected {
			t.Errorf("summ test: %q does not match %q", result, test.expected)
		}
	}
}

func TestMultiply(t *testing.T) {
	for _, test := range multiplyTests {
		result := calc.Multiply(test.args...)
		if result != test.expected {
			t.Errorf("multiply test: %q does not match %q", result, test.expected)
		}
	}
}

var benchSum int64
var benchMultiply int64

func BenchmarkSum10(b *testing.B) {
	// run the Fib function b.N times
	var r int64
	for n := 0; n < b.N; n++ {
		r = calc.Sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	}
	benchSum = r
}

func BenchmarkMultiply10(b *testing.B) {
	// run the Fib function b.N times
	var r int64
	for n := 0; n < b.N; n++ {
		r = calc.Multiply(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	}
	benchMultiply = r
}
