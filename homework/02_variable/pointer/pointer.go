package main

import (
	"fmt"
)

func main() {
	var i int = 10
	//i相当于地名
	//i存放了一个地址，该地址存放了10
	var ptr *int = &i
	//ptr的类型是*int,用&来指向i，ptr也有单独的地址
	fmt.Printf("ptr=%v\n", ptr)  //输出ptr的地址，因指向i，所以与i相同
	fmt.Printf("ptr=%v\n", *ptr) //输出指向地址存放的值
	fmt.Printf("ptr=%v\n", &ptr) //输出ptr本身的地址
}
