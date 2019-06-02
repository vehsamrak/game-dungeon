package app

const ItemTypeCutTree = "cut_tree"

type Item struct {
	itemTypes map[string]bool
}

func (item *Item) HasType(itemType string) bool {
	return item.itemTypes[itemType]
}
