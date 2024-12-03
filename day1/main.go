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
	fmt.Printf("Day 1 - 1: %s\n", day1_1())
	fmt.Printf("Day 1 - 2: %s\n", day1_2())
}

func day1_1() string {
	var lefts  []int
	var rights []int

	file, _ := os.Open("./input-1")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		inputs := strings.Split(line, "   ")
		left, _ := strconv.Atoi(inputs[0])
		right, _ := strconv.Atoi(inputs[1])
		lefts = append(lefts, left)
		rights = append(rights, right)
	}

	sort.Slice(lefts, func(i, j int) bool {
		return lefts[i] < lefts[j]
	})

	sort.Slice(rights, func(i, j int) bool {
		return rights[i] < rights[j]
	})

	var distances []int 

	for i := 0; i < len(lefts); i++ {
		var distance int
		if lefts[i] > rights[i] {
			distance = lefts[i] - rights[i]
		} else if rights[i] > lefts[i] {
			distance = rights[i] - lefts[i]
		} else if lefts[i] == rights[i] {
			distance = 0
		}
		distances = append(distances, distance)
	}

	var totalDistance int

	for _, v := range distances {
		totalDistance = totalDistance + v
	}

	return fmt.Sprint(totalDistance)

}

func day1_2() string {
	var lefts  []int
	var rights []int

	file, _ := os.Open("./input-1")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		inputs := strings.Split(line, "   ")
		left, _ := strconv.Atoi(inputs[0])
		right, _ := strconv.Atoi(inputs[1])
		lefts = append(lefts, left)
		rights = append(rights, right)
	}

	sort.Slice(lefts, func(i, j int) bool {
		return lefts[i] < lefts[j]
	})

	sort.Slice(rights, func(i, j int) bool {
		return rights[i] < rights[j]
	})

	var totalSimScore int
	for _, v := range lefts {
		rightsOccurenceCount := 0
		simScore := 0
		for _, w := range rights {
			if w == v {
				rightsOccurenceCount++
			}
		}

		simScore = v * rightsOccurenceCount
		totalSimScore = totalSimScore + simScore
	}

	return fmt.Sprintf("%d", totalSimScore)
}