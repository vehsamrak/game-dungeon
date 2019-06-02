package app

type RoomMemoryRepository struct {
	rooms []*Room
}

func (repository RoomMemoryRepository) Create() RoomRepository {
	return &RoomMemoryRepository{
		rooms: []*Room{
			{x: 1, y: 1},
		},
	}
}

func (repository *RoomMemoryRepository) FindByXY(x int, y int) *Room {
	for _, room := range repository.rooms {
		if room.x == x && room.y == y {
			return room
		}
	}

	return nil
}
