package basics

import (
	"fmt"
	"sync"
)

func add(wg *sync.WaitGroup, a int, b int) {
	defer wg.Done() // Mark the wait group as done when the function completes
	result := a + b
	fmt.Printf("The sum of %d and %d is %d\n", a, b, result)
}

func go_routine_main() {
	// Go Routine
	// A go routine is a lightweight thread of execution managed by the Go runtime.
	// It allows you to run functions concurrently, enabling efficient use of resources and improved performance.
	// To create a go routine, you can use the `go` keyword followed by a function call. For example:
	// go myFunction()
	// When you start a go routine, it runs concurrently with the main function and other go routines. The Go runtime schedules the execution of go routines, allowing them to run simultaneously.
	// Go routines are particularly useful for tasks that involve I/O operations, such as network requests or file handling, as they can help improve responsiveness and throughput.
	// It's important to note that go routines are not preemptively scheduled like threads in other programming languages. Instead, they are cooperatively scheduled, meaning that they yield control voluntarily when they are waiting for I/O or when they explicitly call `runtime.Gosched()`.
	// To synchronize go routines and ensure that they complete before the main function exits, you can use channels or wait groups from the `sync` package.
	// Overall, go routines are a powerful feature of Go that allows for concurrent programming and can help improve the performance of your applications.

	var wg sync.WaitGroup

	wg.Add(1) // Increment the wait group counter

	go add(&wg, 5, 10) // Start the add function as a go routine

	wg.Wait() // Wait for all go routines to finish
}
