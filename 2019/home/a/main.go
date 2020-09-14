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

		// validate Zuzanka's interrogation
		ok := true
		for n := 1; n <= N; n++ {

			// get a number
			var i int
			fmt.Scan(&i)

			if ok != false {
				if i == 0 {
					continue
				} else if i != n {
					ok = false
				}
			}
		}

		if ok {
			fmt.Fprintln(f, "ANO")
		} else {
			fmt.Fprintln(f, "NE")
		}
	}
}
