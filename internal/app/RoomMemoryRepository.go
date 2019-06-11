package app

type RoomMemoryRepository struct {
	rooms []*Room
}

func (repository RoomMemoryRepository) Create(rooms []*Room) RoomRepository {
	if rooms == nil {
		rooms = []*Room{
			{x: -1, y: 0, z: 0},
			{x: 1, y: 1, z: 0},
		}
	}

	return &RoomMemoryRepository{rooms: rooms}
}

func (repository *RoomMemoryRepository) FindByXYandZ(x int, y int, z int) *Room {
	for _, room := range repository.rooms {
		if room.x == x && room.y == y && room.z == z {
			return room
		}
	}

	return nil
}

func (repository *RoomMemoryRepository) FindByXYZ(XYZ XYInterface) *Room {
	return repository.FindByXYandZ(XYZ.X(), XYZ.Y(), XYZ.Z())
}

func (repository *RoomMemoryRepository) AddRoom(room *Room) {
	// TODO: lock
	repository.rooms = append(repository.rooms, room)
	// TODO: unlock
}
