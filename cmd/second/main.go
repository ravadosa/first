package main

import (
	"fmt"

	"github.com/ravadosa/first/ops"
)

func main() {
	complex := createComplex()
	fmt.Printf("%#v\n", complex.Add())
}

func createComplex() *ops.ComplexNumber {
	complex := ops.ComplexNumber{
		Real:      10,
		Imaginary: 15,
	}
	return &complex
}
