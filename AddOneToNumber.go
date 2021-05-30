/*
	https://www.interviewbit.com/problems/add-one-to-number/
*/

package main

import "fmt"

func main() {
	a := []int{9, 9, 9} //expected output [1, 0, 0, 0]
	fmt.Println(plusOne(a))
}

func plusOne(a []int) []int {

	remainder := 0
	dividend := 0
	i := len(a) - 1
	for i >= 0 && remainder == 0 {
		remainder = (a[i] + 1) % 10
		dividend = (a[i] + 1) / 10
		a[i] = remainder
		i--
	}

	if dividend != 0 {
		return append([]int{dividend}, a...)
	} else {
		i := 0
		for i < len(a) {
			if a[i] != 0 {
				break
			}
			i++
		}
		return a[i:]
	}

}
