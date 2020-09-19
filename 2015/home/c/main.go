package main

import (
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {

	// create output file
	f, err := os.Create("c.out")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// get T
	var T int
	fmt.Scan(&T)

	for t := 0; t < T; t++ {

		// get N
		var N int
		fmt.Scan(&N)

		baguette := make(map[int]int)

		for n := 0; n < N; n++ {

			// get layer info
			var s, e int
			fmt.Scanf("%d %d", &s, &e)

			baguette[s]++
			baguette[e+1]--
		}

		// solve
		s := solve(baguette)
		fmt.Fprintln(f, s)
	}
}

func solve(baguette map[int]int) int {

	// format into slice
	bgt := make([][]int, 0, len(baguette))
	for k, v := range baguette {
		bgt = append(bgt, []int{k, v})
	}

	// sort
	sort.SliceStable(bgt, func(i, j int) bool {
		return bgt[i][0] < bgt[j][0]
	})

	// calculate max
	var max, total int
	for i := 0; i < len(bgt); i++ {
		total += bgt[i][1]
		if total > max {
			max = total
		}
	}

	return max
}
