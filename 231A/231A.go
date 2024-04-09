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
	problems, _ := strconv.Atoi(scanner.Text())

	sureCount := 0

	for x := 0; x < problems; x++ {
		scanner.Scan()
		line := scanner.Text()

		numbers := strings.Split(line, " ")

		petya, _ := strconv.Atoi(numbers[0])
		tonya, _ := strconv.Atoi(numbers[1])
		vasya, _ := strconv.Atoi(numbers[2])

		opinion := petya + vasya + tonya
		if opinion >= 2 {
			sureCount++
		}
	}
	fmt.Println(sureCount)

}
