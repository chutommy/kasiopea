package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	// create an output file
	f, err := os.Create("e.out")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// get T
	var T int
	fmt.Scan(&T)

	// range over problems
	for t := 0; t < T; t++ {

		// get N
		var N int
		fmt.Scan(&N)

		// store lengths
		lengths := make([]int, N)
		for n := 0; n < N; n++ {
			fmt.Scan(&lengths[n])
		}

		// solve
		s := solve(N, lengths)
		fmt.Fprintln(f, s)

		// print status
		fmt.Printf("%d/%d done\n", t+1, T)
	}
}

func solve(N int, lx []int) int {

	// base
	if N < 3 {
		return lx[0]
	}

	// init mem for DP
	best := make([]int, N)
	best[0] = lx[0]
	best[1] = lx[0] + lx[1]

	// range
	for n := 2; n < N; n++ {

		// get last two
		a, b := best[n-2], best[n-1]
		// compare and set
		if a < b {
			best[n] = a + lx[n]
		} else {
			best[n] = b + lx[n]
		}
	}

	// select smaller
	a, b := best[N-1], best[N-2]
	if a < b {
		return a
	}
	return b
}
