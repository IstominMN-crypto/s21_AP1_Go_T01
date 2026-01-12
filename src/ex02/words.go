package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	var k int
	temp := bufio.NewScanner(os.Stdin)
	temp.Scan()
	str := temp.Text()
	for {
		if _, err := fmt.Scan(&k); err != nil {
			fmt.Println("Неверное число, повторите ввод")
			continue
		}
		break
	}

	for _, c := range TopWords(str, k) {
		fmt.Printf("%s ", c)
	}
	fmt.Print("\n")
}

func TopWords(str string, k int) []string {

	s := map[string]int{}
	words := strings.Fields(str)
	for _, word := range words {
		if val, ok := s[word]; ok {
			s[word] = val + 1
		} else {
			s[word] = 1
		}
	}
	keys := make([]string, 0, len(s))
	for c := range s {
		keys = append(keys, c)
	}
	sort.Slice(keys, func(i, j int) bool {
		if s[keys[i]] != s[keys[j]] {
			return s[keys[i]] > s[keys[j]]
		}
		return keys[i] < keys[j]
	})
	if k > len(keys) {
		k = len(keys)
	}
	return keys[:k]
}
