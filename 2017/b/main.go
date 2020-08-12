package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	// prepare output
	f, err := os.Create("out.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// get T
	var T int
	fmt.Scanln(&T)

	// range over problems
	for t := 0; t < T; t++ {

		// get N
		var N int
		fmt.Scanln(&N)

		// flowers
		flowers := make([]int, N)

		// range over heigths to get the records
		for n := 0; n < N; n++ {

			// get the record
			var h int
			fmt.Scanf("%d", &h)
			flowers[n] = h
		}

		shifts := sortable(flowers)
		if other := len(flowers) - shifts; other < shifts {
			shifts = other
		}
		fmt.Fprintln(f, shifts)
	}
}

// sortable returns the numbers of shifts needed to sort the flowers,
// if flowers can not be sorted, it returns -1.
func sortable(ii []int) int {
	l := len(ii)

	if l < 3 {
		return 0
	}

	// preparation
	line := false // the break
	lineN := 0
	sortable := true
	// asc
	for i := 1; i < l; i++ {
		if ii[i-1] > ii[i] {
			if !line {
				line = true
				lineN = i
				continue
			} else {
				sortable = false
				break
			}
		}
	}
	if sortable {
		if line && ii[0] >= ii[l-1] {
			return lineN
		} else if !line {
			return 0
		}
	}

	// preparation
	line = false // the break
	lineN = 0
	sortable = true
	// desc
	for i := 1; i < l; i++ {
		if ii[i-1] < ii[i] {
			if !line {
				line = true
				lineN = i
				continue
			} else {
				sortable = false
				break
			}
		}
	}
	if sortable {
		if line && ii[0] <= ii[l-1] {
			return lineN
		} else if !line {
			return 0
		}
	}

	return -1
}
