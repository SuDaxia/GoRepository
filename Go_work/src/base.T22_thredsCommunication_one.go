package main

import "fmt"

/**
线程间的通信：

java线程间通信有很多种方式：

比如最原始的 wait/notify

到使用juc下高并发线程同步容器，同步队列

到CountDownLatch等一系列工具类
......

甚至是分布式系统不同机器之间的消息中间件，单机的disruptor等等。

Go语言不同，线程间主要的通信方式是Channel。

Channel是实现go语言多个线程（goroutines）之间通信的一个机制。

Channel是一个线程间传输数据的管道，创建Channel必须声明管道内的数据类型是什么
 */
func main() {
	fmt.Println("main server start...")
	flag:=10
	for flag>-1{
		//注意千万不要外面有个变量，里面再:= 这两个就不是同一个的东西了
		ch:=make(chan int)
		fmt.Println(ch)
		flag--
	}
	fmt.Println("main server over!")

}

