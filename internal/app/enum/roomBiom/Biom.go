package roomBiom

import "github.com/vehsamrak/game-dungeon/internal/app/enum/roomFlag"

type Biom string

const (
	Forest   Biom = "forest"
	Sand     Biom = "sand"
	Plain    Biom = "plain"
	Hill     Biom = "hill"
	Mountain Biom = "mountain"
	Water    Biom = "water"
	Cliff    Biom = "cliff"
	Cave     Biom = "cave"
	Air      Biom = "air"
	Swamp    Biom = "swamp"
	Town     Biom = "town"
)

func All() []Biom {
	return []Biom{
		Forest,
		Sand,
		Plain,
		Hill,
		Mountain,
		Water,
		Cliff,
		Cave,
		Air,
		Swamp,
		Town,
	}
}

func (biom Biom) Flags() []roomFlag.Flag {
	flagMap := map[Biom][]roomFlag.Flag{
		Forest:   {roomFlag.Trees},
		Mountain: {roomFlag.CaveProbability},
		Water:    {roomFlag.FishProbability},
		Sand:     {roomFlag.GemProbability},
		Cliff:    {roomFlag.Unfordable},
		Cave:     {roomFlag.OreProbability, roomFlag.CaveProbability, roomFlag.GemProbability},
		Air:      {roomFlag.Unfordable},
	}

	return flagMap[biom]
}

func (biom Biom) String() string {
	return string(biom)
}

func (biom Biom) ExploreProbability() (probabilityPercentage int) {
	probabilityMap := map[Biom]int{
		Plain:    15,
		Hill:     15,
		Forest:   15,
		Water:    10,
		Air:      10,
		Swamp:    7,
		Sand:     5,
		Mountain: 5,
		Cave:     2,
		Cliff:    1,
		Town:     1,
	}

	probabilityPercentage, ok := probabilityMap[biom]

	if !ok {
		return 10
	}

	return
}
