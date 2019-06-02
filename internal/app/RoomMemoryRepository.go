package app

type RoomMemoryRepository struct {
	rooms []*Room
}

func (repository RoomMemoryRepository) Create(rooms []*Room) RoomRepository {
	if rooms == nil {
		rooms = []*Room{
			{x: -1, y: 0},
			{x: 1, y: 1},
		}
	}

	return &RoomMemoryRepository{rooms: rooms}
}

func (repository *RoomMemoryRepository) FindByXY(x int, y int) *Room {
	for _, room := range repository.rooms {
		if room.x == x && room.y == y {
			return room
		}
	}

	return nil
}