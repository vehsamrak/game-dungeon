package app

const ItemFlagCutTree = "cut_tree"
const ItemFlagResourceWood = "resource_wood"

type Item struct {
	flags map[string]bool
}

func (item Item) Create() *Item {
	return &Item{flags: make(map[string]bool)}
}

func (item *Item) AddFlag(flag string) {
	item.flags[flag] = true
}

func (item *Item) HasFlag(flag string) bool {
	return item.flags[flag]
}
