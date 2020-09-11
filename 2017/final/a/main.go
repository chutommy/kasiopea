package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	// create file
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

		// get inputs
		inputs := make([]int, N)
		for n := 0; n < N; n++ {

			fmt.Scan(&inputs[n])
		}

		// solve
		solution := solve(inputs)
		fmt.Fprintf(f, "%d\n", solution)
	}
}

func solve(inp []int) int {
	l := len(inp)

	// implement status array
	status := make([]int, l)
	for i := 0; i < l; i++ {
		value := inp[i]

		// validate
		if !(value >= l) {
			status[value]++
		}
	}

	// count diplomats
	count := status[0]
	for i := 1; i < len(status); i++ {

		if count >= i {
			count += status[i]
		} else {
			break
		}
	}

	return count
}
