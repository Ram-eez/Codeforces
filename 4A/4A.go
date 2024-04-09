package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	scanner := bufio.NewReader(os.Stdin)
	name, _ := scanner.ReadString('\n')
	name = strings.TrimSpace(name)
	num, _ := strconv.Atoi(name)

	if num == 2 {
		fmt.Println("NO")
	} else if num%2 == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
