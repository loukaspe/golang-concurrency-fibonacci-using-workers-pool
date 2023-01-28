package main

import (
	"fmt"
	"github.com/loukaspe/workers-pool/fibonacci/app"
	"time"
)

func main() {
	fibonacciRequestedNumber := 45

	start := time.Now()

	// Solution without workers
	//for i := 1; i <= fibonacciRequestedNumber; i++ {
	//	fmt.Printf("%d ", FibonacciRecursion(i))
	//}

	// Solution with workers
	app.RecursiveFibonacciWithWorkersPool(fibonacciRequestedNumber)

	fmt.Printf("\nExecution took %s", time.Since(start))
}
