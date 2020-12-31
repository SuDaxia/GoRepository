package main

import (
	"fmt"
)

/**
给结构体添加方法，类似于Java class或实例对象的方法
方法的接收器也可以是指针类型
【巨他妈神奇，不管用指针还是变量，进去都统一的变量.属性，到后面访问，用变量调用传指针的还是传变量的都可以】
【然后继续深入，发现用变量的，copy对象进去，不影响外面对象。而传指针进去的会影响外面的对象，难怪要搞两种东西】
【再测试，是不是可以对非指针的，传地址进去是不是就能做到指针那样影响外面的对象了
测试结果：弄成非指针形式的，怎么搞都改变不了外部变量的值，所以还是两大类，无法混用】
【还要加方法，小心传的对象是nil，要注意健壮性,nil.subfunc那一层就是异常，不是subfunc里面】
 */
type Student struct{
	num int
	name string
	class int //班级
	sex bool
}

func (stu Student) learn(book string){
	fmt.Printf("%s学生正在学习%s\n",stu.name,book)
	stu.name+="update1"
}

func (stu *Student) learn22(book string){
	fmt.Printf("%s学生正在学习%s\n",stu.name,book)
	stu.name+="update22"
}
func test1() {
	stu:=Student{num:12,name:"张三"}
	stu.learn("语文")
	fmt.Println(stu)
	stu.learn22("语文")
	fmt.Println(stu)
	fmt.Println("-------")

	stuPoint:=&Student{num:12,name:"王二"}
	stuPoint.learn("数学")
	fmt.Println(stuPoint)
	stuPoint.learn22("数学")
	fmt.Println(stuPoint)
	fmt.Println("-------")

}
func main() {
	//test1()
	//test2()
	testNil()
}

func test2(){
	stu:=Student{num:12,name:"张三"}
	//&stu.learn("测试非指针传地址来调用方法") //不行
	//*stu.learn("测试非指针传地址来调用方法") //不行
	//(*stu).learn("测试非指针传地址来调用方法") //不行
	//(&stu).learn("测试非指针传地址来调用方法") //ok 但是并没有改变外部对象值
	//(*&stu).learn("测试非指针传地址来调用方法") //ok 但是并没有改变外部对象值
	//(&*stu).learn("测试非指针传地址来调用方法") //ok 但是并没有改变外部对象值
	fmt.Println(stu)
	fmt.Println("-------")
}

func (stu Student) learn333(book string){
	defer func() {
		err:=recover()
		if err!=nil{
			fmt.Println("learn333 ERROR:",err)
		}
	}()
	fmt.Printf("%s学生正在学习%s\n",stu.name,book)
}
/**
怎么赋值为nil呢？？奇奇怪怪的东西
哈哈，一般都是指针nil，ok
 */
func testNil(){
	defer func() {
		err:=recover()
		if err!=nil{
			fmt.Println("testNil ERROR:",err)
		}
	}()
	var stu *Student
	stu.learn("没有书可读")
}
