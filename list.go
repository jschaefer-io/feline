package main

import (
	"errors"
	"fmt"
)

type List interface {
	add(value interface{}) int
	find(value interface{}) int
	has(value interface{}) bool
	len() int
}

type ArrayList struct {
	items []interface{}
}

func (list *ArrayList) len() int {
	return len(list.items)
}

func (list *ArrayList) find(value interface{}) int {
	for i, cmp := range list.items {
		if cmp == value {
			return i
		}
	}
	return -1
}

func (list *ArrayList) get(index int) (interface{}, error) {
	if index < 0 || index >= list.len() {
		return nil, errors.New(fmt.Sprintf("undefined index %d", index))
	}
	return list.items[index], nil
}

func (list *ArrayList) has(value interface{}) bool {
	return list.find(value) >= 0
}

func (list *ArrayList) add(value interface{}) int {
	index := list.find(value)
	if index >= 0 {
		return index
	}
	list.items = append(list.items, value)
	return list.len() - 1
}

type LinearList interface {
	List
	push(value interface{})
	pop() (interface{}, error)
}

type Queue struct {
	ArrayList
}

func (queue *Queue) push(value interface{}) {
	queue.items = append(queue.items, value)
}

func (queue *Queue) pop() (interface{}, error) {
	if queue.len() == 0 {
		return nil, errors.New("cannot pop from queue which is empty")
	}
	res := queue.items[0]
	queue.items = queue.items[1:]
	return res, nil
}

type Stack struct {
	ArrayList
}

func (stack *Stack) push(value interface{}) {
	stack.items = append(stack.items, value)
}

func (stack *Stack) pop() (interface{}, error) {
	if stack.len() == 0 {
		return nil, errors.New("cannot pop from stack which is empty")
	}
	res := stack.items[len(stack.items)-1]
	stack.items = stack.items[:len(stack.items)-1]
	return res, nil
}
