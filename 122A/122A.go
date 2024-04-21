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
	n, _ := strconv.Atoi(scanner.Text())

	if n%4 == 0 || n%7 == 0 || n%44 == 0 || n%47 == 0 || n%74 == 0 || n%77 == 0 || n%444 == 0 || n%447 == 0 || n%474 == 0 || n%744 == 0 || n%477 == 0 || n%747 == 0 || n%774 == 0 || n%777 == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

	// var luckyNos = []int{4, 7, 44, 47, 74, 77, 444, 447, 474, 477}

	// for i := range luckyNos {
	// 	if number%luckyNos[i] == 0 {
	// 		fmt.Println("YES")
	// 		return
	// 	}
	// }
	// fmt.Println("NO")
}
