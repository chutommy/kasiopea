package main

import (
	"fmt"
	"os"
	"sort"
	"sync"
)

func main() {

	// create an output file
	f, err := os.Create("d.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// get T
	var T int
	fmt.Scan(&T)

	// set for concurrency
	output := make([]int, T)
	var wg sync.WaitGroup
	wg.Add(T)

	// range over each problem
	for t := 0; t < T; t++ {

		// get N
		var N int
		fmt.Scan(&N)

		// get order
		order := make([]int, N)
		// store input
		for n := 0; n < N; n++ {
			var i int
			fmt.Scan(&i)
			order[n] = i
		}

		// solve
		go func(t int) {
			s := solve(order)
			output[t] = s
			fmt.Printf("%d/%d done\n", t+1, T)
			wg.Done()
		}(t)
	}

	wg.Wait()
	for _, s := range output {
		fmt.Fprintln(f, s)
	}
}

func solve(order []int) int {

	// init connections
	var conns []int

	// range over each element
	for _, elem := range order {

		// calculate position of the elem
		pos := sort.SearchInts(conns, elem)

		if pos == 0 {
			conns = append([]int{elem}, conns...)

		} else if ls := len(conns); pos == ls {
			conns[ls-1] = elem

		} else {
			conns[pos-1] = elem
		}
	}

	return len(conns)
}
