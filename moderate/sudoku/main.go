package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var SIZE_4_GROUPS = [][]int{

	[]int{0, 1, 2, 3},
	[]int{4, 5, 6, 7},
	[]int{8, 9, 10, 11},
	[]int{12, 13, 14, 15},
	[]int{0, 4, 8, 12},
	[]int{1, 5, 9, 13},
	[]int{2, 6, 10, 14},
	[]int{3, 7, 11, 15},
	[]int{0, 1, 4, 5},
	[]int{2, 3, 6, 7},
	[]int{8, 9, 12, 13},
	[]int{10, 11, 14, 15},
}

var SIZE_9_GROUPS = [][]int{
	[]int{0, 1, 2, 3, 4, 5, 6, 7, 8},
	[]int{9, 10, 11, 12, 13, 14, 15, 16, 17},
	[]int{18, 19, 20, 21, 22, 23, 24, 25, 26},
	[]int{27, 28, 29, 30, 31, 32, 33, 34, 35},
	[]int{36, 37, 38, 39, 40, 41, 42, 43, 44},
	[]int{45, 46, 47, 48, 49, 50, 51, 52, 53},
	[]int{54, 55, 56, 57, 58, 59, 60, 61, 62},
	[]int{63, 64, 65, 66, 67, 68, 69, 70, 71},
	[]int{72, 73, 74, 75, 76, 77, 78, 79, 80},
	[]int{0, 9, 18, 27, 36, 45, 54, 63, 72},
	[]int{1, 10, 19, 28, 37, 46, 55, 64, 73},
	[]int{2, 11, 20, 29, 38, 47, 56, 65, 74},
	[]int{3, 12, 21, 30, 39, 48, 57, 66, 75},
	[]int{4, 13, 22, 31, 40, 49, 58, 67, 76},
	[]int{5, 14, 23, 32, 41, 50, 59, 68, 77},
	[]int{6, 15, 24, 33, 42, 51, 60, 69, 78},
	[]int{7, 16, 25, 34, 43, 52, 61, 70, 79},
	[]int{8, 17, 26, 35, 44, 53, 62, 71, 80},
	[]int{0, 1, 2, 9, 10, 11, 18, 19, 20},
	[]int{3, 4, 5, 12, 13, 14, 21, 22, 23},
	[]int{6, 7, 8, 15, 16, 17, 24, 25, 26},
	[]int{27, 28, 29, 36, 37, 38, 45, 46, 47},
	[]int{30, 31, 32, 39, 40, 41, 48, 49, 50},
	[]int{33, 34, 35, 42, 43, 44, 51, 52, 53},
	[]int{54, 55, 56, 63, 64, 65, 72, 73, 74},
	[]int{57, 58, 59, 66, 67, 68, 75, 76, 77},
	[]int{60, 61, 62, 69, 70, 71, 78, 79, 80},
}

func getGroups(size int) (groups [][]int) {
	if size == 4 {
		groups = SIZE_4_GROUPS
	} else if size == 9 {
		groups = SIZE_9_GROUPS
	}
	return
}

func IntEquals(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func isValidSudoku(size int, digits []int) bool {
	groups := getGroups(size)

	var elements = make([]int, size)
	for i := range elements {
		elements[i] = i + 1
	}
	for _, group := range groups {
		values := make([]int, size)

		for i, idx := range group {
			values[i] = digits[idx]
		}
		sort.Ints(values)
		if !IntEquals(values, elements) {
			return false
		}
	}

	return true
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		split_str := strings.Split(scanner.Text(), ";")
		size, _ := strconv.Atoi(split_str[0])
		digits := strings.Split(split_str[1], ",")
		var numdigits = make([]int, len(digits))
		for i := 0; i < len(digits); i++ {
			numdigits[i], _ = strconv.Atoi(digits[i])
		}

		if isValidSudoku(size, numdigits) {
			fmt.Println("True")
		} else {
			fmt.Println("False")
		}

	}
}
