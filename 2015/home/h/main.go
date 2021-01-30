package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	// create output file
	f, err := os.Create("h.out")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// get T
	var T int
	fmt.Scan(&T)

	for t := 0; t < T; t++ {

		// get M, N
		var M, N int
		fmt.Scan(&M, &N)

		// get input
		inp := make([]int, N)
		for n := 0; n < N; n++ {
			var i int
			fmt.Scan(&i)
			inp[n] = i
		}

		// solve
		sol := solve(inp, M)
		fmt.Fprintln(f, sol)
	}
}

func solve(inp []int, M int) int {

	q := inp

	for {
		val := q[0]

		if val%M == 0 {
			return val
		}

		for _, v := range inp {
			q = append(q, val*10+v)
		}

		q = q[1:]
	}
}
