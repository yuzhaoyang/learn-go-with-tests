package channel

import (
	"fmt"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	ch := make(chan string)

	go func() {
		for v := range ch {
			fmt.Println("Received:", v) // 只打印一行 a
			time.Sleep(1 * time.Second)
		}
	}()

	ch <- "a"
	ch <- "b" // 不会被接收处理
}

func TestChannel2(t *testing.T) {
	var ch chan int
	fmt.Println("ch == nil?", ch == nil) // ch == nil? true
	go func() {
		ch <- 1
	}()
	fmt.Println("element of ch: ", <-ch)
}
