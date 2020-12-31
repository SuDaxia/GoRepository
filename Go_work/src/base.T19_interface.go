package main

import "fmt"

/**
Go的多态接口
go语言中的接口是非侵入式接口。

java语言中的接口是侵入式接口。

侵入式接口是指需要显示的在类中写明实现哪些接口。

非侵入式接口是指不要显示的在类中写明要实现哪些接口，只需要方法名同名，参数一致即可。

 */

type Animal interface {
	Eat()
	sleep()
	//move() //如果生命一个变量 var a Animal 那么必须要实现所有方法才行，不然编译报错
}

type Cat struct {

}

type Dog struct {

}

func (cat Cat) Eat(){
	fmt.Println("cat eat ")
}

func (cat Cat) sleep(){
	fmt.Println("cat sleep")
}

func(dog Dog)Eat(){
	fmt.Println("dog eat ")
}

func (dog Dog) sleep(){
	fmt.Println("dog sleep")
}

func main() {
	var cat Animal=Cat{}
	var dog Animal=Dog{}
	cat.Eat()
	cat.sleep()
	dog.Eat()
	dog.sleep()
}