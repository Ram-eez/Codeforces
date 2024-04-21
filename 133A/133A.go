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
	line := strings.Split(scanner.Text(), "")

	for _, letter := range line {
		if letter == "H" || letter == "Q" || letter == "9" {
			fmt.Println("YES")
			return
		}
	}
	fmt.Println("NO")
}
