package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	text := strings.Split(line, "")

	noLetters := "abcdfgijkmnpqrstuvwxyz"
	newSlice := make([]string, len(text))
	for _, letter := range text {
		if !strings.ContainsAny(letter, noLetters) {
			newSlice = append(newSlice, letter)
		}
	}

	fmt.Println(newSlice)

}
