package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	// text := strings.Split(line, "")

	yesLetters := "hello"
	count := 0
	// newSlice := make([]string, len(text))

	for i := range text {
		if text[i] == yesLetters[count] {
			count++
		}
		if count == len(yesLetters) {
			fmt.Println("YES")
			return
		}
	}

	fmt.Println("NO")

}
