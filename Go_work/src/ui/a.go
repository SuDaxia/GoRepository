//这里的ui，也就是package后面的名称尽量和包文件夹的名称一致，不一致也可以
package ui

import "fmt"
var UIName = "UI UP Char start"
var uiname = "ui down char start no use out of this package"
func StartGameUI(){
	fmt.Println("UI go start")
	downCaseIsPrivateOnlyCurrentPackageUse()
}

func downCaseIsPrivateOnlyCurrentPackageUse(){
	fmt.Println("ui downCaseIsPrivateOnlyCurrentPackageUse")
}
