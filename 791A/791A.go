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
	line := strings.Split(scanner.Text(), " ")

	limak, _ := strconv.Atoi(line[0])
	bob, _ := strconv.Atoi(line[1])

	var years int

	for i := 0; limak <= bob; i++ {
		limak *= 3
		bob *= 2
		years++
	}

	fmt.Println(years)
}
