package main

import (
	"fmt"
	"reflect"
)

func main() {
	myInt := 4
	fmt.Println(myInt) //打印初始值
	var myIntPointer *int
	myIntPointer = &myInt //取地址
	fmt.Println(reflect.TypeOf(myIntPointer))
	fmt.Println(myIntPointer) //打印指针地址

	*myIntPointer = 8

	fmt.Println(*myIntPointer) //打印指针指向的值
	fmt.Println(myInt)
}
