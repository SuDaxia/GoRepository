package main

import "fmt"

/**
Go中返回值可以有多个，不像Java中多个值得封装到实体或map返回
//注：【】内的返回值可不写，无返回值直接把返回值部分全部去掉即可。
func 函数名(变量1 变量类型，变量2 变量2类型...)【(返回值1 类型1，返回值2 类型2...)】  {        //注意：这个方法的右中括号必须和func写在同一行才行，否则报错，不能按c语言中的换行写

有多个返回值的时候，因为有可能有些变量不使用编译就不通过，可以使用_符合接住，相当于要抛弃的对象

同时记住 ，大写字母开头的变量和方法 相当于java的public 公共对外可访问的，小写字母的相当于 protected 只能同package下或子类？？使用
*/
func sum(a int,b int) int{
	return a+b
}

func lowerNums(a int)[]int{
	var arr=[]int{}
	for i:=0;i<a;{
		arr = append(arr, i)
		i++
	}
	return arr
}


func returnMuliParams()(int ,
	int){
	return 1,2
}

//变长参数 三个dot ...
func manyArgs(a int ,b int ,str ...string){
	for i ,s :=range str{
		fmt.Println(i,s)
	}
}

func main() {
	fmt.Println(sum(returnMuliParams()))
	manyArgs(1,2,"vsdfsdf","fsefeg")
	manyArgs(1,2,"vsdfsdf","fsefeg","fefsefsssss")
}
