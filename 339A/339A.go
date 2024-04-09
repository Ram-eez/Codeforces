package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()

	sums := (strings.Split(line, "+"))

	newSum := make([]int, len(sums))

	for i, sum := range sums {

		newSum[i], _ = strconv.Atoi(sum)

	}

	sort.Slice(newSum, func(i, j int) bool {

		return newSum[i] < newSum[j]

	})

	easySum := make([]string, len(newSum))
	for i, num := range newSum {
		easySum[i] = strconv.Itoa(num)
	}

	result := strings.Join(easySum, "+")

	fmt.Println(result)

}
