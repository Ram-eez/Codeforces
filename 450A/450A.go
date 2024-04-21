package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := strings.Split(scanner.Text(), " ")

	n, _ := strconv.Atoi(line[0])
	m, _ := strconv.Atoi(line[1])

	scanner.Scan()
	newLine := strings.Split(scanner.Text(), " ")

	candies := make([]int, n)
	for i, number := range newLine {
		candies[i], _ = strconv.Atoi(number)
	}

	iterations := make([]int, len(candies))
	for i := range candies {
		iterations[i] = int(math.Ceil(float64(candies[i]) / float64(m)))
	}

	max := iterations[0]
	index := 0
	for i, candy := range iterations {
		if candy >= max {
			max = candy
			index = i
		}
	}
	fmt.Println(index + 1)
}
