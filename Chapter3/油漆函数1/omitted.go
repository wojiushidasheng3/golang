package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var something int

func main() {
	something = 10
	Printf()

}

func Printf() {
	fmt.Println("输入长")
	reader := bufio.NewReader(os.Stdin)
	input1, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("输入宽")
	input2, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	input1 = strings.TrimSpace(input1)
	input2 = strings.TrimSpace(input2)
	Newinput1, err := strconv.ParseFloat(input1, 64)
	if err != nil {
		log.Fatal(err)
	}
	Newinput2, err := strconv.ParseFloat(input2, 64)
	if err != nil {
		log.Fatal(err)
	}
	area := Newinput1 * Newinput2
	fmt.Printf("油漆面积为：%5.3f\n", area)

	fmt.Printf("公摊物品空间为: %5.3f", area*float64(something))
}
