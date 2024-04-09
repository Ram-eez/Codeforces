package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	number, _ := strconv.Atoi(scanner.Text())

	no := 0
	for i := 0; i < number; i++ {

		scanner.Scan()

		input := scanner.Text()

		if input == "X++" || input == "++X" {
			no++
		} else if input == "X--" || input == "--X" {
			no--
		}

	}
	fmt.Println(no)
}
