package main

import "fmt"

func add(x, y int) (int,int) {
	return x + y, 0
}

func main() {
	_, b := add(23,42)
	fmt.Println(b)
}