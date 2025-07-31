package main

import "fmt"

type StringStack struct {
	data []string
}

func (s *StringStack) push(x string) {
	s.data = append(s.data, x)
}
func (s *StringStack) pop() string {
	if len := len(s.data); len > 0 {
		t := s.data[len-1]
		s.data = s.data[:len-1]
		return t
	}
	panic("pop from Empty Stack")
}

func main() {
	var stack StringStack
	stack.push("sf")
	stack.push("fdg")
	stack.pop()
	stack.push("sffh")
	fmt.Println(stack)
}