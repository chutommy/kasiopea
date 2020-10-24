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

	for t := 0; t < T; t++ {

		// get N
		var M, N int
		fmt.Scanf("%d %d", &M, &N)

		// declare solution
		sol := make([]bool, N)

		// retrieve data
		for m := 0; m < M; m++ {
			for n := 0; n < N; n++ {

				// get atom
				var a int
				fmt.Scan(&a)

				// revert if purple
				if a == 1 {
					sol[n] = !sol[n]
				}
			}
		}

		// create result
		res := make([]string, N)
		for i, v := range sol {
			if v {
				res[i] = "1"
			} else {
				res[i] = "0"
			}
		}

		// store output
		fmt.Fprintln(f, strings.Join(res, " "))
	}
}
