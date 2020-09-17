package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	// create output file
	f, err := os.Create("f.out")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// get T
	var T int
	fmt.Scan(&T)

	// range over problems
	for t := 0; t < T; t++ {

		// get N, K
		var N int
		var K int64
		fmt.Scanf("%d %d", &N, &K)

		// get bilance
		bilance := make([]int64, N)
		for n := 0; n < N; n++ {

			// get int
			var i int64
			fmt.Scan(&i)

			// store
			bilance[n] = i
		}

		// store solution
		s := solve(bilance, N, K)
		fmt.Fprintln(f, s)
	}
}

func solve(bilance []int64, N int, K int64) int64 {

	var count int64
	a := 0

	for a < N {

		var sum int64

		// range from a to b
		for i := a; i < N; i++ {

			// add num
			sum += bilance[i]
			// check
			if sum < K {
				count++
			}
		}

		// move up
		sum -= bilance[a]
		a++
		if sum < K {
			count++
		}

		if a >= N {
			break
		}

		// range from b to a
		for i := N - 1; i > a; i-- {

			// add num
			sum -= bilance[i]
			// check
			if sum < K {
				count++
			}
		}
		a++
	}

	return count
}
