/**
Go有指针概念，Java没有指针概念
指针简单的说就是存储一个【变量地址】的【变量】。


*/
package main

import "fmt"

func main() {
	var numa =100;
	//*类型，表明当前变量是一个指向某种类型的指针，里面存放的是指向变量的地址，所以赋值，需要的是赋值地址
	//&就是取地址
	var point1 *int = &numa
	//然后如果是传point1就是传的地址，但我们使用诗用指针指向的值，所以用*解指针，获取指向的变量值
	fmt.Println(*point1)
	//这个打印的地址
	fmt.Println(point1)
	fmt.Println("----")
	*point1=100
	fmt.Println(*point1)
	fmt.Println(point1)

	fmt.Println("----")
	numa=3000
	fmt.Println(*point1)
	fmt.Println(point1)

	//其实都是copy调用，只不过，传递地址时，方法栈区中的copy变量里面存的是相同的地址，访问内容有，该copy的存储的变量消失，但指针指向的内容没有


}

