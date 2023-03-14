package channel

import (
	"fmt"
	"testing"
)

type Student struct {
}

func (s *Student) printName() {
	fmt.Println("Tom") // Tom
}

func TestStruct(t *testing.T) {

	var s *Student
	fmt.Println("s == nil?", s == nil) // s == nil? true
	s.printName()
}
