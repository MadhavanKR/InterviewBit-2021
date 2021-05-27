package main

import (
	"fmt"
	"math"
)

func main() {
	//large input
	//a := []string{"bgrrrbbgrrbggbrrggrggrrbrgbgbrbggbgrrbbrggrgbgrbbggbrgggrbrbrbgbbbbbbr", "ggrgggrgbgrbrbgrrbgbbgbrrrrgbrrrrggrrrrbbggrrgggbgrrbrbgbbbgbrbbbbgrrg", "rrrbbbggbbbrbgbbrrgbbrrgrbgbbrgggrrbgrggggrgbbrbbggrggbggbrbbbgrbbbrrb", "rrrrbrrgrbrbbbgggrrrrbgrggrbrbrbgrrrgbrrbbbgbbbbgbrgbggbbbbgrgrbrrgbrr", "bbgrgggrrggbgbgbgbrrbgrbrggrggggggbggrbrrgrbrggggggbrbrgbgrrbrggrgrrbg", "rgrbrrgrbrbrrgbgrgbbbrrbbbrbgrrrrrggbrgrbrgrrrggbbrbbrgrgrbbggbbbbbbrg", "brrggrbbggbrrrrggrggrgrrbgbgrbbrgrbrrbgrrrbbrbrgrrgrbgrggrrbrgggrbrbgb", "rrgrgbbbbbbbggbggrrgbgggbrggbbgbbgrgggrgrrbrrggrbbrgbbrggrrgrbrrggbbgg", "rbgbgrbbgbgggbgbrbgggbbbgrbrgbrggbbrgbbrbbbrbgbgrbbgbrbgbgrgbrrrgrrrbg", "rbggbgbrggbgbgrrbbrbgbrbgrgggrrrrggbbbrgrrbgggrrbgrgbrgrrbrrbgbrrrbrgb", "rgrrrrgrbbrbrbggrrgbbbbgbbrbrgbgrggggrrbggrrrbgrgrgbbbgrgbbgbrbrbrbbrg", "rrgbgbggrbbbggrbbgbgrbbbgbgbgbrrbgrbbbgrrrbbgrrgggbbrbgrgbbbrbbgbbgbgr", "bbbggbbrbgggbrrrgrbgbrbrgrrrbbbgbgrgrggbbrbgbbbrrbrbgrggrbrbbgrgrgbrgb", "ggrrrrrgggrggbrrrrrrbbrggggbbgggbrggrbrrbbggrrrrbrbgrrbbrgrgrrrbrgrrbg", "gbrrgrrgbrggbrgrrrgbrrbgbbggbgrbbrgrbrrbbgbbgbgbbggrrggbbbgrrbggbbbggr", "ggbrrggrgbrbgggggrrrrbrbbrbgrggrrbbrgrrgbggbrgbgrbgbbbgrrrrrrgrbgggbrb", "grggrrgbggrggggrgbbggrrggbgrbrrrbbrrggbbgbrgbbgggbbbbbrrrbrrrgbggrrbrg", "rbgbbrgrbgrggbgbbrrrrrrbbrbbgrrbbbgggggrrgrgbbgbbbgrbbbbrbrgrbrbrrggrr", "gbrgbbbgbrbrbbbggrrgbbgrrbrrbrbrrbbrggrgggrbgrrggrggrbrbbbbrgbbbbbgrrg", "bgbggrrbgbrgbgbrrgrbgggrgrbbgrbrbrgrgbrrgbbgbrrbrbrrbbgbrrgrrrrbbrrbrg", "ggbbgrgrggbrgrrbrgbgbggrbrggbrrrbgbgbrbbgrgbbgrrrggrgbbrggrgrbggrrgrgr", "rrbbgrbrrbgrgrrbggrrgrrgrrgrbrggbrbbbbrgrrrbbrbgbbbrbrrbrrbrrggrgggrgb", "rbrrgrgrbbbgrrbggrrggrgggbrbrbrbgbrgrggrgggrrgggbrbrrbrrgbggrgbgrbrrrb", "rbbbrrgbrbrggrbgrgbbrrgrgbrgrgbrbrbrbgrbrbrgbrbgggbbbbrrgbgbrrrrgggrbb", "rrggbgbrrrrrrgrgbgggrbggbrrrbrbggrgrrbrrbrrgrbbrbrbgbbgggbggbrgggrrrbb", "rrbbbbrrbbbbrrgrbrbrggggggrgrrrbgrrrrbrrbbrrgbgbbrrbrbgbbbgrbgrgbggggb", "rbbbbgrrbbgrbgbggrggrggrbbrbrgbgrgggrrggbgrbrbrgggrrrbgrgrgrrbbbrgrbrb", "bbbbrbbrrbrbbgbrbgrgrgggbbgrrgggggrgrgrrrgggggrrrbrbgbgrbrbrbggbrbgbrg", "rrgrrbrrbbgbrgbrggrrbggbbbbggrrrrbbgrrgbrbrrbgggbrrgrrrgbrrgggbggrbrgg", "grggbgrggbgrrrgrggggrbrgrgbbggrggbrbggrggbggbbrggrbbrrrbrrbbrrrgbgrrrg", "bbbbrgbrggbgbrggbgbgrbbrgbbggbgbrgrgggggrbgbbggrggrbggggrbggrbrbrbbbgb", "grrbrgbbbgrrggrgbrggbrrrbrgbgbbggggrrbrbbrrbbbgbgbgbbrbgrbrbbrggbggbbb", "gbgrgrbbrbrbrrgbrbbbrbgrgbrbgrrrgrrgbrggrbbrrgbbbggggrbggrbbgbgrrgrgrg", "grrbrbgbgrrgrgrggrrgrbrggbbrbbgrbggggbggrggrbbrgbbgrgrbrggrrrgrbbrbrrg", "ggggrgbbbbbggbrbrgbrrbrgggggbbgrrgrbrgggbrgbbbgrrgrggbgbrrggbgbrgbrrrb", "gbgrbgrbbrrgbrgggrrbgrgggrgrbrbbbbrrrbrbgrgbbbrbbbbrrgbgrgbgbbrgbrrgbg", "bgrbggrrggbgggrgrbbgbrbrrgbgggggggbbbgbbrbrrrgrbrgbbbrrgrbggrbbggrrbrg", "bgrrrgbgrbrggbrrbgrggrgrbbbbbrbbbrgrbrrbrgbbgggrrgrbrrrbbrbbbbgbbbrrrg", "brgggbgrgrgrrgbrrbrggggrrbrrbbrgbbgrggbggbrbggrbgrgrbbgrrrgrgrbgbgrggr", "rrbrgrrrrrbggrbgrbrrrbbggrbgbgbrgbgrggbggrrrbggrgbbbrrbbrbbrrgrrbrgbrb", "rggggrbrgbrrgrgrrbbgbbgrggrgrbbbgrgbrbbrbbgrrgrbrbggbgbbgggrbgbrbbgrrg", "gbbbrgggbrrrbbgbgbbgrrgggbrrgrrggrrgggrggrgbbrrrrrrgrrrggbbbrbrgrgbgbg", "grrbbgbbbbgbbgbrbgbbrgrgbrrrbrrrbrbggbrbgrrrrrrrbbrbgrrbrggrbbgggbrrrb", "brbbggbgrgbrgrggrrgrrrrrrrrbgrbrrrgrbrrbbbgrbggrgbrrggrbbbbrgggbrbrbrg", "rbgbrrrgrbbgbgrbbgbbbbrgbrbrbbrgrbbrbrbgbgggrbrbgbbgrgrgbbgrrgbbbgrggb", "rrgbggrgrrgbggrggbgbrrgggrrggrbrgggggbbbrgbbbrgbrrggrggggrbbrgbgggbbrb", "bbrrbbbrbbrgbrbrrggrbrrrrbbbggrgrbbrbggrggggbbbrbbbrbrbggrbggrbbbggbrb", "gbbgbbbbggbrggbbgbrggbbbggbgbggggbrbbrrgrgbrrgbgbbggbggrrrrbbggrgbbrrg", "gbrrbrbggrrrbbrbrbgbrrrrbgrgggrbbgbbrgrbrrggrgbgrgggbrbbbgbbbgrbgbbbrr", "grrggbrrgrbgrrgrgrbrgrgrgbbrbbgrrbrgggggggbrggbrgbbbrggggggrbgrrgrrbgr", "brbrrgrgrbggbbrggbgrgggrgbgbbggrbrrbbrrrgrggbbrbrrggbgrbrrbrgbbrggrrrr", "rrgrbbbbbggbggrrggggrggrgbbbbbggrbrbrrrrrrrbgbrrgbrbgbrrggbbrbrrgrgrrr", "brgrgbbbbrrbrgrbbbbrbbrbbrgbrbgggbbrbgbrbgggbrbgggrrbrrggrrrgbrgggbggr", "brgggbbbrbbrrbrggbrbrggbgrbbgbbgbbrrrrbrrgbrbggrrbbggrrbgbggbgrbrgbgbg", "gbrrgbbgrgrgbbggrrbrbbgbgbggggbggrrgbgrgbbgrgbggbbrbggggrbrrrbrbgrgrgr", "rggrgbgbgggbbrrbgbrggbrgbrgrgrgrbbbggrbrbggrrbggggrbrbrrbbrbbbbrbrgrbg", "bggbbbrggrgbgggrrrrgrgbbbgbrggggbrrrrbrbgbrbgbggbbgbgrgbrgrbbbbrgbrrrr", "gbrrbggrbgbbbbbgrrbbggbbbrbgrbgrrrrbgbgbbggrgbbgbbrrggbbgrgrrggbrrbrgb", "rrbggbgbbgrbggbrggrrrbgrgrgbgrgggbrbrbrbgrrggrrgrrggbrgbrgrbbbbrbrbrgg", "gbbbbgbrgbbrrgrgggbbggbrbrrgrgrrgrbbgggrrbgrggrgbbrgrggbgrbbbgbrrrbbrg", "grbgbbrrrgbbrgrgrbgbbbgrrbrbbgbbbbbgggbbbrrggrgrbbbgbbgrrbgbbrrgrrbbgg", "bgrggggggbbrgbrrrrbgrrbrgbbgrgrgrrrrgggrbbrrrgrbgrrrgbggrrgbggrbbbrrgr", "rggbrggggbgbrbrggrrgbrbbbbrbggbbbrbbgggrrbrrbrrbrrgbgbggbggrgggrgbrbgg", "brrbrgggbrgbrrrggbrbrrbrgggrrgrrbgbrrrrrrbggrgrgrrrrbgrggrbbrgbrgbgbrb", "grggbbbgbrrrrggrrbgrrgrgbbrbgrggggrbggrbgggrrrrrbggrbbrbgggbggrrrgrrbg", "rrgrgrggbbbbgbbgrrggrrgrrgrgbbgrbgggbrgrbgrgggggrbbggbrbbgbbrbbgrrrggr", "ggbgbbgbbbbgrbbrggbbggrggggbbrrbrgbrbbbgbrbbbbrgrbbrrgrrbbbgrrrrggrgrr", "bgrgggbbggrbgrrgrrgbbbgbrrbrrrgrgbrgrgrrrgbgbrbrgrbrgbgggrrggrrbbgrgbb", "rgbgggbgrrrgbrbbbgbrgbbgbgbggrggrgbrgbbbggbrbrrgrgbbgggbgrrrgbrbrgggbb", "gbggbgbrrrbgbrrggrbbbbrbgbgbgrggbrrrbbrbggrrbrbrbbrgbbgrbrrbrbbgbrrrrb", "ggrgbrbbrbrbbbrbgbrbbggbbggrrrrrbbggbrgbgrgggbggbgbgbgbrrbrbbgbbbrgrbb", "bbbbrgrrbrrgbggbrrggbgrrrrgrrbggbbbrgbbgrrbgrrrgbbrgrbrbgbbbbbrrbbbggr", "rggggbgrrbrgrbggbbbgbbbbbrbbgrgbggbrrgbgrrgrbggrbggbbgrbbbbbbbbrgrgbrb", "rbrbrbbgrgbgggrbgrrbggrgbggrrrrbbbrbrbrrgbbbbggbrbbbgrbbggbgbgbggbbrrr", "rrbrbbbgggbggrrbgrgbrrbrggggrrrbrrggbrgrgbrgbgrbbbbrgbgggbrrgggrrgbgbg", "rrgrgbrbbbbbbggrgrrgrbbrggbgbbrrbbbbbgrgrgrrbrrrbrrgrbbbrggbrgbgrbgrrr", "gbgbgbbggrgggrbrrrbgrrbggbgbggbgbbgbbbbbbbrrrbgrbrbbgbggbrrgggrrbrbrbb", "bbrbbrggrbbgbrbbbbrbgbgrggbbrggbbgbbgbrrrrrbgrrbrgbbbrbbbgbrgrbrgbrggb", "gggrrggrbrgrbbgrrgrbrgbggbgbgrgbrrbggrgbgrbbrggbgbrgbrgrrbgrgrrrbggrbb", "bgrggrgbrrgbrbbbbbrgbbgggrrrbbrrrggrbbggrgrbgrrgrgrggbrbbgrgrbgrgbrrgg", "rbbbrbggbrrrggrrrgrbgrgrbggbggbrrgbbbbbbrbbrrgbgbbrbrrgrrrbgrrbrrbgrbr", "grgbbrggrgbrggbbbgbrbbbbgbggbrbrgrbrgrbgbbgrbbgrrggrrgrgbgrrrrgbrbbrrg", "bbggrbggrrrggrgrrgbbbgbgbrgbrggbggbbbbrrbgrgrbggggrgrggrbbgrbbbbbrgrgb", "rrgrrrgbrrgrggbgbrbrrbbrbrrrbrrggbrrgbgrrrrbrbgrrbbbbbbggrbrgbgrgrrrbg", "bgggggbbbbgggbbrrrgrgggrggbrrbbrgrgbgbrgggrbrrggbbrrbbgbbrrgrrggrrgrrg", "bbrbggbgrrggbbrbrgrrgrrbbgbrrrbrbrrrbgbrbgrrgggggrrrrbbggggbgbbbrbrbrr", "brbgrrrgbbgbgbgbggrrggbgrbrgrbbbgrrrrggbrrrrbgbbbgrrrrrrrgrbrrbrrrgbrg"}

	//small positive case
	//a := []string{"rrrrr", "rrrrg", "rrrrr", "bbbbb"}
	/*
		r r r r r
		r r r r g
		r r r r r
		b b b b b
	 */


	//small negative case
	//a := []string{"rrr", "rrr", "rrr", "rrr"}

	//test case
	a := []string{"rrrbr", "ggrgb", "bbbrg", "rrgrr", "gbggb", "rbbgr", "rbgrg", "bbbgr", "ggbbb", "bggbr", "ggrbb", "grrrg", "rbrrg", "brrgr", "rrgbg", "bbrgr", "gbbbr", "rrbgb", "brbbr", "bgrrr", "bbggr", "rggbg", "bbggg", "gggbb", "bgbbg", "rrbgr", "rggrb", "grggr", "rggrr"}

	fmt.Println(maximumAreaTriangleV3(a))
}

func maximumAreaTriangleV3(a []string) int {
	numRows := len(a)
	numCols := len(a[0])
	var maxArea float64 = 0
	//left and right stores the left most and right most column in which a color occurs left[0]=1 implies that color red appears left most in column 1
	right := []int{-1, -1, -1}
	left := []int{9999, 9999, 9999}

	//top and bottom are 2d arrays which holds the top most and bottom most occurance of a color in a given column. top[0][0] = 1 implies the color red first occurance in column 0 is at row 1
	top := make([][]int, 0)
	bottom := make([][]int, 0)

	//filling in left and right arrays
	for i := 0; i < numRows; i++ {
		for j := 0; j < numCols; j++ {
			colorCode := getColorCode(a[i][j])
			left[colorCode] = min(left[colorCode], j)
			right[colorCode] = max(right[colorCode], j)
		}
	}

	//fmt.Println("left: ", left)
	//fmt.Println("right: ", right)

	for j := 0; j < numCols; j++ {
		top = append(top, []int{9999, 9999, 9999})
		bottom = append(bottom, []int{-1, -1, -1})
	}

	//fmt.Println("top: ", top)
	//fmt.Println("bottom: ", bottom)
	//filling in top and bottom arrays
	for j := 0; j < numCols; j++ {
		for i := 0; i < numRows; i++ {
			colorCode := getColorCode(a[i][j])
			top[j][colorCode] = min(top[j][colorCode], i)
			bottom[j][colorCode] = max(top[j][colorCode], i)
		}
	}

	//fmt.Println("top: ", top)
	//fmt.Println("bottom: ", bottom)

	//for each column we find the max base and then third vertex and calculate area
	for col := 0; col < numCols; col++ {
		for x := 0; x < 3; x++ {
			for y := 0; y < 3; y++ {
				z := 3 - x - y
				if x != y && top[col][x] != 9999 && bottom[col][y] != -1 && left[z] != 9999 {
					//fmt.Println("colors are: ", x, y, z)
					curArea := 0.5 * (math.Abs(float64(bottom[col][y]-top[col][x]))+1.0) * (math.Abs(float64(col-left[z]))+1.0)
					//fmt.Printf("left: first vertex(%d, %d), second vertex(%d, %d), third vertex in column: %d, area: %f\n", top[col][x], col, bottom[col][y], col, left[z], curArea)
					//fmt.Printf("base: %f, height %f\n", math.Abs(float64(bottom[col][y]-top[col][x]))+1.0, math.Abs(float64(col-left[z]))+1.0)
					if curArea > maxArea {
						maxArea = curArea
					}
				}

				if x != y && top[col][x] != 9999 && bottom[col][y] != -1 && right[z] != -1 {
					//fmt.Println("colors are: ", x, y, z)
					//fmt.Printf("right: first vertex(%d, %d), second vertex(%d, %d), third vertex in column: %d\n", top[col][x], col, bottom[col][y], col, right[z])
					curArea2 := 0.5 * (math.Abs(float64(bottom[col][y]-top[col][x]))+1.0) * (math.Abs(float64(col-right[z]))+1.0)
					if curArea2 > maxArea {
						maxArea = curArea2
					}
				}
			}
		}
	}
	return int(math.Ceil(maxArea))
}

func max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func maximumAreaTriangleV2(a []string) int {
	numRows := len(a)
	numCols := len(a[0])
	var maxArea float64 = 0

	var colorsPresent [1000][1000]int
	for row := 0; row < numRows; row++ {
		for col := 0; col < numCols; col++ {
			colorCode := getColorCode(a[row][col])
			colorsPresent[colorCode][col] = 1
		}
	}

	for col := 0; col < numCols; col++ {
		i := 0
		j := numRows - 1
		var firstVertex, secondVertex []int
		var firstColor, secondColor uint8
		for i < j {
			if a[i][col] != a[j][col] {
				firstVertex = []int{i, col}
				firstColor = a[i][col]
				secondVertex = []int{j, col}
				secondColor = a[j][col]
				break
			}
			i++
			j--
		}
		if i == j {
			continue
		}
		for col := 0; col < numCols; col++ {
			thirdColor := getThirdColor(getColorCode(firstColor), getColorCode(secondColor))
			if colorsPresent[thirdColor][col] == 1 {
				base := math.Abs(float64(secondVertex[0]-firstVertex[0])) + 1
				height := math.Abs(float64(firstVertex[1]-col)) + 1
				curArea := 0.5 * height * base
				if curArea > maxArea {
					maxArea = curArea
				}
			}
		}
	}
	return int(math.Ceil(maxArea))
}

func getThirdColor(firstColor, secondColor int) int {
	return 3 - firstColor - secondColor
}

func getColorCode(color uint8) int {
	switch color {
	case 'r':
		return 0
	case 'g':
		return 1
	case 'b':
		return 2
	}
	return -1
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
	base := math.Abs(float64(secondVertex[0]-firstVertex[0])) + 1
	height := math.Abs(float64(firstVertex[1]-thirdVertex[1])) + 1
	return 0.5 * height * base

}
