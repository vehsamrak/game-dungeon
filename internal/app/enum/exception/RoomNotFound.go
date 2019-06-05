package exception

type RoomNotFound struct {
}

func (error RoomNotFound) Error() string {
	return "room not found"
}
