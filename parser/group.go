package parser

type Group struct {
	Items []Item
}

func (group *Group) add(item Item) {
	group.Items = append(group.Items, item)
}