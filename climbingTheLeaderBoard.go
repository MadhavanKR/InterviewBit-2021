package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'climbingLeaderboard' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY ranked
 *  2. INTEGER_ARRAY player
 */

func climbingLeaderboard(ranked []int32, player []int32) []int32 {
	// Write your code here
	currentRanks := make([]int32, len(ranked))
	var curRank int32 = 1
	var curScore int32 = ranked[0]
	for rankIndex, rankScore := range ranked {
		//fmt.Println("rankIndex: ", rankIndex)
		if rankScore == curScore {
			currentRanks[rankIndex] = curRank
		} else {
			currentRanks[rankIndex] = curRank + 1
			curRank++
			curScore = rankScore
		}
	}
	fmt.Println("Current ranks: ", currentRanks)
	playerRanks := make([]int32, len(player))
	for gameNum, playerScore := range player {
		fmt.Printf("Processing game %d with score %d\n", gameNum, playerScore)
		low := 0
		high := len(ranked) - 1
		fmt.Println("before finding rank: ", playerRanks)
		for low <= high {
			mid := int((low + high) / 2)
			fmt.Printf("low: %d, high: %d mid: %d \n", low, high, mid)
			if ranked[mid] == playerScore {
				playerRanks[gameNum] = currentRanks[mid]
			}
			if (mid-1 >= 0) && (mid+1 < len(ranked)) {
				if (ranked[mid-1] > playerScore) && (ranked[mid] <= playerScore) {
					fmt.Printf("mid is %d\n", mid)
					fmt.Printf("assiging rank %d to score %d since mid-1 is %d and mid+1 is %d\n", currentRanks[mid-1]+1,
						playerScore, ranked[mid-1], ranked[mid])
					playerRanks[gameNum] = currentRanks[mid-1] + 1
				}
			}
			if playerScore < ranked[mid] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
		fmt.Println("after finding rank: ", playerRanks)
		if playerRanks[gameNum] == 0 {
			fmt.Printf("score hasnt been decided yet for %d, high: %d and low: %d\n", playerScore, high, low)
			if low >= len(ranked)-1 {
				if playerScore < ranked[len(ranked)-1] {
					playerRanks[gameNum] = currentRanks[len(ranked)-1] + 1
				} else {
					fmt.Println("am i coming here?")
					playerRanks[gameNum] = currentRanks[len(ranked)-1]
				}
			} else {
				if playerScore < ranked[0] {
					playerRanks[gameNum] = 2
				} else {
					playerRanks[gameNum] = 1
				}
			}
		}
	}
	return playerRanks
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	//stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	//checkError(err)

	//defer stdout.Close()

	//writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	rankedCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	rankedTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var ranked []int32

	for i := 0; i < int(rankedCount); i++ {
		rankedItemTemp, err := strconv.ParseInt(rankedTemp[i], 10, 64)
		checkError(err)
		rankedItem := int32(rankedItemTemp)
		ranked = append(ranked, rankedItem)
	}

	playerCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	playerTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var player []int32

	for i := 0; i < int(playerCount); i++ {
		playerItemTemp, err := strconv.ParseInt(playerTemp[i], 10, 64)
		checkError(err)
		playerItem := int32(playerItemTemp)
		player = append(player, playerItem)
	}

	result := climbingLeaderboard(ranked, player)

	// for i, resultItem := range result {
	// 	fmt.Fprintf(writer, "%d", resultItem)

	// 	if i != len(result)-1 {
	// 		fmt.Fprintf(writer, "\n")
	// 	}
	// }
	fmt.Println(result)
	//fmt.Fprintf(writer, "\n")

	//writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
