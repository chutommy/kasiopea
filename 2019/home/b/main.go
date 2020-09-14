package main

import (
	"fmt"
	"log"
	"os"
	"strings"
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

	// range over problems
	for t := 0; t < T; t++ {

		// get N
		var N int
		fmt.Scan(&N)

		// get M
		var M int
		fmt.Scan(&M)

		status := make([]bool, M)
		// range over molecules
		for n := 0; n < N; n++ {

			// range over the molecule
			for m := 0; m < M; m++ {

				// get the value
				var v int
				fmt.Scan(&v)

				// if 1: switch
				if v == 1 {
					status[m] = !(status[m])
				}
			}
		}

		// build result
		result := make([]string, M)
		for m := 0; m < M; m++ {

			if status[m] {
				result[m] = "1"
			} else {
				result[m] = "0"
			}
		}

		fmt.Fprintln(f, strings.Join(result, " "))
	}
}
