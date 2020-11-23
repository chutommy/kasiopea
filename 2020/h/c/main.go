package main

import (
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {

	// create an output file
	f, err := os.Create("c.out")
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

		// get distances
		dists := make([]int, N)
		for n := 0; n < N; n++ {
			fmt.Scan(&dists[n])
		}

		// solve
		s := solve(N, dists)
		fmt.Fprintln(f, s)

		// print status
		fmt.Printf("%d/%d done\n", t+1, T)
	}
}

func solve(N int, dists []int) float64 {

	// sort
	sort.Ints(dists)

	// get maximal dist diff
	var max int
	for i := 1; i < N; i++ {
		if v := dists[i] - dists[i-1]; v > max {
			max = v
		}
	}

	return float64(max) / 2
}
