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

	num, _ := strconv.Atoi(scanner.Text())

	for i := 0; i <= num; i++ {
		scanner.Scan()
		word := scanner.Text()

		if len(word) > 10 {

			wordLength := len(word) - 2
			fmt.Printf("%s%d%s\n", word[:1], wordLength, word[len(word)-1:])

		} else {

			fmt.Println(word)

		}

	}

}
