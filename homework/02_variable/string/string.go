package main

import "fmt"

func main() {
	//string的基本使用
	var address string = "北京 101 hello"
	fmt.Println(address)
	var str string = "hello" //变量已被定义的字符不能改变,如str[0] = 'a'
	//“”会识别特殊符号，不能输出特殊符号
	//··反引号可以输出特殊符号(键盘波浪线下面的的符号)
	test := `package main

import "fmt"

func main() {
	fmt.Println("姓名\t年龄\t籍贯\t住址\n张三\t20\t河北\t北京")
	var i int = 10
	fmt.Printf("i是 %T", i) //%T表示打印变量的类型,T为typeof
}
`
	fmt.Println(str, test)
	//string字符拼接
	var str1 string = "hello" + "world"
	str1 += "haha!" //str += 1 等价于 str = str + 1
	//多行拼接时加号要放末尾
	var str2 string = "hello" +
		"world"
	fmt.Println(str1, str2)
}
