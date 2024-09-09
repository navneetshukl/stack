package main

import (
	"log"

	stack "github.com/navneetshukl/stack/internal"
)

func main() {
	arr := []int{}
	service := stack.NewService(arr)

	service.Push(10)
	ans, _ := service.Top()
	log.Println("Top element is ", ans)
}
