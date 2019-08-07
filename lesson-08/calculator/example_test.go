package calculator_test

import (
	"fmt"

	"github.com/djumanoff/go-backend/lesson-08/calculator"
)

var clc = calculator.NewCalc()

func ExampleNewCalc() {
	newClc := calculator.NewCalc()
	fmt.Println(newClc.Sum(1, 2, 3))
	// Output:
	// 6
}

func ExampleSum() {
	result := clc.Sum(1, 2, 3)
	fmt.Println(result)
	// Output:
	// 6
}

func ExampleMultiply() {
	result := clc.Multiply(1, 2, 3, 4)
	fmt.Println(result)
	// Output:
	// 24
}
