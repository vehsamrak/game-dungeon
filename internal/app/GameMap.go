package app

type GameMap struct {
}

func (gameMap GameMap) Create() *GameMap {
	return &GameMap{}
}
