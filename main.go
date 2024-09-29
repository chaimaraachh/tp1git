package main

import (
	"fmt"
)

func main() {
var (
s string = "foo"
i int = 5)

fmt.Println(s)
fmt.Println(i)

// Getting user input

var name string
fmt.Print ("Enter your name: ")
fmt.Scanf ("%s", &name)
fmt.Println ("Hey there, ", name)


// using count and err 

var a string
var b int
fmt. Print ("Enter a string and a number: ")
count, err := fmt.Scanf ("%s %d", &a, &b)
fmt.Println("count : ", count)
fmt.Println("error: ", err)
fmt.Println("a: ", a)
fmt.Println("b: ",b)

}