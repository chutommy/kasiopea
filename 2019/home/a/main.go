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

		ok := true
		for n := 1; n <= N; n++ {

			// get a number
			var c int
			fmt.Scan(&c)

			// validate
			if ok {
				if c != n && c != 0 {
					ok = false
				}
			}
		}

		// print out result
		if ok {
			f.WriteString("ANO\n")
		} else {
			f.WriteString("NE\n")
		}
	}
}
