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

	numbers := strings.Split(line, " ")

	n, _ := strconv.Atoi(numbers[0])
	k, _ := strconv.Atoi(numbers[1])

	// fmt.Println(n, k)

	scanner.Scan()
	line1 := scanner.Text()

	scores := strings.Split(line1, " ")

	acScore := make([]int, n)

	for i, score := range scores {
		num, _ := strconv.Atoi(score)

		acScore[i] = num
	}

	for i := 0; i < len(acScore); {
		if acScore[i] <= 0 || acScore[i] < acScore[k-1] {
			acScore = append(acScore[:i], acScore[i+1:]...)
		} else if acScore[i] == acScore[k-1] {
			i++
		} else {
			i++
		}

	}

	fmt.Println(len(acScore))

}
