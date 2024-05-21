package main

import "fmt"

func main() {
	var number int
	var number1 int
	var breakValue int
	array := []int{}
	fmt.Print("enter number :")
	fmt.Scan(&number)
	for i := 0; i < number; i++ {
		fmt.Print("enter total number :")
		fmt.Scan(&number1)
		array = append(array, number1)

	}
	
	fmt.Print("enter the break value :")
	fmt.Scan(&breakValue)
	for _, value := range array {
		if value <= breakValue {
			fmt.Println("minmum number :", value)
		}
	}

}
