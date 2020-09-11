package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	// create file
	f, err := os.Create("b.out")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// get T
	var T int
	fmt.Scan(&T)

	for t := 0; t < T; t++ {

		// get input
		var inp int
		fmt.Scan(&inp)

		// solve
		solution := decomp(inp)
		fmt.Fprintf(f, "%s\n", solution)
	}
}

func decomp(i int) string {

	// base value
	if i == 1 {
		return "1"
	}

	if i%2 == 0 {
		s := decomp(i / 2)
		if s == "1" {
			return fmt.Sprintf("1+1")
		}
		return fmt.Sprintf("(1+1)*(%s)", s)
	}

	return fmt.Sprintf("1+(%s)", decomp(i-1))
}
