package commands

import (
	"github.com/vehsamrak/game-dungeon/internal/app"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/itemFlag"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/timer"
	"time"
)

type Character interface {
	Name() string
	X() int
	Y() int
	Z() int
	Move(x int, y int, z int)
	Inventory() []*app.Item
	AddItems(items []*app.Item)
	AddItem(item *app.Item)
	HasItemFlag(itemFlag itemFlag.Flag) bool
	Health() int
	MaxHealth() int
	LowerHealth(healthPoints int)
	IncreaseHealth(healthPoints int)
	FindItemWithFlag(flag itemFlag.Flag) *app.Item
	DropItem(item *app.Item)
	RestoreHealth()
	Timer(timerName timer.Timer) (timeLeft time.Duration)
	TimerActive(timerName timer.Timer) bool
	HasActiveTimers() bool
	SetTimer(timer timer.Timer, timeDuration time.Duration)
	DropTimer(timer timer.Timer)
	ResetTimer(timer timer.Timer)
	FullHealth() bool
}
