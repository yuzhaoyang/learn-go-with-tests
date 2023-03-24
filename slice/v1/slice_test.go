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

	fmt.Printf("s: %+v\n", s) // s: {Name:Tom Score:60}

}

func changeSlice1(s []*Student) {

	fmt.Printf("s: %+v\n", s) // s: {Name:Tom Score:60}
	tmp := &Student{Name: "yuzhaoyang1", Score: 2}

	s = append(s, tmp)
	fmt.Printf("s: %+v\n", s) // s: {Name:Tom Score:60}
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
