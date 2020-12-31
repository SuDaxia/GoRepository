package main

import "fmt"

/**
Go与Java的数组对比
 */

func main() {
	//一维数组
	var arr1 = []int {1,2,34,6,89,}
	fmt.Println(arr1)
	//二维数组,第二维及后面纬度不能使用...
	var arr2=[...][4]int {{1,2,34,5},{24,5},{2345234}}
	fmt.Println(arr2)

	//三围数组呢
	var arr3=[...][3][4]int {
		{{1,3,4,5},{234,534,2,4}},
		{{1},{2}},
		{{4},{5}}}
	fmt.Println(arr3)
}