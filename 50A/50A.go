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

	line := scanner.Text()

	rectangle := strings.Split(line, " ")

	M, _ := strconv.Atoi(rectangle[0])
	N, _ := strconv.Atoi(rectangle[1])

	maxDomino := (M * N) / 2

	fmt.Printf("%d", maxDomino)
}
