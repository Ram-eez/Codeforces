package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	noSol, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	line := strings.Split(scanner.Text(), " ")

	hSols := make([]int, noSol)
	for i, sol := range line {
		hSols[i], _ = strconv.Atoi(sol)
	}
	maxIndex := largest_ele(hSols)
	lowIndex := smallest_ele(hSols)

	var count int
	if maxIndex > lowIndex {
		count = maxIndex - 1 + (noSol) - lowIndex - 1

	} else {
		count = maxIndex - 1 + (noSol) - lowIndex
	}

	fmt.Println(count)

}

func largest_ele(Slice []int) int {
	var max int
	max = Slice[0]
	var maxIndex int
	for i, ele := range Slice {
		if ele > max {
			max = ele
			maxIndex = i
		}
	}
	return maxIndex

}

func smallest_ele(Slice []int) int {
	var low int
	low = Slice[0]
	var lowIndex int
	for i, ele := range Slice {
		if ele <= low {
			low = ele
			lowIndex = i
		}
	}
	return lowIndex

}
