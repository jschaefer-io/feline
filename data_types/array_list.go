package data_types

import (
	"errors"
	"fmt"
)

type ArrayList struct {
	items []interface{}
}

func (list *ArrayList) Len() int {
	return len(list.items)
}

func (list *ArrayList) Find(value interface{}) int {
	for i, cmp := range list.items {
		if cmp == value {
			return i
		}
	}
	return -1
}

func (list *ArrayList) Get(index int) (interface{}, error) {
	if index < 0 || index >= list.Len() {
		return nil, errors.New(fmt.Sprintf("undefined index %d", index))
	}
	return list.items[index], nil
}

func (list *ArrayList) Has(value interface{}) bool {
	return list.Find(value) >= 0
}

func (list *ArrayList) Add(value interface{}) int {
	index := list.Find(value)
	if index >= 0 {
		return index
	}
	list.items = append(list.items, value)
	return list.Len() - 1
}
