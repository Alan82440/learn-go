package main

import (
	"fmt"
)

func main() {
	// i := 0
	// for {
	// 	n := rand.Intn(100) + 1 // rand.Intn(100)生成[0,100)的随机数
	// 	i++
	// 	if n == 99 {
	// 		fmt.Printf("%v", i)
	// 		break //break可配合标签label结束对应的循环
	// 	}
	// }

	//例题1
	// var i, sum int = 0, 0
	// for ; i < 100; i++ {
	// 	sum += i
	// 	if sum > 20 {
	// 		fmt.Printf("当sum>20时，i为%v", i)
	// 		break
	// 	}
	// }

	//例题2
	var i int = 1
	var name string
	var pwd int
	for ; i <= 3; i++ {
		fmt.Printf("请输入账户名\n")
		fmt.Scan(&name)
		fmt.Printf("请输入密码\n")
		fmt.Scan(&pwd)
		if name == "张无忌" && pwd == 1234 {
			fmt.Printf("登录成功！\n")
			break
		} else {
			fmt.Printf("登录失败，还剩%v机会!\n", 3-i)
		}
	}
	if i == 4 {
		fmt.Printf("登陆失败，账户已锁定！\n")
	}
}
