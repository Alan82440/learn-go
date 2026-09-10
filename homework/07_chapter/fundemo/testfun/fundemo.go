package main

import (
	"fmt"
	"learn-go/homework/07_chapter/fundemo/eng"
)

func main() {
	// 	n := 20
	// 	if n > 10 {
	// 		goto label //goto跳转语句，不建议使用
	// 	}
	// 	fmt.Printf("ok1")
	// label:
	// 	fmt.Printf("ok2")

	//加减乘除的计算器
	var x1, x2, result float64
	var operation string
	fmt.Printf("请输入数字")
	fmt.Scan(&x1, &x2)
	fmt.Printf("请输入运算符")
	fmt.Scan(&operation)
	// result = eng.Cal(x1, x2, operation) //调用函数，将x1,x2,operation的值传给对应形参，用result接受res
	result = eng.Cal(x1, x2, operation) // 调用其他包的函数，包名.函数名(实参)
	fmt.Printf("result=%v", result)
}
