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
	fmt.Scanf("%d\n", &T)

	// iterate over problems
	for t := 0; t < T; t++ {

		// get N
		var N int
		fmt.Scanf("%d\n", &N)

		// load input
		for n := 0; n < N; n++ {

			// scan line
			var i int
			fmt.Scan(&i)
		}
	}
}

func solve() {}
