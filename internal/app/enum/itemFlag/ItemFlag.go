package itemFlag

type Flag string

const (
	CutTreeTool     Flag = "cut tree"
	FishTool        Flag = "fish tool"
	MineTool        Flag = "mine tool"
	ResourceWood    Flag = "resource wood"
	ResourceFish    Flag = "resource fish"
	ResourceOre     Flag = "resource ore"
	Food            Flag = "food"
	IgnoreWaitstate Flag = "ignore waitstate"
	CanFly          Flag = "can fly"
	CliffWalk       Flag = "cliff walk"
)
