package main

import (
	"fmt"
	"time"
)

/**
go中的线程相关的概念是Goroutines(并发)，是使用go关键字开启。

Java中的线程是通过Thread类开启的。

在go语言中，一个线程就是一个Goroutines，主函数就是（主） main Goroutines。

使用go语句来开启一个新的Goroutines
 */

func myFunctin(){
	fmt.Println("myFunction")
}

func goRoutineTestFunc(){
	fmt.Println("Hello,start goroutine")
}

func main() {
	//myFunctin()

	//主线程结束的快，可能都看不到协程中执行方法的打印语句
	go goRoutineTestFunc()
	myFunctin()
	time.Sleep(10*time.Second)//10s

}
