package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Solver struct {
	N       int
	nums    []int
	target  int
	results []string
}

func (S Solver) printResults() {
	fmt.Println("\nResult")
	for i, r := range S.results {
		fmt.Printf("%dth: %s\n", i, r)
	}
}

func (S *Solver) storeNum() {
	var sc = bufio.NewScanner(os.Stdin)
	for i := 0; i < S.N; i++ {
		if sc.Scan() {
			t := sc.Text()
			n, _ := strconv.Atoi(t)
			S.nums = append(S.nums, n)
		}
	}
}

func main() {
	S := Solver{
		N:      4,
		target: 10,
	}
	S.storeNum()
}
