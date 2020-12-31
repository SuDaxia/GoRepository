//主程序必须写成main包名
package main

//单包引入
//import "fmt"
//或者多包引入
import (
	//包含print函数
	"fmt"
)

//全局常量定义
/**
go中的常量和java中的常量含义有一个本质的区别：
go中的常量是指在编译期间就能确定的量(数据)，
而java中的常量是指被赋值一次后就不能修改的量(数据)。
所以两者不一样，因为Java中的常量也是JVM跑起来后赋值的，只不过不允许更改；
go的常量在编译后就确实是什么数值了。
 */
const const_num =10

//全局变量定义
var name string = "super man"

func constSample(){
	//const 常量名常量类型=常量值 显示推断类型
	const name string = "const_sample_visible"
	//隐式推断类型
	const name2 = "const_sample_infer"
	fmt.Println(name)
	fmt.Println(name2)
}

//类型定义
type Person struct{
}

//初始化函数
func init(){
	name = "sdf"
	fmt.Print("hahahah")
}

//定义其他方法
func sample(){
	//var 变量名 变量类型=变量值 :=用于值初始化
	var name string = "aaaaa"
	//方法内部可以直接使用【变量名 := 变量值】 赋值，方法外不可以
	name2:="xixixiixi"
	fmt.Println("name:",name)
	fmt.Println("name2:",name2)
	//多变量同时赋值
	var a,b,c="sdfsdf",85,64
	fmt.Println(a,b,c)

}
//传递地址 类似java引用 或 传递变量(copy)类似Java值类型
func hello(name *string ,name2 string){
	*name+="xxxxx"
	name2+="ooooooo"
	fmt.Println("helllo funct:",name,name2)

}

/**
类似 python的接收多个返回值，但是这里_是丢弃对象，无法获取
 */
func useMultiReturn(){
	//_用于接收并丢弃不要的值，无法通过_来访问变量了
	var _,num1,_,_,num2=1,2,3,4,5
	fmt.Println("测试多个返回",num1,num2)
}

//main函数：主程序入口
func main() {
	//go 关键字开启go routine 协程
	// 默认的println是输出项 standerr(out,err)的，可以看出来是打印的红色的，专门用于调试，官网说后期不一定还可以用
	// 所以很多的源码和生产级别，是用fmt.Pringln来打印日志，并且也支持格式化
	go println("test Hello " + name)
	go fmt.Println("test sdfsdfsd "+name)

	//init()
	var namea="aaa"
	var nameb="bbbb"
	hello(&namea,nameb)
	fmt.Println("看值改变没有：",namea,nameb)
	useMultiReturn()

}

func init() {

}

