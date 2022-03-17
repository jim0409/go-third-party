package main

// Importing sync/atomic, math/rand,
// fmt, sync, and time
import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

// Using sync.WaitGroup in order to
// wait for a collection of
// goroutines to finish
var waittime sync.WaitGroup

// Declaring atomic variable
var atmvar int32

// Defining increment function
func hike(S string) {

	// For loop
	for i := 1; i < 7; i++ {

		// Calling sleep method with its duration
		// and also calling rand.Intn method
		time.Sleep(time.Duration((rand.Intn(5))) * time.Millisecond)

		// Calling AddInt32 method with its
		// parameter
		atomic.AddInt32(&atmvar, 1)

		// Prints output
		fmt.Println(S, i, "count ->", atmvar)
	}

	// Wait completed
	waittime.Done()
}

// Main function
func main() {

	// Calling Add method w.r.to
	// waittime variable
	waittime.Add(2)

	// Calling hike method with
	// values
	go hike("cat: ")
	go hike("dog: ")

	// Calling wait method
	waittime.Wait()

	// Prints the value of last count
	fmt.Println("The value of last count is :", atmvar)
}
