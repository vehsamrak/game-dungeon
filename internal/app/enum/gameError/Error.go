package gameError

type Error string

const (
	NoTool         Error = "no applicable tool"
	OreNotFound    Error = "ore not found"
	RoomNotFound   Error = "room not found"
	RoomUnfordable Error = "room unfordable"
)

func (error Error) Error() string {
	return string(error)
}
