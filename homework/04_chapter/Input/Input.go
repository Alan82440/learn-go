package main

//键盘输入
import (
	"fmt"
)

func main() {
	//一，引用fmt.Scanln
	// var name string
	// var age int
	// var sal float32
	// var isPass bool
	// fmt.Println("请输入名字")
	// fmt.Scanln(&name) //系统等待键盘输入，类似C语言sacaf
	// fmt.Println("请输入年龄")
	// fmt.Scanln(&age)
	// fmt.Println("请输入薪水")
	// fmt.Scanln(&sal)
	// fmt.Println("请输入考试情况")
	// fmt.Scanln(&isPass)
	// fmt.Printf("名字:%s\n 年龄:%d\n 薪水:%f\n 考试情况:%t\n", name, age, sal, isPass)
	//二，使用fmt.Scanf(与C语言类似)
	var Name string
	var Age int
	var Sal float32
	var IsPass bool
	fmt.Println("请输入名字，年龄，薪水，考试情况，用空格隔开")
	fmt.Scanf("%s\n %d\n %f\n %t\n", &Name, &Age, &Sal, &IsPass)
	fmt.Printf("名字:%s\n 年龄:%d\n 薪水:%f\n 考试情况:%t\n", Name, Age, Sal, IsPass)
}
