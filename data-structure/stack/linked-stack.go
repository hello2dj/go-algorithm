package stack

import "errors"

type node struct {
	value interface{}
	next  *node
}

type LinkedStack struct {
	top  *node
	size int
}

func (stack *LinkedStack) Size() int {
	return stack.size
}

func (stack *LinkedStack) Push(value interface{}) {
	stack.top = &node{value, stack.top}
	stack.size++
}

func (stack *LinkedStack) Pop() interface{} {
	if stack.size == 0 {
		panic("Stack is empty.")
	}

	top_node := stack.top
	stack.top = top_node.next
	value := top_node.value
	top_node.next = nil
	top_node.value = nil
	stack.size--
	return value
}

func (stack *LinkedStack) Peek() (interface{}, error) {
	if stack.size == 0 {
		return nil, errors.New("no element")
	}
	return stack.top.value, nil
}
