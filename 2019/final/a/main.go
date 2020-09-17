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

	// range over problems
	for t := 0; t < T; t++ {

		// get N
		var N int
		fmt.Scan(&N)

		// get total
		total := -N + 1
		for n := 0; n < N; n++ {

			// get i
			var i int
			fmt.Scan(&i)

			total += i
		}

		fmt.Fprintln(f, total)
	}
}
