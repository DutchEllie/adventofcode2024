package main

import (
    "fmt"
    "os"
    "strings"
)

func InBoundsStolen[T any](v []T, i int) bool {
    return i >= 0 && i < len(v)
}

func InBounds2DStolen[T any](v [][]T, i, j int) bool {
    return i >= 0 && i < len(v) && j >= 0 && j < len(v[i])
}

func StrReverseStolen(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

type DirectionStolen [2]int

func WordSearchStolen(grid [][]rune, r, c int, dir DirectionStolen, word string) bool {
    if !InBounds2DStolen(grid, r, c) {
        return false
    }

    if grid[r][c] == rune(word[len(word)-1]) {
        word = StrReverseStolen(word)
    }

    for _, char := range word {
        if !InBounds2DStolen(grid, r, c) || grid[r][c] != char {
            return false
        }

        r += dir[0]
        c += dir[1]
    }

    return true
}

func XmasSearchCountStolen(grid [][]rune) int {
    dirs := []DirectionStolen{
        {-1, -1}, { 0, -1}, {-1,  0},
        { 1,  1}, { 0,  1}, { 1,  0},
        {-1,  1},
        { 1, -1},
    }

    count := 0
    for r := 0; r < len(grid); r++ {
        for c := 0; c < len(grid[r]); c++ {
            for _, dir := range dirs {
                if s := WordSearchStolen(grid, r, c, dir, "XMAS"); s {
                    count++
                }
            }
        }
    }

    return count / 2
}

func MasXSearchCountStolen(grid [][]rune) int {
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

func Day04Stolen() (int64, int64) {
    in, err := os.ReadFile("input-4")
    if err != nil {
        fmt.Println(err)
        return -1, -1
    }

    lines := strings.Split(string(in), "\n")
    grid := make([][]rune, len(lines))
    for i, line := range lines {
        grid[i] = []rune(line)
    }

    return int64(XmasSearchCountStolen(grid)), int64(MasXSearchCountStolen(grid)) 
}

func main() {
	ans1, ans2 := Day04Stolen()

	fmt.Println("Stolen:")
 	fmt.Printf("Day 4 - 1: %d\n", ans1)
 	fmt.Printf("Day 4 - 2: %d\n", ans2)
	fmt.Println("Mine (also stolen):")
 	fmt.Printf("Day 4 - 1: %s\n", day4_1())
 	fmt.Printf("Day 4 - 2: %s (INCORRECT I AM DUMB)\n", day4_2())
}