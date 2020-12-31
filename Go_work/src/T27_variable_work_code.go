package main

import "fmt"

/**
变量作用域的区别
	Go语言的变量作用域和Java中的一样，遵循最近原则，逐渐往外层找。
	这个比较简单，就不做过多赘述了。

Go语言和Java语言字符串操作的区别

Go语言和Java语言IO操作的区别

Go语言中有匿名函数，有闭包，Java中没有(高阶函数用法)
	函数也是一种类型，它可以作为一个参数进行传递，也可以作为一个返回值传递。
	Go中可以定义一个匿名函数，并把这个函数赋值给一个变量

【】Go的函数内部是无法再声明一个有名字的函数的，Go的函数内部只能声明匿名函数。

 */
//定义一个匿名函数并赋值给myFun变量
var myFun = func(x,y int) int {
	return x+y
}

func main() {
	//调用myFun
	fmt.Println(myFun(1,2))
	myFunc1()
}

func myFunc1(){
	//编译不通过，方法中智能起匿名方法，然后把它赋值给一个对象，通过对象来调用
	/*func myFunc2(){

	}*/
	var f=func(){
		fmt.Println("hi boy")
	}
	//调用
	f()
	//或者不复制对象，立即调用
	func(str string){
		fmt.Println("hi boy 直接调用",str)
	}("大哥")
}

