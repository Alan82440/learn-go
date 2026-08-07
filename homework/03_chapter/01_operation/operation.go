package main

//算术运算

import (
	"fmt"
)

func main() {
	//用整数计算结果为整数
	fmt.Println(10 / 4)
	var num1 float32 = 10 / 4
	fmt.Println(num1)
	//若想保留小数，需要用浮点数进行计算
	fmt.Println(10.0 / 4.0)
	var num2 float32 = 10.0 / 4.0
	fmt.Println(num2)
	//a % b = a - a / b *b
	fmt.Println("10 % -3 = ", 10%-3)
	//++表示自增，--表示自减
	//自增与自减只能作为独立语句使用
	//++与--只能放在末尾，不能放前面
	var i int = 10
	i++
	fmt.Println(i)
	i--
	fmt.Println(i)
	//实例
	//97天有多少个星期剩多少天
	var day int = 97
	var week int = day / 7
	var leftday int = day % 7
	fmt.Println("week = ", week, "leftday = ", leftday)
	//华氏度转换成摄氏度
	//C = 5/9(F-32)
	var F int = 85
	var C int = 5 / 9 * (F - 32)
	fmt.Println("F = ", F, "C = ", C)
}
