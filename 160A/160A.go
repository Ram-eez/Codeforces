package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	nCoins, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	string := scanner.Text()
	values := strings.Split(string, " ")

	allCoins := make([]int, nCoins)
	for i, value := range values {

		allCoins[i], _ = strconv.Atoi(value)

	}

	sort.Slice(allCoins, func(i int, j int) bool {
		return allCoins[i] > allCoins[j]
	})

	var totalSum int
	for i := range allCoins {

		totalSum += allCoins[i]

	}

	var pickedCoin int
	var count int
	for i := range allCoins {
		pickedCoin += allCoins[i]
		count++
		if pickedCoin > totalSum/2 {
			break
		}
	}

	fmt.Println(count)
}
