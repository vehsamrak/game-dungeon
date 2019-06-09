package gameError

type Error string

const (
	NoTool         Error = "no applicable tool"
	OreNotFound    Error = "ore not found"
	FishNotFound   Error = "fish not found"
	WrongBiom      Error = "biom is not applicable"
	RoomNotFound   Error = "room not found"
	RoomUnfordable Error = "room unfordable"
)

func (error Error) Error() string {
	return string(error)
}
