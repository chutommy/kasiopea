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

	if N < 3 {
		return 1
	}

	// construct left and right
	left, right := make([]int, N), make([]int, N)
	left[0], left[N-1] = 1, 1
	copy(right, left)

	// create left
	for n := 1; n < N; n++ {

		c := inp[n]
		l := left[n-1] + 1

		if c < l {
			left[n] = c
		} else {
			left[n] = l
		}
	}

	// create right
	for n := N - 2; n > 0; n-- {

		c := inp[n]
		r := right[n+1] + 1

		if c < r {
			right[n] = c
		} else {
			right[n] = r
		}
	}

	// merge left and right
	for n := 0; n < N; n++ {

		l := left[n]
		r := right[n]

		if l < r {
			inp[n] = l
		} else {
			inp[n] = r
		}
	}

	// find max
	time := inp[0]
	for n := 1; n < N; n++ {
		if v := inp[n]; v > time {
			time = v
		}
	}

	return time
}
