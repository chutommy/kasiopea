package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	// create output file
	f, err := os.Create("d.out")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// get T
	var T int
	fmt.Scan(&T)
	// range over each problem
	for t := 0; t < T; t++ {

		// get N
		var N int
		fmt.Scan(&N)

		// lamps
		lamps := make([]int, N)
		for n := 0; n < N; n++ {
			// store input
			var i int
			fmt.Scan(&i)
			lamps[n] = i
		}

		// solve
		res := solve(N, lamps)
		fmt.Fprintln(f, res)
	}
}

func solve(N int, lamps []int) int {

	// initialize memodry for DP
	min := make([]int, N)
	min[0] = lamps[0]
	if N > 1 {
		min[1] = lamps[0] + lamps[1]
	}

	// range over lamps and calculate best solution for each
	for n := 2; n < N; n++ {

		// get last two
		a, b := min[n-2], min[n-1]
		// compare and set
		if a < b {
			min[n] = a + lamps[n]
		} else {
			min[n] = b + lamps[n]
		}
	}

	return min[N-1]
}
