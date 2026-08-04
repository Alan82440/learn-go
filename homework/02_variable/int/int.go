package main

import (
	"fmt"
	"unsafe"
)

// 一次性定义多个全局变量
var (
	a     = 100
	b     = 200
	name2 = "tom"
)

func main() {
	//第一种声明变量的方式,变量=变量名+值+类型
	var i int                     //var声明变量，i为变量名，int为变量类型
	i = 10                        //=为赋值符号，将10赋值给变量i
	var j int                     //若j不赋值，则j的默认值为0,int与float类型的默认值为0,string为空
	fmt.Println("i=", i, "j=", j) //打印变量i，j的值
	//第二种声明变量的方式
	var k = 10.01        //若k不指定类型，则根据初始值自动推断类型为float
	fmt.Println("k=", k) //打印变量k的值
	//第三种声明变量的方式
	name1 := "tom" //省略var，注意name不能已被var声明，该赋值等价var name = "tom"
	fmt.Println("name1", name1)
	//多变量声明
	//var n1,n2,n3 int            //n1,n2,n3都为int型
	//var n1,name,age = 100,"tom",18   //n1,age为int,name为strring
	//n1,name,age := 100,"tom",18 //等价上述类型
	fmt.Println("a=", a, "b=", b, "name2=", name2) //打印全局变量a,b,name的值
	//“+”的使用
	var l1 int = 2 //l可在int类型内发生变化，不能变为其他类型
	l1 = 3
	var l2 = 4
	var m = l1 + l2      //都为数值型时，做加法运算
	fmt.Println("m=", m) //m的值为l1和l2的和
	var str1 = "hello"
	var str2 = "world"
	var str3 = str1 + str2 //string类型可以进行拼接运算
	fmt.Println("str3=", str3)
	fmt.Printf("str3是 %T\n", str3) //%T表示打印变量的类型,T为typeof
	//Sizeof()是unsafe包中的函数，表示打印变量的字节数
	fmt.Printf("str3的字节数:%d\n", unsafe.Sizeof(str3)) //unsafe.Sizeof()表示打印变量的字节数
}

//基本变量
//int8,int16,int32,int64,分别为1字节，2字节，4字节，8字节的整数类型,int为int32或int64，取决于操作系统的位数
//unit8,unit16,unit32,unit64,分别为1字节，2字节，4字节，8字节的无符号整数类型
//补充int8表-128~127，unit8表0~2的8次方减1，int32表-2的15次方~2的15次方减1，unit32表0~2的32次方减1
//float32,float64,分别为4字节，8字节的浮点数类型,既单精度型，双精度性
//专门用byte存放一个字母，rune类型存放一个汉字，byte为uint8的别名，rune为int32的别名
//bool类型，存放true或false
