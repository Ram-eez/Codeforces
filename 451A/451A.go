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
	line := strings.Split(scanner.Text(), " ")
	m, _ := strconv.Atoi(line[0])
	n, _ := strconv.Atoi(line[1])

	if min(m, n)%2 == 0 {
		fmt.Println("Malvika")
	} else {
		fmt.Println("Akshat")
	}
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
