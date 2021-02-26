package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// create output file
	f, err := os.Create("out")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var T int
	fmt.Scanf("%d\n", &T)

	for t := 0; t < T; t++ {
		var N int
		fmt.Scanf("%d\n", &N)

		delky := make([]int, N)
		for n := 0; n < N; n++ {
			var i int
			fmt.Scanf("%d", &i)

			delky[n] = i
		}

		solution := solve(N, delky)
		fmt.Fprintln(f, solution)
	}
}

func solve(N int, delky []int) int {
	m := make(map[int]int)
	for _, d := range delky {
		m[d]++
	}

	max := m[delky[0]]
	for _, v := range m {
		if v > max {
			max = v
		}
	}
	return max
}
