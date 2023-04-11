package slice

import (
	"fmt"
	"testing"
)

type Student struct {
	Name  string
	Score int
}

func TestChannel(t *testing.T) {
	s := Student{"Tom", 59}

	students := []Student{s}

	s.Score = 60
	fmt.Printf("s: %+v\n", s) // s: {Name:Tom Score:60}
	fmt.Printf("students[0]: %+v\n", students[0])
}

func TestChannel2(t *testing.T) {
	s := &Student{"Tom", 59}
	students := []*Student{s}
	s.Score = 60
	fmt.Printf("s: %+v\n", s) // s: {Name:Tom Score:60}
	fmt.Printf("students[0]: %+v\n", students[0])
}

func TestSlice(t *testing.T) {

	s := []int{2}
	changeSlice(s)

	fmt.Printf("s: %+v\n", s) // s: {Name:Tom Score:60}

}

func changeSlice(s []int) {

	s = append(s, 1)
}

func TestSlice1(t *testing.T) {

	s := []*Student{{Name: "yuzhaoyang", Score: 1}}
	fmt.Printf("s: %+v\n", s) // s: {Name:Tom Score:60}
	changeSlice1(s)

	fmt.Printf("s: %+v\n", len(s)) // s: {Name:Tom Score:60}

}

func changeSlice1(s []*Student) {

	fmt.Printf("s: %+v\n", s) // s: {Name:Tom Score:60}
	tmp := &Student{Name: "yuzhaoyang1", Score: 2}

	s = append(s, tmp)
	fmt.Printf("s: %+v\n", len(s)) // s: {Name:Tom Score:60}
}

func TestSlice2(t *testing.T) {

	s := []Student{{Name: "yuzhaoyang", Score: 1}}
	fmt.Printf("s: %+v\n", s) // s: {Name:Tom Score:60}
	changeSlice2(s)

	fmt.Printf("s: %+v\n", s) // s: {Name:Tom Score:60}

}

func changeSlice2(s []Student) {

	fmt.Printf("s: %+v\n", s) // s: {Name:Tom Score:60}

	s[0].Name = "tom"
	fmt.Printf("s: %+v\n", s) // s: {Name:Tom Score:60}
}

func TestSlice3(t *testing.T) {

	values := []int{1, 2, 3}

	for i, value := range values {
		fmt.Println(i)
		fmt.Println(&value)
	}

}

func TestSlice4(t *testing.T) {

	var name []string

	name[0] = "1"

	fmt.Println(name)

}

// break 退出内层循环
func TestSlice5(t *testing.T) {

	for i := 0; i < 10; i++ {

		fmt.Println("iiiiiiii")
		fmt.Println(i)

		for j := 0; j < 10; j++ {

			fmt.Println("xxxxxx")
			fmt.Println(j)

			if j == 0 {

				break

			}

		}

	}

}

func TestSlice6(t *testing.T) {

	in := []int{1, 2, 3}

	var out []*int

	for _, v := range in {

		out = append(out, &v)

	}

	fmt.Println("values:", *out[0], *out[1], *out[2])
	fmt.Println("address:", out[0], out[1], out[2])

}

func TestSlice7(t *testing.T) {

	in := []int{1, 2, 3}

	var out []*int

	for _, v := range in {

		tmp := v

		out = append(out, &tmp)

	}

	fmt.Println("values:", *out[0], *out[1], *out[2])
	fmt.Println("address:", out[0], out[1], out[2])

}

func TestSlice8(t *testing.T) {

	s0 := &Student{"name0", 1}
	s1 := &Student{"name1", 2}
	s2 := &Student{"name2", 2}

	in := []*Student{s0, s1, s2}

	var out []*Student

	for _, v := range in {

		out = append(out, v)

	}

	fmt.Println("values:", *out[0], *out[1], *out[2])
	fmt.Println("address:", out[0], out[1], out[2])

}
