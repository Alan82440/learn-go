// string转基本类型
// 字符需对应类型,否则转换失败时为默认值，如“hello”转int变0
package main

import (
	"fmt"
	"strconv"
)

func main() {
	//使用strconv.Parse函数
	var str1 string = "12345"
	var a int64
	a, _ = strconv.ParseInt(str1, 10, 64) //变量+转入类型进制+int类型
	//strconv.ParseInt该函数会转成两个值，一个为int,一个是erro
	//需要a,_忽略erro
	fmt.Printf("a type is %T,a=%d", a, a)

	var str2 string = "h"
	var b bool
	b, _ = strconv.ParseBool(str2) //变量
	//同理，需要忽略erro剩下bool
	fmt.Printf("a type is %T,a=%t", b, b)

	var str3 string = "12.345"
	var c float64
	c, _ = strconv.ParseFloat(str3, 64) //变量+float类型
	//同理
	fmt.Printf("a type is %T,a=%f", c, c)

	//注意.strconv.ParseInt与strconv.ParseFloat默认类型都为64，需手动转32
	var num int32
	num = int32(a)
	fmt.Printf("a type is %T,a=%d", num, num)
}
