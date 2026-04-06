package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	a := скан_число("Введите первый операнд:")
	операция := скан_операция("Введите операцию (+, -, *, /):")
	b := скан_число("Введите второй операнд:")
	switch операция {
	case "+":
		fmt.Println("Ответ: ", a+b)
	case "-":
		fmt.Println("Ответ: ", a-b)
	case "*":
		fmt.Println("Ответ: ", a*b)
	case "/":
		if b == 0 {
			fmt.Println("Ошибка: деление на 0")
			return
		}
		fmt.Printf("Ответ: %.3g\n", a/b)
	}
}

func скан_операция(сообщение string) (in string) {
	for {
		fmt.Println(сообщение)
		fmt.Scanln(&in)
		if in == "+" || in == "-" || in == "*" || in == "/" {
			return
		}
		fmt.Println("Неверное значение")
	}
}

func скан_число(сообщение string) float64 {
	for {
		fmt.Println(сообщение)
		scaner := bufio.NewScanner(os.Stdin)
		scaner.Scan()
		in := scaner.Text()
		parts := strings.Fields(in)
		f, err := strconv.ParseFloat(parts[0], 64)
		if len(parts) > 1 || err != nil {
			fmt.Println("Неверное значение")
			continue
		}

		return f
	}
}
