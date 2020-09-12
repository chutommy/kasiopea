package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	// create output file
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

		// get J
		var J int
		fmt.Scan(&J)

		// get V
		var V int
		fmt.Scan(&V)

		sol := solve(N, J, V)
		fmt.Fprintln(f, sol)
	}
}

func solve(N, J, V int) string {

	valueJ := N * J
	valueV := N / 7 * V

	if valueJ > valueV {
		return "tydenni"
	}
	return "jednodenni"
}
