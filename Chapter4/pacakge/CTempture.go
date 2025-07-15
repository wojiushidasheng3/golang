package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)
	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal(err)
	}

	return number, nil

}

func main() {
	fmt.Print("输入要转换的华氏度")
	num, err := getFloat()
	if err != nil {
		log.Fatal(err)
	}

	C := (num - 32) * 5 / 9
	fmt.Printf("转换后的摄氏度为%.2f", C)
}
