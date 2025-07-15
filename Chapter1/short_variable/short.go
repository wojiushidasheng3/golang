package main

import (
	// 修改为相对模块根目录的路径
	"fmt"
	short_variable "hellogo/short_variable/utils"
	"reflect"
)

func main() {
	quantity := 4
	length, width := 1.2, 2.4
	cust := "hello"

	fmt.Println(quantity)
	fmt.Println(width * length)
	fmt.Println(cust)

	fmt.Println(reflect.TypeOf(quantity))
	fmt.Println(reflect.TypeOf(length))
	fmt.Println(reflect.TypeOf(cust))

	// 调用 utils 包中的 Add 函数
	result := short_variable.Add(2, 3)
	fmt.Println("2 + 3 =", result)
}
