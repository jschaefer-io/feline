package data_types

import "errors"

type Stack struct {
	ArrayList
}

func (stack *Stack) Push(value interface{}) {
	stack.items = append(stack.items, value)
}

func (stack *Stack) Pop() (interface{}, error) {
	if stack.Len() == 0 {
		return nil, errors.New("cannot pop from stack which is empty")
	}
	res := stack.items[len(stack.items)-1]
	stack.items = stack.items[:len(stack.items)-1]
	return res, nil
}
