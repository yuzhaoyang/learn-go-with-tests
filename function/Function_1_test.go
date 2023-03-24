package fucntiontest

import (
	"fmt"
	"testing"
)

func TestFuction(t *testing.T) {

	s := "a"
	fmt.Printf("outer: %v, %v\n", s, &s)
	test_string(s)
	fmt.Printf("outer: %v, %v\n", s, &s)
}

func test_string(s string) {
	fmt.Printf("inner: %v, %v\n", s, &s)
	s = "b"
	fmt.Printf("inner: %v, %v\n", s, &s)
}

func test_map(m map[string]string) {
	fmt.Printf("inner: %v, %p\n", m, m)
	m["a"] = "11"
	fmt.Printf("inner: %v, %p\n", m, m)
}

func TestFuction2(t *testing.T) {

	m := map[string]string{
		"a": "1",
		"b": "2",
		"c": "3",
	}

	fmt.Printf("outer: %v, %p\n", m, m)
	test_map(m)
	fmt.Printf("outer: %v, %p\n", m, m)

}

func test_map2(m map[string]string) {
	fmt.Printf("inner: %v, %p\n", m, m)
	m = make(map[string]string, 0)
	m["a"] = "11"
	fmt.Printf("inner: %v, %p\n", m, m)
}

func TestFuction3(t *testing.T) {

	var m map[string]string //未初始化

	fmt.Printf("outer: %v, %p\n", m, m)
	test_map2(m)
	fmt.Printf("outer: %v, %p\n", m, m)

}

func test_chan2(ch chan string) {
	fmt.Printf("inner: %v, %v\n", ch, len(ch))
	ch <- "b"
	fmt.Printf("inner: %v, %v\n", ch, len(ch))
}

func TestFuction4(t *testing.T) {
	ch := make(chan string, 10)
	ch <- "a"

	fmt.Printf("outer: %v, %v\n", ch, len(ch))
	test_chan2(ch)
	fmt.Printf("outer: %v, %v\n", ch, len(ch))

}

func test_slice(sl []string) {
	fmt.Printf("%v, %p\n", sl, sl)
	sl[0] = "aa"
	//sl = append(sl, "d")
	fmt.Printf("%v, %p\n", sl, sl)
}

func TestFuction5(t *testing.T) {
	sl := []string{
		"a",
		"b",
		"c",
	}

	fmt.Printf("%v, %p\n", sl, sl)
	test_slice(sl)
	fmt.Printf("%v, %p\n", sl, sl)

}

// 浅拷贝
func TestFuction6(t *testing.T) {

	m1 := map[string]string{
		"a": "1",
		"b": "2",
		"c": "3",
	}

	m2 := m1

	m2["a"] = "tmp"

	fmt.Printf("m1: %v\n", m1["a"])

	fmt.Printf("m2: %v\n", m2["a"])

}

// 深拷贝
func TestFuction7(t *testing.T) {

	m1 := map[string]string{
		"a": "1",
		"b": "2",
		"c": "3",
	}

	m2 := map[string]string{}

	for k, v := range m1 {
		m2[k] = v
	}

	m2["a"] = "tmp"

	fmt.Printf("m1: %v\n", m1["a"])

	fmt.Printf("m2: %v\n", m2["a"])

}
