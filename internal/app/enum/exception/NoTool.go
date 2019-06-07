package exception

type NoTool struct {
}

func (error NoTool) Error() string {
	return "no applicable tool"
}
