package main

import "fmt"

func main() {
	fmt.Printf("The %s cost %d cents each.\n", "gumballs", 23)
	fmt.Printf("That will be $%f please.\n", 0.23*5)

	fmt.Printf("A float: %f\n", 3.1415)
	fmt.Printf("An int: %d\n", 15)
	fmt.Printf("A string: %s\n", "hello")
	fmt.Printf("A boolean: %t\n", false)
	fmt.Printf("values: %v %v %v\n", 1.2, "\t", true) //输出任意格式的值

	s := "hello"
	slice := []int{1, 2, 3}
	fmt.Printf("%#v, %#v\n", s, slice) //输出完整名字

	fmt.Printf("Types: %T %T %T\n", 1.2, "\t", true) //输出对应类型 ~reflect.TypeOf("\t")
	fmt.Printf("sign: %%\n")

}
