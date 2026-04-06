package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type PatientNotFoundError struct{}

type massage interface {
	Ошибка() string
}

type vizit struct {
	Specialization string
	Date           time.Time
}

func main() {
	var u vizit
	u = vizit{
		Specialization: "okulist",
		Date:           time.Now(),
	}

	fmt.Println(u)
	person := make(map[string][]vizit)
	scaner := bufio.NewScanner(os.Stdin)
	var m bool = true
	for {
		if m {
			fmt.Println("Menu:\n1. Save\n2. GetHistory\n3. GetLastVisit\n4. Menu\n5. Exit")
		}
		scaner.Scan()
		menu := strings.TrimSpace(scaner.Text())
		switch menu {
		case "1", "Save":
			fio := input(scaner)
			spc := input(scaner)
			dta, _ := time.Parse("2006-01-22", input(scaner))
			person[fio] = append(person[fio], vizit{
				Specialization: spc,
				Date:           dta,
			})
		case "2", "GetHistory":
			fio := input(scaner)
			posehenie, err := findFio(person, fio)
			if err != nil {
				fmt.Println(err.Ошибка())
			} else {
				for _, p := range posehenie {
					fmt.Println(p.Specialization, p.Date.Format("2006-01-02"))
				}
			}
		case "3", "GetLastVisit":
			fio := input(scaner)
			posehenie, err := findFio(person, fio)
			if err != nil {
				fmt.Println(err.Ошибка())
			} else {
				spc := input(scaner)
				for _, p := range posehenie {
					if p.Specialization == spc {
						fmt.Println(p.Date.Format("2006-01-02"))
					}
				}
			}
		case "4", "Menu":
			m = !m
		case "5", "Exit":
			os.Exit(0)
		default:
			fmt.Println("Niht command! Повторите ввод")
		}
		fmt.Println()
	}
}

func input(scaner *bufio.Scanner) string {
	scaner.Scan()
	return strings.TrimSpace(scaner.Text())
}

func findFio(person map[string][]vizit, fio string) ([]vizit, massage) {
	posehenie, ok := person[fio]
	if !ok {
		return nil, PatientNotFoundError{}
	} else {
		return posehenie, nil
	}
}

// func (e PatientNotFoundError) Error() string {
// 	return "patient not found"
// }

func (e PatientNotFoundError) Ошибка() string {
	return "заебало"
}
