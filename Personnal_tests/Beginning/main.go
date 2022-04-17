package main

import (
	"fmt"
)

func main()  {
	x := "Hello"
	fmt.Printf("%T\n",x)
	x = x+" World"
	fmt.Println(x,33)
}