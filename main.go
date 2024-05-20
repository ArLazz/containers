package main

import (
	"fmt"
)

func main() {

	var n int
	fmt.Scan(&n)
	balls := make([][]int, n)
	for i := range balls {
		line := make([]int, n)
		for j := range line {
			fmt.Scan(&line[j])
		}
		balls[i] = line
	}

	checkContainers(balls, n)
}

func checkContainers(balls [][]int, n int) {
	for i := 0; i < n; i++ {
		rowSum := 0
		columnSum := 0
		for j := 0; j < n; j++ {
			rowSum += balls[i][j] 
			columnSum += balls[j][i]
		}
		if rowSum != columnSum {
			fmt.Println("no")
			return
		}
	}	
	fmt.Println("yes")
}
