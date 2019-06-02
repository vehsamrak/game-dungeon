package app

type Character struct {
	name      string
	x         int
	y         int
	inventory []*Item
}

func (character *Character) HasItemFlag(itemFlag string) bool {
	for _, item := range character.inventory {
		if item.HasFlag(itemFlag) {
			return true
		}
	}

	return false
}

func (character *Character) Inventory() []*Item {
	return character.inventory
}

func (character *Character) Name() string {
	return character.name
}

func (character *Character) X() int {
	return character.x
}

func (character *Character) Y() int {
	return character.y
}

func (character *Character) Move(x int, y int) error {
	character.x = x
	character.y = y

	return nil
}

func (character *Character) AddItems(items []*Item) {
	character.inventory = append(character.inventory, items...)
}

func (character *Character) AddItem(item *Item) {
	character.inventory = append(character.inventory, item)
}
