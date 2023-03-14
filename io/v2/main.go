package main

import (
	"errors"
	"fmt"
	"time"
)

type response struct {
	code int
}

func foo(x int) interface{} {
	var resp *response = nil
	if x < 0 {
		return resp
	}

	return &response{0}
}

func main() {
	//server := NewPlayerServer(NewInMemoryPlayerStore())
	//log.Fatal(http.ListenAndServe(":5000", server))

	//var x *int
	//var i interface{}
	//
	//fmt.Printf("%v %v\n", x, x == nil)       // <nil> true
	//fmt.Printf("%T %v %v\n", i, i, i == nil) // <nil> <nil> true
	//
	//i = x
	//
	//fmt.Printf("%T %v %v\n", i, i, i == nil) // *int <nil> false

	var i interface{} = nil

	fmt.Printf("%T %v %v\n", i, i, i == nil) // *int <nil> false

	f := foo(-1)

	fmt.Printf("%T %v %v\n", f, f, f == nil) // *int <nil> false

	a := []int{1, 2, 3}
	for _, x := range a {

		fmt.Println("for  %v\n", x) // *int <nil> false
		go func() {
			fmt.Println(x)

		}()
	}
	time.Sleep(3 * time.Second)

	//mp := map[interface{}]string{
	//	1: "1",
	//	2: "2",
	//}
	//buf, err := json.Marshal(mp)
	//if err != nil {
	//	fmt.Println(err) // 报错：json: unsupported type: map[interface {}]string
	//	return
	//}
	//fmt.Println(string(buf))

	fmt.Println(foo1())

}

func foo1() (err error) {
	if err := test(); err != nil {
		return err // 报错：err is shadowed during return
	}
	return
}

func test() error {
	return errors.New("hello world")
}
