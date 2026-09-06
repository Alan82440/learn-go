package main

//for进阶
import (
	"fmt"
)

func main() {
	//打印实心金字塔
	// var level int
	// fmt.Printf("请输入层数:")
	// fmt.Scan(&level)
	// for i := 1; i <= level; i++ {
	// 	for k := 1; k <= level-i; k++ {
	// 		fmt.Print(" ")
	// 	}
	// 	for j := 1; j <= 2*i-1; j++ {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Printf("\n")
	// }

	//打印空心金字塔
	//打印层数
	// for i := 1; i <= level; i++ {
	// 	//打印*前的空格
	// 	for k := 1; k <= level-i; k++ {
	// 		fmt.Print(" ")
	// 	}
	// 	//打印每层需要的*
	// 	for j := 1; j <= 2*i-1; j++ {
	// 		if j == 1 || j == 2*i-1 || i == level {
	// 			fmt.Print("*")
	// 		} else {
	// 			fmt.Printf(" ")
	// 		}
	// 	}
	// 	fmt.Printf("\n")
	// }

	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v * %v = % v ", i, j, i*j)
		}
		fmt.Printf("\n")
	}
}
