package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	// open file
	f, err := os.Create("out.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// get T
	var T int
	fmt.Scanln(&T)

	// range over problem sets
	for t := 0; t < T; t++ {

		// get N
		var N int
		fmt.Scanln(&N)

		addrs := make([]int, N)
		// range over N
		for n := 0; n < N; n++ {

			// get record
			var r int
			fmt.Scanf("%d", &r)
			addrs[n] = r
		}

		fmt.Fprintln(f, solve(addrs))
	}
}

// solve solves the problem and return the index of the message that can be lost
func solve(addrs []int) int {
	l := len(addrs)

	// base
	if l == 3 {
		return 1
	}

	var skipping int
	var best int
	// find the most worth address to skip
	for i := 1; i < l-1; i++ {

		// calculate how many steps would he save
		// if he skip this one
		{
			// now
			now := abs(addrs[i-1] - addrs[i])
			now += abs(addrs[i+1] - addrs[i])

			// if skip
			after := abs(addrs[i-1] - addrs[i+1])

			// calculate
			spare := abs(now - after)
			if spare >= best {
				best = spare
				skipping = i
			}
		}
	}

	return skipping
}

// abs returns the positive value of the integer
func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
