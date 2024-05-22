package main

import "fmt"

func main() {
	var num int
	var num1 int
	var operation string
	var result int
	fmt.Print("enter first number :")
	fmt.Scan(&num)
	fmt.Print("enter second number :")
	fmt.Scan(&num1)
	fmt.Print("enter the operation :")
	fmt.Scan(&operation)
	switch operation {
	case "+":
		result = num + num1
		fmt.Println("sum :", result)
	case "*":
		result = num * num1
		fmt.Println("multiplication :", result)
	case "-":
		result = num - num1
		fmt.Println("subtract :", result)
	case "/":
		if num1 == 0 {
			fmt.Println("enter a valid number")
			return

		}
		result = num / num1

		fmt.Println("divide", result)
	default:
		for {
			fmt.Println("pls enter valid operation")
		}
	}

}
