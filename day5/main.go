package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Rule struct {
	first int
	second int
}

func NewRule(r string) Rule {
	tmp := strings.Split(r, "|")
	first, _ := strconv.Atoi(tmp[0])
	second, _ := strconv.Atoi(tmp[1])
	rule := Rule{first, second}

	return rule
}

func (r *Rule) IsSorted(l []int) bool {
	firstIndex := slices.Index(l, r.first)
	secondIndex := slices.Index(l, r.second)

	if firstIndex == -1 || secondIndex == -1 {
		// fmt.Println("Doesn't contain rule nums")
		return true
	}

	if firstIndex < secondIndex {
		// Everything is fine
		// fmt.Println("Already sorted")
		return true
	}

	// fmt.Println("not sorted")

	return false
}

func (r *Rule) CorrectList(l []int) []int {
	// Returns the list corrected according to the rule

	// // Check if the list even contains the numbers we're the rule of
	// if !(slices.Contains(l.list, r.first) && slices.Contains(l.list, r.second)) {
	// 	// The list doesn't contain our rule numbers
	// 	return l.list
	// }

	// Now check if it's already in order
	firstIndex := slices.Index(l, r.first)
	secondIndex := slices.Index(l, r.second)

	if firstIndex == -1 || secondIndex == -1 {
		// fmt.Println("Doesn't contain rule nums")
		return l
	}

	if firstIndex < secondIndex {
		// Everything is fine
		// fmt.Println("Already sorted")
		return l
	}

	// List is not in order
	// for i := 0; i < len(l.list); i++ {
	// fmt.Println(r.first)
	// fmt.Println(r.second)
	// fmt.Println("trying to delete the first from the list")
	listWithoutFirst := append(l[:firstIndex], l[firstIndex+1:]...)
	// listWithFirstAddedBackIn := append(listWithoutFirst[:secondIndex], r.first)
	// final := append(listWithFirstAddedBackIn, listWithoutFirst[secondIndex+1:]...)
	// fmt.Println("trying to add first back in at the end")
	final := append(listWithoutFirst, r.first)

	// fmt.Println(l.list)
	// fmt.Println(final)
		
	// }
	
	
	return final
}

type List struct {
	list []int
}

func NewList(l string) List {
	tmp := strings.Split(l, ",")
	list := make([]int, 0)
	for _, val := range tmp {
		i, _ := strconv.Atoi(val)
		list = append(list, i)
	}

	return List{list}
}

// func contains(l []int, num int) bool {
// 	for _, a := range e
// }

// Check if c is part of s in order
// func subslice(s []int, c []int) bool {
// 	if len(c) > len(s) { return false }

// 	for _, e := range s {
// 	}
// }

func (l *List) IsSorted(rules []Rule) bool {
	// fmt.Println("Checking a list")
	tmp := l.list
	for _, r := range rules {
		s := r.IsSorted(tmp)
		if s {
			// fmt.Printf("List is sorted: true\n")
		} else {
			// fmt.Printf("List is sorted: false\n")
			return false
		}
	}

	// fmt.Println("Returning a true")
	return true
}

func CheckAndCorrect(list []int, rules []Rule) []int {
	// Iterate all the rules and check them against the list
	// tmp := make([]int, 0)
	// for _, r := range rules {
	// 	// fmt.Println(tmp)
	// 	tmp := r.CorrectList(list)
		
	// 	// fmt.Println(tmp)
	// }

	if len(rules) == 1 {
		tmp := rules[0].CorrectList(list)
		return tmp
	}

	rule, theRest := rules[0], rules[1:]
	tmp := rule.CorrectList(list)

	// tmp2 := List{tmp}
	return CheckAndCorrect(tmp, theRest)
} 

func part1() (p1, p2 int) {
	file, _ := os.Open("./input-5")
	defer file.Close()

	full := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		full = append(full, scanner.Text())
	}

	// First is the rules, then a newline, then the lists, ended with a newline
	stringed := strings.Join(full, "\n")
	sections := strings.Split(stringed, "\n\n")

	rulesString := sections[0]
	listsString := sections[1]

	rulesSlice := strings.Split(rulesString, "\n")
	listsSlice := strings.Split(listsString, "\n")

	rules := make([]Rule, 0)
	for _, r := range rulesSlice {
		rules = append(rules, NewRule(r))
	}
	lists := make([]List, 0)
	for _, l := range listsSlice {
		lists = append(lists, NewList(l))
	}

	// Now we have the lists and rules
	// Check every single list and sort it
	// for _, list := range lists {
	// 	l := CheckAndCorrect(list.list, rules)
	// 	list.list = l
	// }

	bruh := make([]List, 0)
	for _, list := range lists {
		listissortedalready := list.IsSorted(rules)
		if listissortedalready {
			// fmt.Println("Added list to bruh")
			bruh = append(bruh, list)
		}
	}

	// fmt.Printf("Original list length: %d\n", len(lists))
	// fmt.Printf("Correct list length: %d\n", len(bruh))

	p1 = 0
	for _, list := range bruh {
		// Index is always uneven, hopefully
		// If 0 ... 4, there are 5 elements
		// len() will then be 5
		// 5 / 2 = 2.5, floor that and you get 2
		// 0 1 2 3 4
		//     ^
		//   Middle
		// 
		// Thank you for the tip, Tim!

		middle := len(list.list) / 2
		mid := list.list[middle]
		p1 += mid
	}

	p2 = 0
	incorrectLists := make([]List, 0)
	for _, l := range lists {
		if !l.IsSorted(rules) {
			incorrectLists = append(incorrectLists, l)
		}
	}

	fmt.Printf("len(incorrectLists): %v\n", len(incorrectLists))
	
	fmt.Printf("(len(incorrectLists) + len(bruh)): %v\n", (len(incorrectLists) + len(bruh)))
	fmt.Printf("len(lists): %v\n", len(lists))

	correctedIncorrectLists := make([]List, 0)
	for _, l := range incorrectLists {
		// list := CheckAndCorrect(l.list, rules)
		
		// debug
		// checkaaah := func(list []int, rules []Rule) []int {
		// 	tmp := CheckAndCorrect(list, rules)
		// 	l := 
		// }
		// isSortedDebug := func(list *[]int, rules []Rule) bool {
		// 	l := List{*list}
		// 	return l.IsSorted(rules)
		// }

		// yeet := sort(l.list, rules)
		yeet := l.list
		fmt.Printf("yeet: %v\n", yeet)

		slices.SortFunc(yeet, func(i, j int) int {
			for _, r := range rules {
				if r.first == i && r.second == j { 
					return -1 
				} else if r.first == j && r.second == i {
					return 1
				}
			}

			return 0


			firstIndex := slices.Index(l.list, i)
			secondIndex := slices.Index(l.list, j)

			if firstIndex == -1 || secondIndex == -1 {
				// fmt.Println("Doesn't contain rule nums")
				return 0
			}

			if firstIndex < secondIndex {
				// Everything is fine
				// fmt.Println("Already sorted")
				return -1
			}

			// fmt.Println("not sorted")

			return 1
		})

		fmt.Printf("yeet: %v\n", yeet)


		// fmt.Println(l.list)
		// fmt.Println(list)
		// fmt.Println("had a good list this time")
		correctedIncorrectLists = append(correctedIncorrectLists, List{yeet})
	}

	fmt.Printf("Length of correctedIncorrectLists: %d\n", len(correctedIncorrectLists))

	for _, testList := range correctedIncorrectLists {
		if !testList.IsSorted(rules) {
			fmt.Println("Not sorted!")
		}
	}

	for _, l := range correctedIncorrectLists {
		middle := len(l.list) / 2
		mid := l.list[middle]
		p2 += mid
	}

	return
}

func sort(list []int, rules []Rule) []int {
	tmp := CheckAndCorrect(list, rules)
	l := List{tmp}
	if !l.IsSorted(rules) {
		sort(tmp, rules)
	}

	return tmp
}


func main() {
	p1, p2 := part1()
	fmt.Printf("Day 5 - 1: %d\nDay 5 - 2: %d\n", p1, p2)

}