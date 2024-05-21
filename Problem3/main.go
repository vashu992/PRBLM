package main

import "fmt"

func main() {
	var num int
	sum := 0

	for i := 0; i <= 1; i++ {
		fmt.Print("enter the two number:")
		fmt.Scan(&num)
		sum = sum + num
		if sum < 0 {
			fmt.Println("pls enter valid number")
			return
		}
	}
	fmt.Println("sum", sum)
}
