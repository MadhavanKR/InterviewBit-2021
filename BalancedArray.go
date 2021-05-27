/*
https://www.interviewbit.com/problems/balance-array/
*/
package main

import "fmt"

func main() {
	//small positive input
	a := []int{2, 1, 6, 4}
	fmt.Println("special elements: ", balanacedArray(a))
}

func balancedArray(a []int) int {
	n := len(a)
	preEven := make([]int, n)
	preOdd := make([]int, n)
	preEven[0] = a[0]
	preOdd[0] = 0
	for i := 1; i < n; i++ {
		if i%2 == 0 {
			preEven[i] = preEven[i-1] + a[i]
			preOdd[i] = preOdd[i-1]
		} else {
			preOdd[i] = preOdd[i-1] + a[i]
			preEven[i] = preEven[i-1]
		}
	}
	sufEven := make([]int, n+1)
	sufOdd := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		if i%2 == 0 {
			sufEven[i] = sufEven[i+1] + a[i]
			sufOdd[i] = sufOdd[i+1]
		} else {
			sufEven[i] = sufEven[i+1]
			sufOdd[i] = sufOdd[i+1] + a[i]
		}
	}

	count := 0
	for i := 0; i < n; i++ {
		/*
		   evenSum := even[i]
		   oddSum := odd[i]
		   if i%2==0 {
		       evenSum -= a[i]
		   } else {
		       oddSum -= a[i]
		   }
		   for j:=i+1; j<n; j++ {
		       if j%2==0 {
		           oddSum += a[j]
		       } else {
		           evenSum += a[j]
		       }
		   }
		*/
		evenSum := 0
		oddSum := 0
		if i%2 == 0 {
			evenSum = preEven[i] - a[i] + sufOdd[i+1]
			oddSum = preOdd[i] + sufEven[i+1]
		} else {
			evenSum = preEven[i] + sufOdd[i+1]
			oddSum = preOdd[i] - a[i] + sufEven[i+1]
		}
		if evenSum == oddSum {
			count++
		}
	}
	return count
}
