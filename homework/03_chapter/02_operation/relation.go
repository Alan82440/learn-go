package main

//关系运算

import (
	"fmt"
)

func main() {
	//关系运算
	var n1 int = 8
	var n2 int = 9
	fmt.Println(n1 == n2) //false
	fmt.Println(n1 != n2) //true
	fmt.Println(n1 >= n2) //false
	fmt.Println(n1 <= n2) //true
	fmt.Println(n1 > n2)  //false
	fmt.Println(n1 < n2)  //true
	var result bool
	result = n1 < n2
	fmt.Println(result)
}
