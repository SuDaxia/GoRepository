package main

/**
Go中的反射与Java中的反射对比

		整体概述：反射是一个通用的概念，是指在程序运行期间获取到变量或者对象，结构体的元信息，比如类型信息，并且能够取出其中变量的值，调用对应的方法。

​      首先我们先来回顾一下Java语言用到反射的场景有哪些？

​      1.比如说我们的方法参数不能确定是什么类型，是Object类型，我们就可以通过反射在运行期间获取其真实的类型，然后做对应的逻辑处理。

​	  2.比如动态代理，我们需要在程序运行时，动态的加载一个类，创建一个类，使用一个类。

​	  3.比如在想要强行破解获取程序中被private的成员。

​      4.Java的各种框架中用的非常多，框架中用反射来判断用户自定义的类是什么类型，然后做区别处理。

​

​	  Go中的反射大概也是相同的，

		比如，go中有一个类型 interface,interface类型相当于Java中的Object类，
		当以interface作为参数类型时，可以给这个参数传递任意类型的变量。
	--------------------------------------------------------------
		那么第一种应用场景就出现了，当我们在go中想实现一个函数/方法，
		这个函数/方法的参数类型在编写程序的时候不能确认，
		在运行时会有各种不同的类型传入这个通用的函数/方法中，
		我们需要对不同类型的参数做不同的处理，那么我们就得能获取到参数是什么类型的，
		然后根据这个类型信息做业务逻辑判断。
		反射我们需要调用reflect包模块,使用reflect.typeOf()可以获取参数的类型信息对象，
		再根据类型信息对象的kind方法，获取到具体类型
 */

import (
	"fmt"
	"reflect"
)

//interface{}代表任意类型
func testAllType (data interface{}){
	fmt.Println(data)
}
func main() {
	testAllType(1)
	testAllType("Go")

	//获得类型
	handleType(10)
	handleType("haha")

	//或得值
	handleValue(1000)
	handleValue("wahahah")

}

func handleType(data interface{}) {
	//获取类型对象
	d := reflect.TypeOf(data)
	//kind方法是获取类型
	fmt.Println(d.Kind())
	switch d.Kind() {
	case reflect.Invalid:
		//无效类型逻辑处理
		fmt.Println("无效类型")
	case reflect.Int,reflect.Int8,reflect.Int16,reflect.Int32,reflect.Int64:
		fmt.Println("整形")
	case reflect.Bool:
		fmt.Println("bool类型")
	default:
		fmt.Println("懒得输出了")
	}

}

func handleValue(data interface{}) {
	//获取类型对象
	d := reflect.ValueOf(data)
	//kind方法是获取类型
	fmt.Println(d.Kind())
	switch d.Kind() {
	case reflect.Invalid:
		//无效类型逻辑处理
		fmt.Println("无效类型")
	case reflect.Int,reflect.Int8,reflect.Int16,reflect.Int32,reflect.Int64:
		//取出值
		var myNum = d.Int()
		fmt.Println(myNum)
	case reflect.Bool:
		//取出bool值
		var myBool = d.Bool()
		fmt.Println(myBool)
	default:
		fmt.Println("懒得搞value")
	}

}


