package main

import "fmt"

func main() {
	var c1 byte = 'a' //byte类型存放一个字母
	var c2 byte = '0'
	fmt.Println("c1=", c1, "c2=", c2) //输出的是字符对应的ASCII码值
	//若希望输出对应字符，应格式化输出
	fmt.Printf("c1=%c c2=%c\n", c1, c2) //%c表示输出字符
	//若字符变量超出byte范围（ASCII码值），可用更大的类型
	var c3 int = '北'          //可通过utf-8编码表查找对应的码值，北的码值为21271
	fmt.Printf("c3=%c\n", c3) //%c表示输出字符
	var c4 int = 21271
	fmt.Printf("c4=%c\n", c4) //输出UTF-8编码对应的字符
	//字符类型可使用码值计算
	var c5 int = 'a' + 1 //97+1=98
	fmt.Printf("c5=%c\n", c5)
}
