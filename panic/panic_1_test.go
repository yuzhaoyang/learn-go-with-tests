package channel

import (
	"fmt"
	"testing"
	"time"
)

func Testpanic1(t *testing.T) {
	// 主线程中的defer函数并不会执行，因为子协程 panic后，主线程中的defer并不会执行
	defer println("in main")

	go func() {
		defer println("in goroutine")
		fmt.Println("子协程running")
		panic("子协程崩溃")
	}()

	time.Sleep(1 * time.Second)
}

func Testpanic2(t *testing.T) {
	defer fmt.Println("in main")
	if err := recover(); err != nil {
		fmt.Println(err)
	}

	panic("unknown err")
}

func Testpanic3(t *testing.T) {
	defer fmt.Println("in main")
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("occur error")
			fmt.Println(err)
		}
	}()

	panic("unknown err")
}

func Testpanic4(t *testing.T) {
	defer fmt.Println("in main")
	defer func() {
		defer func() {
			panic("panic again and again")
		}()
		panic("panic again")
	}()

	panic("panic once")
}
