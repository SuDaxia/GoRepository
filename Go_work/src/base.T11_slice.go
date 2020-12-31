package main

import "fmt"

func main() {
	//2.2) 有两个参数，第二个参数2代表此时代表缓存队列的长度为2
	//创建一个传输int类型数据的通道,缓存为2
	var chan2 = make(chan int,2)
	fmt.Println("无缓存chan2\t",chan2)

	//3.通过make创建slice切片
	//有两种方式，一种是两个参数，一种是三个参数
	//我们只有在创建一个空的切片时才会使用make
	//如果通过一个已有的数组创建切片往往是下面的形式
	//创建一个底层数组
	var arr = []int{0,1,2,3,4,5,6}
	var slice1 = arr[2:4]
	fmt.Println("slice1t\t",slice1)

	//我们如果是想创建一个空的slice,则用make创建切片
	//如下形式 make(int[],num1,num2)
	//num1 = 切片的长度(默认分配内存空间的元素个数)
	//num2 = 切片的容量(解释：底层数组的长度/切片的容量，超过底层数组长度append新元素时会创建一个新的底层数组，
	//不超过则会使用原来的底层数组)
	var llen = 10
	var capacity = 15
	//int数组类型，长度len，容量capactiy
	var slice2 =make([]int,llen,capacity)
	fmt.Println("空 slice2\t",slice2,"地址\t",&slice2)

	var data= []int{1,2,3,4,5,6,7}
	for _,e := range data{
		slice2 = append(slice2, e)
	}
	fmt.Println("append超容量 slice2\t",slice2, len(slice2), cap(slice2),"地址\t",&slice2)
}
