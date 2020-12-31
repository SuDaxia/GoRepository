//package other //此种状态下，如何使用最外面src目录下Go文件中定义的interface结构体，还不知道
package main

import (
	"fmt"
	"otherinterface"
)

/**
测试别的包继承
 */

type Snake struct {

}

func(snake Snake) Nothing() {
	fmt.Println("snake nothing ")
}


func test1() {
	var snake otherinterface.BaseInterface = Snake{}
	snake.Nothing()
}
func main() {
	//var sanke main.Animal =Snake{}
	//	test1()

}