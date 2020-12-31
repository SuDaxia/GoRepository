package main
/**
结构体中的属性和方法怎么获取呢？

获取结构体属性的个数是先ValueOf获取结构体值对象v后，用v.NumField()获取该结构体有几个属性，
通过v.Field(i)来获取对应位置的属性的元类型。
 */

