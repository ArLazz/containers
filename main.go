package main

import (
	"fmt"
)

func main() {

	var SizeOfMatrix int
	fmt.Scan(&SizeOfMatrix)
	ballsInContainers := make([][]int, SizeOfMatrix)
	for row := range ballsInContainers {
		line := make([]int, SizeOfMatrix)
		for j := range line {
			fmt.Scan(&line[j])
		}
		ballsInContainers[row] = line
	}

	checkContainers(ballsInContainers, SizeOfMatrix)
}

func checkContainers(ballsInContainers [][]int, SizeOfMatrix int) {
	for i := 0; i < SizeOfMatrix; i++ {
		rowSum := 0
		columnSum := 0
		for j := 0; j < SizeOfMatrix; j++ {
			rowSum += ballsInContainers[i][j] 
			columnSum += ballsInContainers[j][i]
		}
		if rowSum != columnSum {
			fmt.Println("no")
			return
		}
	}	
	
	fmt.Println("yes")
}
