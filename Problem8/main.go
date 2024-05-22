package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var num string
	var operation string
	const (
		proceed = "proceed"
	)
	count := 0
	sum := 0
	min := math.MaxInt64
	max := math.MinInt64

	for {
		fmt.Print("enter the number :")
		fmt.Scan(&num)

		if num == proceed {
			break

		}
		num1, _ := strconv.Atoi(num)

		count = count + 1
		sum += num1
		if num1 < min {
			min = num1

		}
		if num1 > max {
			max = num1
			fmt.Println(".........>", min)

		}

	}

	fmt.Print("enter the operation :")
	fmt.Scan(&operation)
	switch operation {

	case "count":
		fmt.Println("count of number :", count)
	case "mean":
		fmt.Println("-------->", sum)
		result := sum / count
		fmt.Println("mean number :", result)
	case "min":
		// if num1 > 0 {
		fmt.Println("min :", min)
	// }
	case "max":
		fmt.Println("max :", max)

	}
}
