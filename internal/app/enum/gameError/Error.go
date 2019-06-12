package gameError

type Error string

const (
	NoTool                 Error = "no applicable tool"
	OreNotFound            Error = "ore not found"
	FishNotFound           Error = "fish not found"
	WrongBiom              Error = "biom is not applicable"
	RoomNotFound           Error = "room not found"
	RoomUnfordable         Error = "room unfordable"
	CaveNotFound           Error = "cave not found"
	RoomAlreadyExist       Error = "room already exist"
	CommandNotFound        Error = "command not found"
	WrongCommandAttributes Error = "wrong command attributes"
	WrongDirection         Error = "wrong direction"
)

func (error Error) Error() string {
	return string(error)
}
