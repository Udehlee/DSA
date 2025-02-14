package main

import "fmt"

type stack struct {
	arr []int
}

func (s *stack) add(v int) {
	s.arr = append(s.arr, v)

}

// push adds value to the stack
func (s *stack) Push(value int) {
	s.arr = append(s.arr, value)
}

// pop removes value from the stack
func (s *stack) Pop() int {
	toRemove := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return toRemove
}

func (s *stack) Print() {
	for i := 0; i < len(s.arr); i++ {
		fmt.Println(s.arr[i])
	}
}

func main() {

	s := &stack{}
	s.Push(2)
	s.Push(3)
	s.Push(4)

	s.Print()

	fmt.Println("------------")

	s.Pop()
	s.Print()

}
