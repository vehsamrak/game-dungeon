package mapGenerator

type Factory struct {
}

func (factory Factory) Create() *Factory {
	return &Factory{}
}

func (factory *Factory) CreateGenerator(name Name) (generator MapGenerator) {
	generators := map[Name]MapGenerator{
		Water: &WaterGenerator{},
	}

	generator, _ = generators[name]

	return
}
