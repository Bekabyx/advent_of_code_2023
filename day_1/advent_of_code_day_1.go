package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"unicode"
)

func main() {

	fmt.Println("Solution 1 answer:" + strconv.Itoa(solution_1()))
	fmt.Println("Solution 2 answer:" + strconv.Itoa(solution_2()))
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

func solution_2() (total int) {
	file, _ := os.Open("input_1_1.txt")
	bscanner := bufio.NewScanner(file)
	number_regex, _ := regexp.Compile(`(\d|one|two|three|four|five|six|seven|eight|nine)`)

	for bscanner.Scan() {
		line := bscanner.Text()
		var matches []string

		for i := 0; i < len(line)-1; i++ {
			submatches := (number_regex.FindAllString(line[i:], -1))
			// had a fun time learning about variadic functions :)
			matches = append(matches, submatches...)
		}

		length := len(matches)
		first_number, _ := strconv.Atoi(matches[0])
		last_number, _ := strconv.Atoi(matches[length-1])

		if first_number == 0 {
			first_number = getDigitFromString(matches[0])
		}

		if last_number == 0 {
			if matches[0] == matches[length-1] {
				last_number = first_number
			} else {
				last_number = getDigitFromString(matches[length-1])
			}
		}

		total += (first_number * 10) + last_number
	}

	return total
}

func getDigitFromString(number_string string) (digit int) {
	switch digit := number_string; digit {
	case "one":
		return 1
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	default:
		return 0
	}
}
