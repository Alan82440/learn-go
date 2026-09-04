package main

import (
	"fmt"
)

func main() {
	// 	var second float64
	// 	fmt.Printf("请输入成绩")
	// 	fmt.Scan(&second)
	// 	if second <= 10 {
	// 		var gender string
	// 		fmt.Printf("请输入性别")
	// 		fmt.Scan(&gender)
	// 		if gender == "男" {
	// 			fmt.Printf("进入男子组")
	// 		} else {
	// 			fmt.Printf("进入女子组")
	// 		}
	// 	} else {
	// 		fmt.Printf("out...")
	// 	}
	var season, age int
	var price float64 = 60
	fmt.Printf("请输入月份")
	fmt.Scan(&season)
	fmt.Printf("请输入年龄")
	fmt.Scan(&age)
	if season > 4 && season < 10 {
		if age > 60 {
			fmt.Printf("price is %v", price/2)
		} else if age >= 18 && age <= 60 {
			fmt.Printf("price is %v", price)
		} else {
			fmt.Printf("price is %v", price/3)
		}
	} else {
		if age >= 18 {
			fmt.Printf("price is %v", price)
		} else {
			fmt.Printf("price is %v", price/2)
		}
	}
}
