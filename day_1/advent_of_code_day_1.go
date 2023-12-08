package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {

	fmt.Printf("Solution 1 answer:" + strconv.Itoa(solution_1()))
}

func solution_1() (total int) {
	file, _ := os.Open("input_1_1.txt")
	bscanner := bufio.NewScanner(file)

	for bscanner.Scan() {
		line := bscanner.Text()
		result := []int{}

		for _, v := range line {
			if unicode.IsDigit(v) {
				result = append(result, int(v)-'0')
			}
		}

		length := len(result)

		if length == 1 {
			total += (result[0] * 10) + result[0]
		} else if length > 1 {
			total += (result[0] * 10) + result[length-1]
		}
	}

	return total
}
