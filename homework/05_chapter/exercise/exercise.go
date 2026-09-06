package main

import (
	"fmt"
)

func main() {
	//水仙花数题目
	// var data, num, n int
	// var sum int
	// fmt.Printf("请输入一个数字")
	// fmt.Scan(&num)
	// temp := num //初始化temp
	// for temp > 0 {
	// 	temp = temp / 10 //通过除以10累加n可得出这个数的位数
	// 	n++
	// }
	// sum = 0
	// temp = num //为防止上述temp干扰，再次初始化
	// for temp > 0 {
	// 	data = temp % 10                                //通过除以10求余数可得出每一位的数字
	// 	sum += int(math.Pow(float64(data), float64(n))) //累加每一位位数的n次方的和
	// 	//math.Pow(a,b)表示a的b次方，需要括号内都是float64类型
	// 	temp = temp / 10 //通过不断除以10减少已经计算的位数，并进入循环计算下一位
	// }
	// if sum == num {
	// 	fmt.Printf("%v是水仙花数", num)
	// } else {
	// 	fmt.Printf("%v不是水仙花数", num)
	// }

	// var month, year int
	// fmt.Printf("请输入年份,月份")
	// fmt.Scan(&year, &month)
	// switch month {
	// case 1, 3, 5, 7, 8, 10, 12:
	// 	fmt.Printf("%v月份有31天", month)
	// case 4, 6, 9, 11:
	// 	fmt.Printf("%v月份有30天", month)
	// case 2:
	// 	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
	// 		fmt.Printf("%v月份有29天", month)
	// 	} else {
	// 		fmt.Printf("%v月份有28天", month)
	// 	}
	// default:
	// 	fmt.Printf("输入错误")
	// }

	var a, b, c, temp int
	fmt.Printf("请输入3个数字")
	fmt.Scan(&a, &b, &c)
	if a > b {
		temp = a
		b = a
		a = temp
	}
	if a > c {
		temp = a
		a = c
		c = temp
	}
	if b > c {
		temp = b
		b = c
		c = temp
	}
	fmt.Printf("%v<%v<%v", a, b, c)
}
