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

		// get N, K, P
		var N, K, P int
		fmt.Scanf("%d %d %d", &N, &K, &P)

		inpN := make([]int, N)
		for n := 0; n < N; n++ {
			var i int
			fmt.Scan(&i)
			inpN[n] = i
		}

		inpK := make([]int, K)
		for k := 0; k < K; k++ {
			var i int
			fmt.Scan(&i)
			inpK[k] = i
		}

		inpP := make([]int, P)
		for p := 0; p < P; p++ {
			var i int
			fmt.Scan(&i)
			inpP[p] = i
		}

		// solve
		s := solve(N, K, P, inpN, inpK, inpP)
		ls := len(s)

		// format solution
		res := make([]string, ls)
		for i := 0; i < ls; i++ {
			res[i] = fmt.Sprint(s[i])
		}

		// store solution
		fmt.Fprintln(f, strings.Join(res, " "))
	}
}

func solve(N, K, P int, inpN, inpK, inpP []int) []int {

	// get f,s
	var first, second int
	for k := 0; k < K; k++ {

		if inpK[k] == 1 {
			first++
		}
	}
	second = K - first

	f := (first % 2) == 1
	s := (second % 2) == 1

	// make the moves
	if f {
		temp := make([]int, N)
		for i := 0; i < N; i++ {
			temp[i] = inpN[N-i-1]
		}
		copy(inpN, temp)
	}
	if s {
		inpN[0], inpN[N-1] = inpN[N-1], inpN[0]
	}

	// read favorites
	fav := make([]int, P)
	for p := 0; p < P; p++ {
		fav[p] = inpN[inpP[p]-1]
	}

	return fav
}
