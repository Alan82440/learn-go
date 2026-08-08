package main

import (
	"fmt"
)

func main() {
	var age int
	fmt.Println("输入年龄")
	fmt.Scan(&age)
	//单分支
	if age >= 18 {
		fmt.Println("已成年")
	}
	fmt.Printf("\n")
	//双分支
	if age >= 18 {
		fmt.Println("已成年")
	} else {
		fmt.Println("未成年")
	}
	fmt.Printf("\n")
	var n1 int
	var n2 int
	fmt.Println("请输入n1,n2,用空格隔开")
	fmt.Scan(&n1, &n2)
	//输入前加一个空格，避免上一个人回车残留导致跳过代码
	if n1+n2 >= 50 {
		fmt.Println("hello world")
	}
	fmt.Printf("\n")
	var n3 float64
	var n4 float64
	fmt.Println("请输入数字")
	fmt.Scan(&n3, &n4)
	if (n3+n4 > 10.0) && (n3+n4 < 20.0) {
		fmt.Println("两数之和为", n3+n4)
	}
	fmt.Printf("\n")
	var year int
	fmt.Println("输入年份")
	fmt.Scan(&year)
	if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
		fmt.Println("该年时闰年")
	}
}
