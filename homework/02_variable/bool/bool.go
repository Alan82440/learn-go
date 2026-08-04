package main

import "fmt"

func main() {
	//bool类型，只存放true或false
	//bool默认值为false
	var isMarry bool
	fmt.Printf("isMaarry=%v", isMarry) //%v表示按变量值输出
	//bool只占用一个字节
	var b1 bool = true
	var b2 bool = false
	fmt.Println("b1=", b1, "b2=", b2)
}
