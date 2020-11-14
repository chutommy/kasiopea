package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	// create output file
	f, err := os.Create("e.out")
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
			var i int
			fmt.Scan(&i)
			inp[n] = i
		}

		// solve
		s := solve(inp, N)
		fmt.Fprintln(f, s)
		fmt.Printf("%d/%d\n", t+1, T)
	}
}

func solve(inp []int, N int) int {

	// set sides
	if inp[0] != 0 {
		inp[0] = 1
	}
	if inp[N-1] != 0 {
		inp[N-1] = 1
	}

	// range over the array from both sides
	for n := 1; n < N-1; n++ {
		m := N - n - 1

		// process n
		if i := inp[n-1] + 1; i < inp[n] {
			inp[n] = i
		}

		// process m
		if i := inp[m+1] + 1; i < inp[m] {
			inp[m] = i
		}
	}

	// get max
	max := 0
	for _, v := range inp {
		if v > max {
			max = v
		}
	}

	return max
}
