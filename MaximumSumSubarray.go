/*
	https://www.interviewbit.com/problems/max-sum-contiguous-subarray/
*/

package main

import "fmt"

func main() {
	a := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println("maximum sum: ", maxSubArray(a))
}

func maxSubArray(a []int) int {
	/* n := len(a)
	 maxSumArray := make([]int, n)
	 maxSumArray[0] = a[0]
	 maxSum := a[0]
	 for i:=1; i<n ; i++ {
		 if a[i] > maxSumArray[i-1] + a[i] {
			 maxSumArray[i] = a[i]
		 } else {
			 maxSumArray[i] = maxSumArray[i-1] + a[i]
		 }
		 if maxSumArray[i] > maxSum {
			 maxSum = maxSumArray[i]
		 }
	 }
	 return maxSum */
	return maxSubArraySpaceComp1(a)
}

func maxSubArraySpaceComp1(a []int) int {
	n := len(a)
	var maxSumTillNow int = a[0]
	maxSum := a[0]
	for i := 1; i < n; i++ {
		if a[i] > maxSumTillNow+a[i] {
			maxSumTillNow = a[i]
		} else {
			maxSumTillNow = maxSumTillNow + a[i]
		}
		if maxSumTillNow > maxSum {
			maxSum = maxSumTillNow
		}
	}
	return maxSum
}
