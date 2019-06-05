package exception

type RoomUnfordable struct {
}

func (error RoomUnfordable) Error() string {
	return "room unfordable"
}
