package app

import "github.com/vehsamrak/game-dungeon/internal/app/enum/itemFlag"

type Character struct {
	name      string
	x         int
	y         int
	z         int
	health    int
	inventory []*Item
}

// Create new character
func (character Character) Create(name string) *Character {
	return &Character{name: name, health: 100}
}

// Name of character
func (character *Character) Name() string {
	return character.name
}

// X coordinate of character
func (character *Character) X() int {
	return character.x
}

// Y coordinate of character
func (character *Character) Y() int {
	return character.y
}

// Z coordinate of character (up and down)
func (character *Character) Z() int {
	return character.z
}

// Move character to given X and Y coordinates
func (character *Character) Move(x int, y int, z int) {
	character.x = x
	character.y = y
	character.z = z
}

// Inventory of character
func (character *Character) Inventory() []*Item {
	return character.inventory
}

// AddItems adds several items to characters inventory
func (character *Character) AddItems(items []*Item) {
	character.inventory = append(character.inventory, items...)
}

// AddItem adds one item to characters inventory
func (character *Character) AddItem(item *Item) {
	character.inventory = append(character.inventory, item)
}

// HasItemFlag checks character inventory to have given item flag
func (character *Character) HasItemFlag(itemFlag itemFlag.Flag) bool {
	for _, item := range character.inventory {
		if item.HasFlag(itemFlag) {
			return true
		}
	}

	return false
}
func (character *Character) Health() int {
	return character.health
}
