package main

import (
	"fmt"
	"sync"
	"time"
)

func unicornSleep(num int, wg *sync.WaitGroup) {
	defer wg.Done()

	// sleep `num` milliseconds
	time.Sleep(time.Duration(num) * 300 * time.Nanosecond)
	fmt.Println(num)
}

func main() {
	numbers := []int{21, 1, 54321, 321, 4321}

	// pointer to a WaitGroup used to synchronize goroutines.
	var wg sync.WaitGroup
	// call
	wg.Add(len(numbers))
	for _, i := range numbers {
		// launching a goroutine for each
		go unicornSleep(i, &wg)
	}
	wg.Wait()
}
