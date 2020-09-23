package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	// create output file
	f, err := os.Create("c.out")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// get T
	var T int
	fmt.Scan(&T)

	for t := 0; t < T; t++ {

		// get N, M
		var N, M int
		fmt.Scan(&N, &M)

		// get dragons
		drags := make([]int, N)
		for n := 0; n < N; n++ {
			var i int
			fmt.Scan(&i)
			drags[n] = i
		}

		// get knights
		knis := make([]int, M)
		for m := 0; m < M; m++ {
			var i int
			fmt.Scan(&i)
			knis[m] = i
		}

		// solve
		s := solve(N, M, drags, knis)
		fmt.Fprintln(f, s)
	}
}

func solve(N, M int, drags, knis []int) int {

	var n, m int
	var sum int

	for n < N && m < M {

		d := drags[n]
		k := knis[m]

		if d <= k {
			sum += k
			n++
			m++
		} else {
			m++
		}
	}

	if n < N {
		return -1
	} else {
		return sum
	}
}
