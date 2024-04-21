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
	nFriends, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	line := strings.Split(scanner.Text(), " ")

	eachFinger := make([]int, nFriends)
	for i, finger := range line {
		eachFinger[i], _ = strconv.Atoi(finger)
	}

	var sum int
	for i := range eachFinger {
		sum += eachFinger[i]
	}

	dina := []int{1, 2, 3, 4, 5}

	var count int

	for _, finger := range dina {
		totalCount := sum + finger
		if totalCount%(nFriends+1) != 1 {
			count++
		}

	}
	fmt.Println(count)
}
