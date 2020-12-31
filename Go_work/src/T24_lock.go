package main

import (
	"fmt"
	"sync"
	"time"
)

/**
除了chan缓存
还可以有互斥锁，类似java中的Lock接口
 */

//多线程并发下使用channel改造
//金额
var moneyC  = 1000
//var wait sync.WaitGroup
var lock sync.Mutex
//添加金额
func subtractMoneyLock(subMoney int) {
	lock.Lock()
	time.Sleep(3*time.Second)
	moneyC-=subMoney
	lock.Unlock()
}

//查询金额
func getMoneyLock() int {
	lock.Lock()
	result := moneyC
	lock.Unlock()
	return result;
}


func main() {
	//添加查询金额
	go func() {
		if(getMoneyLock()>200) {
			subtractMoneyLock(200)
			fmt.Printf("200元扣款成功，剩下：%d元\n",getMoneyLock())
		}else {
			fmt.Println("余额不足，无法扣款")
		}
	}()

	//添加查询金额
	go func() {
		if(getMoneyLock()>900) {
			subtractMoneyLock(900)
			fmt.Printf("900元扣款成功，剩下：%d元\n",getMoneyLock())
		}else {
			fmt.Println("余额不足，无法扣款")
		}
	}()
	//正常
	time.Sleep(5*time.Second)
	fmt.Println(getMoneyLock())
}
