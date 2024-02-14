package main

import "fmt"

func Hello() string {
	fmt.Println("hello, world!")
	return "hello"
}

func main() {
	msg := Hello()
	fmt.Printf("the function from Hello says %v", msg)
}
