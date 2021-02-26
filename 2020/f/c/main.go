package main

import (
	"fmt"
	"log"
	"os"
)

type ured struct {
	key   int
	tasks int
	proc  int
	emp   []int
}

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
		var N, W int
		fmt.Scanf("%d %d\n", &N, &W)

		urednici = make([]*ured, N+1)
		for n := 0; n <= N; n++ {
			urednici[n] = &ured{}
			urednici[n].emp = []int{}
		}

		urednici[0].key = 0
		urednici[0].tasks = 0
		urednici[0].proc = W

		for n := 1; n <= N; n++ {
			var k, proc, tasks int
			fmt.Scanf("%d %d %d", &k, &proc, &tasks)

			urednici[n].key = n
			urednici[n].proc = proc
			urednici[n].tasks = tasks
			urednici[k].emp = append(urednici[k].emp, n)
		}

		s := solve(N, W)
		fmt.Fprintln(f, s)
	}
}

var need []int
var urednici []*ured

func solve(N, W int) int {
	need = make([]int, N+1)
	dfs(urednici[0], W)

	min := need[0]
	for _, v := range need {
		if v < min {
			min = v
		}
	}

	return min
}

func dfs(u *ured, treba int) {
	jeste := u.proc - u.tasks
	treba -= u.tasks

	podstr := 0
	if treba > jeste {
		podstr = treba
	} else {
		podstr = jeste
	}

	need[u.key] = podstr
	for _, pod := range u.emp {
		dfs(urednici[pod], podstr)
	}
}
