package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

// checks if the slice is sorted
func isSorted(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			fmt.Println("Nope!\n")
			return false
		}
	}

	return true
}

// shuffle is Fisher-Yates shuffle algorithm, an efficient and unbiased way to shuffle elements.
func shuffle(arr []int) {
	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
		fmt.Printf("[%d,%d] = %v\n", i, j, arr)
	})
}

func bogoSort(arr []int) {
	for !isSorted(arr) {
		shuffle(arr)
		fmt.Printf("Is %v sorted? ", arr)
	}
	fmt.Println("Yep!")
}
func main() {
	// Use New(NewSource()) instead of deprecated Seed
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	rand.Seed(r.Int63())

	numbers := []int{3, 1, 2, 5}

	fmt.Println("Sort the numbers:", numbers)

	// Add timeout to prevent infinite runs
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	go func() {
		bogoSort(numbers)
		cancel()
	}()

	<-ctx.Done()

	if isSorted(numbers) {
		fmt.Printf("\nSorted successfully: %v\n", numbers)
	} else {
		fmt.Println("\nFailed to sort in 5 seconds (try running again)")
	}
}
