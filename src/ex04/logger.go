package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type vizit struct {
	Specialization string
	Date           time.Time
}

func main() {
	person := make(map[string][]vizit)
	scaner := bufio.NewScanner(os.Stdin)
outerLoop:
	for {
		fmt.Println("Menu:\n1. Save\n2. GetHistory\n3. GetLastVisit\n4. Exit")
		scaner.Scan()
		menu := strings.TrimSpace(scaner.Text())
		switch menu {
		case "1", "Save":
			fio := input(scaner)
			spc := input(scaner)
			dta, _ := time.Parse("2006-01-02", input(scaner))
			person[fio] = append(person[fio], vizit{
				Specialization: spc,
				Date:           dta,
			})
		case "2", "GetHistory":
			fio := input(scaner)
			posehenie, _ := person[fio]
			for _, p := range posehenie {
				fmt.Println(p.Specialization, p.Date)
			}
			fmt.Println()
		case "3", "GetLastVisit":
			fio := input(scaner)
			spc := input(scaner)
			posehenie, _ := person[fio]
			for _, p := range posehenie {
				if p.Specialization == spc {
					fmt.Println(p.Date)
				}
			}
		case "4", "Exit":
			break outerLoop
		}
		fmt.Println()
	}
	fmt.Print("Итить колотить")
}

func input(scaner *bufio.Scanner) string {
	scaner.Scan()
	return strings.TrimSpace(scaner.Text())
}
