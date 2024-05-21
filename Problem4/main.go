package main

import "fmt"

func main() {
	var num int
	var num1 int
	var operation string
	var result int
	fmt.Print("enter first number:")
	fmt.Scan(&num)
	fmt.Print("enter second number:")
	fmt.Print(&num1)
	fmt.Print("enter the operation:")
	fmt.Scan(&operation)
	if operation == "+" {
		result = num + num1
		fmt.Println("sum of number:", result)

	} else if operation == "*" {
		result = num * num1
		fmt.Println("multiplication of number:", result)
	} else if operation == "-" {
		result = num - num1
		fmt.Println("subtract of number:", result)
	} else if operation == "/" {
		result = num / num1
		fmt.Println("divide of number:", result)
	}

}
