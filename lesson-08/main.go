package main

import (
	"github.com/djumanoff/go-backend/lesson-08/calculator"
)

func main() {
	calc := calculator.NewCalc()
	calc.Sum(111, 2, 22)
}
