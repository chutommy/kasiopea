package main

import "fmt"

func main() {

	// get the number of problems
	var T int
	fmt.Scanln(&T)

	// range over problems
	for t := 0; t < T; t++ {

		// get the number of records
		var N int
		fmt.Scanln(&N)

		result := true

		var previous int
		// range over the records
		for n := 0; n < N; n++ {

			// get the n-th record
			var record int
			fmt.Scan(&record)

			// first time defines the previous
			if n == 0 {
				previous = 0
			} else if previous+3 < record {
				result = false
				continue
			}
			previous = record
		}

		// print the result
		if result == true {
			fmt.Println("ano")
		} else {
			fmt.Println("ne")
		}
	}
}
