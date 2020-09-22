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

		var pos, max int
		for n := 0; n < N; n++ {

			// get i
			var i int
			fmt.Scan(&i)

			if i > max {
				max = i
				pos = n
			}
		}

		fmt.Fprintln(f, pos+1)
	}
}
