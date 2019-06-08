package exception

type OreNotFound struct {
}

func (error OreNotFound) Error() string {
	return "ore not found"
}
