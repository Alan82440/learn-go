package eng

//同一个文件夹下，只能有一个main包,若有其他包，需要包名一致
import (
	"fmt"
)

func Cal2(n1 float64, n2 float64, operation string) float64 {
	var res float64
	switch operation {
	case "+":
		res = n1 + n2
	case "-":
		res = n1 - n2
	case "*":
		res = n1 * n2
	case "/":
		res = n1 / n2
	case "%":
		res = float64(int(n1) % int(n2))
	default:
		fmt.Println("Invalid operation")
	}
	return res
}
