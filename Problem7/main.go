package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num string
	const (
		proceed = "proceed"
	)
	sum := 0
	for {
		fmt.Print("enter the number:")
		fmt.Scan(&num)
		num1, _:= strconv.Atoi(num)
		if num == proceed {
			break
		}
		sum += num1
	}
	fmt.Println("sum:",sum)
}