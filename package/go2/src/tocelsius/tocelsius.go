package main

import (
	"fmt"
	"hellogo/src/keyboard"
	"log"
)

func main() {
	fmt.Print("输入要转换的华氏度")
	num, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}

	C := (num - 32) * 5 / 9
	fmt.Printf("转换后的摄氏度为%.2f", C)
}
