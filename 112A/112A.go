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

	line1 := strings.ToLower(scanner.Text())

	scanner.Scan()

	line2 := strings.ToLower(scanner.Text())

	for i := 0; i < len(line1); i++ {

		if line1[i] < line2[i] {
			fmt.Println("-1")
			break

		} else if line1[i] > line2[i] {
			fmt.Println("1")
			break
		}

	}
	if line1 == line2 {
		fmt.Println("0")
	}

}
