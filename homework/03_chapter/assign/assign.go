package main

//赋值运算
import (
	"fmt"
)

func test() int {
	return 90
}
func main() {
	a := 10
	b := 20
	fmt.Printf("a = %d,b = %d", a, b)
	//引入第三变量
	t := a
	a = b
	b = t
	//a与b值互换
	//不引入第三变量
	fmt.Printf("a = %d,b = %d", a, b)
	//不引入第三变量
	c := 10
	d := 20
	c = c + d                         //将a重新赋值，此时a = 30
	d = c - d                         //将b重新赋值,此时b = 10
	c = c - d                         //再将a重新赋值,此时a = 20
	fmt.Printf("c = %d,d = %d", c, d) //i的值为最终赋值结果

	var i int = 10
	i += 5 //计算从右往左
	i /= 2 //左边只能是变量
	i *= 2
	i ^= 2
	i += test()    //右边可以不是变量
	fmt.Println(i) //i的值为最终赋值结果
	//注：单目运算与赋值运算时从右到左，其余相反
}
