package main

import (
	"fmt"
	"log"
	"math"
	"os"
)

func main() {

	// create output file
	f, err := os.Create("out.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// get T
	var T int
	fmt.Scan(&T)

	// range over problem set
	for t := 0; t < T; t++ {

		// get N
		var N int
		fmt.Scan(&N)

		// get DNAs
		var dna1, dna2 string
		fmt.Scanln(&dna1)
		fmt.Scanln(&dna2)

		c := solve(dna1, dna2, N)
		fmt.Fprintln(f, c)
		fmt.Println(c)
	}
}

// solve solves the problem.
func solve(dna1, dna2 string, l int) int {
	count := 0
	swap := true

	for i := 0; ; {

		for swap == (dna1[i] == dna2[i]) {
			i++
			if i == l {
				result := math.Ceil(float64(count) / 2)
				return int(result)
			}
		}

		swap = !swap
		count++
	}
}
