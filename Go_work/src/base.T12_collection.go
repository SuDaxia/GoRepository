package main

import "fmt"

/**
Go中有的数据结构：数组，切片，map，双向链表，环形链表，堆
Go自己的类库中没有set,没有集合(List)，但是第三方库有实现。
Java中有： Map,Set,List,Queue,Stack,数组
Java中没有切片的概念。
Go中的数组打印格式是[1,2,3,4,5]
Go中的切片打印格式是[[1,2,3]]
Go中切片的概念：切片是数组的一个子集，就是数组截取某一段。
Go的map和Java的map大致相同
 */

func main() {
	arr:=[]int{1,2,3,4,5}
	slice:=make([]int,4)
	for _,e :=range  arr{
		slice = append(slice, e)
	}
	fmt.Println(slice)

	//map初始化与使用
	//方式一：
	//这只是声明，这样做完了还是需要使用make进行初始化的
	//【注意使用make后，key不允许为nil了】
	//mapa:= map[int]int //此时mapa是nil的
	//需要使用make进行是初始化才能使用
	mapa :=make(map[int]int,4)
	//使用
	for _,e:=range arr{
		mapa[e]=e^2
	}
	//mapa[5]=nil //编译报错 make后map不允许nil
	//mapa[nil]=4 //运行报错，不允许nil
	fmt.Println(mapa)

	//map初始化方式二，必须初始化+赋值一体化必须要加逗号
	mapb:=map[int]int{
		1:2,3:3,
		//4:nil, //运行报错，不允许nil
		//nil:8, //运行报错，不允许nil
	}
	fmt.Println(mapb)

	//查找key是否有
	if v,ok:=mapb[5];ok{
		fmt.Println(v)
	}else{
		fmt.Println("key not found")
	}

	//便利key
	for k,v:=range mapb{
		fmt.Println(k,v)
	}

}