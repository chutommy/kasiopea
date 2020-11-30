package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	// create an output file
	f, err := os.Create("d.out")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// get T
	var T int
	fmt.Scanf("%d\n", &T)

	// iterate over problems
	for t := 0; t < T; t++ {

		// get
	}
}
