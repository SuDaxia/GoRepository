package main

import "fmt"

/**
Go中有闭包的功能。(闭包是一个通用的编程概念，一些语言有，一些没有，javascript中就有这个概念，Java中没有)

闭包,通俗易懂的讲，就是你有一个A函数，A函数有一个a参数，然后在A函数内部再定义或者调用或者写一个B函数，这个B函数叫做闭包函数。B函数内部的代码可以访问它外部的A函数的a参数，正常A函数调用返回完毕，a参数就不能用了，可是闭包函数B函数仍然可以访问这个a参数，B函数能不受A函数的调用生命周期限制可以随时访问其中的a参数，这个能访问的状态叫做已经做了闭包，闭包闭的是把a参数封闭到了B函数中，不受A函数的限制。

也就是说，我们用程序实现一个闭包的功能，实质上就是写一个让外层的函数参数或者函数内变量封闭绑定到内层函数的功能。
 */
//我们来看看实现闭包
func main() {
	var f = f1(100)
	f(100) //print 200
	f(100) //print 300
	f(100) //print 400
	f(0)
	f(0)
	f(0)
	f(0)
}

func f1(x int) func(int){
	//此时即使f1函数执行完毕，x也不会消失
	//x在func(y int)这个函数内一直存在并且叠加，
	//这里把x的值封闭到func(y int)这个返回函数中，使其函数一直能使用x的值
	//就叫做闭包，把x变量封闭到fun(y int)这个函数包内。
	return func(y int){
		x+=y
		fmt.Printf("x=%d\n",x)
	}
}
/**
做下闭包的总结，如何实现一个闭包：

1.定义一个A函数，此函数返回一个匿名函数。（定义一个返回匿名函数的A函数）

2.把在A函数的b参数或A函数代码块中的b变量，放入匿名函数中，进行操作。

3.这样我们调用A函数返回一个函数，这个函数不断的调用就可以一直使用之前b参数，b变量，并且b值不会刷新，
有点像在匿名函数外部自定义了一个b的成员变量（成员变量取自Java中类的相关概念）
 */