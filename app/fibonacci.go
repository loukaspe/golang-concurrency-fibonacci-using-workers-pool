package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/loukaspe/workers-pool/fibonacci/concurrency"
)

func fibonacciWithRecursionExecutionFunction(
	ctx context.Context,
	args interface{},
) (interface{}, error) {
	number, ok := args.(int)
	if !ok {
		return nil, errors.New("wrong argument number")
	}

	if number <= 1 {
		return number, nil
	}

	result := FibonacciRecursion(number)
	fmt.Printf("%d ", result)

	return result, nil
}

func FibonacciRecursion(n int) int {
	if n <= 1 {
		return n
	}
	return FibonacciRecursion(n-1) + FibonacciRecursion(n-2)
}

func RecursiveFibonacciWithWorkersPool(n int) int {
	result := 0
	workersCount := 4

	jobs := make([]concurrency.Job, n)
	for i := 0; i < n; i++ {
		jobs[i] = concurrency.Job{
			Descriptor: concurrency.JobDescriptor{
				ID: concurrency.JobID(fmt.Sprintf("%v", i)),
			},
			ExecFn: fibonacciWithRecursionExecutionFunction,
			Args:   i,
		}
	}

	workerPool := concurrency.NewWorkerPool(workersCount)
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	go workerPool.GenerateFrom(jobs)
	go workerPool.Run(ctx)

	for {
		select {
		case r, ok := <-workerPool.Results():
			if !ok {
				continue
			}

			result = r.Value.(int)
		case <-workerPool.Done:
			return result
		default:
		}
	}
}
