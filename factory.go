package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

var factory OperatorFactory
var x, y int
var op Operation

func init() {
	factory = OperatorFactory{}
}

func main() {

	if len(os.Args) != 4 {
		fmt.Println("go run factory.go <int> <operator> <int>")
		os.Exit(0)
	}
	// Converting input
	x, _ = strconv.Atoi(os.Args[1])
	oper := os.Args[2]
	y, _ = strconv.Atoi(os.Args[3])
	op := factory.ConvertOperator(oper)

	// getting operation
	operation := factory.GetOperator(op)

	result, err := operation.Calculate(x, y)

	fmt.Println("The result for", x, op, y, "is", result, ". Error?", err)
}

// Operator This is the main interface
// ----------------
type Operator interface {
	Calculate(x, y int) (int, error)
}

// Opertion
// ----------------
type Operation int

const (
	Sum Operation = iota
	Subtraction
	Multiplication
	Division
)

// OperatorFactory the one that will give me the real implementation
// ----------------
type OperatorFactory struct{}

// GetOperator will return a Operator according to the operation type
func (OperatorFactory) GetOperator(operation Operation) Operator {
	switch operation {
	case Sum:
		return Summator{}
	case Subtraction:
		return Subtractor{}
	case Multiplication:
		return Multiplicator{}
	case Division:
		return Divisor{}
	default:
		return NotImplemented{}
	}
}

// Summator Operator
// ----------------
type Summator struct{}

func (Summator) Calculate(x, y int) (int, error) {
	return x + y, nil
}

// Subtractor Operator
// ----------------
type Subtractor struct{}

func (Subtractor) Calculate(x, y int) (int, error) {
	return x - y, nil
}

// Multiplier Operator
// ----------------
type Multiplicator struct{}

func (Multiplicator) Calculate(x, y int) (int, error) {
	return x * y, nil
}

// Divisor Operator
// ----------------
type Divisor struct{}

func (Divisor) Calculate(x, y int) (int, error) {
	return x / y, nil
}

// NotImplemented
// ----------------
type NotImplemented struct{}

func (NotImplemented) Calculate(x, y int) (int, error) {
	return 0, errors.New("Not implemented")
}

/*****/

func (op Operation) String() string {
	switch op {
	case Sum:
		return "+"
	case Subtraction:
		return "-"
	case Multiplication:
		return "*"
	case Division:
		return "/"
	default:
		return "err"
	}
}

func (OperatorFactory) ConvertOperator(op string) Operation {
	switch op {
	case "+":
		return Sum
	case "-":
		return Subtraction
	case "*":
		return Multiplication
	case "/":
		return Division
	}
	return Operation(999)
}
