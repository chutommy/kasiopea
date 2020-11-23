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

	// range over problems
	for t := 0; t < T; t++ {

		// get N
		var N int
		fmt.Scan(&N)

		// store temperatures
		temps := make([]int, N)
		for n := 0; n < N; n++ {
			// store integer
			var i int
			fmt.Scan(&i)
			temps[n] = i
		}

		// solve
		s := solve(N, temps)
		fmt.Fprintln(f, s)
	}
}

func solve(N int, temps []int) int {

	// set to the first value
	min := top(temps[:3])

	// range over temperatures
	for i := 1; i < N-2; i++ {
		if v := top(temps[i : i+3]); v < min {
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
