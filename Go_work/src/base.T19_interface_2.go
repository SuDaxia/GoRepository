package main
/**
如果是别的类里面实现 接口呢？？
同package，直接拿来用，，如果用 var aa InterfaceStruct = ImplementStruct{}，接口中方法要全实现
不同package，要引用，从src为根目录开始弄
要注意，外package能用方法 属性，都要大写开头
 */
import (
	"fmt"
	"otherinterface"
)
type Tiger struct {

}

func (tiger Tiger) Nothing() {
	fmt.Println("tiger nothing")
}

func (tiger Tiger)Eat(){

}
func (tiger Tiger)sleep(){

}

func main() {
/*	var tiger Animal=Tiger{}
	tiger.Eat()*/
	var tiger otherinterface.BaseInterface = Tiger{}
	tiger.Nothing()

}
