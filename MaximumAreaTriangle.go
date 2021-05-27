package main

import (
	"fmt"
	"math"
)

func main() {
	a := []string{"bgrrrbbgrrbggbrrggrggrrbrgbgbrbggbgrrbbrggrgbgrbbggbrgggrbrbrbgbbbbbbr", "ggrgggrgbgrbrbgrrbgbbgbrrrrgbrrrrggrrrrbbggrrgggbgrrbrbgbbbgbrbbbbgrrg", "rrrbbbggbbbrbgbbrrgbbrrgrbgbbrgggrrbgrggggrgbbrbbggrggbggbrbbbgrbbbrrb", "rrrrbrrgrbrbbbgggrrrrbgrggrbrbrbgrrrgbrrbbbgbbbbgbrgbggbbbbgrgrbrrgbrr", "bbgrgggrrggbgbgbgbrrbgrbrggrggggggbggrbrrgrbrggggggbrbrgbgrrbrggrgrrbg", "rgrbrrgrbrbrrgbgrgbbbrrbbbrbgrrrrrggbrgrbrgrrrggbbrbbrgrgrbbggbbbbbbrg", "brrggrbbggbrrrrggrggrgrrbgbgrbbrgrbrrbgrrrbbrbrgrrgrbgrggrrbrgggrbrbgb", "rrgrgbbbbbbbggbggrrgbgggbrggbbgbbgrgggrgrrbrrggrbbrgbbrggrrgrbrrggbbgg", "rbgbgrbbgbgggbgbrbgggbbbgrbrgbrggbbrgbbrbbbrbgbgrbbgbrbgbgrgbrrrgrrrbg", "rbggbgbrggbgbgrrbbrbgbrbgrgggrrrrggbbbrgrrbgggrrbgrgbrgrrbrrbgbrrrbrgb", "rgrrrrgrbbrbrbggrrgbbbbgbbrbrgbgrggggrrbggrrrbgrgrgbbbgrgbbgbrbrbrbbrg", "rrgbgbggrbbbggrbbgbgrbbbgbgbgbrrbgrbbbgrrrbbgrrgggbbrbgrgbbbrbbgbbgbgr", "bbbggbbrbgggbrrrgrbgbrbrgrrrbbbgbgrgrggbbrbgbbbrrbrbgrggrbrbbgrgrgbrgb", "ggrrrrrgggrggbrrrrrrbbrggggbbgggbrggrbrrbbggrrrrbrbgrrbbrgrgrrrbrgrrbg", "gbrrgrrgbrggbrgrrrgbrrbgbbggbgrbbrgrbrrbbgbbgbgbbggrrggbbbgrrbggbbbggr", "ggbrrggrgbrbgggggrrrrbrbbrbgrggrrbbrgrrgbggbrgbgrbgbbbgrrrrrrgrbgggbrb", "grggrrgbggrggggrgbbggrrggbgrbrrrbbrrggbbgbrgbbgggbbbbbrrrbrrrgbggrrbrg", "rbgbbrgrbgrggbgbbrrrrrrbbrbbgrrbbbgggggrrgrgbbgbbbgrbbbbrbrgrbrbrrggrr", "gbrgbbbgbrbrbbbggrrgbbgrrbrrbrbrrbbrggrgggrbgrrggrggrbrbbbbrgbbbbbgrrg", "bgbggrrbgbrgbgbrrgrbgggrgrbbgrbrbrgrgbrrgbbgbrrbrbrrbbgbrrgrrrrbbrrbrg", "ggbbgrgrggbrgrrbrgbgbggrbrggbrrrbgbgbrbbgrgbbgrrrggrgbbrggrgrbggrrgrgr", "rrbbgrbrrbgrgrrbggrrgrrgrrgrbrggbrbbbbrgrrrbbrbgbbbrbrrbrrbrrggrgggrgb", "rbrrgrgrbbbgrrbggrrggrgggbrbrbrbgbrgrggrgggrrgggbrbrrbrrgbggrgbgrbrrrb", "rbbbrrgbrbrggrbgrgbbrrgrgbrgrgbrbrbrbgrbrbrgbrbgggbbbbrrgbgbrrrrgggrbb", "rrggbgbrrrrrrgrgbgggrbggbrrrbrbggrgrrbrrbrrgrbbrbrbgbbgggbggbrgggrrrbb", "rrbbbbrrbbbbrrgrbrbrggggggrgrrrbgrrrrbrrbbrrgbgbbrrbrbgbbbgrbgrgbggggb", "rbbbbgrrbbgrbgbggrggrggrbbrbrgbgrgggrrggbgrbrbrgggrrrbgrgrgrrbbbrgrbrb", "bbbbrbbrrbrbbgbrbgrgrgggbbgrrgggggrgrgrrrgggggrrrbrbgbgrbrbrbggbrbgbrg", "rrgrrbrrbbgbrgbrggrrbggbbbbggrrrrbbgrrgbrbrrbgggbrrgrrrgbrrgggbggrbrgg", "grggbgrggbgrrrgrggggrbrgrgbbggrggbrbggrggbggbbrggrbbrrrbrrbbrrrgbgrrrg", "bbbbrgbrggbgbrggbgbgrbbrgbbggbgbrgrgggggrbgbbggrggrbggggrbggrbrbrbbbgb", "grrbrgbbbgrrggrgbrggbrrrbrgbgbbggggrrbrbbrrbbbgbgbgbbrbgrbrbbrggbggbbb", "gbgrgrbbrbrbrrgbrbbbrbgrgbrbgrrrgrrgbrggrbbrrgbbbggggrbggrbbgbgrrgrgrg", "grrbrbgbgrrgrgrggrrgrbrggbbrbbgrbggggbggrggrbbrgbbgrgrbrggrrrgrbbrbrrg", "ggggrgbbbbbggbrbrgbrrbrgggggbbgrrgrbrgggbrgbbbgrrgrggbgbrrggbgbrgbrrrb", "gbgrbgrbbrrgbrgggrrbgrgggrgrbrbbbbrrrbrbgrgbbbrbbbbrrgbgrgbgbbrgbrrgbg", "bgrbggrrggbgggrgrbbgbrbrrgbgggggggbbbgbbrbrrrgrbrgbbbrrgrbggrbbggrrbrg", "bgrrrgbgrbrggbrrbgrggrgrbbbbbrbbbrgrbrrbrgbbgggrrgrbrrrbbrbbbbgbbbrrrg", "brgggbgrgrgrrgbrrbrggggrrbrrbbrgbbgrggbggbrbggrbgrgrbbgrrrgrgrbgbgrggr", "rrbrgrrrrrbggrbgrbrrrbbggrbgbgbrgbgrggbggrrrbggrgbbbrrbbrbbrrgrrbrgbrb", "rggggrbrgbrrgrgrrbbgbbgrggrgrbbbgrgbrbbrbbgrrgrbrbggbgbbgggrbgbrbbgrrg", "gbbbrgggbrrrbbgbgbbgrrgggbrrgrrggrrgggrggrgbbrrrrrrgrrrggbbbrbrgrgbgbg", "grrbbgbbbbgbbgbrbgbbrgrgbrrrbrrrbrbggbrbgrrrrrrrbbrbgrrbrggrbbgggbrrrb", "brbbggbgrgbrgrggrrgrrrrrrrrbgrbrrrgrbrrbbbgrbggrgbrrggrbbbbrgggbrbrbrg", "rbgbrrrgrbbgbgrbbgbbbbrgbrbrbbrgrbbrbrbgbgggrbrbgbbgrgrgbbgrrgbbbgrggb", "rrgbggrgrrgbggrggbgbrrgggrrggrbrgggggbbbrgbbbrgbrrggrggggrbbrgbgggbbrb", "bbrrbbbrbbrgbrbrrggrbrrrrbbbggrgrbbrbggrggggbbbrbbbrbrbggrbggrbbbggbrb", "gbbgbbbbggbrggbbgbrggbbbggbgbggggbrbbrrgrgbrrgbgbbggbggrrrrbbggrgbbrrg", "gbrrbrbggrrrbbrbrbgbrrrrbgrgggrbbgbbrgrbrrggrgbgrgggbrbbbgbbbgrbgbbbrr", "grrggbrrgrbgrrgrgrbrgrgrgbbrbbgrrbrgggggggbrggbrgbbbrggggggrbgrrgrrbgr", "brbrrgrgrbggbbrggbgrgggrgbgbbggrbrrbbrrrgrggbbrbrrggbgrbrrbrgbbrggrrrr", "rrgrbbbbbggbggrrggggrggrgbbbbbggrbrbrrrrrrrbgbrrgbrbgbrrggbbrbrrgrgrrr", "brgrgbbbbrrbrgrbbbbrbbrbbrgbrbgggbbrbgbrbgggbrbgggrrbrrggrrrgbrgggbggr", "brgggbbbrbbrrbrggbrbrggbgrbbgbbgbbrrrrbrrgbrbggrrbbggrrbgbggbgrbrgbgbg", "gbrrgbbgrgrgbbggrrbrbbgbgbggggbggrrgbgrgbbgrgbggbbrbggggrbrrrbrbgrgrgr", "rggrgbgbgggbbrrbgbrggbrgbrgrgrgrbbbggrbrbggrrbggggrbrbrrbbrbbbbrbrgrbg", "bggbbbrggrgbgggrrrrgrgbbbgbrggggbrrrrbrbgbrbgbggbbgbgrgbrgrbbbbrgbrrrr", "gbrrbggrbgbbbbbgrrbbggbbbrbgrbgrrrrbgbgbbggrgbbgbbrrggbbgrgrrggbrrbrgb", "rrbggbgbbgrbggbrggrrrbgrgrgbgrgggbrbrbrbgrrggrrgrrggbrgbrgrbbbbrbrbrgg", "gbbbbgbrgbbrrgrgggbbggbrbrrgrgrrgrbbgggrrbgrggrgbbrgrggbgrbbbgbrrrbbrg", "grbgbbrrrgbbrgrgrbgbbbgrrbrbbgbbbbbgggbbbrrggrgrbbbgbbgrrbgbbrrgrrbbgg", "bgrggggggbbrgbrrrrbgrrbrgbbgrgrgrrrrgggrbbrrrgrbgrrrgbggrrgbggrbbbrrgr", "rggbrggggbgbrbrggrrgbrbbbbrbggbbbrbbgggrrbrrbrrbrrgbgbggbggrgggrgbrbgg", "brrbrgggbrgbrrrggbrbrrbrgggrrgrrbgbrrrrrrbggrgrgrrrrbgrggrbbrgbrgbgbrb", "grggbbbgbrrrrggrrbgrrgrgbbrbgrggggrbggrbgggrrrrrbggrbbrbgggbggrrrgrrbg", "rrgrgrggbbbbgbbgrrggrrgrrgrgbbgrbgggbrgrbgrgggggrbbggbrbbgbbrbbgrrrggr", "ggbgbbgbbbbgrbbrggbbggrggggbbrrbrgbrbbbgbrbbbbrgrbbrrgrrbbbgrrrrggrgrr", "bgrgggbbggrbgrrgrrgbbbgbrrbrrrgrgbrgrgrrrgbgbrbrgrbrgbgggrrggrrbbgrgbb", "rgbgggbgrrrgbrbbbgbrgbbgbgbggrggrgbrgbbbggbrbrrgrgbbgggbgrrrgbrbrgggbb", "gbggbgbrrrbgbrrggrbbbbrbgbgbgrggbrrrbbrbggrrbrbrbbrgbbgrbrrbrbbgbrrrrb", "ggrgbrbbrbrbbbrbgbrbbggbbggrrrrrbbggbrgbgrgggbggbgbgbgbrrbrbbgbbbrgrbb", "bbbbrgrrbrrgbggbrrggbgrrrrgrrbggbbbrgbbgrrbgrrrgbbrgrbrbgbbbbbrrbbbggr", "rggggbgrrbrgrbggbbbgbbbbbrbbgrgbggbrrgbgrrgrbggrbggbbgrbbbbbbbbrgrgbrb", "rbrbrbbgrgbgggrbgrrbggrgbggrrrrbbbrbrbrrgbbbbggbrbbbgrbbggbgbgbggbbrrr", "rrbrbbbgggbggrrbgrgbrrbrggggrrrbrrggbrgrgbrgbgrbbbbrgbgggbrrgggrrgbgbg", "rrgrgbrbbbbbbggrgrrgrbbrggbgbbrrbbbbbgrgrgrrbrrrbrrgrbbbrggbrgbgrbgrrr", "gbgbgbbggrgggrbrrrbgrrbggbgbggbgbbgbbbbbbbrrrbgrbrbbgbggbrrgggrrbrbrbb", "bbrbbrggrbbgbrbbbbrbgbgrggbbrggbbgbbgbrrrrrbgrrbrgbbbrbbbgbrgrbrgbrggb", "gggrrggrbrgrbbgrrgrbrgbggbgbgrgbrrbggrgbgrbbrggbgbrgbrgrrbgrgrrrbggrbb", "bgrggrgbrrgbrbbbbbrgbbgggrrrbbrrrggrbbggrgrbgrrgrgrggbrbbgrgrbgrgbrrgg", "rbbbrbggbrrrggrrrgrbgrgrbggbggbrrgbbbbbbrbbrrgbgbbrbrrgrrrbgrrbrrbgrbr", "grgbbrggrgbrggbbbgbrbbbbgbggbrbrgrbrgrbgbbgrbbgrrggrrgrgbgrrrrgbrbbrrg", "bbggrbggrrrggrgrrgbbbgbgbrgbrggbggbbbbrrbgrgrbggggrgrggrbbgrbbbbbrgrgb", "rrgrrrgbrrgrggbgbrbrrbbrbrrrbrrggbrrgbgrrrrbrbgrrbbbbbbggrbrgbgrgrrrbg", "bgggggbbbbgggbbrrrgrgggrggbrrbbrgrgbgbrgggrbrrggbbrrbbgbbrrgrrggrrgrrg", "bbrbggbgrrggbbrbrgrrgrrbbgbrrrbrbrrrbgbrbgrrgggggrrrrbbggggbgbbbrbrbrr", "brbgrrrgbbgbgbgbggrrggbgrbrgrbbbgrrrrggbrrrrbgbbbgrrrrrrrgrbrrbrrrgbrg"}
	fmt.Println(maximumAreaTriangle(a))
}

func maximumAreaTriangle(a []string) int {
	numRows := len(a)
	numCols := len(a[0])
	var maxArea float64 = 0
	for row := 0; row < numRows; row++ {
		for col := 0; col < numCols; col++ {
			firstVertex := []int{row, col}
			for row2 := row + 1; row2 < numRows; row2++ {
				secondVertex := []int{row2, col}
				curMaxArea := findMaxArea(firstVertex, secondVertex, a)
				if curMaxArea > maxArea {
					maxArea = curMaxArea
				}
			}
		}
	}
	return int(math.Ceil(maxArea))
}

func findMaxArea(firstVertex []int, secondVertex []int, a []string) float64 {
	if a[firstVertex[0]][firstVertex[1]] == a[secondVertex[0]][secondVertex[1]] {
		return 0
	}

	numRows := len(a)
	numCols := len(a[0])
	var maxArea float64 = 0
	for row := 0; row < numRows; row++ {
		for col := 0; col < numCols; col++ {
			if a[row][col] != a[firstVertex[0]][firstVertex[1]] && a[row][col] != a[secondVertex[0]][secondVertex[1]] {
				thirdVertex := []int{row, col}
				curArea := areaTriangle(firstVertex, secondVertex, thirdVertex)
				//fmt.Printf("first: %v, second: %v, third: %v, area: %v\n", firstVertex, secondVertex, thirdVertex, curArea)
				if curArea > maxArea {
					maxArea = curArea
				}
			}
		}
	}
	return math.Ceil(maxArea)
}

func areaTriangle(firstVertex, secondVertex, thirdVertex []int) float64 {

	//return math.Abs(0.5 * (float64((firstVertex[0]+1)*(secondVertex[1]-thirdVertex[1])) + float64((secondVertex[0]+1)*(thirdVertex[1]-firstVertex[1])) + float64((thirdVertex[0]+1)*(firstVertex[1]-secondVertex[1]))))
	base := math.Abs(float64(secondVertex[0] - firstVertex[0])) + 1
	height := math.Abs(float64(firstVertex[1] - thirdVertex[1])) + 1
	return 0.5 * height * base

}

