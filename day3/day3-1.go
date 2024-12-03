package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/repeale/fp-go"
)

func day3_1() string {
	file, _ := os.Open("./input-3")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	
	var fullText []string
	for scanner.Scan() {
		line := scanner.Text()
		fullText = append(fullText, line)
	}

	everything := strings.Join(fullText, "")

	re := regexp.MustCompile(`mul\(\d+,\d+\)`)

	allMuls := re.FindAllString(everything, -1)


	result := fp.Reduce(func(result int, current string) int {
		numRe := regexp.MustCompile(`\d+`)
		nums := numRe.FindAllString(current, -1)
		if len(nums) > 2 {
			fmt.Println("More than two nums found?")
		}

		leftStr, rightStr := nums[0], nums[1]
		left, _ := strconv.Atoi(leftStr)
		right, _ := strconv.Atoi(rightStr)

		return result + (left * right)
	}, 0)(allMuls)


	return fmt.Sprintf("%d", result)
}