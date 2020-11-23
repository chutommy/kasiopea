package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	// create an output file
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

		// get input
		var W, WO, P int
		fmt.Scanf("%d %d %d", &W, &WO, &P)

		// solve
		if WO < W-P {
			fmt.Fprintln(f, "REKLAMU")
		} else {
			fmt.Fprintln(f, "NE REKLAMU")
		}
	}
}
