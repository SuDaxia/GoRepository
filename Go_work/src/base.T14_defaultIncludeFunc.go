package main

import (
	"fmt"
)

/**
Go的内置函数 类似 Java的默认导入包 java.lang.*
为了在Java中快速开发，Java语言的创造者把一些常用的类和接口都放到到java.lang包下，lang包下的特点就是不用写import语句导入包就可以用里面的程序代码。

Go中也有类似的功能，叫做Go的内置函数，Go的内置函数是指不用导入任何包，直接就可以通过函数名进行调用的函数。

Go中的内置函数有：

close          关闭channel

len            求长度

make	      创建slice,map,chan对象

append	      追加元素到切片(slice)中

panic         抛出异常，终止程序,相当于java的 throw new Exception()

recover       尝试恢复异常，必须写在defer相关的代码块中,
			才能再捕获异常的时候有作用，而且，子协程中只能处理子协程的异常，不能往外交给主defer，各自自管自己的协程
 */
func sampleOne(){
	arr:=[]int{1,2,3,4,5}
	str:="myname"
	fmt.Println(len(str))
	fmt.Println(len(arr))
	fmt.Println(len(arr[1:]))

	//make创建channel示例
	intChan:=make(chan int,1)

	myMap:=make(map[string]interface{})

	mySlice:=make([]int,5,10)
	fmt.Println(intChan)
	fmt.Println(myMap)
	fmt.Println(mySlice)

	//关闭管道
	close(intChan)
	//注意，加元素，如果不用切面，那么原始的数组也会改变??测试结果不会改变
	//arr2:=append(arr[:],7)
	arr2:=append(arr,7)
	fmt.Println("原始arr改变么？",arr)
	fmt.Println("新arr",arr2)

	//new 案例,new一个int类型的对象出来
	num:=new(int)
	fmt.Println(num)
	fmt.Println("-------")
}
func main() {

	func1()
	//func2()
	func22()
	func3()
}
/**

 */
func func1(){
	fmt.Println("1")
}

func func2(){
	//刚刚打开某资源
	//先定义异常后要做什么，后面再些泡
	defer func() {
		fmt.Println("释放资源1")
		err:=recover()
		fmt.Println(err)
		fmt.Println("释放资源2")
		fmt.Println("释放资源3")
		fmt.Println("释放资源over")
	}()
	//通过注释来发现，上面定义的还是会直接执行啊
	//panic("抛出异常，相当于java的 throw new Exception(message)")
	var a int
	fmt.Println("请输出不等与0的数字")
	fmt.Scanf("%d",&a)
	fmt.Print(2 / a)//出异常，然后后面的没有去执行，然在
	//devideZero()
	fmt.Println("2222")
}

/**
错误实例，卸载后面，程序崩溃直接推出，defer定义的就相当于 catch和finnaly要放前面一样
 */
func func22(){
	//通过注释来发现，上面定义的还是会直接执行啊
	//panic("抛出异常，相当于java的 throw new Exception(message)")
	var a int
	fmt.Println("请输出不等与0的数字")
	fmt.Scanf("%d",&a)
	fmt.Print(2 / a)//出异常，然后后面的没有去执行，然在
	//devideZero()
	fmt.Println("2222")
	defer func() {
		fmt.Println("释放资源1")
		if err:=recover();err!=nil{
			fmt.Println("出现异常",err)
		}
		fmt.Println("释放资源2")
		fmt.Println("释放资源3")
		fmt.Println("释放资源over")
	}()
}

func func3(){
	fmt.Println("3")
}

func devideZero()string{
	var a int
	//fmt.Scan类似C的输入
	//返回的是长度，以及错误信息，不是变量值，变量值在里面就赋值了
	fmt.Println("请输入一个不等于0的数字")
	length,err:=fmt.Scanf("%d",&a)
	if err!=nil{
		fmt.Println("error:",err.Error())
		return err.Error()
	}
	fmt.Println("read a:\t",a,"\tand data len:",length)
	fmt.Println("2/a",20/a)
	return "ok"
}

