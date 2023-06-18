package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	slice := generateSlice(20)
	fmt.Println("\n--- Unsorted --- \n\n", slice)
	quicksort(slice)
	fmt.Println("\n--- Sorted ---\n\n", slice)
}

func quicksort(a []int) []int {
	// already sorted
	if len(a) < 2 {
		return a
	}

	// index left and right
	left, right := 0, len(a)-1

	// generates a valid random pivot
	pivot := rand.Int() % len(a)

	// swap right most with pivot
	// swap pivot with right most
	a[pivot], a[right] = a[right], a[pivot]

	// range over slice
	for i, _ := range a {
		// if i is less than right (which is now the pivot)
		if a[i] < a[right] {
			// swap left with i
			// swap i with left
			a[left], a[i] = a[i], a[left]
			// increment left
			left++
		}
	}
	// swap left and right
	a[left], a[right] = a[right], a[left]

	// qucksort the left half
	quicksort(a[:left])
	// quciksort the right half
	quicksort(a[left+1:])

	return a
}

// Generates a slice of size, size filled with random numbers
func generateSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}
