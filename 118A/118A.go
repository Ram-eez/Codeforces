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

	letters := strings.Split(line, "")

	vovels := "AOYEUIaoyeui"

	newSlice := make([]string, len(letters))

	for _, letter := range letters {

		if !strings.ContainsAny(letter, vovels) {

			newSlice = append(newSlice, letter)

		}

	}
	snew := strings.Join(newSlice, "")

	var result string

	for i, char := range snew {

		if i >= 0 {
			result += "."
		}
		result += string(char)
	}
	answer := strings.ToLower(result)
	fmt.Println(answer)

}
