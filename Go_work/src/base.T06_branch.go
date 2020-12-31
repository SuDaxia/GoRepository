/**
	分支结构
*/
package main

import (
	"fmt"
)

//main函数：主程序入口
func main() {
	var numa =10;

	//if 不需要括号
	if numa>10{
		fmt.Println(numa)
	}else if numa >0{
		fmt.Println(numa+10000)
	}else{
		fmt.Println(numa-10000)
	}

	//switch
	switch numa {
	//可以case多个值
	case 1,2,3,4,5,6,7:
		fmt.Println("10")
	case 00:
		fmt.Println("100")
	default:
		fmt.Println("other")
	}

	//swithc 表达式
	switch  {
	case (numa > 100) && (numa < 1000):
		fmt.Println("100")
	}

	//switch fallthrough下面的case  实测只执行下一行
	var namea="ll"
	switch namea {
	case "la":
		fmt.Println("la")
	case "lb":
		fmt.Println("lb")
	case "ll":
		fmt.Println("ll")

		fallthrough
	case "haha":
		fmt.Println("haha")
	case "hahb":
		fmt.Println("hahb")
	case "hahc":
		fmt.Println("hahc")
	}

	var numc=10
	switch  {
	case numc==11:
		fmt.Println("aa11")
	case numc==10:
		fmt.Println("aa10")
		fallthrough
	case numc==13:
		fmt.Println("aa13")
		case numc==14:
		fmt.Println("aa14")
	default:
		fmt.Println("aa other")
	}

	//for 同java省略小括号,使用:= 就不用在外面创建下标i
	fmt.Println("for start")
	for i:=1;i<10;i++{
		fmt.Println(i)
	}
	fmt.Print("for end")

	//没有while关键字，for可以充当while，与java相同，break,contine同java
	i:=0
	for i<10{
		fmt.Println(i)
		if i==5{
			break
		}
		i++
	}

	//死循环，就是for三项都不写，或for true
	//for true{
	/*for {
		fmt.Print("死循环")
	}*/

	//range 循环
	var data = []int{1,2,3,4,5,6,11,9}
	//var data []int = []int{1,2,3,4,5,6,11,9}
	for i ,numa:=range data {
		fmt.Println(i,numa)
	}

	// goto 关键字使用，跳转，这个坏习惯也继承了吗？不怕跑飞了？？，但是不能跨函数使用，只能在一个函数里面使用
	fmt.Println("goto start")
	goto my_location
	fmt.Print("看会不会跳过")
	my_location:
		fmt.Println("okkkkk location")
	fmt.Println("goto end")

}

