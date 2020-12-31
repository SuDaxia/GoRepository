package main

import (
	"fmt"
	"sync"
	"time"
)

/**
还有读写锁  sync.RWMutex
 */
var moneyA2=1000
var rwLock sync.RWMutex
var wait2 sync.WaitGroup

func subtractMoney2(subMoney int){
	rwLock.Lock()
	time.Sleep(3*time.Second)
	moneyA2-=subMoney
	rwLock.Unlock()
}

func getMoney2()int{
	rwLock.RLock()
	result:=moneyA2
	rwLock.RUnlock()
	return result
}

func main() {
	wait1:=sync.WaitGroup{}
	wait1.Add(2)
	go func() {
		defer wait1.Done()
		if getMoney2()>200{
			subtractMoney2(200)
			fmt.Printf("200元扣款成功,剩下：%d元\n",getMoney2())
		}
	}()

	go func() {
		defer wait1.Done()
		if getMoney2()>900{
			subtractMoney2(900)
			fmt.Printf("900元扣款成功，剩下：%d元\n",getMoney2())
		}
	}()
	wait1.Wait()
	fmt.Printf("main最终:%d元",getMoney2())
}
