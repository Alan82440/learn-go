// 基本类型转string
package main

import (
	"fmt"
	"strconv"
)

func main() {
	//fmt.Sprintf转换法(建议)
	var a int = 99
	var b float64 = 23.456
	var c bool = true
	var d byte = 'h'
	var str string
	//目标变量 = fmt.Sprintf("原变量输出形式","原变量")
	str = fmt.Sprintf("%d", a) //%d表十进制输出，%q表用引号引起输出值
	fmt.Printf("str type is %T,str = %q\n", str, str)

	str = fmt.Sprintf("%f", b) //%f表小数形式输出
	fmt.Printf("str type is %T,str = %q\n", str, str)

	str = fmt.Sprintf("%t", c) //%t表bool类型的true和false形式输出
	fmt.Printf("str type is %T,str = %q\n", str, str)

	str = fmt.Sprintf("%c", d) //%c表字符输出
	fmt.Printf("str type is %T,str = %q\n", str, str)

	//使用strconv包的Format函数
	str = strconv.FormatInt(int64(a), 10)
	//变量类型(要精准)+基本输出形式(如2表2进制，10表10进制)
	fmt.Printf("str type is %T,str = %q\n", str, str)

	str = strconv.FormatFloat(b, 'f', 10, 64)
	//变量(b)+输出形式('f')+精度(10,小数后10位)+float类型(32或64)
	fmt.Printf("str type is %T,str = %q\n", str, str)

	str = strconv.FormatBool(c)
	//直接bool类型变量
	fmt.Printf("str type is %T,str = %q\n", str, str)

	//使用strconv包的Itoa函数
	var num5 int = 999
	var str1 string
	str1 = strconv.Itoa(num5) //若int是精确类型,则 var num5 int 8 str = strconv,Itoa(int8(num5))
	fmt.Printf("str1 type is %T,str = %q\n", str1, str1)
}
