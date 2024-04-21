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
	n, _ := strconv.Atoi(scanner.Text())

	var sumX, sumY, sumZ int
	for i := 0; i < n; i++ {
		scanner.Scan()
		line := strings.Split(scanner.Text(), " ")

		for i, force := range line {
			forces, _ := strconv.Atoi(force)
			if i == 0 {
				sumX += forces
			}
			if i == 1 {
				sumY += forces
			}
			if i == 2 {
				sumY += forces
			}
		}
	}

	if sumX == 0 && sumY == 0 && sumZ == 0 {
		fmt.Println("YES")
		return
	}
	fmt.Println("NO")

}
