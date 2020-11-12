package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	// create an output file
	f, err := os.Create("c.out")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// get T
	var T int
	fmt.Scan(&T)

	// range over each problem
	for t := 0; t < T; t++ {

		// get N
		var N int
		fmt.Scan(&N)

		// store heights
		hs := make([]int, N)
		for n := 0; n < N; n++ {
			var h int
			fmt.Scan(&h)
			hs[n] = h
		}

		// solve
		s := solve(hs)
		fmt.Fprintln(f, s)
	}
}

func solve(hs []int) int {
	l := len(hs)

	// from left to right
	ltr := make([]int, l)
	// set first element
	ltr[0] = hs[0]

	// query from left to right
	for i := 1; i < l; i++ {

		// increment
		if v := hs[i]; v > ltr[i-1] {
			ltr[i] = v
			continue
		}

		ltr[i] = ltr[i-1] + 1
	}

	// from right to left
	rtl := make([]int, l)
	// set last
	rtl[l-1] = hs[l-1]

	// query from right to left
	for i := l - 2; i >= 0; i-- {

		// increment
		if v := hs[i]; v > rtl[i+1] {
			rtl[i] = v
			continue
		}

		rtl[i] = rtl[i+1] + 1
	}

	// query results
	var max int
	for i := 0; i < l; i++ {

		// get smaller
		n := ltr[i]
		if ltr[i] > rtl[i] {
			n = rtl[i]
		}

		// compare to max
		if n > max {
			max = n
		}
	}

	return max
}
