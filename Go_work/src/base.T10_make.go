package main

import "fmt"

/**
go中的make是用来创建slice(切片)，map(映射表)，chan(线程通信管道)这三个类型的对象的，返回的就是对应类型的对象，他的作用就相当于Java中new一个ArrayList，new一个HashMap时候的new的作用，只不过是go语法规定用make来创建slice(切片)，map(映射表)，chan(线程通信管道)。
 */

func main() {
	//make只能为map,channel,slice申请分配内存，只有这三种，没有第四种
	//所有通过make创建的这三种类型都是引用类型，传递参数时虽然是引用值传递，
	//但是对方法内引用变量参数的修改可以影响到外部的引用变量
	//1.通过make创建map对象  如下代码类似于Java中 Map<String,Integer> myMap = new HashMap<>();
	//在这里make就是申请分配map的内存，和java中创建map的new一样
	var map1 =  make(map[string]int)
	map1["hello"]=20
	fmt.Println("make Map:",map1)

	//2.通过make创建channel,make函数内可以有一个参数，也可以有两个参数，有两个参数时第二个参数
	//为通道的缓存队列的长度
	//2.1) 只有一个参数，通道的缓存队列长度此时为0，也就是无缓存。
	//创建一个传输int类型数据的通道 ,没有第二个参数代表无缓存
	var chan1 = make(chan int)
	fmt.Println("无缓存chan1",chan1)




}
