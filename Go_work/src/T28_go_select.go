package main

import "fmt"

/**
Go中 利用chan 缓存，写满到一定缓存，那么就会触发阻塞，
无法满足读取多个chan的场景，因为被阻塞了
怎么解决呢？java中是用select，GO中也有

 */

func main() {
	myChan :=make(chan int ,1)
	for i:=1;i<=100;i++{
		//select 的用法是，从上到下依次判断case 是否可执行，如果可执行，则执行完毕跳出select,如果不能执行，尝试下一个执行
		//这里的可执行是指的不阻塞，也就是说，select从上到下开始挑选一个不阻塞的case执行，执行完毕后跳出，
		//如果所有case都阻塞，则执行default
		//如下输出结果，i=奇数的时候走case   myChan<-i:，把奇数放入mychan
		//走偶数的时候因为myChan中有数据了，则把上一个奇数打印出来。
		//所以结果是 1  3  5  7  ...
		select{
			/**
			不管 是写入case在前，还是读取case在前，结果都是一样的，因为没有读到数据的就是阻塞，然后后面的去写，写完了一个的就阻塞，等到另一个读消费掉
			 */
			case data:=<-myChan:
			fmt.Println(data)
			case myChan<-i:
				fmt.Println("写入数据")
			default:
				fmt.Println("default...")
		}
	}
}
