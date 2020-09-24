package main

import "fmt"

func main() {

	// p 作为左值
	var p int  = 10

	var a int

	// p作为又值
	a  = p

	fmt.Println(a)
	test()
}

func swap(a, b *int) {
	fmt.Println("-----------------")
	fmt.Println(&a)
	fmt.Println(&b)
	a,b = b,a // 等同如下
	//*a,*b = *b,*a
	fmt.Println("-----------------")
	fmt.Println(&a)
	fmt.Println(&b)
	println(*a) // 2
	println(*b)  // 1

	/** 地址值
		0xc00000a100
		0xc00000a0e8
		-----------------
		0xc000006030
		0xc000006038
	 */

}

func test() {
	a := 1
	b := 2
	//fmt.Println(&b)
	//fmt.Println(&a)
	swap(&a,&b)
}
