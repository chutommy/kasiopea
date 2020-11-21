package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	// create an output file
	f, err := os.Create("b.out")
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

		// store problem
		arr := make([]int, N)
		for n := 0; n < N; n++ {

			// get integer
			var i int
			fmt.Scan(&i)
			arr[n] = i
		}

		// solve
		s := solve(N, arr)
		fmt.Fprintln(f, s)
	}
}

func solve(N int, arr []int) int {
	min := top(arr[:3])

	for i := 1; i < N-2; i++ {
		v := top(arr[i : i+3])
		if v < min {
			min = v
		}
	}

	return min
}

func top(arr []int) int {
	a, b, c := arr[0], arr[1], arr[2]
	if c > b {
		b = c
	}
	if b > a {
		a = b
	}
	return a
}
