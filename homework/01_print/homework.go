// 这是go语言的包声明，表示该文件属于main包
package main

//这里调用函数fmt包中的Println函数，打印Hello, World!到控制台
import (
	"fmt"
)

// 这是一个函数入口，主函数为main函数，程序从这里开始执行
func main() { //“{”只能放在()的后面，不能换行
	fmt.Println("Hello, World!")
	fmt.Println("性别\t年龄\t籍贯\t住址\njohn\t12\t河北\t北京") //\t表示制表符，\n表示换行符
	fmt.Println("tom says: \"Hello, World!\"")      //\表示转义字符，表示后面的字符是特殊字符
	fmt.Println("鸣潮666\r原神")                        //\r表示回车符，表示将光标移动到当前行的开头
	fmt.Println("hello,world",
		"hello,world") //函数参数可以换行，换行后需要缩进
	/*块注释，块注释不能再嵌套注释*/ //ctrl+/表示注释当前行，再次ctrl+/表示取消注释
	//shift+tab表示缩进，shift+tab表示取消缩进
	//值类型：int,float,bool,string,数组，结构体，通常存放于栈区
	//引用类型：指针，slice,map,chan,interface，通常存放于堆区
}
