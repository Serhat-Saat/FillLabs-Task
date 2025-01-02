package main

import "fmt"

//Stack
type Stack struct {
	items []int
}

// Push adds an element to the stack
func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

// Pop, returns the top element of the stack
func (s *Stack) Pop() int {
	if len(s.items) == 0 {
		return -1 // Return -1 if the stack is empty
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

// Recursive function processes the number and adds it to the stack
func processNumberRecursive(n int, stack *Stack) {
	if n < 2 {
		return
	}

	// First add n to the stack
	stack.Push(n)

	if n%2 == 0 {
		n = n / 2
	} else {
		n = n - 1
		n = n / 2
	}
	processNumberRecursive(n, stack)
}

func main() {
	n := 9
	stack := &Stack{items: []int{}}

	processNumberRecursive(n, stack)

	// To print the elements in the stack in reverse order
	for len(stack.items) > 0 {
		// Pop from stack and print
		fmt.Println(stack.Pop())
	}
}
