package data_types

import "errors"

type Queue struct {
	ArrayList
}

func (queue *Queue) Push(value interface{}) {
	queue.items = append(queue.items, value)
}

func (queue *Queue) Pop() (interface{}, error) {
	if queue.Len() == 0 {
		return nil, errors.New("cannot pop from queue which is empty")
	}
	res := queue.items[0]
	queue.items = queue.items[1:]
	return res, nil
}
