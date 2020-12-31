package main

import "fmt"

/**
嵌入结构体来达到继承的目的，精简代码

需要注意的是，Go中可以嵌入多个结构体，但是多个结构体不能有相同的方法，
如果有参数和方法名完全相同的方法，在编译的时候就会报错。所以Go不存在嵌入多个结构体后，
嵌入的几个结构体有相同的方法，最后不知道选择执行哪个方法的情况，多个结构体方法相同时，直接编译就会报错。

 */

type Computer struct {
	screen string
	keyboard string
}

func (cp Computer) compute(num1,num2 int) int{
	return num1+num2
}

type NoteBookComputer struct {
	Computer
	whireless_network_adapter string
}

func main_1() {
	cp1 :=NoteBookComputer{}
	cp1.screen="屏幕"
	cp1.keyboard="cherry keyboard"
	cp1.whireless_network_adapter="1KM net"
	fmt.Println(cp1)
	fmt.Println(cp1.compute(1, 2))
}
func main() {
	c1:=C{}
	fmt.Println(c1)
	//c1.sameFunc() //不写调用同名方法的代码，编译不会报错，但是如果写了，要调用的时候会编译报错
				//因为不知道调用那个成员的方法
	c1.A.sameFunc()//这样调用就没有问题了
	c1.B.sameFunc()
	c1.a.sameFunc()
}

type A struct {
}
type B struct {
}
func (a A)sameFunc(){
	fmt.Println("a")
}
func (b B)sameFunc(){
	fmt.Println("b")
}
type C struct {
	A
	B
	a A
}
