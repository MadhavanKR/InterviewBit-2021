package main

import "fmt"

func main() {
	//medium input
	/*
	a := [][]int{
		{1, 1, 1, 1, 1},
		{2, 2, 2, 2, 2},
		{3, 8, 6, 7, 3},
		{4, 4, 4, 4, 4},
		{5, 5, 5, 5, 5},
	}
	b:=3
	*/
	//small input
	a := [][]int {
		{2,2},
		{2,2},
	}
	b:=2
	fmt.Println(maxSumSubMatrix(a, b))
}

func maxSumSubMatrix(a [][]int, B int) int {

	n := len(a)
	maxSum := 0
	for i := 0; i <= n-B; i++ {
		for j := 0; j <= n-B; j++ {
			curSum := calculate(i, j, i+B, j+B, a)
			if curSum > maxSum {
				//fmt.Printf("submatrix between (%d, %d) and (%d, %d) has sum of %d\n", i, j, i+B-1, j+B-1, curSum)
				maxSum = curSum
			}
		}
	}
	return maxSum
}

func calculate(startX, startY, endX, endY int, a [][]int) int {
	sum := 0
	for i := startX; i < endX; i++ {
		for j := startY; j < endY; j++ {
			sum += a[i][j]
		}
	}
	return sum
}
