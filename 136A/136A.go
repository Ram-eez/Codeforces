package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Index(s []int, v int) int {
	for i, vs := range s {
		if vs == v {
			return i
		}
	}

	return -1
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	line := strings.Split(scanner.Text(), " ")

	friends := make([]int, n)

	for i := range line {
		friends[i], _ = strconv.Atoi(line[i])

	}

	gifts := make([]int, n)
	for i, friend := range friends {
		if friend == i+1 {
			gifts[i] = i + 1
		} else {
			gifts[i] = Index(friends, i+1) + 1
		}

	}
	answer := make([]string, len(gifts))
	for i := range gifts {
		answer[i] = strconv.Itoa(gifts[i])

	}

	ans := strings.Join(answer, " ")

	fmt.Println(ans)

}
