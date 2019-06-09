package roomBiom

import "github.com/vehsamrak/game-dungeon/internal/app/enum/roomFlag"

type Biom string

const (
	Forest   Biom = "forest"
	Sand     Biom = "sand"
	Plain    Biom = "plain"
	Hill     Biom = "hill"
	Mountain Biom = "mountain"
	Sea      Biom = "sea"
)

func All() []Biom {
	return []Biom{
		Forest,
		Sand,
		Plain,
		Hill,
		Mountain,
		Sea,
	}
}

func (biom Biom) Flags() []roomFlag.Flag {
	flagMap := map[Biom][]roomFlag.Flag{
		Forest:   {roomFlag.Trees},
		Mountain: {roomFlag.OreProbability},
	}

	return flagMap[biom]
}
