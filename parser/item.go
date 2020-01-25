package parser

type Item interface{}

type ItemGroup interface {
	add(item Item)
}