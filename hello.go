package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("hello world,世界")
	var a = 78
	b := string(a)
	fmt.Println(b)
	var c string
	c = strconv.Itoa(a)
	fmt.Println(c)

	bc, _ := strconv.Atoi(c)
	fmt.Println(bc)

}
