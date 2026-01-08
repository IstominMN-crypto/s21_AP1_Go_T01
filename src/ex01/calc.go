package main

import (
	"fmt"
)

func main() {
	a := скан_число("Введите первый операнд:")
	операция := скан_операция("Введите операцию (+, -, *, /):")
	b := скан_число("Введите второй операнд:")
	fmt.Print("Ответ: ")
	switch операция {
	case "+":
		fmt.Println(a + b)
	case "-":
		fmt.Println(a - b)
	case "*":
		fmt.Println(a * b)
	case "/":
		if b == 0 {
			fmt.Println("Ошибка: деление на 0")
			return
		}
		fmt.Printf("%.3f\n", a/b)
	}
}

func скан_операция(сообщение string) string {
	for {
		fmt.Println(сообщение)
		var in string
		fmt.Scanln(&in)
		if in == "+" || in == "-" || in == "*" || in == "/" {
			return in
		}
		fmt.Println("Неверное значение")
	}
}

func скан_число(сообщение string) float64 {
	for {
		fmt.Println(сообщение)
		var in string
		fmt.Scanln(&in)
		var value float64
		_, err := fmt.Sscan(in, &value)
		if err != nil {
			fmt.Println("Неверное значение")
			continue
		}
		return value
	}
}
