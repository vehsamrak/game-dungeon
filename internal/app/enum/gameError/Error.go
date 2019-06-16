package gameError

type Error string

const (
	NoTool                 Error = "no applicable tool"
	OreNotFound            Error = "ore not found"
	FishNotFound           Error = "fish not found"
	WrongBiom              Error = "biom is not applicable"
	RoomNotFound           Error = "room not found"
	RoomUnfordable         Error = "room unfordable"
	CantMoveInWater        Error = "can't move in water"
	CaveNotFound           Error = "cave not found"
	RoomAlreadyExist       Error = "room already exist"
	CommandNotFound        Error = "command not found"
	WrongCommandAttributes Error = "wrong command attributes"
	WrongDirection         Error = "wrong direction"
	LowHealth              Error = "health is too low"
	FoodNotFound           Error = "food not found"
	WaitState              Error = "wait before next command"
	HealthFull             Error = "health already full"
)

func (error Error) Error() string {
	return string(error)
}
