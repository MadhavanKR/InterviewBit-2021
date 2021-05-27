package main

import (
	"fmt"
	"reflect"
)

func main() {
	set := make(map[int]bool)
	set[5] = true
	set[1] = true
	set[2] = true
	fmt.Println(reflect.ValueOf(set).MapKeys()[0].Int())
}

func maximumSumTriplet(A []int )  (int) {
	n:=len(A)

	//first we calculate right max suffix array
	rightMax := make([]int, n+1)
	for i:=n-1; i>0; i-- {
		if rightMax[i+1] > A[i] {
			rightMax[i] = rightMax[i+1]
		} else {
			rightMax[i] = A[i]
		}
	}
	var sum int = 0
	set := make(map[int]bool)
	_ = set
	_ = sum
	//now fixing the jth element we calculate the triplet sum
	for j:=0; j<n-1;j++ {

	}
	return 0
}

func getLowerBound(set map[int]bool, bound int) int {
	keys := reflect.ValueOf(set).MapKeys()
	var low int = 0
	var high int = len(keys)
	var mid int
	for ;low<high; {
		mid = (high - low)/2
		if int(keys[mid].Int()) < bound {
			return int(keys[mid].Int())
		}
	}
	return -1
}

