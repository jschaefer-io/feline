package data_types

type List interface {
	Add(value interface{}) int
	Find(value interface{}) int
	Has(value interface{}) bool
	Len() int
}

type LinearList interface {
	List
	push(value interface{})
	pop() (interface{}, error)
}
