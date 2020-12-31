package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/**
Go中的AtomicXXX原子操作类(类似于Java中的AtocmicInteger之类的)
Go中的atomic包里面的功能和Java中的Atomic一样，原子操作类，原理也是cas,甚至提供了cas的api函数

Go中的WaitGroup(类似于Java中的CountDownLatch)
 */
func main() {
	var num int64 =20
	wait3:=sync.WaitGroup{}
	wait3.Add(3)
	go func() {
		defer wait3.Done()
		atomic.AddInt64(&num,2)
	}()
	go func() {
		defer wait3.Done()
		atomic.AddInt64(&num,2)
	}()
	go func() {
		defer wait3.Done()
		atomic.AddInt64(&num,2)
	}()
	wait3.Wait()
	fmt.Println(num)
}
