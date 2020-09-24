package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var s, seq string
	for i := 0; i < len(os.Args); i++ {
		s = seq + os.Args[i]
		seq = " "
	}

	fmt.Printf("len of args : %v\n", len(os.Args))

	// 拼接字符串效率低可以使用下面方法
	fmt.Printf("join results : %s\n", strings.Join(os.Args[1:], " "))

	fmt.Println("--------------------------------------------------------------")
	fmt.Printf("result:%s\n", s)

}
