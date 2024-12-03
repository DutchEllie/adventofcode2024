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

func day3_2() string {
	file, _ := os.Open("./input-3")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var fullText []string
	for scanner.Scan() {
		line := scanner.Text()
		fullText = append(fullText, line)
	}

	everything := strings.Join(fullText, "")

	allFuncsRe := regexp.MustCompile(`(mul\(\d+,\d+\))|(do\(\))|(don\'t\(\))`)
	allInstructions := allFuncsRe.FindAllString(everything, -1)

	type pair struct{
		enabled bool
		result int
	}

	initPair := pair{true, 0}
	result := fp.Reduce(func(result pair, current string) pair {
		mulRe := regexp.MustCompile(`mul\(\d+,\d+\)`)
		mul := mulRe.FindString(current)

		doRe := regexp.MustCompile(`do\(\)`)
		do := doRe.FindString(current)

		dontRe := regexp.MustCompile(`don\'t\(\)`)
		dont := dontRe.FindString(current)

		// Instruction is mul
		if mul != "" {
			if !result.enabled {
				return result
			}

			numRe := regexp.MustCompile(`\d+`)
			nums := numRe.FindAllString(current, -1)
			if len(nums) > 2 {
				fmt.Println("More than two nums found?")
			}

			leftStr, rightStr := nums[0], nums[1]
			left, _ := strconv.Atoi(leftStr)
			right, _ := strconv.Atoi(rightStr)

			return pair{result.enabled, result.result + (left * right)}
		}

		// Instruction is do
		if do != "" {
			return pair{true, result.result}
		}

		if dont != "" {
			return pair{false, result.result}
		}

		fmt.Println("You probably shouldn't get here...")
		return result

	}, initPair)(allInstructions)

	return fmt.Sprintf("%d", result.result)
}
