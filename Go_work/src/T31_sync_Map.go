package main

import (
	"fmt"
	"sync"
)

/**
Go中的并发安全Map(类似于CurrentHashMap)
Go中自己通过make创建的map不是线程安全的，具体体现在多线程添加值和修改值下会报如下错误:
		fatal error : concurrent map writes
	这个错类似于java中多线程读写线程不安全的容器时报的错。
Go为了解决这个问题，专门给我们提供了一个并发安全的map，这个并发安全的map不用通过make创建，拿来即可用，并且他提供了一些不同于普通map的操作方法。

 */
var myConcurrentMap = sync.Map{}
var myRangeMap = sync.Map{}
func rangeFunc(){
	myRangeMap.Store(1,"xiaoming")
	myRangeMap.Store(2,"xiaoli")
	myRangeMap.Store(3,"xiaoke")
	myRangeMap.Store(4,"xiaolei")
	myRangeMap.Range(func(k,v interface{}) bool{
		fmt.Println("data_key_value=: ",k,v)
		//return ture 代表继续遍历下一个，return false代表结束遍历操作
		return true
	})
}
func main() {
	myConcurrentMap.Store(1,"li_ming")
	name,ok:=myConcurrentMap.Load(1)
	if(!ok){
		fmt.Println("不存在")
		return
	}
	fmt.Println(name)
	//有值则不设置，返回原来值和true；如果无该key，添加操作，返回新值和false，代表返回第二个数就是有没有旧值
	name2,ok2:=myConcurrentMap.LoadOrStore(1,"xiao_hong")
	if ok2{
		fmt.Println("有原值为：",name2)
	}else{
		fmt.Println("无原值，设置新值，并返回新值和false：",name2)
	}

	name3,ok3:=myConcurrentMap.LoadOrStore(10,"xiao_hsssssg")
	if ok3{
		fmt.Println("有原值为：",name3)
	}else{
		fmt.Println("无原值，设置新值，并返回新值和false：",name3)
	}

	//标记删除值
	myConcurrentMap.Delete(1)
	name4,ok4:=myConcurrentMap.LoadOrStore(1,"xiao_hccccccccc")
	if ok4{
		fmt.Println("有原值为：",name4)
	}else{
		fmt.Println("无原值，设置新值，并返回新值和false：",name4)
	}

	//遍历数据
	rangeFunc()
}