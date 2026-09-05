package main

//for的使用
import (
	"fmt"
)

func main() {
	//传统循环
	// 	for i := 1;i < 3;i++{
	// 		fmt.Printf("hello,world")
	// 	}

	// i:=1
	// for;i<3;i++{
	// 	fmt.Printf("hello,world")
	// }

	// i := 1
	// for i < 3 {
	// 	fmt.Printf("hello,world")
	// 	i++
	// }

	// i := 1
	// for { //for{} 等价于for;;{},一般配合break使用
	// 	if i < 3 {
	// 		fmt.Printf("hello,world")
	// 	} else {
	// 		break
	// 	}
	// 	i++
	// }

	//遍历带中文的字符数组需[]rune
	// var str string = "hello,world!北京"
	// str2 := []rune(str) //遍历中文需要添加
	// for i := 0; i < len(str2); i++ {
	// 	fmt.Printf("%c", str2[i])
	// }

	//for-range，遍历数组一般使用
	// var str string = "hello,world!北京" //使用for-range遍历带中文的数组不需要[]rune
	// for index, char := range str {  //如果不需要遍历其中某个变量，可用“_”省略
	// 	fmt.Printf("index=%d,char=%c\n", index, char)  //index,char可写成其他变量名
	// } //index和char是变量名，用于遍历字符字节与字符，一个中文占3字节
	//例子：
	// var str string = "hello,world!"
	// for _, c := range str {
	// 	fmt.Printf("str=%c\n", c)
	// }

	//例题1
	// var i int = 1
	// var sum int = 0
	// for ; i <= 100; i++ {
	// 	if i%9 == 0 {
	// 		sum += i
	// 		fmt.Printf("9的倍数%d\n", i)
	// 	}
	// }
	// fmt.Printf("9的倍数和%d\n", sum)

	//例题2
	// var num int = 6
	// for i := 0; i <= num; i++ {
	// 	fmt.Printf("%v + %v = %v\n", i, num-i, num)
	// }

	//循环嵌套
	var stu int = 5
	var class int = 3
	var score, sum float64
	var grade float64 = 0.0
	fmt.Printf("请输入成绩与所在班级")
	for i := 1; i <= class; i++ {
		sum = 0.0
		var peo int = 0
		for j := 1; j <= stu; j++ {
			fmt.Printf("第%v个班,第%v名同学的成绩\n", i, j)
			fmt.Scan(&score)
			sum += score
			if score >= 60 {
				peo++
			}
		}
		fmt.Printf("第%v个班的平均分是%v,及格人数是%v\n", i, sum/5, peo)
		grade += sum
		if i == class {
			fmt.Printf("%v个班的平均分是%v", i, grade/3)
		}
	}
}
