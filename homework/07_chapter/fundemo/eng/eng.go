package eng

import (
	"fmt"
)

// 自定义函数
// func+函数名+形参+返回值类型
// 函数名大写公有，小写私有，公有可被其他包调用
func Cal(n1 float64, n2 float64, operation string) float64 {
	var res float64
	switch operation {
	case "+":
		res = n1 + n2
	case "-":
		res = n1 - n2
	case "*":
		res = n1 * n2
	case "/":
		res = n1 / n2
	default:
		fmt.Printf("非法运算\n")
	}
	return res //结果返回给res
}
