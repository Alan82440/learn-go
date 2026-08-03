package main

import "fmt"

func main() {
	//浮点数=符号+指数位+小数位
	var f1 float32 = 3.14
	var f2 float32 = -0.09
	fmt.Println("f1=", f1, "f2=", f2)
	//浮点可能有精度损失
	var f3 float32 = 123.0000901
	var f4 float64 = 123.0000901
	fmt.Println("f3=", f3, "f4=", f4)
	//浮点数的默认类型为float64
	//十进制浮点数必须带小数点，科学计数法可以不带小数点
	var f5 = 0.123
	var f6 = .123 //=> var f6 = 0.123
	fmt.Println("f5=", f5, "f6=", f6)
	//科学计数法表示浮点数
	var f7 = 1.23e2  //=> 1.23*10的2次方
	var f8 = 1.23e-2 //=> 1.23*10的-2次方
	fmt.Println("f7=", f7, "f8=", f8)
}
