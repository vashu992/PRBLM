package main

import "fmt"

func main() {
	var num int

	var number int
	fmt.Println("enter value")
	fmt.Scan(&num)
	num1 := []int{}


	sum := 0

	for i := 0; i < num; i++ {
		fmt.Println("enter the total number:")
		fmt.Scan(&number)
		num1 = append(num1, number)
	}

	for _, value := range num1 {
		fmt.Println("--------->",value)
		total := value * value * value
		sum += total
	}

	fmt.Println("total squre count",sum)
}