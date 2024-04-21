package main

import (
	"fmt"
	"math"
)

func main() {

	matrix := make([][]int, 5)
	for i := range matrix {
		matrix[i] = make([]int, 5)
	}

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Scan(&matrix[i][j])

		}
	}
	// print a matrix
	// for i := 0; i < 5; i++ {
	// 	for j := 0; j < 5; j++ {
	// 		fmt.Printf("%d ", matrix[i][j])
	// 	}
	// 	fmt.Println()
	// }

	var rowIndex, colIndex int

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if matrix[i][j] == 1 {
				rowIndex = i
				colIndex = j
				break
			}
		}
	}

	// var Steps int
	// for range matrix {
	// 	if rowIndex == 2 && colIndex == 2 {
	// 		break

	// 	}
	// 	if rowIndex != 2 && rowIndex > 2 {
	// 		rowIndex--
	// 		Steps++
	// 	}
	// 	if rowIndex != 2 && rowIndex < 2 {
	// 		rowIndex++
	// 		Steps++
	// 	}
	// 	if colIndex != 2 && colIndex > 2 {
	// 		colIndex--
	// 		Steps++
	// 	}
	// 	if colIndex != 2 && colIndex < 2 {
	// 		colIndex++
	// 		Steps++
	// 	}
	// }

	//solution 2

	Steps := math.Abs(float64(2-rowIndex)) + math.Abs(float64(2-colIndex))

	fmt.Println(Steps)

}
