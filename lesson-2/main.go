package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	var a, b, res float64
	var op string
	var operator string

	fmt.Println("Введите первое число")
	_, err := fmt.Scan(&a)
	if err == nil {

	} else {
		fmt.Println("Произошла ошибка: ", err)
	}

	fmt.Println("Введите второе число")
	_, err = fmt.Scan(&b)
	if err == nil {

	} else {
		fmt.Println("Произошла ошибка: ", err)
	}

	fmt.Println("Введите оператор")
	_, err = fmt.Scan(&op)
	if err == nil {

	} else {
		fmt.Println("Произошла ошибка: ", err)
	}

	switch op {
	case "+":
		res = a + b
		operator = "Сложение"
	case "-":
		res = a - b
		operator = "Вычитания"
	case "*":
		res = a * b
		operator = "Умножения"
	case "/":
		res = a / b
		operator = "Деления"
	case "%":
		res = a / 100 * b
		operator = "Процент"
	case "^":
		res = math.Pow(a, b)
		operator = "Возведения в степень"
	default:
		fmt.Println("Операция выбрана неверно")
		os.Exit(1)
	}
	fmt.Println("Вы выбрали оператор:", operator)
	fmt.Printf("Результат выполнения операции: %.0f\n", res)
}