package main

import (
	"log"

	"github.com/navneetshukl/stack/stack"
)

func main() {
	arr := []int{}
	service := stack.NewService(arr)

	service.Push(10)
	ans, _ := service.Top()
	log.Println("Top element is ", ans)
}
