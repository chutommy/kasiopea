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

		// get N
		var N int
		fmt.Scan(&N)

		// get K
		var K int
		fmt.Scan(&K)

		// get M
		var M int
		fmt.Scan(&M)

		s := solve(N, K, M)
		fmt.Fprintln(f, s)
	}
}

func solve(N, K, M int) int {

	// calculate the count
	count := N * K

	if M > N {
		M = N
	}

	// calculate the solution
	res := count / M
	if count%M != 0 {
		res++
	}

	return res
}
