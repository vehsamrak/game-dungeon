package app

const ItemTypeCutTree = "cut_tree"
const ItemTypeResourceWood = "resource_wood"

type Item struct {
	itemTypes map[string]bool
}

func (item Item) Create() *Item {
	return &Item{itemTypes: make(map[string]bool)}
}

func (item *Item) AddType(itemType string) {
	item.itemTypes[itemType] = true
}

func (item *Item) HasType(itemType string) bool {
	return item.itemTypes[itemType]
}
