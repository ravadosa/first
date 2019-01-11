package main

import (
	"fmt"

	"github.com/ravadosa/first/ops"
)

func main() {
	fmt.Println("Hello from first!")
	sum, diff, err := ops.PerformSumDiff(10, 15)
	if err == nil {
		fmt.Printf("Sum: %d, Diff: %d\n", sum, diff)
	} else {
		fmt.Println(err.Error())
	}
}
