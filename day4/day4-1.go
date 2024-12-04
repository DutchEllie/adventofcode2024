package main

import (
	"bufio"
	"fmt"
	"os"
)

// type key struct {
// 	y int
// 	x int
// }

// I am literally just stealing from this guy: https://github.com/thisRedH/AdventOfCode/blob/main/2024/days/day04.go

// Not sure if this works yet, we'll see
// func InBounds[T any](grid [][]T, x, y int) bool {
// 	return x >= 0 && x < len(grid[y]) && y >= 0 && y < len(grid)
// }

func InBounds[T any](v [][]T, y, x int) bool {
    return y >= 0 && y < len(v) && x >= 0 && x < len(v[y])
}



// Stolen
func StrReverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func WordSearch(grid [][]rune, x, y int, dirDelta DirectionD, word string) bool {
	// x and y are the position of the start of the word, so the X in XMAS
	// dir is a [2]int with the delta of the direction
	// 			// ###################################
	//			// UPDATE: WE'RE NOT USING SUBSLICES ANYMORE, BUT THIS GRAPHIC STILL HELPS TO VISUALIZE
	// 			// Subslice is now a 2d slice with the X in the middle
	// 			// .......
	// 			// .......
	// 			// .......
	// 			// ...X...
	// 			// .......
	// 			// .......
	// 			// .......
	// 			// But it could also have another location, depending on if the slice is against a corner or wall
	// 			// ###################################
		// if !InBounds(grid, x, y) {
		if !InBounds(grid, y, x) {
			// fmt.Println("Not in bounds!")
			return false
		}

		// if grid[y][x] == rune(word[len(word)-1]) {
    //     word = StrReverse(word)
    // }

	// fmt.Println("We're on a new one ======================")
	// printWhichDirection(dirDelta)
	for _, letter := range word {
		// If it ain't inbound, it ain't gonna work
		if !InBounds(grid, x, y) {
			return false
		}

		// If it ain't the letter, it ain't the word
		if grid[x][y] != letter {
			return false
		}

		// fmt.Printf("Found %s", string(letter))
		// fmt.Printf("   x=%d, y=%d\n", x, y)

		// Check next letter, so change coordinates before we go again
		x += dirDelta[1]
		y += dirDelta[0]
	}

	return true
}

type DirectionD [2]int // { y,  x }

func CountXMAS(grid [][]rune) int {
	dirDeltas := []DirectionD{
		// left, 	right, 		up, 			down, 		upleft, 	upright, 	downleft,	downright
		{ 0, -1}, { 0,  1}, {-1,  0}, { 1,  0}, {-1, -1}, {-1,  1}, { 1, -1}, { 1,  1},
	}

	count := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			for _, dir := range dirDeltas {
				if WordSearch(grid, x, y, dir, "XMAS") {
					count++
				}
			}
		}
	}

	return count
}

func CountMASXes(grid [][]rune) int {
	dirDeltas := []DirectionD {
	// upleft, 	upright, 	downleft,	downright
		{-1, -1}, {-1,  1}, { 1, -1}, { 1,  1},
	}

	count := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			for _, dir := range dirDeltas {
				// Now check the something..
				
				// Stolen
				// I just flipped x and y in this???
				// huh??!?!
				if WordSearch(grid, y, x, dir, "MAS") {
					newY := y + 2 * dir[0]
					newDir := DirectionStolen{dir[0] * -1, dir[1]}
					if WordSearchStolen(grid, newY, x, newDir, "MAS") {
						count++
					}
				}
			}
		}
	}

	// HUH WTF???
	return count / 4
}

func MasXSearchCount(grid [][]rune) int {
    dirs := []DirectionStolen{
        {-1, -1},
        { 1,  1},
        {-1,  1},
        { 1, -1},
    }

    count := 0
    for r := 0; r < len(grid); r++ {
        for c := 0; c < len(grid[r]); c++ {
            for _, dir := range dirs {
                if WordSearchStolen(grid, r, c, dir, "MAS") {
                    nr := r + 2 * dir[0]
                    nd := DirectionStolen{dir[0] * -1, dir[1]}
                    if WordSearchStolen(grid, nr, c, nd, "MAS") {
                        count++
                    }
                }
            }
        }
    }

    return count / 4
}

func readThatInputYo() [][]rune {
	file, _ := os.Open("./input-4")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}
	
	return grid
}

func day4_2() string {
	grid := readThatInputYo()

	return fmt.Sprintf("%d", CountMASXes(grid))

}

func day4_1() string {
	grid := readThatInputYo()

	return fmt.Sprintf("%d", CountXMAS(grid))

	// for y := 0; y < len(byteArray); y++ {
	// 	for x := 0; x < len(byteArray[y]); x++ {
	// 		// fmt.Println(byteArray[y][x])
	// 		// fmt.Printf("_ = byteArray[C(%d-3, 0, %d):C(%d+3, 0, %d)][C(%d-3, 0, %d-1):C(%d+3, 0, %d-1)]", y, len(byteArray), y, len(byteArray), x, len(byteArray), x, len(byteArray))

	// 		// fmt.Println(byteArray[34])
			



	// 		// fmt.Printf("_ = byteArray[%d:%d][%d:%d]\n", C(y-3, 0, len(byteArray)), C(y+3, 0, len(byteArray)), C(x-3, 0, len(byteArray[y])-1), C(x+3, 0, len(byteArray[y])-3))
	// 		// fmt.Printf("_ = byteArray[%d:%d][%d:%d]\n", C(y-3, 0, 140), C(y+3, 0, 140), C(x-3, 0, 140-3), C(x+3, 0, 140-3))

	// 		// fmt.Printf("byteArray[%d][%d]\n", len(byteArray), len(byteArray[y]))


	// 		// fmt.Println("")
	// 		// _ = byteArray[C(y-3, 0, len(byteArray)):C(y+3, 0, len(byteArray))][C(x-3, 0, len(byteArray[y])-3):C(x+3, 0, len(byteArray[y])-3)]
	// 		if string(byteArray[y][x]) == "X" {
	// 			verticalSlice := byteArray[C(y-3, 0, len(byteArray)) : C(y+4, 0, len(byteArray))]
	// 			var subSlice [][]byte
	// 			for i := 0; i < len(verticalSlice); i++ {
	// 				subSlice = append(subSlice, verticalSlice[i][C(x-3, 0, len(verticalSlice[i])):C(x+4, 0, len(verticalSlice[i]))])
	// 			}

	// 			// Find X location in the subslice
	// 			/*
	// 			ASXMA
	// 			XSMSS
	// 			ASAMX
	// 			SMMXA <-- die X hier
	// 			XASMS
	// 			SMMAA
	// 			MAMSM
	// 			*/

	// 			// ###################################
	// 			// Subslice is now a 2d slice with the X in the middle
	// 			// .......
	// 			// .......
	// 			// .......
	// 			// ...X...
	// 			// .......
	// 			// .......
	// 			// .......
	// 			// But it could also have another location, depending on if the slice is against a corner or wall
	// 			// ###################################

	// 			// ===================
	// 			// YOU WERE HERE
	// 			// ===================
				


	// 			// for q := 0; q < len(subSlice); q++ {
	// 			// 	for p := 0; p < len(subSlice[q]); p++ {
	// 			// 		fmt.Printf("%s", string(subSlice[q][p]))
	// 			// 	}
	// 			// 	fmt.Printf("\n")
	// 			// }
	// 			// fmt.Println()
	// 		}

	// 		// _ = byteArray[34:40][131:137]
	// 		// fmt.Println()
	// 	}
	// }





	// for y := 0; y < lines; y++ {
	// 	for x := 0; x < len(maybeBetter[y]); x++ {
	// 		// fmt.Println(maybeBetter[y][x])
	// 		letter := string(maybeBetter[y][x])

	// 		// Make quadrants
	// 		topLeft := make(map[int]map[int]string)

	// 		// Check directions for making submap
	// 		if x-3 <= 0 {
	// 			// We cannot go left
	// 		}

	// 		if letter == "X" {

	// 		}

	// 	}
	// }

}

// Clamp
func C(f, lowest, highest int) int {
	if f < lowest {
		return lowest
	}
	if f > highest {
		return highest
	}
	return f
}

func printWhichDirection(dir DirectionD) {
	switch dir {
		// left, 	right, 		up, 			down, 		upleft, 	upright, 	downleft,	downright
		case DirectionD{ 0, -1}:
			fmt.Println("left")
		case DirectionD{ 0,  1}: 
			fmt.Println("right")
		case DirectionD{-1,  0}:
			fmt.Println("up")
		case DirectionD{ 1,  0}:
			fmt.Println("down")
		case DirectionD{-1, -1}:
			fmt.Println("upleft")
		case DirectionD{-1,  1}:
			fmt.Println("upright")
		case DirectionD{ 1, -1}:
			fmt.Println("downleft")
		case DirectionD{ 1,  1}:
			fmt.Println("downright")
	}
}

// func checkForXMAS(bruh *map[int]map[int]string, letter int) {
// 	// var word []byte = []byte{'X', 'M', 'A', 'S'}
// 	// val, ok := bruh[]
// }

// // littleSlice :=
// // leftValid := true
// // rightValid := true
// // upValid := true
// // downValid := true

// validDirections := []string{}
// if x != 0 {
// 	// We can go left
// 	validDirections = append(validDirections, "left")
// }
// if x != len(maybeBetter[y]) - 1 {
// 	// We can go right
// 	// rightValid = false
// 	validDirections = append(validDirections, "right")
// }
// if y != 0 {
// 	// We can go up
// 	// upValid = false
// 	validDirections = append(validDirections, "up")
// }
// if y != lines - 1 {
// 	// We can go down
// 	// downValid = false
// 	validDirections = append(validDirections, "down")
// }

// for _, dir := range validDirections {
// 	switch dir {
// 	case "left":
// 		// We're going left, check one left and call yourself again
// 		if maybeBetter[y][x - 1] == "M" {

// 		}
// 	}
// }
