package main

import "fmt"

func Calculator(num1 int, num2 int, CMD func(int, int) int) int {
	return CMD(num1, num2)
}

func Add(num1 int, num2 int) int {
	return num1 + num2
}

func Minus(num1 int, num2 int) int {
	return num1 - num2
}

func Multiply(num1 int, num2 int) int {
	return num1 * num2
}

func Divide(num1 int, num2 int) int {
	if num2 == 0 {
		print("除数不能为零")
	}
	return num1 / num2
}

func main() {
	num1 := 4
	num2 := 2

	AddResult := Calculator(num1, num2, Add)
	fmt.Println("加法运算结果为", AddResult)

	MinusResult := Calculator(num1, num2, Minus)
	fmt.Println("减法运算结果为", MinusResult)

	MultiplyResult := Calculator(num1, num2, Multiply)
	fmt.Println("乘法运算结果为", MultiplyResult)

	DivideResult := Calculator(num1, num2, Divide)
	fmt.Println("除法运算结果为", DivideResult)
}
