package main

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	x := map[int]int{}
	for i := 0; i < 10000; i++ {
		x[i] = i
	}
	fmt.Println("初始化后,长度:", len(x))

	// 遍历时删除所有的偶数
	for k := range x {
		if k%2 == 0 {
			delete(x, k)
		}
	}
	fmt.Println("删除所有的偶数后,长度:", len(x))

	// 遍历时删除所有的元素
	for k := range x {
		delete(x, k)
	}
	fmt.Println("删除所有的元素后,长度:", len(x))
}
