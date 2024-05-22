package main

import "fmt"

func main() {
	var num int
	sum := 0

	fmt.Print("Enter the Lenth of Number")
	fmt.Scan(&num)
	num1 := []int{}

	for i:= 0; i < num; i++ {
		fmt.Print("Enter the Total Number:")
		var num2 int
		fmt.Print(&num2)
		num1 = append(num1, num2)
	}
	fmt.Println(sum)
}