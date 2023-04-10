package typetest

import (
	"fmt"
	"testing"
)

func foo() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	var i interface{} = "abc"
	v, ok := i.([]string)

	fmt.Println(v, ok)
}

func TestType(t *testing.T) {
	foo()
	fmt.Println("exit")
}
