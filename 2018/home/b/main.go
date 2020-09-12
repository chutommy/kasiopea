package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	// create output file
	f, err := os.Create("b.out")
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

		// get K
		var K int
		fmt.Scan(&K)

		result := solve(N, K)
		fmt.Fprintln(f, result)
	}
}

func solve(N, K int) int {

	// create a status holder
	status := make([]bool, K)

	for n := 0; n < N; n++ {

		// get Item
		var i int
		fmt.Scan(&i)
		i--

		if status[i] {
			status[i] = false
		} else {
			status[i] = true
		}
	}

	var count int
	// count unpaired items
	for k := 0; k < K; k++ {
		if status[k] {
			count++
		}
	}

	return count
}
