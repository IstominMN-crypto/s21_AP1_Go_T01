package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	slice_1 := input()
	slice_2 := input()
	rez := []int{}
	rez = intersection(slice_1, slice_2)
	if len(rez) == 0 {
		fmt.Println("Empty intersection")
	} else {
		for _, v := range rez {
			fmt.Print(v, " ")
		}
		fmt.Print("\n")
	}
}

func input() []int {
	scaner := bufio.NewScanner(os.Stdin)
	scaner.Scan()
	words := strings.Fields(scaner.Text())
	rez := make([]int, 0, len(words))
	for _, w := range words {
		n, err := strconv.Atoi(w)
		if err != nil {
			fmt.Println("Invalid input")
			os.Exit(0)
		}
		rez = append(rez, n)
	}
	return rez
}

func intersection(slice_1 []int, slice_2 []int) []int {
	rez := []int{}
	used := make(map[int]bool)
	for _, v := range slice_1 {
		for _, vol := range slice_2 {
			if v == vol && !used[v] {
				rez = append(rez, vol)
				used[v] = true
			}
		}
	}
	return rez
}
