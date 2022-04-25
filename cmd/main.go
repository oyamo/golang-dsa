package main

import (
	"dsa-golang/stack"
	"fmt"
)

func main() {
	s := stack.Stack[int]{}

	for i := 0; i < 10; i++ {
		s.Append(i)
	}

	fmt.Println(s.Pop().Value)
	fmt.Println(s.Pop().Value)
	fmt.Println(s.Pop().Value)

}
