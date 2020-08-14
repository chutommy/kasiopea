package main

import (
	"fmt"
	"log"
	"os"
	"sync"
)

func main() {

	// create output file
	f, err := os.Create("out.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// get T
	var T int
	fmt.Scan(&T)

	// enable goroutines management
	var manager []chan struct{}
	for i := 0; i < T; i++ {
		manager = append(manager, make(chan struct{}))
	}
	var wg sync.WaitGroup

	// range over problems
	for t := 0; t < T; t++ {

		// get N, M
		var N, M int
		fmt.Scanf("%d %d\n", &N, &M)

		// get votes
		votes := make([]int, N)
		for n := 0; n < N; n++ {

			// save record
			var r int
			fmt.Scan(&r)
			// insert
			votes[n] = r
		}

		go func(t int) {

			var results []int
			// solve
			results = solve(votes, N, M)

			// block
			manager[t] <- struct{}{}

			// print the solution
			for i, r := range results {
				if i == 0 {
					fmt.Fprintf(f, "%d", r)
				} else {
					fmt.Fprintf(f, " %d", r)
				}
			}
			fmt.Fprintf(f, "\n")
			log.Printf("Problem %d solved.\n", t)
			wg.Done()
		}(t)
	}

	// print to output
	for i := 0; i < T; i++ {
		_ = <-manager[i]
		wg.Add(1)
		wg.Wait()
		// time.Sleep(1 * time.Second)
	}
}

// solve solves one problem set
func solve(votes []int, n int, m int) []int {

	// prepare fields
	results := make([]int, n)

	usage := make([]float64, n)
	// fill usage with 1
	for i := 0; i < n; i++ {
		usage[i] = 1.0
	}

	status := make([]float64, n)
	// fill statuses
	for i := 0; i < n; i++ {
		status[i] = float64(votes[i])
	}

	// find the best
	for ; m > 0; m-- {

		// find highest record
		next := 0
		highest := status[0]
		for i := 1; i < n; i++ {
			if curr := status[i]; curr > highest {
				next = i
				highest = curr
			}
		}

		// add point
		results[next]++

		// update status
		status[next] *= usage[next]
		usage[next]++
		status[next] /= usage[next]
	}

	return results
}
