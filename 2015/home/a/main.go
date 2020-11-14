package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	// create output file
	f, err := os.Create("a.out")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// get T
	var T int
	fmt.Scan(&T)

	for t := 0; t < T; t++ {

		// get N
		var N int
		fmt.Scan(&N)

		// get input
		inp := make([]int, N)
		for n := 0; n < N; n++ {

			// get int
			var i int
			fmt.Scan(&i)
			// store
			inp[n] = i
		}

		// solve
		s := solve(inp, N)
		fmt.Fprintln(f, s)
	}
}

func solve(inp []int, N int) int {

	count := 0
	for n := 0; n < N-1; {
		n += inp[n]
		count++
	}

	return count
}
