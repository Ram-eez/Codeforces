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

	line[0] = strings.ToUpper(line[0])

	ans := strings.Join(line, "")

	fmt.Println(ans)
}
