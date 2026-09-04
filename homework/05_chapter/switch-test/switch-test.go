package main

import (
	"fmt"
)

func main() {
	// var score float64
	// fmt.Printf("请输入分数")
	// fmt.Scan(&score)
	// switch {
	// case score >= 60 && score < 100:
	// 	fmt.Printf("成绩合格")
	// case score < 60 && score >= 0:
	// 	fmt.Printf("成绩不合格")
	// default:
	// 	fmt.Printf("输入错误")
	// }

	// var season int
	// fmt.Printf("请输入月份")
	// fmt.Scan(&season)
	// switch season {
	// case 3, 4, 5:
	// 	fmt.Printf("春季")
	// case 6, 7, 8:
	// 	fmt.Printf("夏季")
	// case 9, 10, 11:
	// 	fmt.Printf("秋季")
	// case 12, 1, 2:
	// 	fmt.Printf("冬季")
	// default:
	// 	fmt.Printf("错误输入")
	// }

	// var day string
	// fmt.Printf("请输入星期")
	// fmt.Scan(&day)
	// switch day {
	// case "星期一":
	// 	fmt.Printf("干煸豆角")
	// case "星期二":
	// 	fmt.Printf("醋溜土豆")
	// case "星期三":
	// 	fmt.Printf("红烧狮子头")
	// case "星期四":
	// 	fmt.Printf("油炸花生米")
	// case "星期五":
	// 	fmt.Printf("蒜蓉扇贝")
	// case "星期六":
	// 	fmt.Printf("东北乱炖")
	// case "星期日":
	// 	fmt.Printf("大盘鸡")
	// }

	//fallthrough,无判断执行下一个case
	var n1 int = 1
	switch n1 {
	case 1:
		fmt.Printf("ok1\n")
		fallthrough
	case 2:
		fmt.Printf("ok2")
	default: //default可省略
		fmt.Printf("无")
	}

}
