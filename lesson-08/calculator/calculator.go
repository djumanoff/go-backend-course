// Package calculator provides simple math operations
package calculator

// Calculator interface describing all the operations supported by our calculator
type Calculator interface {
	Sum(...int64) int64

	Multiply(...int64) int64
}

type defaultCalc struct{}

// Summarize all numbers passed in
func (clc *defaultCalc) Sum(args ...int64) int64 {
	var result int64

	for _, num := range args {
		result = result + num
	}

	return result
}

// Multiplies all numbers passed in
func (clc *defaultCalc) Multiply(args ...int64) int64 {
	var result int64 = 1

	for _, num := range args {
		result = result * num
	}
	return result
}

// Returns default calculator implementation
func NewCalc() Calculator {
	return &defaultCalc{}
}

// Summarize all numbers passed in
func Sum(args ...int64) int64 {
	var result int64

	for _, num := range args {
		result = result + num
	}

	return result
}

// Multiplies all numbers passed in
func Multiply(args ...int64) int64 {
	var result int64 = 1

	for _, num := range args {
		result = result * num
	}

	return result
}
