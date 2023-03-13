package main

import (
	"fmt"
	"strings"
)

func main() {
	var sayHello string = "selamat malam"
	char := strings.Split(sayHello, "")

	count := make(map[string]int)

	for _, v := range char {
		fmt.Printf("%v \n", v)
	}
	for _, v := range char {
		count[v]++
	}
	fmt.Println(count)

}
