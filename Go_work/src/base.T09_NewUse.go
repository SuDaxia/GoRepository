package main

import "fmt"

/**
Go语言的中new,make和Java中的new对象有什么区别？

 Go中的创建一个struct结构体的对象，是不需要用new关键字的，参考【20】中有代码示例讲解如何创建结构体对象。

  Go中new的概念是和内存相关的，我们可以通过new来为基础数据类型申请一块内存地址空间，然后把这个把这个内存地址空间赋值给

  一个指针变量上。（new主要就是为基础数据类型申请内存空间的，当我们需要一个基础数据类型的指针变量，并且在初始化这个基础指针变量时，不能确定他的初始值，此时我们才需要用new去内存中申请一块空间，并把这空间绑定到对应的指针上，之后可以用该指针为这块内存空间写值。new关键字在实际开发中很少使用，和java很多处用new的情况大不相同）

 */

func main() {
	/*
	var numa *int
	//此处numa 是 nil，访问会空指针异常
	fmt.Println(*numa)
	*/

	var numa = new(int)
	//ok ,很少用到关键字new
	fmt.Println(numa)
	*numa=123
	fmt.Println(*numa)

}
