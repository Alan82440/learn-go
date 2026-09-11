package main

import (
	"fmt"
)

func test(n1 int) {
	n1 += 1
	fmt.Printf("test() = %v\n", n1)
	//结果为21，这里接收了n1的值，但这里的n1与main中的n1不同
	//结果不同，结果没返回给调用者
}

func getsum(n1 int, n2 int) int {
	var sum int
	sum = n1 + n2
	fmt.Printf("test()sum = %v\n", sum)
	return sum
	//这里将结果返回了sum，相当于把结果返回给调用者
}

func getsumAndsub(n1 int, n2 int) (int, int) {
	var sum, sub int
	sum = n1 + n2
	sub = n1 - n2
	return sum, sub
	//Go中，可返回多个结果
}

func main() {
	var n1 int = 20
	test(n1)
	fmt.Printf("main() = %v\n", n1)
	//打印结果为20，只是将n1传给函数test,结果没有返回来
	var sum int
	sum = getsum(1, 2)
	fmt.Printf("main()sum = %v\n", sum)
	var sum1, sub1 int
	sum1, sub1 = getsumAndsub(1, 2)
	//对于不用的结果可用"_"屏蔽
	//如 _,sub1 = getsumAndsub(1,2)只获得sub
	fmt.Printf("sum1 = %v,sub1 = %v\n", sum1, sub1)
}
