package main

//逻辑运算

import (
	"fmt"
)

// 自定义函数test
func test() bool {
	//函数为bool型
	fmt.Println("true")
	//函数打印ture
	return true
	//结果返回true
} //函数输出为true

func main() {
	//&&逻辑与，两个同时满足为true
	//||逻辑或，满足其中一个为true
	//！逻辑非，取反
	var age int = 40
	if age > 30 && age < 40 {
		fmt.Println("one")
	}
	if age > 30 && age < 50 {
		fmt.Println("two")
	}
	var i int = 10
	if i > 9 && test() {
		//当前一个条件为真，继续执行后一个条件
		fmt.Println("ok1")
	}
	if i < 9 && test() {
		//当前一个条件以及错误时，后一个条件不执行
		fmt.Println("two2")
	}
	if i > 9 || test() {
		//当前一个条件为真时，后一个条件不执行
		fmt.Println("two2")
	}
	if i < 9 || test() {
		//当前一个条件以及错误时，后一个条件继续执行
		fmt.Println("two2")
	}
}
