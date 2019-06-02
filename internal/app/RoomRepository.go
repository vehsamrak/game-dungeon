package app

type RoomRepository struct {
}

func (repository RoomRepository) Create() *RoomRepository {
	return &RoomRepository{}
}
