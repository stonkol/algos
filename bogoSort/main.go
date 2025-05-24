package main

import (
	"fmt"
	"math/rand"
	"time"
)

func isSorted(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			fmt.Println("Nope!\n")
			return false
		}
	}
	fmt.Println("Yeap!")
	return true
}

func shuffle(arr []int) {
	for i := range arr {
		j := rand.Intn(len(arr))
		arr[i], arr[j] = arr[j], arr[i]
		fmt.Printf("[%d,%d] = %v\n", i, j, arr)
	}
}

func bogoSort(arr []int) {
	for !isSorted(arr) {
		shuffle(arr)
		fmt.Printf("- Is %v sorted? ", arr)
	}
}
func main() {
	rand.Seed(time.Now().UnixNano()) // seed random number generator

	numbers := []int{3, 1, 2, 5}

	fmt.Println("Sort the numbers:", numbers)
	bogoSort(numbers)
}
