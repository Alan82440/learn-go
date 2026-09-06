package main

import (
	"fmt"
	"math/rand"
)

func main() {
	i := 0
	for {
		n := rand.Intn(100) + 1 // rand.Intn(100)生成[0,100)的随机数
		i++
		if n == 99 {
			fmt.Printf("%v", i)
			break
		}
	}
}
