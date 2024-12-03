package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Day 2 - 1: %s\n", day2_1())
	fmt.Printf("Day 2 - 2: %s\n", day2_2())
}

func day2_1() string {
	var input [][]int

	file, _ := os.Open("./input-2")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		inputs := strings.Split(line, " ")

		var inputsInts []int
		for _, v := range inputs {
			vi, _ := strconv.Atoi(v)
			inputsInts = append(inputsInts, vi)
		}

		input = append(input, inputsInts)
		count++
	}

	// fmt.Println(input)
	numSafe := 0

	for _, report := range input {
		isSortedIncreasing := sort.SliceIsSorted(report, func(i, j int) bool {
			return report[i] < report[j]
		})
		isSortedDecreasing := sort.SliceIsSorted(report, func(i, j int) bool {
			return report[i] > report[j]
		})

		if !isSortedDecreasing && !isSortedIncreasing {
			// failed
			continue
		}

		isUnsafe := false

		distance := 0
		if isSortedIncreasing {
			for i := 0; i < len(report); i++ {
				if i + 1 == len(report) { continue }
				distance = report[i + 1] - report[i]
				// fmt.Println(distance)
				if distance < 1 || distance > 3 {
					isUnsafe = true
				}
			}
		} else if isSortedDecreasing {
			for i := 0; i < len(report); i++ {
				if i + 1 == len(report) { continue }
				distance = report[i] - report[i + 1]
				// fmt.Println(distance)
				if distance < 1 || distance > 3 {
					isUnsafe = true
				}
			}
		}

		if !isUnsafe {
			numSafe++

		}

	}

	return fmt.Sprintf("%d", numSafe)
}

func day2_2() string {
	var input [][]int

	file, _ := os.Open("./input-2")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		inputs := strings.Split(line, " ")

		var inputsInts []int
		for _, v := range inputs {
			vi, _ := strconv.Atoi(v)
			inputsInts = append(inputsInts, vi)
		}

		input = append(input, inputsInts)
		count++
	}

	numSafe := 0

	for _, report := range input {
		if isSafeWithDampener(report) { numSafe++ }
	}

	return fmt.Sprintf("%d", numSafe)
}

func isSafeWithDampener(report []int) bool {
	if isSafe(report) {
		return true
	} 

	for i := range report {
		portion := slices.Delete(slices.Clone(report), i, i+1)
		if isSafe(portion) {
			return true
		}
	}

	return false
}

func isSafe(report []int) bool {
	// Reverse to increasing
	if report[0] > report[1] {
		slices.Reverse(report)
	}
	isSortedIncreasing := sort.SliceIsSorted(report, func(i, j int) bool {
		return report[i] < report[j]
	})

	if !isSortedIncreasing {
		// failed
		return false
	}

	for i := 0; i < len(report); i++ {
		if i + 1 == len(report) { continue }
		distance := report[i + 1] - report[i]
		if distance < 1 || distance > 3 {
			// failed
			return false
		}

	}

	return true // safe
}