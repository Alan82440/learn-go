package main

import (
	"fmt"
)

func main() {
	// var i int = 0
	// for ; i <= 100; i++ {
	// 	if i%2 == 0 {
	// 		continue  //可配合label进入跳出对应循环
	// 	}
	// 	fmt.Printf("100内的奇数%v\n", i)
	// }

	// var i int = 0
	// var num int
	// var count1 int = 0
	// var count2 int = 0
	// for ; i <= 100; i++ {
	// 	fmt.Printf("请输入一个数\n")
	// 	fmt.Scan(&num)
	// 	if num > 0 {
	// 		count1++
	// 	} else if num < 0 {
	// 		count2++
	// 	} else if num == 0 {
	// 		break
	// 	}
	// }
	// fmt.Printf("正整数有%v个,负整数有%v个\n", count1, count2)

	var money float64 = 100000 //总现金
	var num int = 0            //路口
	for {
		if money > 50000 { //条件一
			money = money - money*0.05
			num++
		} else if money <= 50000 && money >= 1000 { //条件二且不小于1000
			money -= 1000
			num++
		}
		if money < 1000 { //退出循环
			break
		}
	}
	fmt.Printf("路口有%v,剩余%v", num, money)
}
