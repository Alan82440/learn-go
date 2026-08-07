package main

//指针
import (
	"fmt"
)

func main() {
	var num int = 10
	//num相当于地名
	//num存放了一个地址，该地址存放了10
	var ptr *int = &num //*ptr只能接收地址
	//ptr的类型是*int,用&来指向num，ptr也有单独的地址
	fmt.Printf("num address is %v\n", &num) //num地址
	fmt.Printf("ptr=%v\n", ptr)             //输出ptr的地址，因指向i，所以与i相同
	fmt.Printf("ptr=%v\n", *ptr)            //输出指向地址存放的值,加*号一般可以直接取值
	fmt.Printf("ptr=%v\n", &ptr)            //输出ptr本身的地址
	//通过指针修改值
	//指针类型与变量类型不匹配不能取
	var p *int
	var a, b int
	a = 300
	b = 400
	p = &a   //p指向a
	*p = 100 //将a的值修改为100
	p = &b   //p指向b
	*p = 200 //将b的值修改为200
	fmt.Printf("a = %d,b = %d,p = %d\n", a, b, *p)
}
