package calculator

import "log"

type loggingCalc struct {
	calc Calculator
}

// Summarize all numbers passed in and logs arguments and result
func (clc *loggingCalc) Sum(args ...int64) int64 {
	result := clc.calc.Sum(args...)
	log.Println("Sum", args, result)
	return result
}

// Multiplies all numbers passed in and logs arguments and result
func (clc *loggingCalc) Multiply(args ...int64) int64 {
	result := clc.calc.Multiply(args...)
	log.Println("Multiply", args, result)
	return result
}

// Returns logging calculator implementation
func NewLoggingCalc(calc Calculator) Calculator {
	return &loggingCalc{calc}
}
