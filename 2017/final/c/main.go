package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	// create a file
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

		// get first row
		var l1 int
		fmt.Scan(&l1)
		fr := make([]byte, l1)
		for i := 0; i < l1; i++ {
			var ch byte
			fmt.Scanf("%c", &ch)
			fr[i] = ch
		}

		solution := solve(fr, N)
		fmt.Fprintln(f, solution)
		fmt.Printf("%d/%d\n", t+1, T)
	}
}

func solve(fr []byte, N int) int {

	min := -1

	for n := 0; n < N; n++ {

		// get lenght
		var L int
		fmt.Scan(&L)

		q := []byte{}

		// range over the queue (fr)
		for l := 0; l < len(fr); l++ {

			q = append(q, fr[l])
			lq := len(q)

			if lq > 1 {
				if q[lq-1] == q[lq-2] {
					q = q[:lq-2]
				}
			}
		}

		// range over the queue
		for l := 0; l < L; l++ {

			// get char
			var ch byte
			fmt.Scanf("%c", &ch)

			q = append(q, ch)
			lq := len(q)

			if lq > 1 {
				if q[lq-1] == q[lq-2] {
					q = q[:lq-2]
				}
			}
		}

		lq := len(q)
		if lq < min || min == -1 {
			min = lq
		}
	}

	return min
}
