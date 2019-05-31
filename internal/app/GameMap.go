package app

type GameMap struct {
	height int
	width  int
}

func (gameMap GameMap) Create(height int, width int) *GameMap {
	return &GameMap{height: height, width: width}
}

func (gameMap *GameMap) GetHeightAndWidth() (height int, width int) {
	return gameMap.height, gameMap.width
}
