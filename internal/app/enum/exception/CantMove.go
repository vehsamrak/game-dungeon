package exception

type CantMove struct {
}

func (error *CantMove) Error() string {
	return "cant move"
}
