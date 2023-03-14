package channel

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
