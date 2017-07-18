package main

import (
	"fmt"
	"container/list"
	"sort"
)

func main() {
	x := list.New()

	x.PushBack(1)
	x.PushBack(2)
	x.PushBack(3)

	sort.Sort(x)
	for e := x.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(int))
	}
}
