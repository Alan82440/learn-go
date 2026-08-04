package main

import "fmt"

func main() {
	//T(v)类型将变量v值转化为T类型
	//转化范围可从大到小，也可从小到大
	var i int32 = 100
	var j float32 = float32(i)
	var k int8 = int8(i) //高精度与低精度之间不能直接转换
	fmt.Printf("i=%v,j=%v,k=%v", i, j, k)
	fmt.Printf("i type is %T", i) //转化只是转化本身数据(值)，数据类型不会改变
	//当从大到小的数据超出范围时，编译不出错，但数据溢出会改变结果
	var num1 int64 = 999
	var num2 int8 = int8(num1)
	fmt.Println(num2)
	//只用同类型才能进行计算
	var num3 int32 = 12
	var num4 int8
	num4 = int8(num3) + 100
	//若相加结果超过int8范围，编译能过，但结果出现溢出错误
	//若相加的数字超出int8范围(>127),则编译不通过，如num4 = int8(num3) + 128
	fmt.Println(num4)
}
