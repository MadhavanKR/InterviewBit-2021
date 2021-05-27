// https://www.geeksforgeeks.org/array-rotation/

package main

import (
	"fmt"
	"math"
)

func main() {
	input1 := []int{1, 2, 3, 4, 5, 6}
	d := 1
	fmt.Printf("input: \n array: %v \n d: %d \n", input1, d)
	fmt.Println("rotateUsingTempArray result: ", rotateUsingTempArray(input1, d))
	input2 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("rotateUsingTempArrayV2 result: ", rotateUsingTempArrayV2(input2, d))
	input3 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("rotateUsingOneByOneRotation result: ", rotateUsingOneByOneRotation(input3[:], d))
	input4 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("rotateUsingJugglingAlgorithm result: ", rotateUsingJugglingAlgorithm(input4[:], d))
}

// time complexity - O(n) and space complexity O(n) where n is array size
func rotateUsingTempArray(a []int, d int) []int {
	var tempArray []int
	d = int(math.Mod(float64(d), float64(len(a))))
	for i := d; i < len(a); i++ {
		tempArray = append(tempArray, a[i])
	}
	for i := 0; i < d; i++ {
		tempArray = append(tempArray, a[i])
	}

	return tempArray
}

// time complexity - O(n) and space complexity O(d) where n is array size and d is rotation number
func rotateUsingTempArrayV2(a []int, d int) []int {
	var tempArray []int
	d = int(math.Mod(float64(d), float64(len(a))))
	//copy first d elements into temp array
	for i := 0; i < d; i++ {
		tempArray = append(tempArray, a[i])
	}
	//shift left original array by len(a) - d positions
	for i := 0; i < len(a)-d; i++ {
		a[i] = a[i+d]
	}
	//copy tempArray to last d indices of original array
	tempIndex := 0
	for i := len(a) - d; i < len(a); i++ {
		a[i] = tempArray[tempIndex]
		tempIndex++
	}
	return a
}

func rotateUsingOneByOneRotation(a []int, d int) []int {
	for i := 0; i < d; i++ {
		for j := 0; j < len(a)-1; j++ {
			temp := a[j]
			a[j] = a[j+1]
			a[j+1] = temp
		}
	}
	return a
}

func rotateUsingJugglingAlgorithm(a []int, d int) []int {
	d = d % len(a)
	gcd := gcd(len(a), d)
	numSets := int(len(a) / gcd)
	for setElemIndex := 0; setElemIndex < gcd; setElemIndex++ {
		for setNum := 0; setNum < numSets-1; setNum++ {
			temp := a[setElemIndex+setNum*gcd]
			a[setElemIndex+setNum*gcd] = a[setElemIndex+(setNum+1)*gcd]
			a[setElemIndex+(setNum+1)*gcd] = temp
		}
	}
	return a
}

func gcd(a int, b int) int {
	gcd := 1
	if a == 0 {
		return b
	} else if b == 0 {
		return a
	}
	for i := 1; i <= a && i <= b; i++ {
		if a%i == 0 && b%i == 0 {
			gcd = i
		}
	}
	return gcd
}
