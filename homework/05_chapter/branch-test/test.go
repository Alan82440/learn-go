package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, x1, x2 float64
	var m float64
	fmt.Println("ax^2+bx+c=?")
	fmt.Println("请输入a,b,c的值")
	fmt.Scan(&a, &b, &c)
	m = b*b - 4*a*c
	if a == 0 {
		fmt.Println("不是一元二次方程")
	}
	if m > 0 {
		x1 = (-b + math.Sqrt(m)) / 2
		x2 = (-b - math.Sqrt(m)) / 2
		fmt.Println("x1=", x1, "x2=", x2)
	} else if m == 0 {
		x1 = (-b + math.Sqrt(m)) / 2
		x2 = x1
		fmt.Println("x1=", x1, "x2=", x2)
	} else {
		fmt.Println("方程无解")
	}
}
