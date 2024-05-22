package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num string
	var operation string
	odd := 0
	even := 0
	const (
		proceed = "proceed"
	)
	for {
		fmt.Print("enter the number :")
		fmt.Scan(&num)
		if num == proceed {
			break
		}
		num1, _ := strconv.Atoi(num)
		if num1%2 != 0 {
			odd = num1
		}
		if num1%2 == 0 {
			even = num1
		}
	}
	fmt.Print("enter the operation :")
	fmt.Scan(&operation)
	switch operation {
	case "countodd":
		fmt.Println("odd :", odd)
	case "counteven":
		fmt.Println("even :", even)
	}
}
