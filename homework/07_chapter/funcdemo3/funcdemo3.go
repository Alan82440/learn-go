package main

//递归调用

import (
	"fmt"
)

func test(n int) {
	if n > 2 {
		n--
		test(n) //一个函数可申请多个空间
	}
	fmt.Printf("test n=%v\n", n) //位于if语句外，必须执行
}

func test2(n int) {
	if n > 2 {
		n--
		test2(n)
	} else {
		fmt.Printf("test2 n=%v\n", n) //位于if语句中，满足条件才能执行
	}
}

// 斐波那契
func fibonacci(n int) int {
	if n <= 0 {
		return 0 //当输入一个小于1的数，return 0直接结束当前函数
	}
	if n == 1 || n == 2 {
		return 1 //当输入1或2时，结果返回给1
	} else {
		return fibonacci(n-1) + fibonacci(n-2) //当输入的数大于2时，前两项相加
	}
}

// 猴子吃桃，第10天剩一个桃子
// 每天吃一半加一个
func peach(n int) int {
	if n > 10 || n < 1 {
		return 0
	}
	if n == 10 {
		return 1
	} else {
		return (peach(n+1) + 1) * 2
	}
}

// f(1)=3,f(n)=2*f(n-1)+1
func f(n int) int {
	if n == 1 {
		return 3
	} else {
		return 2*f(n-1) + 1
	}
}

func main() {
	test(4)
	//结果为n=2,n=2,n=3,结果由里到外
	test2(4)
	//结果为n=2
	fmt.Printf("fbn n=%v\n", fibonacci(2)) //第二项为1
	fmt.Printf("fbn n=%v\n", fibonacci(3)) //第三项为2
	fmt.Printf("fbn n=%v\n", fibonacci(4)) //第四项为3
	fmt.Printf("fbn n=%v\n", fibonacci(5)) //第五项为5
	fmt.Printf("f(1) = %v\n", f(1))        //3
	fmt.Printf("f(1) = %v\n", f(2))        //7
	fmt.Printf("f(1) = %v\n", f(3))        //15
	var n int
	fmt.Printf("请输入天数")
	fmt.Scan(&n)
	fmt.Printf("第%v天的桃子数为%v\n", n, peach(1))
}
