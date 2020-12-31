package main

import (
	"fmt"
	"time"
)

/**
无缓存的channel可以起到一个多线程间线程数据同步锁安全的作用。

缓存的channel创建方式是

make(chan string,缓存个数)

缓存个数是指直到多个数据没有消费或者接受后才进行阻塞。

类似于java中的synchronized和lock

可以保证多线程并发下的数据一致性问题。

 */

/**
首先我们看一个线程不安全的代码示例：↓
 */
var moneyA int =1000

func subtractMoney(subMoney int){
	time.Sleep(3*time.Second)
	moneyA-=subMoney
}
func getMoney() int {
	return moneyA
}
func noSafe() {
	go func() {
		if getMoney() > 200 {
			subtractMoney(200)
			fmt.Printf("200取现成功，剩下:%d元\n", getMoney())
		}
	}()

	//添加金额
	go func(){
		if getMoney()>900{
			subtractMoney(900)
			fmt.Printf("900取现成功，剩下:%d元\n", getMoney())
		}
	}()

	//正常逻辑只够扣款成功一单，现在扣出负数来了；怎么解决呢，看用缓存chan 起了锁作用
	time.Sleep(5*time.Second)
	fmt.Println(getMoney())
}
/**
测试证明，还是尽量用缓存的chan，获得不到锁，老排队，都不知道是不是死锁
 */
func chanSafe(){
	var synchLock = make(chan int,1) //缓存，有上锁效果;Astart/Bstart有，end处都注释有上锁效果
	//var synchLock = make(chan int) //不缓存的,而且仅需Astart/Bstart中有一处有另一处注释即有上锁效果，
									// Astart/Bstart两处都上锁end处注释，竞争，知道main结束都不见取款
	go func() {
		synchLock<-1 //Astart 这里注释掉，就算其他三处没注释掉，也是不安全的【只要两处协程方法里开头用都算上锁了】
		if(getMoney()>200) {
			subtractMoney(200)
			fmt.Printf("200元扣款成功，剩下：%d元\n",getMoney())
		}
		//奇奇怪怪的操作，不太懂这里哦
		<-synchLock //Aend
	}()

	//添加查询金额
	go func() {
		synchLock<-10 //Bstart 这里注释掉，就算其他三处没注释掉，也是不安全的
		if(getMoney()>900) {
			subtractMoney(900)
			fmt.Printf("900元扣款成功，剩下：%d元\n",getMoney())
		}
		synchLock<-9 //Bend
	}()
	//这样类似于java中的Lock锁，不会扣多
	time.Sleep(5*time.Second)
	fmt.Println(getMoney())
}

func main() {
	//noSafe()
	moneyA=1000
	fmt.Println("----------")
	chanSafe()
}