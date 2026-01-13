package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	slice_1, ok_1 := input()
	if ok_1 == false {
		return
	}
	slice_2, ok_2 := input()
	rez := []int{}
	if ok_1 && ok_2 {
		rez = inter(slice_1, slice_2)
		if len(rez) == 0 {
			fmt.Println("Empty intersection")
		} else {
			for _, v := range rez {
				fmt.Print(v, " ")
			}
			fmt.Print("\n")
		}
	}
}

func input() ([]int, bool) {
	temp := bufio.NewScanner(os.Stdin)
	temp.Scan()
	words := strings.Fields(temp.Text())
	Zzz := true
	rez := make([]int, 0, len(words))
	for _, w := range words {
		n, err := strconv.Atoi(w)
		if err != nil {
			fmt.Println("Invalid input")
			Zzz = false
		}
		rez = append(rez, n)
	}
	return rez, Zzz
}

func inter(slice_1 []int, slice_2 []int) []int {
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
