package main

import (
	"fmt"
	"log"
	"os"
)

type node struct {
	val         int
	left, right *node
}

type solution struct {
	i int
	s string
}

func main() {
	// create output file
	f, err := os.Create("d.out")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var T int
	fmt.Scanf("%d\n", &T)

	// implement channels
	solutions := make(chan solution)
	output := make([]string, T)

	for t := 0; t < T; t++ {
		// load problem input
		var N int
		fmt.Scanf("%d\n", &N)

		inp := make([]int, N)
		for n := 0; n < N; n++ {
			fmt.Scanf("%d", &inp[n])
		}

		fmt.Println("stuck")

		go func(t int) {
			sol := solve(N, inp)
			s := fmt.Sprint(sol)
			s = s[:len(s)-1]
			s = s[1:]

			solutions <- solution{
				i: t,
				s: s,
			}
		}(t)
	}

	for t := 0; t < T; t++ {
		s := <-solutions
		output[s.i] = s.s
	}

	// print solution
	for t := 0; t < T; t++ {
		fmt.Fprintln(f, output[t])
	}

}

func solve(N int, inp []int) []int {
	out := make([]int, N)
	root := &node{val: inp[0]}

	for n := 1; n < N; n++ {
		out[n] = insert(root, inp[n])
	}

	return out
}

func insert(root *node, n int) int {
	steps := 1
	r := root

	for {
		if n > r.val {
			if r.right == nil {
				r.right = &node{val: n}
				return steps
			}

			r = r.right
		} else {
			if r.left == nil {
				r.left = &node{val: n}
				return steps
			}

			r = r.left
		}

		steps++
	}
}
