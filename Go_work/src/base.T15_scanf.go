package main

import (
	"bufio"
	"fmt"
	"os"
)
/**
Go的标准格式化输出库fmt和java的输出打印库对比
Java的标准输出流工具类是java.lang包下的System类，具体是其静态成员变量PrintStream类。

他有静态三个成员变量:

分别是PrintStream类型的out,in,err

我们常见System.out.println(),实际上调用的就是PrintStream类对象的println方法。
-----------------------------------------------------------

Go中的格式化输出输入库是fmt模块。

fmt在Go中提供了输入和输出的功能，类型Java中的Scanner和PrintStream(println)。

它的使用方法如下：
Print:   原样输出到控制台，不做格式控制。
Println: 输出到控制台并换行
Printf : 格式化输出(按特定标识符指定格式替换)
Sprintf：格式化字符串并把字符串返回，不输出，有点类似于Java中的拼接字符串然后返回。
Fprintf：来格式化并输出到 io.Writers 而不是 os.Stdout
 */

/**

读取控制台输入对象的方法
方法一：fmt.Scanf
方法二：bufio.NewScanner(os.Stdin)

 */

func getStringInputByFmt() string{
	var str string
	var str2 string
	var str3 string
	//命名返回的是读取变量的个数
	counts,err:=fmt.Scanf("%s%s%s",&str,&str2,&str3)
	//fmt.Scanf("%s",&str)
	if err!=nil{
		fmt.Println("[ERROR]\t",err.Error())
	}
	//这里内置的len方法的string长度，是看编码格式的，默认UTF-8中文字符3个byte，返回的byte的长度，并不是展示的字符个数
	fmt.Println("Read str:\t",str,"\t counts:",counts,"\tlen:\t", len(str))
	return str
}
/**
使用os.Stdin开启输入流,不只有Scan，还有Read、ReadByte ReadBytes ReadLine
ReadRune ReadSlice ReadString等等
len(string)求出的都是byte/int8的长度
字符串类各字符个数需要专程 rune 类型 len([]rune(str))
*/
func getStringInputByScanner()string{
	fmt.Println("请输入字符:")
	var str string
	in:=bufio.NewScanner(os.Stdin)
	if in.Scan(){
		str=in.Text()
	}else{
		str="Find input error"
	}
	fmt.Println("os.Strdin read str:\t",str,"\tlen:\t",len(str),"\tbytelens:\t",len([]byte(str)),"\t字符个数:\t",len([]rune(str)))
	return str
}

func showDisplayCharNums(str string){
	//"str:\t",str,"\tDisplayCharNums:\t"
}

func main() {
	getStringInputByFmt()
	getStringInputByScanner()
}