package main

import (
	"container/list"
	"fmt"
)

func main() {
	var data *list.List = list.New()

	data.PushBack("muhammad")
	data.PushBack("ivan")
	data.PushBack("satria")

	var head *list.Element = data.Front()
	fmt.Println(head.Value)

	next := head.Next()
	fmt.Println(next.Value)

	next = next.Next()
	fmt.Println(next.Value)

	fmt.Println(data.Len())

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
