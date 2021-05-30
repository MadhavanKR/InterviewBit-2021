/*
https://www.interviewbit.com/problems/largest-number/
 */
package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	a := []int{3, 30, 34, 5, 9} //expected output: 9534330
	fmt.Println(largestNumber(a))
}

func largestNumber(a []int) string {
	str := make([]string, 0)
	zeroCount := 0
	for _, num := range a {
		if num == 0 {
			zeroCount++
		}
		str = append(str, strconv.Itoa(num))
	}
	if zeroCount == len(a) {
		return "0"
	}
	sort.Slice(str, func(i, j int) bool {
		ab := str[i]+str[j]
		ba := str[j]+str[i]
		if ab>=ba {
			return true
		} else {
			return false
		}
	})
	result := ""
	for _, curStr := range str {
		result+=curStr
	}
	return result
}


