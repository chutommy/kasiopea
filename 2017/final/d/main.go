package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Coor struct {
	x int
	y int
}

type Car struct {
	A, B Coor
}

type State struct {
	car     Car
	history []int
}

func main() {
	// create output file
	f, err := os.Create("d.out")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// load input
	var T int
	fmt.Scanf("%d\n", &T)

	for t := 0; t < T; t++ {
		fmt.Scanf("%d %d\n", &N, &M)

		// initialize grid
		grid = make([][]int, N)
		for n := 0; n < N; n++ {
			grid[n] = make([]int, M)
		}

		// load grid
		for n := 0; n < N; n++ {
			for m := 0; m < M; m++ {
				var c string
				fmt.Scanf("%1s", &c)

				switch c {
				case "#":
					grid[n][m] = 1
					continue
				case "A":
					start.A = Coor{m, n}
					continue
				case "B":
					start.B = Coor{m, n}
					continue
				}

				if n == 0 || n == N-1 || m == 0 || m == M-1 {
					exit = Coor{m, n}
				}
			}
			fmt.Scanf("\n")
		}

		visited = make(map[Car]struct{})

		s := solve()

		// sol(s)

		var sb strings.Builder
		for _, c := range s {
			switch c {
			case 0:
				sb.WriteRune('^')
			case 1:
				sb.WriteRune('v')
			case 2:
				sb.WriteRune('<')
			case 3:
				sb.WriteRune('>')
			}
		}

		fmt.Fprintln(f, sb.String())
	}
}

func sol(s []int) {
	grid[start.A.y][start.A.x] = 44
	grid[start.B.y][start.B.x] = 88

	car := start

	for _, d := range s {

		car = move(d, car)

		grid[car.A.y][car.A.x] = 4
		grid[car.B.y][car.B.x] = 8
	}
	fmt.Println()
	fmt.Println("======================")
	fmt.Println()
	for _, row := range grid {
		fmt.Println(row)
	}
}

var N, M int
var start Car
var exit Coor
var grid [][]int
var visited map[Car]struct{}

func solve() []int {
	q := []State{{start, []int{}}}

	for {
		this := q[0]
		q = q[1:]

		if this.car.A == exit && this.car.B == exit {
			return this.history
		}

		nq := unvisitedAdj(this)
		q = append(q, nq...)
	}
}

func unvisitedAdj(c State) []State {
	var cc []State
	l := len(c.history)

	for d := 0; d < 4; d++ {
		adjC := move(d, c.car)
		if _, ok := visited[adjC]; !ok {
			h := make([]int, l+1)
			copy(h, c.history)
			h[l] = d

			s := State{adjC, h}
			cc = append(cc, s)
			visited[adjC] = struct{}{}
		}
	}

	return cc
}

func move(d int, c Car) Car {
	after := c

	switch d {
	case 0:
		// up
		after.A.y--
		after.B.y--
	case 1:
		// down
		after.A.y++
		after.B.y++
	case 2:
		// left
		after.A.x--
		after.B.x--
	case 3:
		// right
		after.A.x++
		after.B.x++
	}

	if c.A == exit || grid[after.A.y][after.A.x] == 1 {
		after.A = c.A
	}
	if c.B == exit || grid[after.B.y][after.B.x] == 1 {
		after.B = c.B
	}

	return after
}
