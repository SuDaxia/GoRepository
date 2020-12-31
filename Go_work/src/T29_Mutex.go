package main

import (
	"fmt"
	"sync"
)

/**
前面有讲过 chan 缓存不缓存都可以实现锁的机制还有没有额外的呢？

GO中也有互斥锁sync.Mutex，类似Java中的 ReentrantLock
	也有任务计数器sync.WaitGroup,类似Java中的 CountDownLatch
	也有读写锁sync.RWMutext,类似Java （读可以重入锁，但是写必须要等待所有读或单个写释放锁）

 */

var numb int
var wait1 sync.WaitGroup
func myAddNoSafe(){
	defer wait1.Done()
	for i:=0;i<10000;i++{
		numb+=1
	}
}

func test11() {
	wait1.Add(5)
	go myAddNoSafe()
	go myAddNoSafe()
	go myAddNoSafe()
	go myAddNoSafe()
	go myAddNoSafe()
	wait1.Wait()
	fmt.Printf("num = %d\n",numb)
}

var lock1 sync.Mutex
func myAddSafe(){
	defer wait1.Done()
	for i:=0;i<10000;i++{
		lock1.Lock()
		numb+=1
		lock1.Unlock()
	}
}

func test22(){

	wait1.Add(5)
	go myAddSafe()
	go myAddSafe()
	go myAddSafe()
	go myAddSafe()
	go myAddSafe()
	wait1.Wait()
	fmt.Printf("num = %d\n",numb)
}
func main() {
	test11()
	numb=0
	test22()
}

