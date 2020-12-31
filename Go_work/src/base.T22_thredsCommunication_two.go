package main

import (
	"fmt"
	"time"
)

/**
多协程 chan通信 client
 */
//channel发送数据和接受数据用 <-表示,是发送还是接受取决于chan在  <-左边还是右边
//创建一个传输字符串数据类型的管道
var chanStr=make(chan string)

func main() {
	fmt.Print("main goroutine start...")
	//默认channel是没有缓存的，阻塞的，也就是说，发送端发送后直到接受端接受到才会施放阻塞往下面走。
	//同样接收端如果先开启，直到接收到数据才会停止阻塞往下走
	//开启新线程发送数据
	go startNewGoroutineOne()

	go startNewGoroutineTwo()

	//不等待直接就结束了
	time.Sleep(100*time.Second)

}

func startNewGoroutineOne(){
	fmt.Println("send channel print ...")
	//管道发送数据 <-
	chanStr<-"Hello!"
}
func startNewGoroutineTwo(){
	fmt.Print("receive channel print ...")
	strVar:=<-chanStr
	fmt.Println(strVar)
}