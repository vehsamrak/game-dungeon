package app

import "github.com/vehsamrak/game-dungeon/internal/app/enum/itemFlag"

type Character struct {
	name      string
	x         int
	y         int
	z         int
	maxHealth int
	health    int
	inventory []*Item
}

// Create new character
func (character Character) Create(name string) *Character {
	maxHealth := 100

	return &Character{name: name, maxHealth: maxHealth, health: maxHealth}
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
func (character *Character) HasItemFlag(flag itemFlag.Flag) bool {
	return character.FindItemWithFlag(flag) != nil
}

// Health points of character
func (character *Character) Health() int {
	return character.health
}

// MaxHealth is maximum health points of character
func (character *Character) MaxHealth() int {
	return character.maxHealth
}

// RestoreHealth sets character health to maximum
func (character *Character) RestoreHealth() {
	character.health = character.maxHealth
}

// LowerHealthOnError lowers character health
func (character *Character) LowerHealth(healthPoints int) {
	character.health -= healthPoints
}

// IncreaseHealth increases character health
func (character *Character) IncreaseHealth(healthPoints int) {
	character.health += healthPoints
}

// FindItemWithFlag returns item with flag from character inventory
func (character *Character) FindItemWithFlag(flag itemFlag.Flag) *Item {
	for _, item := range character.inventory {
		if item.HasFlag(flag) {
			return item
		}
	}

	return nil
}

// DropItem drops item from character inventory
func (character *Character) DropItem(item *Item) {
	for i, inventoryItem := range character.inventory {
		if inventoryItem == item {
			character.inventory = append(character.inventory[:i], character.inventory[i+1:]...)
		}
	}
}
